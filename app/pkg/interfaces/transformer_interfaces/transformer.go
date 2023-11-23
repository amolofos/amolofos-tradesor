package transformer_interfaces

import (
	"amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"amolofos/tradesor/pkg/models/models_csv"
)

type TransformerInterface struct{}

func (t *TransformerInterface) TransformToCsv(xmlDoc *tradesor_models.Xml) (csvDoc *models_csv.Csv, err error) {
	return nil, nil
}
