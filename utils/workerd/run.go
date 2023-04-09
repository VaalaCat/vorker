package workerd

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"voker/defs"
)

const (
	workerdPath = "/workspaces/vorker/bin/workerd"
)

func Run(workerdDir string, argv []string) {
	args := []string{"serve",
		filepath.Join(workerdDir, defs.CapFileName),
		"--watch", "--verbose"}
	args = append(args, argv...)
	cmd := exec.Command(workerdPath, args...)
	cmd.Dir = workerdDir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
}
