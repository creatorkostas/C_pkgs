package core

import (
	"cpk/internal/configs"
	"cpk/internal/configs/global"
	"fmt"
	"os/exec"
)

func Setup(command string) string {

	var command_to_run = "export TEST=" + configs.Cpks_Settings.Install_dir + "libs/ ;"
	if global.DEV_MODE {
		command_to_run = command_to_run + "env ;" + command
	} else {
		command_to_run = command_to_run + command
	}

	return command_to_run
}

func Run(command string) {
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
