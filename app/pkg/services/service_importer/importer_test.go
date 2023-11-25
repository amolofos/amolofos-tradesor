package service_importer

import (
	"amolofos/tradesor/pkg/conf"
	"amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"path/filepath"

	"github.com/stretchr/testify/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {

	args := m.Called(number)
	return args.Bool(0), args.Error(1)

}

func TestHttpDownloadFails(t *testing.T) {
	testCases := []struct {
		enabled         bool
		name            string
		catalog         string
		isErrorExpected bool
		expectedDoc     *tradesor_models.Xml
	}{
		{
			enabled:         true,
			name:            "Importer, non existent remote",
			catalog:         "https://unknown.unknown/",
			isErrorExpected: true,
			expectedDoc:     nil,
		},
		{
			enabled:         true,
			name:            "Importer, empty response from remote",
			catalog:         "https://google.com",
			isErrorExpected: false,
			expectedDoc: &tradesor_models.Xml{
				Tradesor: tradesor_models.XmlTradesor{
					CreatedAt: "",
					Products: tradesor_models.XmlProducts{
						ProductList: nil,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		if !tc.enabled {
			return
		}

		t.Run(tc.name, func(t *testing.T) {
			var importer = &Importer{}
			importer.Init()
			importer.localFile = filepath.Join("../../../../", conf.LOCAL_BUILD_DIR, conf.LOCAL_FILE)

			doc, errImport := importer.Import(tc.catalog)
			if tc.isErrorExpected {
				assert.NotNil(t, errImport)
			} else {
				assert.Nil(t, errImport)
			}

			assert.Equal(t, tc.expectedDoc, doc)
		})
	}
}
