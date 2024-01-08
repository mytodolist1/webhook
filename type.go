package webhook

type Request struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
