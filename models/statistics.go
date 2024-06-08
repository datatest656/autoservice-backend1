package models

// Statistics представляет собой базовую статистику
type Statistics struct {
	TotalVisits       int `json:"total_visits" example:"1000"`
	TotalAppointments int `json:"total_appointments" example:"50"`
}
