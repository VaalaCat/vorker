package main

import (
	"embed"
	"vorker/exec"
	"vorker/services"

	"github.com/sirupsen/logrus"
)

//go:embed all:www/out/*
var fs embed.FS

func main() {
	defer exec.ExecManager.ExitAllCmd()
	logrus.SetLevel(logrus.DebugLevel)
	services.Run(fs)
}
