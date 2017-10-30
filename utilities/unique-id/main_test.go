package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	os.Setenv("MY_IP", "0.0.0.0")
	r = newEngine()
	os.Exit(m.Run())
}

func TestUniqueID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "they should be equal")

	schema := gojsonschema.NewStringLoader(`{"id": "string"}`)
	body := w.Body.String()
	doc := gojsonschema.NewStringLoader(body)
	if result, err := gojsonschema.Validate(schema, doc); err != nil {
		t.Fatalf("invalid response: %s", body)
	} else {
		if !result.Valid() {
			t.Fatalf("invalid schema: %s", body)
		}
	}
}
