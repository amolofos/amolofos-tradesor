package transformer_interfaces

import (
	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
)

type Transformer interface {
	Init()
	Transform(xmlDoc *tradesor_models.Xml) (doc canonical_models.CanonicalModel, err error)
}
