package hooker

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ddavison/emotes2emoji/utils"
)

var pwd string

// set the path of working directory so we have a context
func init() {
	pwd, _ = os.Getwd()
}

// Will hook this into the commit-msg hook on a github repository to change
// a git commit message automatically to the emoji conversions
func SellSelf() {
	if IsGitRepository() {
		// add the prepare-commit-msg hook
		hook := "# turn emotes to emoji https://github.com/ddavison/emotes2emoji\nemotes2emoji \"`cat $1`\" > \"$1\""
		f, err := os.OpenFile(pwd+"/.git/hooks/prepare-commit-msg", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0667)
		defer f.Close()

		_, werr := f.WriteString(hook)
		if werr != nil || err != nil {
			log.Fatal("Couldn't add the prepare-commit-msg hook!", werr, err)
		}
	} else {
		fmt.Println(pwd, "is not a git repository")
	}
}

// Will uninstall the hook from the commit-msg hooks
func GiveRefund() {

}

func IsGitRepository() bool {
	return utils.FileExists(pwd + "/.git")
}

func execCommand(cmd string, cmdArgs ...string) string {
	command := exec.Command(cmd, cmdArgs...)
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}
