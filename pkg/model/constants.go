package model

const (
	ModelBaseUrl          = "http://localhost:"
	ModelName             = "AI_MODEL_NAME"
	ModelPort             = "AI_PORT"
	ModelPrompt           = "AI_PROMPT"
	ModelTemp             = "AI_TEMPERATURE"
	ModelDefaultPort      = "11434"
	ModelDefaultName      = "tinyllama"
	ModelMaxPromptLength  = "120"
	ModelMinimumTempValue = "0.1"
	ModelMaximunTempValue = "1.0"
	ErrEmptyModelName     = "value of \"AI_MODEL_NAME\" must not be empty"
	ErrInvalidModelName   = "value of \"AI_MODEL_NAME\" must be \"tinyllama\""
	ErrEmptyModelPort     = "value of \"AI_PORT\" must not be empty"
	ErrInvalidModelPort   = "value of \"AI_PORT\" must not be " + ModelDefaultPort
	ErrEmptyModelPrompt   = "value of \"AI_PROMPT\" must not be empty"
	ErrInvalidModelPrompt = "length of \"AI_PROMPT\" be less than " + ModelMaxPromptLength
	ErrEmptyModelTemp     = "value of \"AI_TEMPERATURE\" must not be empty"
	ErrInvalidModelTemp   = "value of \"AI_TEMPERATURE\" must be a valid number between " + ModelMinimumTempValue + " and " + ModelMaximunTempValue
)
