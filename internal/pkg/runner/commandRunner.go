package runner

import (
	"log"
	"os/exec"
	"projectAutomation/internal/common"
	"projectAutomation/internal/config"
)

func CommandRunner(commands map[string]common.Command, execDir string) []error {
	var errs []error

	// Get shell of current running computer
	shell, argFlag := config.GetShell()

	for _, cmd := range commands {
		execCmd := exec.Command(shell, argFlag, cmd.Cmd)
		execCmd.Dir = execDir
		err := execCmd.Run()
		if err != nil {
			log.Printf("Command %s run with this error: %v", execCmd, err)
			errs = append(errs, err)
		}
	}
	return errs
}
