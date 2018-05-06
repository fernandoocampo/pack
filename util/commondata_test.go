package util_test

import (
	"testing"

	"github.com/fernandoocampo/pack/util"
)

// TestNewHealthStatus tests instances a HealthStatus struct
func TestNewHealthStatus(t *testing.T) {
	expresult := createExpHealthStatus(true)
	// GIVEN data to build a health status
	dbs := make([]util.DBHealth, 1)
	dbs[0] = util.DBHealth{
		Name:    "mongo",
		Status:  true,
		Message: "ok",
	}
	service := "Pruebillo"

	// When we need to get a health status struct
	result := util.NewHealthStatus(service, true, dbs)

	// THEN system returns a good Health Status
	if result == nil {
		t.Fatalf("Expected a HealthStatus struct with data but got nil")
	}

	if result.ServiceName != expresult.ServiceName {
		t.Fatalf("Expected Result#ServiceName %s but got %s", expresult.ServiceName, result.ServiceName)
	}

	if result.DBClients == nil {
		t.Fatalf("Expected Result#DBClients to be not nil but got nil")
	}

	if len(result.DBClients) != 1 {
		t.Fatalf("Expected Result#DBClients to contain 1 member but got a different value")
	}

}

// TestAddDBToHealthStatus tests adding a new DBHealth to HealthStatus
func TestAddDBToHealthStatus(t *testing.T) {
	expresultdb := createExpHealthStatus(true)
	expresult := createExpHealthStatus(false)
	// GIVEN data to build a health status
	db1 := util.NewDBHealth("mongo", true, "ok")
	db2 := util.NewDBHealth("es", true, "ok")
	db3 := util.NewDBHealth("mysql", false, "ko")
	// THEN system add db health status to db
	expresult.AddDBToHealthStatus(db3)
	expresultdb.AddDBToHealthStatus(db2)
	expresultdb.AddDBToHealthStatus(db1)
	// THEN system returns a good Health Status
	if len(expresult.DBClients) != 1 {
		t.Fatalf("Expected expresult#DBClients to contain 1 member but got a different value")
	}

	if len(expresultdb.DBClients) != 3 {
		t.Fatalf("Expected expresultdb#DBClients to contain 3 members but got a different value")
	}

}

func createExpHealthStatus(withdb bool) *util.HealthStatus {
	health := new(util.HealthStatus)
	health.ServiceName = "Pruebillo"
	health.ServiceStatus = true
	if withdb {
		dbs := make([]util.DBHealth, 1)
		dbs[0] = util.DBHealth{
			Name:    "mongo",
			Status:  true,
			Message: "ok",
		}
		health.DBClients = dbs
	}
	return health
}
