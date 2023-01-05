//
// Copyright 2023 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	//"encoding/json"
	"net/http"
	"testing"

	fuzz "github.com/AdaLogics/go-fuzz-headers"
	"github.com/spf13/viper"
	//"github.com/go-openapi/loads"

	"github.com/sigstore/rekor/pkg/generated/restapi/operations/entries"
	//"github.com/sigstore/rekor/pkg/generated/restapi/operations"
	"github.com/sigstore/rekor/pkg/generated/models"
)

var (
	conf := var yamlDeepNestedSlices = []byte(`TV:
- title: "The Expanse"
  title_i18n:
    USA: "The Expanse"
    Japan: "エクスパンス -巨獣めざめる-"
`)
)

func init() {
	ConfigureAPI(uint(1337))
}

func FuzzSearchLogQueryHandler(f *testing.F) {
	f.Fuzz(func(t *testing.T, data, httpBody []byte) {
		v := viper.New()
		slq := &models.SearchLogQuery{}
		ff := fuzz.NewConsumer(data)
		ff.GenerateStruct(slq)
		/*if !json.Valid(jsonData) {
			t.Skip()
		}
		rawJson := json.RawMessage(jsonData)
		if err != nil {
			t.Skip()
		}
		document, err := loads.Analyzed(rawJson, "")
		if err != nil {
			t.Skip()
		}
		rekorServerApi := operations.NewRekorServerAPI(document)*/
		params := entries.SearchLogQueryParams{}
		req, err := http.NewRequest("GET", "", bytes.NewReader(httpBody))
		if err != nil {
			t.Skip()
		}
		params.HTTPRequest = req
		params.Entry = slq
		/*h := SearchLogQueryHandlerFunc(func(params entries.SearchLogQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.SearchLogQuery has not yet been implemented")
		})*/
		//params.Entry = entries.NewSearchLogQuery(rekorServerApi.Context(), rekorServerApi.EntriesSearchLogQueryHandler)
		_ = SearchLogQueryHandler(params)
	})
}
