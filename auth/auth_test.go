package auth

import (
	"testing"

	"github.com/adamrdrew/mosh/config"

	"github.com/stretchr/testify/assert"
)

var conf config.Config

func TestGetAuthorizer(t *testing.T) {
	auth := GetAuthorizer(&conf)
	assert.False(t, auth.Authorized)

}
