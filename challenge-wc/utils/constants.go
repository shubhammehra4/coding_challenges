package utils

const CmdName = "ccwc"

const (
	EndOfLine = '\n'
)

type ReadStrategy int

const (
	DefaultStrategy ReadStrategy = iota
	ChunkedStrategy
)

func GetStrategy(strategy string) ReadStrategy {
	switch strategy {
	case "default":
		return DefaultStrategy
	case "chunked":
		return ChunkedStrategy
	default:
		return DefaultStrategy
	}
}

const (
	DefaultChunkSize = 100 * 1024
)
