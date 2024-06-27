package utils

import (
	"os"
	"os/exec"
)

func Run_command(command string) {
	cmd := exec.Command("bash", "-c", command)

	// cmd := exec.Command("vault", "login", "-method=okta", "-format=json", "username=abc")
	// cmd.Stdin = os.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	// stdout, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println(string(stdout))
	// 	fmt.Println(err)
	// 	return
	// }

	// // Print the output
	// fmt.Println(string(stdout))
}
