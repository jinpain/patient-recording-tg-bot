package main

import (
	"fmt"

	"github.com/jinpain/patient-recording-tg-bot/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
