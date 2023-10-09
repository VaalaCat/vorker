package tunnel

import (
	"context"

	"github.com/sirupsen/logrus"
)

func logger(ctx context.Context, method string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"module": "tunnel", "method": method})
}
