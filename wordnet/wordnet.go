package wordnet

import (
	"log"

	"github.com/advancedlogic/go-freeling/models"
	"github.com/fluhus/gostuff/nlp/wordnet"
)

type WN struct {
	wn *wordnet.WordNet
}

type partOfSpeech struct {
	short string
	long  string
}

func getPOS(p string) (pos *partOfSpeech) {

	pos = new(partOfSpeech)

	switch p {

	case "JJ", "JJR", "JJS":
		pos.short = "a" //adjective
		pos.long = "adjective"
		break

	case "NNS", "NN", "NNP", "NP00000", "NP", "NP00G00", "NP00O00", "NP00V00", "NP00SP0", "NNPS":
		pos.short = "n" //noun
		pos.long = "noun"
		break

	case "RB", "RBR", "RBS", "WRB":
		pos.short = "r" //adverb
		pos.long = "adverb"
		break

	case "MD", "VBG", "VB", "VBN", "VBD", "VBP", "VBZ":
		pos.short = "v" //verb
		pos.long = "verb"
		break

	default:
		return nil
	}
	return pos
}

func NewWordNet() *WN {
	wn, err := wordnet.Parse("/data/dict")
	instance := new(WN)

	if err != nil {
		log.Printf("There was an error during parsing WordNet database: %+v", err)
		return nil
	}

	instance.wn = wn
	return instance
}

func (this *WN) Annotate(word string, pos string) []*models.Annotation {
	if this.wn == nil {
		return nil
	}

	wnPOS := getPOS(pos)

	if wnPOS == nil {
		return nil
	}

	result := this.wn.Search(word)[wnPOS.short]

	var annotation []*models.Annotation
	for _, synset := range result {
		annotation = append(annotation, &models.Annotation{wnPOS.long, synset.Word, synset.Gloss})
	}

	return annotation
}
