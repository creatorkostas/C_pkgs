package yamls

type Instructions struct {
	Version      float32 `yaml:"Version"`
	AppName      string  `yaml:"AppName"`
	AppRepo      string  `yaml:"AppRepo"`
	MakefileName string  `yaml:"MakefileName"`
	Dependencies []struct {
		Name              string `yaml:"Name"`
		FromPacageManager string `yaml:"FromPacageManager"`
		URL               string `yaml:"Url"`
		Build             bool   `yaml:"Build"`
	} `yaml:"Dependencies"`
	Build struct {
		EnvParams string `yaml:"EnvParams"`
		Command   string `yaml:"Command"`
	} `yaml:"Build"`
}

type Settings struct {
	Install_dir string `yaml:"install_dir"`
}
