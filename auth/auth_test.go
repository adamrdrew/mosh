package auth

import (
	"testing"

	"github.com/adamrdrew/mosh/config"
	"github.com/stretchr/testify/assert"
)

var conf config.Config

/* The rest of auth can't be tested wiht my
current level of test fu. It requires actually
talkiung to Plex. In order to mock that out
I'd need to break the whole API up to fit that
*/

func TestGetAuthorizer(t *testing.T) {
	auth := GetAuthorizer(&conf)
	assert.False(t, auth.Authorized)

}
