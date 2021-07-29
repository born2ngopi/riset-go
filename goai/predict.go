package goai

import "math"

type MultinomialPredict struct {
	Evaluator *TermFrequenceInverseDocumentFrequence
}

func NewPredict(freqIDF *TermFrequenceInverseDocumentFrequence) *MultinomialPredict {
	return &MultinomialPredict{
		Evaluator: freqIDF,
	}
}

func (mp *MultinomialPredict) Predict(data interface{}) ([]string, error) {

	var predicted []string
	probabilities, err := mp.PredictProbabilities(data)
	if err != nil {
		return nil, err
	}

	for _, prob := range probabilities {
		var (
			highestProb   float64 = 0
			selectedClass string
		)

		for key, value := range prob {
			if highestProb < value {
				selectedClass = key
				highestProb = value
			}
		}

		predicted = append(predicted, selectedClass)
	}

	return predicted, nil
}

func (mp *MultinomialPredict) PredictProbabilities(data interface{}) ([]map[string]float64, error) {

	evaluatedDatas, err := mp.Evaluator.EvaluateData(data)
	if err != nil {
		return nil, err
	}

	var allPrediciton []map[string]float64

	for _, evalidatedData := range evaluatedDatas {
		var (
			predictedClass         = make(map[string]float64)
			denominator    float64 = 0
		)

		for corpusesClass, _ := range mp.Evaluator.GetTrainedData() {

			var (
				predictedClassValue float64 = 1
				totalValueForClass          = mp.Evaluator.GetSumDataOfClass(corpusesClass)
				dictionaryLength            = float64(len(mp.Evaluator.GetDictionary()))
			)

			for i, val := range mp.Evaluator.GetSumVectorDataOfClass(corpusesClass) {
				predictedWordValue := math.Pow((val+1)/(totalValueForClass/dictionaryLength), evalidatedData[i])
				predictedClassValue *= predictedWordValue
			}

			denominator += predictedClassValue
			predictedClass[corpusesClass] = predictedClassValue
		}

		allPrediciton = append(allPrediciton, predictedClass)
	}

	return allPrediciton, nil
}
