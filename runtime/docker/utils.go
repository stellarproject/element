package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/sirupsen/logrus"
)

func getContainerEndpoint(c types.Container) string {
	logrus.Debug("getting container endpoint")
	endpoint := ""
	if len(c.Ports) > 0 {
		for _, p := range c.Ports {
			logrus.WithFields(logrus.Fields{
				"port": fmt.Sprintf("%+v", p),
			}).Debug("checking container port")

			if p.IP != "" && p.PublicPort != 0 {
				endpoint = fmt.Sprintf("%s:%d", p.IP, p.PublicPort)
				break
			}
		}
	}

	return endpoint
}
