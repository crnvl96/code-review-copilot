package model

type Instanciator interface {
	Run() error
}

type Loader interface {
	Load() (ModelArgs, error)
}

type Parser interface {
	Parse() (ModelParams, error)
}

type Validator interface {
	Validate() (ModelArgs, error)
}
