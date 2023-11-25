package wordpress_services

import (
	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"github.com/amolofos/tradesor/pkg/models/models_csv"
)

type Wordpress struct{}

func (w *Wordpress) Init() {}

func (w *Wordpress) TransformToCsv(xmlDoc *tradesor_models.Xml) (csvDoc *models_csv.Csv, err error) {
	return nil, nil
}
