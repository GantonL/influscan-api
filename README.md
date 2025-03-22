# Influscan API

A high-performance Go API service powering the Influscan platform - enabling businesses to identify and analyze their customers' influential footprint based on purchase data.

## 🚀 Features

- Secure authentication via Clerk
- Data persistence with Supabase
- RESTful API design
- High-performance architecture
- Comprehensive logging and monitoring
- Rate limiting and security measures

## 🏗️ Architecture

```
influscan-api/
├── cmd/
│   └── api/           # Application entry point
├── internal/
│   ├── auth/          # Authentication logic
│   ├── domain/        # Business logic and models
│   ├── handlers/      # HTTP request handlers
│   ├── middleware/    # HTTP middleware
│   ├── repository/    # Data access layer
│   └── service/       # Business logic layer
├── pkg/               # Public packages
├── api/              # API documentation and schemas
└── scripts/          # Build and deployment scripts
```

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Authentication**: Clerk
- **Database**: Supabase
- **Deployment**: Hetzner
- **API Style**: RESTful

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker (optional)
- Make (optional)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/influscan-api.git
   cd influscan-api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## 🔧 Configuration

The application is configured via environment variables:

- `PORT`: Server port (default: 8080)
- `ENV`: Environment (development/production)
- `CLERK_SECRET_KEY`: Clerk authentication key
- `SUPABASE_URL`: Supabase project URL
- `SUPABASE_KEY`: Supabase API key


## 🧪 Testing

Run tests:
```bash
go test ./...
```


