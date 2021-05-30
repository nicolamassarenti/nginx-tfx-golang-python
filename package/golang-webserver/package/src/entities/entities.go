package entities

type Sentence struct {
	Text      string
	Sentiment string
	Stats     SentenceStats
}

type SentenceStats struct {
	NumWords   int
	Length     int
	IsQuestion int
}
