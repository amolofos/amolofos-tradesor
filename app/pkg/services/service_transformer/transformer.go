package service_transformer

import (
	"github.com/amolofos/tradesor/pkg/features/facebook/facebook_services"
	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"github.com/amolofos/tradesor/pkg/features/wordpress/wordpress_services"
	"github.com/amolofos/tradesor/pkg/models/models_csv"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type Transformer struct{}

func (t *Transformer) Init() {}

func (t *Transformer) Transform(xmlDoc *tradesor_models.Xml, outputFormat models_outputFormat.OutputFormat) (csvDoc *models_csv.Csv, err error) {
	switch outputFormat {
	case models_outputFormat.Facebook:
		f := &facebook_services.Facebook{}
		f.Init()
		return f.TransformToCsv(xmlDoc)
	case models_outputFormat.Wordpress:
		w := &wordpress_services.Wordpress{}
		w.Init()
		return w.TransformToCsv(xmlDoc)
	}

	return nil, nil
}
