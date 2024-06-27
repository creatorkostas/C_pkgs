package configs

import (
	"cpk/internal/utils"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Install_dir string `yaml:"install_dir"`
}

var Cpks_Settings Settings

func Load_settings() {
	data, err := os.ReadFile("./configs/settings.yaml")

	utils.Log_debug("\n" + string(data) + "\n")

	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &Cpks_Settings); err != nil {
		panic(err)
	}

	// print the fields to the console
	utils.Log_debug("settings:")
	utils.Log_debug(Cpks_Settings)

}
