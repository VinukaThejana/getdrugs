package lib

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

var (
	// ErrFileType is used to indicate that the file type is invalid
	ErrFileType = errors.New("only PNG files are allowed")
	// ErrFailedToResetFilePointer is used to indicate that the file pointer could not be reset
	ErrFailedToReadFileHeader = errors.New("failed to read the file header")
	// ErrFileDoesNotImplementSeeker is used to indicate that the file does not implement the seeker
	ErrFileDoesNotImplementSeeker = errors.New("file does not implement the seeker")
	// ErrFailedToResetFilePointer is used to indicate that the file pointer could not be reset
	ErrFailedToResetFilePointer = errors.New("failed to reset the file pointer")
)

// ValidateFileType is used to validate the file type of the uploaded file
func ValidateFileType(file io.Reader) error {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		return ErrFailedToReadFileHeader
	}

	seeker, ok := file.(io.Seeker)
	if !ok {
		return ErrFileDoesNotImplementSeeker
	}
	_, err = seeker.Seek(0, io.SeekStart)
	if err != nil {
		return ErrFailedToResetFilePointer
	}

	filetype := http.DetectContentType(buff)
	if !strings.HasPrefix(filetype, "image/png") {
		return ErrFileType
	}

	return nil
}
