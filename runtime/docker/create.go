package docker

import (
	"github.com/ehazlett/element/runtime"
	"github.com/sirupsen/logrus"
)

func (d *Docker) Create(spec *runtime.Spec) error {
	logrus.WithFields(logrus.Fields{
		"spec": spec,
	}).Debug("creating container")
	return nil
}
