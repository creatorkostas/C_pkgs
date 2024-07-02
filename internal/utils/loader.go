package utils

import (
	"cpk/internal/yamls"
	"os"

	"gopkg.in/yaml.v3"
)

func Load_yaml_to_struct[T yamls.Instructions | yamls.Settings](yaml_path string, struct_data *T) *T {
	data, err := os.ReadFile(yaml_path)

	Log_debug("data:\n" + string(data) + "\n")

	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, struct_data); err != nil {
		panic(err)
	}

	// print the fields to the console
	Log_debug("instructions:")
	Log_debug(*struct_data)
	return struct_data
}
