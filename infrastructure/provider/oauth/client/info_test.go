package client_test

import (
	"testing"

	"github.com/nurmanhabib/go-oauth2-server/infrastructure/provider/oauth/client"
	"github.com/stretchr/testify/assert"
)

func TestApplicationInfo_GetDomain(t *testing.T) {
	var clientInfo *client.Info

	assert.Empty(t, clientInfo.GetDomain())
}
