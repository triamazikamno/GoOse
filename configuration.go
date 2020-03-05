package goose

import (
	"strings"
	"time"
)

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7"

// Configuration is a wrapper for various config options
type Configuration struct {
	localStoragePath        string //not used in this version
	imagesMinBytes          int    //not used in this version
	targetLanguage          string
	imageMagickConvertPath  string //not used in this version
	imageMagickIdentifyPath string //not used in this version
	browserUserAgent        string
	debug                   bool
	extractPublishDate      bool
	additionalDataExtractor bool
	enableImageFetching     bool
	useMetaLanguage         bool

	//path to the stopwords folder
	stopWordsPath string
	stopWords     StopWords
	parser        *Parser

	timeout time.Duration
}

// GetDefaultConfiguration returns safe default configuration options
func GetDefaultConfiguration(args ...string) Configuration {
	if len(args) == 0 {
		return Configuration{
			localStoragePath:        "",   //not used in this version
			imagesMinBytes:          4500, //not used in this version
			enableImageFetching:     true,
			useMetaLanguage:         true,
			targetLanguage:          "en",
			imageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
			imageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
			browserUserAgent:        defaultUserAgent,
			debug:                   false,
			extractPublishDate:      false,
			additionalDataExtractor: false,
			stopWordsPath:           "resources/stopwords",
			stopWords:               NewStopwords(), //TODO with path
			parser:                  NewParser(),
			timeout:                 time.Duration(5 * time.Second),
		}
	}
	return Configuration{
		localStoragePath:        "",   //not used in this version
		imagesMinBytes:          4500, //not used in this version
		enableImageFetching:     true,
		useMetaLanguage:         true,
		targetLanguage:          "en",
		imageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
		imageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
		browserUserAgent:        defaultUserAgent,
		debug:                   false,
		extractPublishDate:      false,
		additionalDataExtractor: false,
		stopWordsPath:           "resources/stopwords",
		stopWords:               NewStopwords(), //TODO with path
		parser:                  NewParser(),
		timeout:                 time.Duration(5 * time.Second),
	}
}

var (
	jaSeparators = []rune{'。', '！', '？', '\n', '、', ' '}
)

func split(s string, separators []rune) []string {
	f := func(r rune) bool {
		for _, s := range separators {
			if r == s {
				return true
			}
		}
		return false
	}
	return strings.FieldsFunc(s, f)
}

func (c Configuration) Words(str string) []string {
	switch c.targetLanguage {
	case "ja":
		words := split(str, jaSeparators)
		for i := range words {
			words[i] = strings.TrimSpace(words[i])
		}
		return split(str, jaSeparators)
	}
	return strings.Split(strings.ToLower(str), " ")
}
