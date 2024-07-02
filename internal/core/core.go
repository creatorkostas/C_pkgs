package core

import (
	"cpk/internal/configs"
	"cpk/internal/utils"
	"cpk/internal/yamls"
	"os"
)

func Exec_and_move(command string, instructions *yamls.Instructions) {
	utils.Run_command(command)
	os.MkdirAll(configs.Cpks_Settings.Install_dir+"headers/"+instructions.AppName, 0700)
	os.MkdirAll(configs.Cpks_Settings.Install_dir+"libs/"+instructions.AppName, 0700)
	os.MkdirAll(configs.Cpks_Settings.Install_dir, 0700)
	utils.Run_command(
		"find " +
			configs.Cpks_Settings.Install_dir +
			"git/ -type f -regex \".*\\.\\(h\\)\" -exec cp -v -t " +
			configs.Cpks_Settings.Install_dir +
			"headers/" + instructions.AppName +
			" {} \\;")

	utils.Run_command(
		"find " +
			configs.Cpks_Settings.Install_dir +
			"git/ -type f -regex \".*\\.\\(o\\|a\\|so\\)\" -exec cp -v -t " +
			configs.Cpks_Settings.Install_dir +
			"libs/" + instructions.AppName +
			" {} \\;")
}
