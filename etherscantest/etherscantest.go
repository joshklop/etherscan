package etherscantest

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func New(t *testing.T) string {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryMap, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		t.Log(queryMap)

		module := queryMap["module"][0]
		var entity string
		switch module {
		case "contract":
			entity = strings.ToLower(queryMap["address"][0])
		}
		action := queryMap["action"][0]
		testdataPath := os.Getenv("ETHERSCANTEST_TESTDATA")

		path := filepath.Join(testdataPath, module, entity, action+".json")
		t.Logf(path)
		read, err := os.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(read)
		require.NoError(t, err)
	}))
	t.Cleanup(srv.Close)
	return srv.URL
}
