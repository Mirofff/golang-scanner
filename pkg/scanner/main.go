package scanner

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

type Scanner struct {
	start_port    int    `default: 1`
	end_port      int    `default: 65535`
	host          string `default: "localost"`
	threads_count int    `default: 100`
}

func (s *Scanner) Scan() {
	for port := s.start_port; port < s.end_port; port++ {
		logrus.Info("port: %d", port)
		checkPort(s.host, port)
	}
}

func (s *Scanner) SelectScan(ports []int) {
	for port := range ports {
		checkPort(s.host, port)
	}
}

func checkPort(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	logrus.Info("address: %s", address)
	if conn, err := net.Dial("tcp", address); err != nil {
		defer conn.Close()
		fmt.Println(address, " is open!")
	} else {
		fmt.Println(address, " is close...")
	}
}

// func (s *Scanner) ValidateValidate() error {
// 	return fmt.Errorf("")
// }

// func (s *Scanner) CheckAddress() error {

// }

// func validateAddress(address string) error {
// 	_, err := idna.Lookup.ToASCII(address)
// 	return fmt.Errorf("You pass unresolved address: %v", err)
// }
