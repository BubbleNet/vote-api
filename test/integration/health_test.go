//go:build integration

package integration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/health")
	if assert.NoError(t, err) {
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	}
	assert.True(t, true)
}
