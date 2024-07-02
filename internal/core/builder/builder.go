package builder

import (
	"cpk/internal/configs"
	"cpk/internal/core"
	"cpk/internal/downloader"
	"cpk/internal/utils"
	"cpk/internal/yamls"
)

func Build(instructions yamls.Instructions) {
	var command string
	downloader.Download(instructions.AppRepo)
	for _, depend := range instructions.Dependencies {
		if depend.FromPacageManager != "None" {
			switch depend.FromPacageManager {
			case "apt":
				core.Setup("apt install " + depend.Name)
			case "nix-env":
				core.Setup("nix-env -iA nixos." + depend.Name)
			case "nix-shell":
				core.Setup("nix-shell -p " + depend.Name)
			}
		} else if depend.Build && depend.URL != "" {
			utils.Run_command("git clone " + instructions.AppRepo + " " + configs.Cpks_Settings.Install_dir + "git/" + instructions.AppName)
			command = core.Setup("cd ~/.cpks/packages/git; make")
			core.Exec_and_move(command, &instructions)
		}
	}

	if instructions.MakefileName != "" {
		utils.Run_command("git clone " + instructions.AppRepo + " " + configs.Cpks_Settings.Install_dir + "git/" + instructions.AppName)
		command = core.Setup("cd ~/.cpks/packages/git; make")
	} else {
		command = core.Setup(instructions.Build.EnvParams + ";" + instructions.Build.Command)
	}

	core.Exec_and_move(command, &instructions)
}
