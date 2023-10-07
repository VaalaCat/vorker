package main

import (
	"embed"
	"vorker/services"

	"github.com/sirupsen/logrus"
)

//go:embed all:www/out/*
var fs embed.FS

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	services.Run(fs)
}
