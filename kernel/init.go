package kernel

import "github.com/sirupsen/logrus"

var (
	logger *logrus.Logger

	language string
)

// SetLog : set log
func SetLog(log *logrus.Logger) {
	logger = log
}
