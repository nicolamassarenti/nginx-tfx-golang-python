package entities

type Sentence struct {
	Text      string
	NLP NLPAnalysis
	Stats     SentenceStats
}

type NLPAnalysis struct {
	Sentiment float32
	Embedding []float32
	Tokens    []Token
}

type Token struct {
	Lemma string
	POS 	string
}

type SentenceStats struct {
	NumWords   int
	Length     int
	IsQuestion int
}

type NLPAnalyzer interface {
	Sentiment(sentence string) (sentiment float32, err error)
	Tokens(sentence string) (tokens []Token, err error)
	Embedding(sentence string) (embedding []float32, err error)
}