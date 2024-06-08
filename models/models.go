package models

import (
	"time"
)

// SuccessResponse используется для успешных ответов API
type SuccessResponse struct {
	Data interface{} `json:"data" example:"{}"` // Поле Data может содержать любые данные
}

// ErrorResponse is a generic error response
type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

// Appointment represents an appointment
type Appointment struct {
	ID              uint      `json:"id" gorm:"primary_key" example:"1"`
	ClientID        uint      `json:"client_id" example:"1"`
	EmployeeID      uint      `json:"employee_id" example:"1"`
	ServiceID       uint      `json:"service_id" example:"1"`
	AppointmentDate time.Time `json:"appointment_date" example:"2024-06-08T12:00:00Z"`
	Status          string    `json:"status" example:"Scheduled"`
	CreatedAt       time.Time `json:"created_at"`
}

// Client represents a client
type Client struct {
	ID          uint      `json:"id" gorm:"primary_key" example:"1"`
	Name        string    `json:"name" example:"John Doe"`
	PhoneNumber string    `json:"phone_number" example:"123456789"`
	Email       string    `json:"email" example:"johndoe@example.com"`
	CreatedAt   time.Time `json:"created_at"`
}

// Contact represents contact information
type Contact struct {
	ID          uint      `json:"id" gorm:"primary_key" example:"1"`
	Address     string    `json:"address" example:"123 Main St"`
	PhoneNumber string    `json:"phone_number" example:"123456789"`
	Email       string    `json:"email" example:"contact@example.com"`
	CreatedAt   time.Time `json:"created_at"`
}

// Employee represents an employee
type Employee struct {
	ID          uint      `json:"id" gorm:"primary_key" example:"1"`
	Name        string    `json:"name" example:"Jane Smith"`
	PhoneNumber string    `json:"phone_number" example:"987654321"`
	Email       string    `json:"email" example:"janesmith@example.com"`
	Role        string    `json:"role" example:"моторист"`
	PhotoURL    string    `json:"photo_url" example:"http://example.com/photo.jpg"`
	CreatedAt   time.Time `json:"created_at"`
}

// Feedback represents feedback from a client
type Feedback struct {
	ID        uint      `json:"id" gorm:"primary_key" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"johndoe@example.com"`
	Message   string    `json:"message" example:"Great service!"`
	CreatedAt time.Time `json:"created_at"`
}

// News represents a news item
type News struct {
	ID        uint      `json:"id" gorm:"primary_key" example:"1"`
	Title     string    `json:"title" example:"New Service Available"`
	Content   string    `json:"content" example:"We are now offering oil change services."`
	CreatedAt time.Time `json:"created_at"`
}

// Service represents a service offered by the auto service
type Service struct {
	ID          uint      `json:"id" gorm:"primary_key" example:"1"`
	Name        string    `json:"name" example:"Oil Change"`
	Description string    `json:"description" example:"Full oil change service"`
	Price       float64   `json:"price" example:"49.99"`
	CreatedAt   time.Time `json:"created_at"`
}
