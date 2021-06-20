package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {

	router := gin.Default()
	request, err := http.NewRequest(http.MethodPost, "/v1/workedHours", bytes.NewBufferString(`{"TemperatureCelsius":10, "Summary":"Worked Hours Test"}`))
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.EqualValues(recorder.Code, http.StatusOK)
}
