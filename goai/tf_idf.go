package goai

import (
	"errors"
	"math"
)

type TermFrequenceInverseDocumentFrequence struct {
	Smooth                   bool
	InverseDocumentFrequence []float64
	DocumentFrequence        []uint64
	Data                     map[string][][]float64
	SumVectorDataPerClass    map[string][]float64
	SumDataPerClass          map[string]float64
	TotalDocument            uint64
	NormalizeType            string
	Normalizer               []float64
	CountVectorizer          *TermFrequence
}

const (
	UnequalDocumentLength = "UnequalDocumentLength"

	EuclideanSumSquare = "EuclideanSumSquare"
	EuclideanSum       = "EuclideanSum"
)

func NewTermFrequenceIDF(smooth bool, normalizeType string, freq *TermFrequence) (*TermFrequenceInverseDocumentFrequence, error) {

	var (
		totalDocument, documentLength uint64
	)

	for _, corpuses := range freq.VectorizedCounter() {
		for _, corpus := range corpuses {
			totalDocument += 1
			if documentLength == 0 {
				documentLength = uint64(len(corpus))
			} else {
				if documentLength != uint64(len(corpus)) {
					return nil, errors.New(UnequalDocumentLength)
				}
			}
		}
	}

	return &TermFrequenceInverseDocumentFrequence{
		Smooth:                   smooth,
		InverseDocumentFrequence: make([]float64, documentLength),
		DocumentFrequence:        make([]uint64, documentLength),
		Data:                     make(map[string][][]float64),
		SumVectorDataPerClass:    make(map[string][]float64),
		SumDataPerClass:          make(map[string]float64),
		TotalDocument:            totalDocument,
		NormalizeType:            normalizeType,
		CountVectorizer:          freq,
	}, nil
}

func (freqIDF *TermFrequenceInverseDocumentFrequence) Fit() error {

	for _, corpuses := range freqIDF.CountVectorizer.VectorizedCounter() {
		for _, corpus := range corpuses {
			for i, word := range corpus {
				if word != 0 {
					freqIDF.DocumentFrequence[i] += 1
				}
			}
		}
	}

	for i, _ := range freqIDF.InverseDocumentFrequence {
		var (
			numerator   = float64(freqIDF.TotalDocument)
			denominator = float64(freqIDF.DocumentFrequence[i])
		)

		if freqIDF.Smooth {
			numerator += 1
			denominator += 1
		}

		freqIDF.InverseDocumentFrequence[i] = math.Log(float64(numerator/denominator)) + 1
	}

	for corpusesClass, corpuses := range freqIDF.CountVectorizer.VectorizedCounter() {
		var (
			sumValue       float64 = 0
			sumVectorValue         = make([]float64, len(freqIDF.DocumentFrequence))
		)

		for _, corpus := range corpuses {
			newFreqIDF, normalize := freqIDF.Normalize(corpus)

			for i, _ := range corpus {
				newFreqIDF[i] = newFreqIDF[i] / normalize

				sumVectorValue[i] += newFreqIDF[i]
				sumValue += newFreqIDF[i]
			}

			freqIDF.Data[corpusesClass] = append(freqIDF.Data[corpusesClass], newFreqIDF)
		}

		freqIDF.SumVectorDataPerClass[corpusesClass] = sumVectorValue
		freqIDF.SumDataPerClass[corpusesClass] = sumValue
	}

	return nil
}

func (freqIDF *TermFrequenceInverseDocumentFrequence) Normalize(corpus []uint64) ([]float64, float64) {
	var (
		normalize  float64 = 0
		newFreqIDF         = make([]float64, len(corpus))
	)

	for i, word := range corpus {
		newFreqIDF[i] = float64(word) * freqIDF.InverseDocumentFrequence[i]

		if freqIDF.NormalizeType == EuclideanSumSquare {
			normalize += math.Pow(newFreqIDF[i], 2)
		} else if freqIDF.NormalizeType == EuclideanSum {
			normalize += math.Abs(newFreqIDF[i])
		}
	}

	if normalize == 0 {
		normalize = 1
	}

	if freqIDF.NormalizeType == EuclideanSumSquare {
		normalize = math.Sqrt(normalize)
	}

	return newFreqIDF, normalize
}

func (freqIDF *TermFrequenceInverseDocumentFrequence) EvaluateData(data interface{}) ([][]float64, error) {

	var (
		evaluatedData [][]float64
		convertedData = data.([]string)
	)

	vectorizeInput, err := freqIDF.CountVectorizer.Vectorize(convertedData)
	if err != nil {
		return nil, err
	}

	for _, corpus := range vectorizeInput {
		newFreqIDF, normalize := freqIDF.Normalize(corpus)

		for i, _ := range corpus {
			newFreqIDF[i] = newFreqIDF[i] / normalize
		}

		evaluatedData = append(evaluatedData, newFreqIDF)
	}

	return evaluatedData, nil

}

func (freqIDF *TermFrequenceInverseDocumentFrequence) GetTrainedData() map[string][][]float64 {
	return freqIDF.Data
}

func (freqIDF *TermFrequenceInverseDocumentFrequence) GetSumDataOfClass(class string) float64 {
	if val, exist := freqIDF.SumDataPerClass[class]; exist {
		return val
	}

	return 0
}

func (freqIDF *TermFrequenceInverseDocumentFrequence) GetSumVectorDataOfClass(class string) []float64 {
	if val, exist := freqIDF.SumVectorDataPerClass[class]; exist {
		return val
	}

	return nil
}

func (freqIDF *TermFrequenceInverseDocumentFrequence) GetDictionary() map[string]uint64 {
	return freqIDF.CountVectorizer.GetDictionary()
}
