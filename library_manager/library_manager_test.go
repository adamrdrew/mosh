package library_manager

import (
	"testing"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/server"
	"github.com/stretchr/testify/assert"
)

func setup() LibrarySelector {
	conf := config.GetConfig()
	serv := server.GetServer(&conf)
	return LibrarySelector{
		Config: &conf,
		Server: serv,
	}
}

/*
This is kind of a weird case.
library_manager only has one public method: SetLibraryInConf()
But we can't test SetLibraryInConf() because its mutative.

So, instead we're gonna test private methods which I'm usually against.
*/

func TestIsLibInConfNegative(t *testing.T) {
	ls := setup()
	got := ls.isLibInConf()
	assert.False(t, got)
}

func TestIsLibInConfPositive(t *testing.T) {
	ls := setup()
	ls.Config.Library = "3333"
	got := ls.isLibInConf()
	assert.True(t, got)
}
