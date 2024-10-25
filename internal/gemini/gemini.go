package gemini

import (
	"context"

	env "github.com/VinukaThejana/getdrugs/internal/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// NewClient creates a new client for the Gemini API
func NewClient(ctx context.Context, e *env.Env) (client *genai.Client, err error) {
	client, err = genai.NewClient(ctx, option.WithAPIKey(e.GeminiAPIKey))
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewModel creates a new model for the Gemini API
func NewModel(client *genai.Client) (model *genai.GenerativeModel) {
	model = client.GenerativeModel("gemini-1.5-flash-002")
	model.SetTemperature(0)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	return model
}
