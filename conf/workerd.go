package conf

type Workerd struct {
	StartPort int `json:"start_port"`
	EndPort   int `json:"end_port"`
}
