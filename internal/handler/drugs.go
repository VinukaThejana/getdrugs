package handler

import (
	"errors"
	"fmt"
	"net/http"

	env "github.com/VinukaThejana/getdrugs/internal/config"
	"github.com/VinukaThejana/getdrugs/internal/gemini"
	"github.com/VinukaThejana/getdrugs/pkg/lib"
	"github.com/VinukaThejana/getdrugs/pkg/prompts"
	"github.com/VinukaThejana/getdrugs/pkg/types"
	"github.com/bytedance/sonic"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
)

// GetDrugs is used to extract drugs, dosage, sickness type from the doctor given prescription
func GetDrugs(
	w http.ResponseWriter,
	r *http.Request,
	e *env.Env,
) {
	const (
		// 15 MB in bytes
		maxRequestBodySize = 15 << 20
	)
	err := r.ParseMultipartForm(maxRequestBodySize)
	if err != nil {
		http.Error(w, "request body too large", http.StatusRequestEntityTooLarge)
		return
	}
	defer r.MultipartForm.RemoveAll()

	client, err := gemini.NewClient(r.Context(), e)
	if err != nil {
		log.Error().Err(err).Msg("failed to create a gemini client")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "invalid file", http.StatusBadRequest)
			return
		}
		filename := "prescription"
		defer file.Close()

		err = lib.ValidateFileType(file)
		if err != nil {
			if errors.Is(err, lib.ErrFileType) {
				http.Error(w, "invalid file type, only PNG files are allowed", http.StatusBadRequest)
				return
			}
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		uploadedFile, err := client.UploadFile(r.Context(), filename, file, &genai.UploadFileOptions{})
		if err != nil {
			log.Error().Err(err).Msg("failed to upload the file to gemini")
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}
		defer client.DeleteFile(r.Context(), filename)

		url = uploadedFile.URI
	}

	model := gemini.NewModel(client)

	resp, err := model.GenerateContent(r.Context(), genai.FileData{URI: url}, genai.Text(prompts.ReadDoctorPrescriptionV1))
	if err != nil {
		log.Error().Err(err).Msg("failed to generate content")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}

	var data string
	for _, part := range resp.Candidates[0].Content.Parts {
		data += fmt.Sprintf("%v\n", part)
	}

	var payload types.Prescription
	err = sonic.UnmarshalString(data, &payload)
	if err != nil {
		log.Error().Err(err).Str("data", data).Msg("failed to unmarshal the data")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		sonic.ConfigDefault.NewEncoder(w).Encode(types.Prescription{
			Medicine: []types.Medicine{},
			Causes:   []string{},
			Fatality: -1,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	sonic.ConfigDefault.NewEncoder(w).Encode(payload)
	return
}
