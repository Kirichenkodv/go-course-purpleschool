package api

import (
	"config"
	"fmt"
)

func Hello(cfg config.Config) {

	key := cfg.Key

	fmt.Printf("%v", key)

}
