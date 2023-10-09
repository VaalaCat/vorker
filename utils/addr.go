package utils

import (
	"fmt"
	"strings"
	"vorker/conf"
)

func NodeHostPrefix(nodeName, nodeID string) string {
	return fmt.Sprintf("%s%s", nodeName, nodeID)
}

func NodeHost(nodeName, nodeID string) string {
	suffix := strings.Trim(conf.AppConfigInstance.WorkerURLSuffix, ".")
	return fmt.Sprintf("%s.%s", NodeHostPrefix(nodeName, nodeID), suffix)
}

func WorkerHostPrefix(workerName string) string {
	return workerName
}

func WorkerHost(workerName string) string {
	suffix := strings.Trim(conf.AppConfigInstance.WorkerURLSuffix, ".")
	return fmt.Sprintf("%s.%s", WorkerHostPrefix(workerName), suffix)
}
