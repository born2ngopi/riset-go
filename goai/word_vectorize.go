package goai

import (
	"regexp"
	"strings"
)

type RegexReplacer struct {
	Pattern  string
	Replacer string
}

type WordVectorize struct {
	Lower           bool
	Data            map[string]uint64
	CleanedCorpuses map[string][]string
	RegexReplacers   []RegexReplacer
}

func NewWordVectorize(lower bool) *WordVectorize {
	return &WordVectorize{
		Lower: lower,
		Data: make(map[string]uint64),
		CleanedCorpuses: make(map[string][]string),
		RegexReplacers: []RegexReplacer{
			{Pattern: `[^a-zA-Z0-9 ]+`, Replacer: ``},
			{Pattern: `\s+`, Replacer: ` `},
		},
	}
}

func (vec *WordVectorize) Learn(corpuses map[string][]string) error {

	for corpusesClass, corpus := range corpuses {
		for _, document := range corpus {

			cleanedDocument, err := vec.Normalize(document)
			if err != nil {
				return err
			}

			tokenizeWords := strings.Split(cleanedDocument, " ")
			for _, word := range tokenizeWords {
				if _, exist := vec.Data[word]; !exist {
					vec.Data[word] = uint64(len(vec.Data))
				}
			}

			vec.CleanedCorpuses[corpusesClass] = append(vec.CleanedCorpuses[corpusesClass], cleanedDocument)
		}
	}

	return nil
}


func (vec *WordVectorize) Normalize(document string) (string, error) {
	if vec.Lower {
		document = strings.ToLower(document)
	}

	for _, regexReplacer := range vec.RegexReplacers {
		reg, err := regexp.Compile(regexReplacer.Pattern)
		if err != nil {
			return "", err
		}

		document = reg.ReplaceAllString(document, regexReplacer.Replacer)
	}

	return document, nil
}

func (vec *WordVectorize) GetVectorizeWord() map[string]uint64 {
	return vec.Data
}

func (vec *WordVectorize) GetCleanedCorpuses() map[string][]string {
	return vec.CleanedCorpuses
}