package tests

import (
	"os"
	"testing"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
	"github.com/amolofos/tradesor/pkg/models/models_outputType"
	"github.com/amolofos/tradesor/pkg/services/service_exporter"
	"github.com/amolofos/tradesor/pkg/services/service_importer"
	"github.com/amolofos/tradesor/pkg/services/service_transformer"

	"github.com/stretchr/testify/assert"
)

func TestEnd2End(t *testing.T) {
	testCases := []struct {
		enabled         bool
		name            string
		catalog         string
		outputTo        string
		outputType      models_outputType.OutputType
		expectedDataDir string
		failFast        bool
	}{
		{
			enabled:         true,
			name:            "Facebook, empty catalog",
			catalog:         "./data/tradesor_data-no_product.xml",
			outputTo:        "../../build/test/tests_tradesor_data-no_product",
			outputType:      models_outputType.Facebook,
			expectedDataDir: "./data/tradesor_data-no_product_expected_csvs",
			failFast:        false,
		},
		{
			enabled:         true,
			name:            "Facebook, one product catalog",
			catalog:         "./data/tradesor_data-one_product.xml",
			outputTo:        "../../build/test/tests_tradesor_data-one_product",
			outputType:      models_outputType.Facebook,
			expectedDataDir: "./data/tradesor_data-one_product_expected_csvs",
			failFast:        false,
		},
		{
			enabled:         true,
			name:            "Facebook, full catalog",
			catalog:         "./data/tradesor_data-full_catalog.xml",
			outputTo:        "../../build/test/tests_tradesor_data-full_catalog",
			outputType:      models_outputType.Facebook,
			expectedDataDir: "./data/tradesor_data-full_catalog_expected_csvs",
			failFast:        false,
		},
	}

	t.Parallel()
	for _, tc := range testCases {
		if !tc.enabled {
			return
		}

		t.Run(tc.name, func(t *testing.T) {

			os.RemoveAll(tc.outputTo)
			os.MkdirAll(tc.outputTo, 0744)

			var importer = &service_importer.Importer{}
			importer.Init()

			var transformer = &service_transformer.Transformer{}
			transformer.Init()

			var exporter = &service_exporter.Exporter{}
			exporter.Init()

			doc, errImport := importer.Import(tc.catalog)
			assert.Nil(t, errImport)

			out, errTransform := transformer.Transform(doc, tc.outputType)
			assert.Nil(t, errTransform)

			errExport := exporter.Export(out, models_outputFormat.CSV, tc.outputTo)
			assert.Nil(t, errExport)

			isContentTheSame, errCompare := isContentTheSame(tc.outputTo, tc.expectedDataDir, tc.failFast)
			assert.Nil(t, errCompare)
			assert.True(t, isContentTheSame)
		})
	}
}
