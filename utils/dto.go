package utils

type Event struct {
	Message   string `json:"message"`
	File      string `json:"file"`
	Operation string `json:"operation"`
}
