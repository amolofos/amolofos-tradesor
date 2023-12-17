package services

import (
	"path/filepath"

	"github.com/amolofos/tradesor/pkg/conf"
	"github.com/amolofos/tradesor/pkg/features/tradesor"

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
		enabled           bool
		name              string
		catalog           string
		isErrorExpected   bool
		expectedNProducts int
		expectedDoc       *tradesor.ModelXml
	}{
		{
			enabled:           true,
			name:              "Importer, non existent remote",
			catalog:           "https://unknown.unknown/",
			isErrorExpected:   true,
			expectedNProducts: 0,
			expectedDoc:       nil,
		},
		{
			enabled:           true,
			name:              "Importer, empty response from remote",
			catalog:           "https://google.com",
			isErrorExpected:   false,
			expectedNProducts: 0,
			expectedDoc: &tradesor.ModelXml{
				Tradesor: tradesor.Xml{
					CreatedAt: "",
					Products: tradesor.Products{
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
			var importer = NewImporter()
			importer.localFile = filepath.Join("../../../../", conf.LOCAL_BUILD_DIR, conf.LOCAL_FILE)

			nProducts, doc, errImport := importer.Import(tc.catalog)
			if tc.isErrorExpected {
				assert.NotNil(t, errImport)
			} else {
				assert.Nil(t, errImport)
			}

			assert.Equal(t, tc.expectedNProducts, nProducts)
			assert.Equal(t, tc.expectedDoc, doc)
		})
	}
}
