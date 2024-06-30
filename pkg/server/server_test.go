package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateServer(t *testing.T) {
	e := CreateServer()
	routes := e.Routes()
	assert.Len(t, routes, 1)
	assert.Equal(t, routes[0].Method, "GET")
	assert.Equal(t, routes[0].Path, "/")
}
