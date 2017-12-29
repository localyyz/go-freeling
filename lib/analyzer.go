package lib

import (
	"github.com/advancedlogic/go-freeling/engine"
	"github.com/advancedlogic/go-freeling/models"
)

type Analyzer struct {
	context *engine.Context
}

func NewAnalyzer() *Analyzer {
	context := engine.NewContext()
	context.InitNLP()
	instance := new(Analyzer)
	instance.context = context

	return instance
}

func (this *Analyzer) AnalyzeText(document *models.DocumentEntity) *models.DocumentEntity {
	ch := make(chan *models.DocumentEntity)
	defer close(ch)

	go this.context.Engine.NLP.Workflow(document, ch)
	output := <-ch

	return output
}
