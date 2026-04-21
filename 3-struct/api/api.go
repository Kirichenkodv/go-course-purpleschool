package api

import (
	"3-struct/config"
	"config"
	"fmt"
)

func Hello() {
	cfg := config.New()

	key := cfg.Key

	fmt.Printf("%v", key)

}
