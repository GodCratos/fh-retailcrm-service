package structures

type Message struct {
	ID  string
	New bool
}

type ResponseError struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errorMsg"`
}

type Response struct {
	Success bool `json:"success"`
}
