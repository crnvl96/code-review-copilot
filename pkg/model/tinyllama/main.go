package tinyllama

func NewTinyLlamaInstance() *modelInstanciator {
	tinyLlamaArgsLoader := newTinyLlamaArgsLoader()
	tinyLlamaArgsValidator := newTinyLlamaArgsValidator(tinyLlamaArgsLoader)
	tinyLlamaArgsParser := newTinyLlamaArgsParser(tinyLlamaArgsValidator)

	return newModelInstanciator(tinyLlamaArgsParser)
}
