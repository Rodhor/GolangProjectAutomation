package runner

import (
	"fmt"
	"os/exec"
	"projectAutomation/internal/common"
	"projectAutomation/internal/config"
	"projectAutomation/internal/pkg/project"
)

func CommandsOrganiser(p *project.Project, r common.RunTime) []error {
	var errors []error
	// Get shell of current running computer
	shell := GetShell()

	for _, cmd := range p.Language.Commands {
		if cmd.ActualRunTime == r {
			execCmd := AdjustDynamicCommands(cmd.Cmd, p)
			err := CommandRunner(execCmd, shell, p.RootDir)
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func CommandRunner(cmd string, shell config.Shell, execDir string) error {
	execCmd := exec.Command(shell.Shell, shell.ArgFlag, cmd)
	execCmd.Dir = execDir
	err := execCmd.Run()
	if err != nil {
		return fmt.Errorf("command %s run with this error: %v", execCmd, err)
	}
	return nil
}
