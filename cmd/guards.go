package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/moshd"
)

//Many activities need to be guarded on the daemon being alive
func guardDaemonIsRunning() bool {
	proc := moshd.CheckDaemonStatus()
	return proc != nil
}

//Many acitivities need to be guarded on setup being complete
func guardSetupCompleted() bool {
	conf := config.GetConfig()
	setupDone := conf.Token != config.UNINITIALIZED
	if !setupDone {
		fmt.Println("Plex token not found. Please run setup.")
	}
	return setupDone
}
