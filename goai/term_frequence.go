package goai

import "strings"

type TermFrequence struct {
	Data     map[string][][]uint64
	IsBinary bool
	*WordVectorize
}

func NewTermFrequence(isBinary bool, workVetorize *WordVectorize) *TermFrequence {
	return &TermFrequence{
		Data:          make(map[string][][]uint64),
		IsBinary:      isBinary,
		WordVectorize: workVetorize,
	}
}

func (freq *TermFrequence) Learn(corpuses map[string][]string) {
	for corpusClass, corpus := range corpuses {
		for _, document := range corpus {
			fr := freq.CountWord(document)
			freq.Data[corpusClass] = append(freq.Data[corpusClass], fr)
		}
	}
}

func (freq *TermFrequence) CountWord(document string) (result []uint64) {

	vectorizeWord := freq.WordVectorize.GetVectorizeWord()

	result = make([]uint64, len(vectorizeWord))

	tokenizeWord := strings.Split(document, "")
	for _, word := range tokenizeWord {
		if _, exist := vectorizeWord[word]; exist {
			if freq.IsBinary {
				result[vectorizeWord[word]] = 1
			} else {
				result[vectorizeWord[word]] += 1
			}
		}
	}

	return
}

func (freq *TermFrequence) Vectorize(corpus []string) ([][]uint64, error) {

	var result [][]uint64

	for _, document := range corpus {
		cleanedDocument, err := freq.WordVectorize.Normalize(document)
		if err != nil {
			return nil, err
		}

		slice := freq.CountWord(cleanedDocument)
		result = append(result, slice)
	}

	return result, nil
}

func (freq *TermFrequence) VectorizedCounter() map[string][][]uint64 {
	return freq.Data
}

func (freq *TermFrequence) GetDictionary() map[string]uint64 {
	return freq.WordVectorize.GetVectorizeWord()
}
