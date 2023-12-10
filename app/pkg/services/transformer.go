package services

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/interfaces/transformer_interfaces"
	"github.com/amolofos/tradesor/pkg/models/models_outputType"

	"github.com/amolofos/tradesor/pkg/features/facebook"
	"github.com/amolofos/tradesor/pkg/features/tradesor"
	"github.com/amolofos/tradesor/pkg/features/woocommerce"
)

type Transformer struct{}

func NewTransformer() (t *Transformer) {
	t = &Transformer{}
	return
}

func (t *Transformer) Transform(xmlDoc *tradesor.ModelXml, outputType models_outputType.OutputType) (nProducts int, doc canonical_models.CanonicalModel, err error) {
	var transformer transformer_interfaces.Transformer

	switch outputType {
	case models_outputType.Facebook:
		transformer, err = facebook.NewFacebookService()
	case models_outputType.Woocommerce:
		transformer, err = woocommerce.NewWoocommerceService()
	default:
		errStr := fmt.Sprintf("Output type %s is not supported.", outputType)
		err = errors.New(errStr)
	}

	if err != nil {
		slog.Error(err.Error())
		return
	}
	return transformer.CanonicalModel(xmlDoc)
}
