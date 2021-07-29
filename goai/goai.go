package goai

type Config struct {
	IsLower bool
	IsBinary bool
	Smoot bool
	NormalizeType string
	Corpuses map[string][]string
}

type GOAI interface {
	Learn(corpuses map[string][]string)
	Predict(data interface{}) ([]string, error)
}

func New(cfg Config) GOAI {
	return &cfg
}

func (cfg *Config) Learn(corpuses map[string][]string) {
	cfg.Corpuses = corpuses
}

func (cfg *Config) Predict(data interface{}) ([]string, error) {

	wordVectorize := NewWordVectorize(cfg.IsLower)

	if err := wordVectorize.Learn(cfg.Corpuses); err != nil {
		return nil, err
	}

	frequence := NewTermFrequence(cfg.IsBinary, wordVectorize)
	frequence.Learn(wordVectorize.GetCleanedCorpuses())

	freqIDF, err := NewTermFrequenceIDF(cfg.Smoot, cfg.NormalizeType, frequence)
	if err != nil {
		return nil,  err
	}

	if err := freqIDF.Fit(); err != nil {
		return nil, err
	}


	predictable := NewPredict(freqIDF)
	return predictable.Predict(data)
}


