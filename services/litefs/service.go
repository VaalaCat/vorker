package litefs

import (
	"context"
	"os"
	"os/exec"
	"time"
	"vorker/conf"

	"github.com/sirupsen/logrus"
)

func RunService() {
	if !conf.AppConfigInstance.LitefsEnabled {
		return
	}
	for {
		ctx := context.Background()
		args := []string{"mount"}
		cmd := exec.CommandContext(ctx,
			conf.AppConfigInstance.LitefsBinPath, args...)
		cmd.Dir = conf.AppConfigInstance.LitefsDirPath
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = os.Environ()
		if err := cmd.Run(); err != nil {
			logrus.WithError(err).Errorf("run litefs error, sleep 3s and retry")
			time.Sleep(3 * time.Second)
		} else {
			logrus.Infof("run litefs success, sleep 5s and retry")
			time.Sleep(5 * time.Second)
		}
	}
}
