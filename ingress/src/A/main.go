package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

const serviceName = "AAAAAAAAAA"

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
		str := fmt.Sprintf("%s: Hostname: %s, IP: %s, path: %s", serviceName, hostname, ip, r.URL)
		tpl := `
            <h3>` + str + `</h3>
            <a href="/service-a/liveness">Liveness</a>
            <br />
            <a href="/service-a/readiness">Readiness</a>
            <br />
            <a href="/service-a/a-to-b">A-To-B</a>
        `
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprintf(w, tpl)
	})

	// now := time.Now()
	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("HELLO: %s FROM Hostname: %s, IP: %s, path: %s\n", serviceName, hostname, ip, r.URL.Path)

		// if time.Now().Sub(now) < time.Second*8 {
		// 	log.Println("Liveness Endpoint: ", str+": IS NOT ALIVE")
		// 	w.WriteHeader(http.StatusBadGateway)
		// 	return
		// }

		log.Println("Liveness Endpoint: ", str+": IS ALIVE")

		w.Write([]byte(str + ": IS ALIVE"))
	})

	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("HELLO: %s FROM Hostname: %s, IP: %s, path: %s\n", serviceName, hostname, ip, r.URL.Path)

		// if time.Now().Sub(now) < time.Second*11 {
		// 	log.Println("Liveness Endpoint: ", str+": IS NOT READY")
		// 	w.WriteHeader(http.StatusBadGateway)
		// 	return
		// }

		log.Println("Readiness Endpoint: ", str+": IS READY")

		w.Write([]byte(str + ": IS READY"))
	})

	http.HandleFunc("/a-to-b", func(w http.ResponseWriter, r *http.Request) {
		str := fmt.Sprintf("%s FROM Hostname: %s, IP: %s, path: %s\n", serviceName, hostname, ip, r.URL.Path)

		str = str + fmt.Sprintf(" IS CALLING \n")

		rsp, err := http.Get("http://service-b.ingresslabs.svc.cluster.local/api")
		if err != nil {
			log.Printf("A-to-B: err: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("A-to-B: err: %s", err)))
			return
		}
		defer rsp.Body.Close()

		c, _ := io.ReadAll(rsp.Body)

		w.Write([]byte(str + string(c)))
	})

	log.Printf("Start Service %s and Listen At: 8080\n", serviceName)

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
