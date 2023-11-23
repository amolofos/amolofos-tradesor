package wordpress_services

import (
	"amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"amolofos/tradesor/pkg/models/models_csv"
)

type Wordpress struct{}

func (w *Wordpress) Init() {}

func (w *Wordpress) TransformToCsv(xmlDoc *tradesor_models.Xml) (csvDoc *models_csv.Csv, err error) {
	return nil, nil
}
