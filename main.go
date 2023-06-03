package main

import (
	"embed"
	"voker/services"
)

//go:embed all:www/out/*
var fs embed.FS

func main() {
	services.Run(fs)
}
