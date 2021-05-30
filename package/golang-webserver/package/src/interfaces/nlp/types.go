package nlp

type NLPAnalyzer struct {
	a string
}

type NLPInteractor struct {
	Logger   Logger
	Analyzer NLPAnalyzer
}

type Logger interface {
	Log(message string) error
}
