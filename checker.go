package main

import (
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"
)

type ServerResult struct {
	Name    string
	Address string
	IsAlive bool
	Latency time.Duration
}

func processParallel(configs []string, workerCount int) []ServerResult {
	jobs := make(chan string, len(configs))
	results := make(chan ServerResult, len(configs))
	var wg sync.WaitGroup

	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for _, cfg := range configs {
		jobs <- cfg
	}

	close(jobs)

	wg.Wait()
	close(results)

	var finalResults []ServerResult
	for res := range results {
		finalResults = append(finalResults, res)
	}

	return finalResults
}

func worker(jobs <-chan string, results chan<- ServerResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for link := range jobs {
		res, _ := checkServer(link)
		results <- res
	}
}

func checkServer(rawCfg string) (ServerResult, error) {
	result := ServerResult{Name: "Unknown"}

	parsed, err := url.Parse(rawCfg)
	if err != nil {
		result.Name = "Invalid URI"
		return result, fmt.Errorf("invalid link: %w", err)
	}

	name, unescapeErr := url.QueryUnescape(parsed.Fragment)
	if unescapeErr != nil || name == "" {
		result.Name = "Unknown server name"
	} else {
		result.Name = name
	}

	host := parsed.Hostname()
	port := parsed.Port()

	if host == "" {
		return result, fmt.Errorf("%s server error: host/IP not found", result.Name)
	}

	if port == "" {
		switch parsed.Scheme {
		case "https", "vless", "trojan":
			port = "443"

		case "ss", "vmess", "shadowsocks":
			result.Name = fmt.Sprintf("Error: port is not specified for protocol %s", parsed.Scheme)
			return result, fmt.Errorf("protocol %s does not specify a port", parsed.Scheme)

		default:
			port = "80"
		}
	}

	address := net.JoinHostPort(host, port)
	result.Address = address
	start := time.Now()

	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		result.IsAlive = false
		return result, err
	}

	defer conn.Close()

	_ = conn.SetDeadline(time.Now().Add(1500 * time.Millisecond))

	PingPayload := []byte{0x05, 0x01, 0x00}
	_, err = conn.Write(PingPayload)
	if err != nil {
		result.IsAlive = false
		return result, fmt.Errorf("server did not accept data: %w", err)
	}

	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		result.IsAlive = false
		return result, fmt.Errorf("server is silent (no response): %w", err)
	}

	result.Latency = time.Since(start)
	result.IsAlive = true
	return result, nil
}
