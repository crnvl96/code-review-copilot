package spec

import (
	"reflect"
	"testing"

	"github.com/omega-energia/code-review-copilot/internal/env"
	"github.com/omega-energia/code-review-copilot/pkg/validation"
)

type MockEnvLoader struct{}

func NewMockEnvLoader() *MockEnvLoader {
	return &MockEnvLoader{}
}

func (m *MockEnvLoader) Load() {}

type MockValidator struct{}

func NewMockValidator() *MockValidator {
	return &MockValidator{}
}

func (m *MockValidator) ModelName(input string) (string, error) {
	return input, nil
}

func (m *MockValidator) ModelPort(input string) (string, error) {
	return input, nil
}

func (m *MockValidator) ModelPrompt(input string) (string, error) {
	return input, nil
}

func (m *MockValidator) ModelTemperature(input string) (string, error) {
	return input, nil
}

func (m *MockValidator) ModelBaseURL(input string) (string, error) {
	return input, nil
}

func TestSpec_FromEnv(t *testing.T) {
	type fields struct {
		validator validation.ValidatorInterface
		envLoader env.EnvLoaderInterface
	}

	tests := []struct {
		env    func()
		name   string
		fields fields
		want   SpecConstants
	}{
		{
			name: "Should successfully return the EnvConstants",
			fields: fields{
				validator: NewMockValidator(),
				envLoader: NewMockEnvLoader(),
			},
			want: SpecConstants{
				AiModelName:   "tinyllama",
				AiPort:        "1111",
				AiPrompt:      "prompt",
				AiTemperature: "0.5",
				AiBaseUrl:     "http://localhost:",
			},
			env: func() {
				t.Setenv("AI_MODEL_NAME", "tinyllama")
				t.Setenv("AI_PORT", "1111")
				t.Setenv("AI_PROMPT", "prompt")
				t.Setenv("AI_TEMPERATURE", "0.5")
				t.Setenv("AI_BASE_URL", "http://localhost:")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.env()

			e := &Spec{
				validator: tt.fields.validator,
				envLoader: tt.fields.envLoader,
			}

			if got := e.FromEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spec.FromEnv():\n\tGot: %v\n\tWant: %v", got, tt.want)
			}
		})
	}
}
