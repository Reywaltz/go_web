package bus

type Bus struct {
	From    string  `json:"from"`
	To      string  `json:"to"`
	Subject string  `json:"subject"`
	Data    *string `json:"data"`
}
