#  Fuster — Fast Server & Protocol Availability Checker

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
  <img src="https://img.shields.io/badge/License-MIT-blue.style=for-the-badge" alt="License" />
  <img src="https://img.shields.io/badge/Speed-Parallel%20Goroutines-brightgreen?style=for-the-badge" alt="Goroutines" />
</p>

**Fuster** is a high-performance CLI utility written in Go for quick checking and pinging modern configuration links (including subscription links from web sources).

---

##  What Problem Does It Solve?

Most popular clients and desktop tools suffer from the same issues:
*  **Slow downloads & laggy UI:** When dealing with hundreds or thousands of configs, traditional GUIs freeze or take minutes to finish.
*  **Inaccurate ping data:** Standard ICMP or basic TCP checks often yield false results.
*  **Long timeouts:** Unnecessary waiting times for non-responsive endpoints.

**Fuster** solves this by performing **parallel TCP Byte-Handshake checks** using Go goroutines — giving you realistic latency measurements in milliseconds without UI overhead.

---

##  Features

-  **High Speed:** Concurrent checks powered by Go channels & goroutines.
-  **Accurate Handshake Testing:** Real byte-level handshake evaluation (not basic socket connection checks).
-  **Subscription Support:** Direct URL loading and automatic Base64 parsing.
-  **Multi-Protocol Ready:** Designed for modern protocols (VLESS, Trojan, Shadowsocks, etc.).
-  **Smart Sorting:** Fast alive nodes ranked at the top, dead nodes listed at the bottom.
-  **Clear Diagnostics:** Informative error logging so you know exactly what went wrong.

---

##  Output Example

```text
[+] Server: 🇳🇱 Нидерланды #3 | Status: ALIVE | Latency: 917ms
[+] Server: 🇫🇷 Франция | WI-FI | Status: ALIVE | Latency: 967ms
[-] Server: 🇳🇱 Нидерланды | WI-FI | Status: DEAD
[-] Server: 🇳🇱 Нидерланды #2 | WI-FI | Status: DEAD
```

##  Quick Start

### 1. Clone the repository
```bash
git clone https://github.com/Crombles/Fast-server-config-checker.git checker
cd checker
```

### 2. Run directly
```bash
go run .
```

### 3. Build executable
```bash
go build -o checker .
```

### 4. Run executable
```bash
./checker
```
