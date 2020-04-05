package utils

type response struct {
	Success	bool	`json:"success"`
	Payload interface{}	`json:"payload"`
}