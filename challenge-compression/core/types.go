package core

type MODE string

const (
	ENCODE MODE = "compressed"
	DECODE MODE = "decompressed"
)

type CompressOptions struct {
	mode       MODE
	filePath   string
	outputPath string
	showStats  bool
}

func NewCompressOptions(filePath string) *CompressOptions {
	return &CompressOptions{
		mode:       ENCODE,
		filePath:   filePath,
		outputPath: "",
	}
}

func (o *CompressOptions) WithMode(mode MODE) *CompressOptions {
	o.mode = mode
	return o
}

func (o *CompressOptions) WithOutputPath(outputPath string) *CompressOptions {
	o.outputPath = outputPath
	return o
}

func (o *CompressOptions) WithShowStats(showStats bool) *CompressOptions {
	o.showStats = showStats
	return o
}
