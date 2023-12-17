package tests

import (
	"os"
	"testing"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
	"github.com/amolofos/tradesor/pkg/models/models_outputType"
	"github.com/amolofos/tradesor/pkg/services"

	"github.com/stretchr/testify/assert"
)

func TestEnd2End(t *testing.T) {
	testCases := []struct {
		enabled                    bool
		name                       string
		catalog                    string
		outputTo                   string
		outputType                 models_outputType.OutputType
		expectedDataDir            string
		expectedNProductsImport    int
		expectedNProductsTransform int
		failFast                   bool
	}{
		{
			enabled:                    true,
			name:                       "Facebook, empty catalog",
			catalog:                    "./data/tradesor_data-no_product.xml",
			outputTo:                   "../../build/test/tests_tradesor_data-no_product",
			outputType:                 models_outputType.Facebook,
			expectedDataDir:            "./data/tradesor_data-no_product_expected_csvs",
			expectedNProductsImport:    0,
			expectedNProductsTransform: 0,
			failFast:                   false,
		},
		{
			enabled:                    true,
			name:                       "Facebook, one product catalog",
			catalog:                    "./data/tradesor_data-one_product.xml",
			outputTo:                   "../../build/test/tests_tradesor_data-one_product",
			outputType:                 models_outputType.Facebook,
			expectedDataDir:            "./data/tradesor_data-one_product_expected_csvs",
			expectedNProductsImport:    1,
			expectedNProductsTransform: 1,
			failFast:                   false,
		},
		{
			enabled:                    true,
			name:                       "Facebook, full catalog",
			catalog:                    "./data/tradesor_data-full_catalog.xml",
			outputTo:                   "../../build/test/tests_tradesor_data-full_catalog",
			outputType:                 models_outputType.Facebook,
			expectedDataDir:            "./data/tradesor_data-full_catalog_expected_csvs",
			expectedNProductsImport:    11162,
			expectedNProductsTransform: 11162,
			failFast:                   false,
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

			var importer = services.NewImporter()

			var transformer = services.NewTransformer()

			var exporter = services.NewExporter()

			nProductsImport, docXml, errImport := importer.Import(tc.catalog)
			assert.Nil(t, errImport)
			assert.Equal(t, nProductsImport, tc.expectedNProductsImport)

			nProductsTransform, docCanonical, errTransform := transformer.Transform(docXml, tc.outputType)
			assert.Nil(t, errTransform)
			assert.Equal(t, nProductsTransform, tc.expectedNProductsTransform)

			errExport := exporter.Export(docCanonical, models_outputFormat.CSV, tc.outputTo)
			assert.Nil(t, errExport)

			isContentTheSame, errCompare := isContentTheSame(tc.outputTo, tc.expectedDataDir, tc.failFast)
			assert.Nil(t, errCompare)
			assert.True(t, isContentTheSame)
		})
	}
}
