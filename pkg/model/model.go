package model

const (
	BaseUrl          = "http://localhost:"
	Name             = "AI_MODEL_NAME"
	Port             = "AI_PORT"
	Prompt           = "AI_PROMPT"
	Temp             = "AI_TEMPERATURE"
	DefaultPort      = "11434"
	DefaultName      = "tinyllama"
	MaxPromptLen     = "120"
	MinTempVal       = "0.1"
	MaxTempVal       = "1.0"
	ErrEmptyName     = "value of \"AI_MODEL_NAME\" must not be empty"
	ErrInvalidName   = "value of \"AI_MODEL_NAME\" must be \"tinyllama\""
	ErrEmptyPort     = "value of \"AI_PORT\" must not be empty"
	ErrInvalidPort   = "value of \"AI_PORT\" must not be " + DefaultPort
	ErrEmptyPrompt   = "value of \"AI_PROMPT\" must not be empty"
	ErrInvalidPrompt = "length of \"AI_PROMPT\" be less than " + MaxPromptLen
	ErrEmptyTemp     = "value of \"AI_TEMPERATURE\" must not be empty"
	ErrInvalidTemp   = "value of \"AI_TEMPERATURE\" must be a valid number between " + MinTempVal + " and " + MaxTempVal
)

type Params struct {
	Name   string
	Port   string
	Prompt string
	Temp   float64
}

type Args struct {
	Name   string
	Port   string
	Prompt string
	Temp   string
}

type modelInstanciator interface {
	Run() error
}

type model struct {
	Instance modelInstanciator
}

func NewModel(instance modelInstanciator) *model {
	return &model{Instance: instance}
}
