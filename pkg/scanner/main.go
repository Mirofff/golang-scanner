package scanner

type Scanner struct {
	start_port    int `default: "0"`
	end_port      int `default: "65535"`
	threads_count int `default: "100"`
}
