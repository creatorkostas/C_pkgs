package configs

import (
	"cpk/internal/utils"
	"cpk/internal/yamls"
)

var Cpks_Settings yamls.Settings

func Load_settings() {
	utils.Load_yaml_to_struct("./configs/settings.yaml", &Cpks_Settings)
}
