package contexts

import "context"

type ModelContextInterface interface {
	Start() context.Context
}

type ModelContext struct{}

func NewModelContext() *ModelContext {
	return &ModelContext{}
}

func (m *ModelContext) Start() context.Context {
	return context.Background()
}
