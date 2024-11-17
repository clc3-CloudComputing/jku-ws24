package format

import (
	"os"
	"strings"
)

const OUTPUT_FORMAT_ENV_VAR_NAME = "FORMAT"

type OutputFormat int

const (
	Plain OutputFormat = iota
	Html
)

func GetOutputFormat() OutputFormat {
	return parseString(os.Getenv(OUTPUT_FORMAT_ENV_VAR_NAME))
}

var (
	outputFormatMap = map[string]OutputFormat{
		"plain": Plain,
		"html":  Html,
	}
)

func parseString(str string) OutputFormat {
	return outputFormatMap[strings.ToLower(str)]
}
