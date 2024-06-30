package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Set up echo
func setup() (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

func TestNewHandler(t *testing.T) {
	_ = NewHandler()
	assert.True(t, true)
}

func TestGetHealth(t *testing.T) {
	rec, c := setup()
	h := NewHandler()

	if assert.NoError(t, h.GetHealth(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"status\":\"OK\"}\n", rec.Body.String())
	}
}
