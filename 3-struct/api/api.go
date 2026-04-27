package api

import (
	"3-struct/config"
	"fmt"
)

func Hello(cfg config.Config) {

	key := cfg.Key

	fmt.Printf("%v", key)

}
