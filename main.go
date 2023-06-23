package main

import (
	"embed"
	"vorker/services"
)

//go:embed all:www/out/*
var fs embed.FS

func main() {
	InitCache()
	services.Run(fs)
}
