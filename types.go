package GoSpeedTest

// Result contains data from CLI output
type Result struct {
	Type       string    `json:"type"`
	Timestamp  string    `json:"timestamp"`
	Ping       Ping      `json:"ping"`
	Download   Measure   `json:"download"`
	Upload     Measure   `json:"upload"`
	PacketLoss int       `json:"packetLoss"`
	ISP        string    `json:"isp"`
	Interface  Interface `json:"interface"`
	Server     Server    `json:"server"`
	Servers    *[]Server `json:"servers"`
	Result     SResult   `json:"result"`
}

// Ping contains data for ping
type Ping struct {
	Jitter  float32 `json:"jitter"`
	Latency float32 `json:"latency"`
}

// Measure contains data for Download and Upload
type Measure struct {
	Bandwidth int `json:"bandwidth"`
	Bytes     int `json:"bytes"`
	Elapsed   int `json:"elapsed"`
}

// Interface contains data about the interface used to run the speedtest
type Interface struct {
	InternalIP string `json:"internalIp"`
	Name       string `json:"name"`
	MacAddress string `json:"macAddr"`
	IsVPN      bool   `json:"isVpn"`
	ExternalIP string `json:"externalIp"`
}

// Server contains data about the server used to run the speedtest
type Server struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Country  string `json:"country"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	IP       string `json:"ip"`
}

// SResult contains data for speedtest share
type SResult struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
