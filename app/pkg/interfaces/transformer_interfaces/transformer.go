package transformer_interfaces

import (
	"github.com/amolofos/tradesor/pkg/features/tradesor"
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
)

type Transformer interface {
	Init()
	CanonicalModel(xmlDoc *tradesor.ModelXml) (nProducts int, doc canonical_models.CanonicalModel, err error)
}
