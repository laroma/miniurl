package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	api "github.com/laroma/miniurl/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI_AddUrl(t *testing.T) {
	const (
		payload            = "{'url': 'https://pkg.go.dev/golang.org/x/tools/cmd/godoc'}"
		expectedBody       = "{'url': 'https://pkg.go.dev/golang.org/x/tools/cmd/godoc'}"
		expectedStatusCode = http.StatusOK
	)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/url", strings.NewReader((payload)))
	rr := httptest.NewRecorder()

	r := httprouter.New()
	api.Bind(r, nil)
	r.ServeHTTP(rr, req)

	assert.Equal(t, expectedStatusCode, rr.Result().StatusCode)
	body, err := io.ReadAll(rr.Result().Body)
	require.NoError(t, err)
	assert.Equal(t, expectedBody, string(body))
}
