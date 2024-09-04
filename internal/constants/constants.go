package constants

const (
	LlmContainerBaseUrl = "http://localhost:"
	LlmName             = "AI_MODEL_NAME"
	LlmPort             = "AI_PORT"
	LlmPrompt           = "AI_PROMPT"
	LlmTemp             = "AI_TEMPERATURE"
	LlmDefaultPort      = "11434"
	LlmDefaultName      = "tinyllama"
	LlmMaxPromptLen     = "360"
	LlmMinTempVal       = "0.1"
	LlmMaxTempVal       = "1.0"
)

const (
	ErrEmptyLlmName     = "value of \"AI_MODEL_NAME\" must not be empty"
	ErrInvalidLlmName   = "value of \"AI_MODEL_NAME\" must be \"tinyllama\""
	ErrEmptyLlmPort     = "value of \"AI_PORT\" must not be empty"
	ErrInvalidLlmPort   = "value of \"AI_PORT\" must not be " + LlmDefaultPort
	ErrEmptyLlmPrompt   = "value of \"AI_PROMPT\" must not be empty"
	ErrInvalidLlmPrompt = "length of \"AI_PROMPT\" be less than " + LlmMaxPromptLen
	ErrEmptyLlmTemp     = "value of \"AI_TEMPERATURE\" must not be empty"
	ErrInvalidLlmTemp   = "value of \"AI_TEMPERATURE\" must be a valid number between " + LlmMinTempVal + " and " + LlmMaxTempVal
)
