package main

import (
	"cpk/internal/configs"
	"cpk/internal/configs/global"
	"cpk/internal/core"
	"cpk/internal/downloader"
	"cpk/internal/utils"
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

	fmt.Println(global.TEST)
	fmt.Println(global.DEV_MODE)

	if make {
		var temp = strings.Split(package_url, "/")
		var name = strings.Split(temp[len(temp)-1], ".")
		if package_url != "" || name[1] != "git" {
			// https://github.com/KDesp73/httpd.git
			utils.Run_command("git clone " + package_url + " " + configs.Cpks_Settings.Install_dir + "git/" + name[0])
			// core.Run(command)

			if command != "" {
				command = core.Setup(command)
				utils.Run_command(command)
				utils.Run_command(
					"find " +
						configs.Cpks_Settings.Install_dir +
						"git/ -type f -regex \".*\\.\\(h\\)\" -exec cp -v -t " +
						configs.Cpks_Settings.Install_dir +
						"headers/" + name[0] +
						" {} \\;")

				utils.Run_command(
					"find " +
						configs.Cpks_Settings.Install_dir +
						"git/ -type f -regex \".*\\.\\(o\\|a\\|so\\)\" -exec cp -v -t " +
						configs.Cpks_Settings.Install_dir +
						"libs/" + name[0] +
						" {} \\;")
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
