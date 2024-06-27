package utils

import (
	"cpk/internal/configs/global"
	"fmt"
)

func Log_debug(message any) {
	if global.DEV_MODE {
		fmt.Println(message)
	}
}
