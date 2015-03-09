package hooker

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

var pwd string

// set the path of working directory so we have a context
func init() {
	cmd := exec.Command("pwd")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	pwd = out.String()
}

// Will hook this into the commit-msg hook on a github repository to change
// a git commit message automatically to the emoji conversions
func SellSelf() {
	gitStatusCmd := exec.Command("git", "status")
	var out bytes.Buffer
	err := gitStatusCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}

// Will uninstall the hook from the commit-msg hooks
func GiveRefund() {

}

func isGitRepository() bool {
	return false
}
