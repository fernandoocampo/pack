package util

// DBHealth contains data about database connection status
type DBHealth struct {
	Name    string `json:"name"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// HealthStatus contains data to report the state of a
// service
type HealthStatus struct {
	ServiceName   string     `json:"name"`
	ServiceStatus bool       `json:"status"`
	DBClients     []DBHealth `json:"dbhealth"` // the db connections.
}

// NewHealthStatus build a new HealthStatus struct
func NewHealthStatus(servicename string, servicestatus bool, dbclients []DBHealth) *HealthStatus {
	status := new(HealthStatus)
	status.ServiceName = servicename
	status.DBClients = dbclients
	status.ServiceStatus = servicestatus
	return status
}

// NewDBHealth build a new DBHealth struct
func NewDBHealth(dbname string, status bool, message string) *DBHealth {
	db := new(DBHealth)
	db.Name = dbname
	db.Status = status
	db.Message = message
	return db
}

// AddDBToHealthStatus add a new DBHealth to the given health status.
func (h *HealthStatus) AddDBToHealthStatus(dbhealth *DBHealth) bool {
	if dbhealth == nil {
		return false
	}
	if h.DBClients == nil {
		h.DBClients = make([]DBHealth, 1)
		h.DBClients[0] = *dbhealth
		return true
	}
	h.DBClients = append(h.DBClients, *dbhealth)
	return true
}
