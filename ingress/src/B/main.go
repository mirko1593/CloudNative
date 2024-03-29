package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const serviceName = "BBBBBBBBBB"

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(serviceName + ": Cannot Get Hostname")
	}
	ip, err := getIP()
	if err != nil {
		log.Fatalln(serviceName + ": Cannot Get IP")
	}

	now := time.Now()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("%s: Hostname: %s, IP: %s", serviceName, hostname, ip)

		w.Write([]byte(str))
	})

	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("%s: Hostname: %s, IP: %s", serviceName, hostname, ip)

		if time.Now().Sub(now) < time.Second*8 {
			log.Println("Liveness Endpoint: ", str+": IS NOT ALIVE")
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		log.Println("Liveness Endpoint: ", str+": IS ALIVE")

		w.Write([]byte(str))
	})

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("%s: Hostname: %s, IP: %s", serviceName, hostname, ip)

		if time.Now().Sub(now) < time.Second*11 {
			log.Println("Liveness Endpoint: ", str+": IS NOT READY")
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		log.Println("Readiness Endpoint: ", str+": IS READY")

		w.Write([]byte(str))
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("HELLO: %s FROM Hostname: %s, IP: %s", serviceName, hostname, ip)

		w.Write([]byte(str))
	})

	log.Printf("Start Service %s and Listen At: 8080\n", serviceName)

	server := getServer()

	// server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}

func getServer() *http.Server {
	// data, _ := os.ReadFile("./minica.pem")
	// cp, _ := x509.SystemCertPool()
	// cp.AppendCertsFromPEM(data)

	tls := &tls.Config{
		// RootCAs: cp,
	}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tls,
	}

	return server
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
