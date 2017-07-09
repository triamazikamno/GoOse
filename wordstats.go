package goose

//some word statistics
type wordStats struct {
	//total number of stopwords or good words that we can calculate
	stopWordCount int
	//total number of words on a node
	wordCount int
	//holds an actual list of the stop words we found
	stopWords map[string]struct{}
}

func (w *wordStats) getStopWords() map[string]struct{} {
	return w.stopWords
}

func (w *wordStats) setStopWords(stopWords map[string]struct{}) {
	w.stopWords = stopWords
}

func (w *wordStats) getStopWordCount() int {
	return w.stopWordCount
}

func (w *wordStats) setStopWordCount(stopWordCount int) {
	w.stopWordCount = stopWordCount
}

func (w *wordStats) getWordCount() int {
	return w.wordCount
}

func (w *wordStats) setWordCount(wordCount int) {
	w.wordCount = wordCount
}
