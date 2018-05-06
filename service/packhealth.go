package service

import (
	"github.com/fernandoocampo/pack/dao"
	"github.com/fernandoocampo/pack/util"
)

// PackHealth is the implementation of goforallIHealthService
type PackHealth struct {
}

// Health implements goforall IHealthService.Health capability.
func (c *PackHealth) Health() *util.HealthStatus {
	esstatus := dao.IsRunning()
	var esdb *util.DBHealth
	if esstatus != nil {
		esdb = util.NewDBHealth("packmongo", false, esstatus.Error())
	} else {
		esdb = util.NewDBHealth("packmongo", true, "ok")
	}
	servicename := "pack"
	servicestatus := true
	health := util.NewHealthStatus(servicename, servicestatus, nil)
	health.AddDBToHealthStatus(esdb)
	return health
}
