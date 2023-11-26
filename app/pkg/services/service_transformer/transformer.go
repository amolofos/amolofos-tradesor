package service_transformer

import (
	"github.com/amolofos/tradesor/pkg/features/facebook/facebook_services"
	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"github.com/amolofos/tradesor/pkg/features/woocommerce/woocommerce_services"
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/interfaces/transformer_interfaces"
	"github.com/amolofos/tradesor/pkg/models/models_outputType"
)

type Transformer struct{}

func (t *Transformer) Init() {}

func (t *Transformer) Transform(xmlDoc *tradesor_models.Xml, outputFormat models_outputType.OutputType) (doc canonical_models.CanonicalModel, err error) {
	var transformer transformer_interfaces.Transformer

	switch outputFormat {
	case models_outputType.Facebook:
		transformer = &facebook_services.Facebook{}
	case models_outputType.Woocommerce:
		transformer = &woocommerce_services.Woocommerce{}
	}

	transformer.Init()
	return transformer.Transform(xmlDoc)
}
