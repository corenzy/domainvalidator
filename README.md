# 🔍 Domain Validator & DNS Lookup API

[![Go Version](https://img.shields.io/github/go-mod/go-version/corenzy/domainvalidator)](https://golang.org)
[![Fiber Framework](https://img.shields.io/badge/framework-Fiber_v2-blue)](https://gofiber.io)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker Support](https://img.shields.io/badge/docker-ready-blue.svg)](https://www.docker.com/)

A high-performance, concurrent DNS lookup API built with **Go** and **GoFiber**. It provides detailed DNS records (A, AAAA, MX, NS, TXT, CNAME) and automatically detects the DNS provider (Cloudflare, AWS, Google, etc.).

## 🚀 Key Features

-   **⚡ Concurrent Lookups:** Queries all DNS record types in parallel using Goroutines for massive performance gains.
-   **🏢 Provider Detection:** Identifies ~50+ different DNS providers based on nameserver patterns.
-   **🛡️ Reliability:** Uses a custom resolver (Cloudflare 1.1.1.1) to bypass local OS resolution issues and ensure consistent low latency.
-   **🐳 Docker Ready:** Includes a multi-stage Dockerfile for a minimal (~15MB) production image.
-   **✨ Clean Architecture:** Modular and maintainable project structure, perfect for production or learning.

---

## 🛠️ Installation

### Using Go

```bash
# Clone the repository
git clone https://github.com/corenzy/domainvalidator.git
cd domainvalidator

# Install dependencies
go mod tidy

# Run the server
go run main.go
```

### Using Docker

```bash
# Build and run with Docker Compose
docker-compose up --build -d
```

---

## 📖 API Documentation

### `POST /lookup`

Performs a full DNS lookup for a given domain.

**Request Body:**
```json
{
  "domain": "google.com"
}
```

**Successful Response:**
```json
{
  "success": true,
  "message": "DNS lookup completed for google.com",
  "lookup_time_ms": 145,
  "result": {
    "domain": "google.com",
    "provider": {
      "name": "Google Cloud DNS",
      "website": "https://cloud.google.com/dns"
    },
    "nameservers": ["ns1.google.com", "ns2.google.com"],
    "a_records": [
       { "type": "A", "value": "142.250.184.238" }
    ],
    "mx_records": [
       { "type": "MX", "value": "smtp.google.com", "priority": 10 }
    ],
    "txt_records": ["v=spf1 include:_spf.google.com ~all"]
  }
}
```

### `GET /`

Returns the API status and minimal health check information.

---

## 📂 Project Structure

```bash
.
├── main.go              # Entrypoint (Fiber setup, routes, middleware)
├── handlers/
│   └── lookup.go        # HTTP request handling & validation
├── services/
│   └── dns.go           # Core concurrent DNS resolution logic
├── providers/
│   └── providers.go     # Registry of ~50+ DNS providers
├── models/
│   └── models.go        # Shared request/response data structures
├── Dockerfile           # Multi-stage production build
└── docker-compose.yml   # Local orchestration
```

---

## 🤝 Contributing

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the `LICENSE` file for details.
