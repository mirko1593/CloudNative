package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const serviceName = "EEEEEEEEEE"

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(serviceName + ": Cannot Get Hostname")
	}
	ip, err := getIP()
	if err != nil {
		log.Fatalln(serviceName + ": Cannot Get IP")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("%s: Hostname: %s, IP: %s", serviceName, hostname, ip)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		tpl := `
        <html>
        <body>
        <h1>` + str + `</h1>
        <a href="/service-a">Service A</a>
        <br />
        <a href="/service-B">Service B</a>
        </body>

        </html>
        `
		fmt.Fprintf(w, tpl)
	})

	log.Println("Start Service Listen On :8080")

	http.ListenAndServe(":8080", nil)
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("No IP Addr Found")
}
