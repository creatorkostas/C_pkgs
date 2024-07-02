package builder

import (
	"cpk/internal/utils"
	"cpk/internal/yamls"
)

func Load_Instructions(instructions_path string) yamls.Instructions {
	var instructions yamls.Instructions
	utils.Load_yaml_to_struct(instructions_path, &instructions)
	return instructions
}
