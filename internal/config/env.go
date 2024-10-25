package env

import environ "github.com/VinukaThejana/env"

// Env is the struct that holds the environment variables
type Env struct {
	GeminiAPIKey string `mapstructure:"GEMINI_API_KEY" validate:"required"`
	Port         string `mapstructure:"PORT" validate:"required"`
	Env          string `mapstructure:"ENV" validate:"required,oneof=dev stg prod"`
}

// Load loads the environment variables from the given path
func (e *Env) Load(path ...string) {
	environ.Load(e, path...)
}
