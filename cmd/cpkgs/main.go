package main

import (
	"cpk/internal/configs"
	"cpk/internal/configs/global"
	"cpk/internal/core"
	"cpk/internal/downloader"
	"cpk/internal/utils"
	"cpk/internal/yamls"
	"flag"
	"fmt"
	"os"
	"strings"
)

func cpks_init() {
	configs.Load_settings()
	os.MkdirAll(configs.Cpks_Settings.Install_dir, 0700)
	os.MkdirAll(configs.Cpks_Settings.Install_dir+"headers/", 0700)
	os.MkdirAll(configs.Cpks_Settings.Install_dir+"libs/", 0700)
	os.MkdirAll(configs.Cpks_Settings.Install_dir+"git/", 0700)
}

func main() {
	cpks_init()

	var package_url string = ""
	var command string = ""
	// var dir string = ""
	var make bool = false

	flag.BoolVar(&global.DEV_MODE, "D", global.DEV_MODE, "enable dev mode")
	flag.BoolVar(&global.TEST, "t", global.TEST, "for test")
	flag.BoolVar(&make, "M", make, "use the make file")
	flag.StringVar(&package_url, "url", package_url, "package to install")
	flag.StringVar(&command, "run", command, "run command with cpks")
	// flag.StringVar(&dir, "run", dir, "run command with cpks")
	flag.Parse()

	fmt.Print("global.TEST: ")
	fmt.Println(global.TEST)
	fmt.Print("global.DEV_MODE: ")
	fmt.Println(global.DEV_MODE)

	// builder.Load_Instructions("./tests/software_config.yaml")

	if make {
		var temp = strings.Split(package_url, "/")
		var name = strings.Split(temp[len(temp)-1], ".")
		if package_url != "" || name[1] != "git" {
			utils.Run_command("git clone " + package_url + " " + configs.Cpks_Settings.Install_dir + "git/" + name[0])

			var instructions yamls.Instructions
			instructions.AppName = name[0]
			if command != "" {
				command = core.Setup(command)
				core.Exec_and_move(command, &instructions)
			} else {
				command = core.Setup("cd " + configs.Cpks_Settings.Install_dir + "git/" + name[0] + " ; " + "make")
				core.Exec_and_move(command, &instructions)
			}

		} else {
			fmt.Println("Empty package url or the url is not a git repo!")
		}

	} else {
		if package_url != "" {
			downloader.Download(package_url)
		} else if command != "" {
			command = core.Setup(command)
			utils.Run_command(command)
			// core.Run(command)
		}
	}

}
