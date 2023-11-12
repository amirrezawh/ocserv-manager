package generator

type Data []struct {
	Username    string `json:"Username"`
	State       string `json:"State"`
	RX          string `json:"RX"`
	TX          string `json:"TX"`
	ConnectedAt string `json:"Connected at"`
}
