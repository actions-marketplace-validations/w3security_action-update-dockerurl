package dockerurl

import (
	"github.com/w3security/action-update/actions/updateaction"
	"github.com/w3security/action-update/updater"
)

type Environment struct {
	updateaction.Environment
}

func (c *Environment) NewUpdater(root string) updater.Updater {
	u := NewUpdater(root)
	u.pathFilter = c.Ignored
	return u
}
