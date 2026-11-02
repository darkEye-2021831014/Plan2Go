<div align="center">

# ✈️ Plan2Go

**AI-powered travel planning and local tour guide platform for Bangladesh.**
Tell it where you're going, how long you'll stay, and what you like. Plan2Go builds the trip.

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-18-61DAFB?logo=react&logoColor=black)
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3-38BDF8?logo=tailwindcss&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-database-4479A1?logo=mysql&logoColor=white)
![Gemini](https://img.shields.io/badge/Google_Gemini-AI-8E75B2?logo=googlegemini&logoColor=white)

</div>

---

## Table of Contents

- [The Problem](#the-problem)
- [Features](#features)
- [Project Status](#project-status)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [API Reference](#api-reference)
- [Project Structure](#project-structure)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [Author](#author)

---

## The Problem

Travelers in unfamiliar areas run into the same issues again and again:

- They don't know which places are actually worth visiting.
- They don't know what to pack for the weather or the type of location.
- They waste time hunting for hotels, restaurants, transport, and hospitals nearby.
- Most apps show raw data without personalizing it.
- There is no easy way to find and book a local guide.

Plan2Go brings AI trip recommendations, weather awareness, and a tour guide marketplace into one app.

## Features

### 🎯 AI Travel Recommendations
Choose a **division**, **district**, **trip length**, **budget** (economical, normal, luxury), and **location type** (chill, nature, urban, mountains). The backend sends your preferences to Google Gemini and returns a structured plan with:

- 4 to 7 recommended places, each with a description, activities, rating, and packing suggestions
- Nearby hotels and restaurants with ratings
- A day-by-day itinerary with the expected weather for each day

### 🌤️ Weather Data
A weather endpoint returns current conditions and a forecast for any coordinates, powered by OpenWeatherMap: temperature, feels-like, humidity, wind, precipitation, UV index, sunrise, and sunset.

### 👥 Tour Guide System
A two-sided marketplace:

- **Guides** register a profile with their city, hourly fee, languages, and years of experience.
- **Travelers** browse guides, filter by division, and start a booking.

### 🔐 Accounts and Security
- Registration with **email OTP verification**
- Login with **JWT** authentication (24-hour tokens)
- Passwords hashed with bcrypt
- Profile editing and password change from a settings page
- Protected and public route guards on the frontend

### 📊 Tour Monitor
A dashboard page for tracking active, upcoming, and completed tours.

## Project Status

Plan2Go is under active development. Here is an honest picture of what is wired up end to end and what is still in progress.

| Area | Status |
| --- | --- |
| Registration, OTP verification, login, profile, password change | ✅ Working |
| AI trip planning (Gemini) with daily itinerary | ✅ Working |
| Guide registration and guide listing | ✅ Working |
| Weather API endpoint | ✅ Backend ready, not yet used by the UI |
| Guide booking | 🚧 UI and database table exist, backend booking endpoint pending |
| Tour Monitor | 🚧 UI built with sample data, not yet connected to the backend |
| Nearby Essentials Map | 🗺️ Planned |
| Place Explorer, reviews and ratings | 🗺️ Planned |

## Tech Stack

| Layer | Technology |
| --- | --- |
| **Frontend** | React 18, React Router 6, Tailwind CSS 3, React Hook Form, React Hot Toast, React Toastify |
| **Backend** | Go (standard `net/http` with Go 1.22+ method routing), JWT (`golang-jwt`), `godotenv` |
| **Database** | MySQL / MariaDB (TLS is enabled automatically for remote hosts) |
| **AI** | Google Gemini via the `google.golang.org/genai` SDK (`gemini-2.0-flash`) |
| **Weather** | OpenWeatherMap API |
| **Email** | SMTP for OTP delivery |

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.24 or newer
- [Node.js](https://nodejs.org/) 18 or newer, with npm
- MySQL or MariaDB
- A [Google Gemini API key](https://aistudio.google.com/app/apikey)
- An [OpenWeatherMap API key](https://openweathermap.org/api)

### 1. Clone the repository

```bash
git clone https://github.com/darkEye-2021831014/Plan2Go.git
cd Plan2Go
```

### 2. Set up the database

Create a database and import the schema:

```bash
mysql -u root -p -e "CREATE DATABASE plan2go;"
mysql -u root -p plan2go < database/plan2go.sql
```

This creates four tables: `users`, `guides`, `guide_bookings`, and `email_verification`.

### 3. Run the backend

Create a `.env` file inside `backend/` (see [Configuration](#configuration)), then:

```bash
cd backend
go mod download
go run main.go
```

The API starts on **http://localhost:8080** by default.

### 4. Run the frontend

In a second terminal:

```bash
cd frontend
npm install
npm start
```

The app opens at **http://localhost:3000**.

> The frontend currently calls the API at `http://localhost:8080`. If you deploy the backend elsewhere, update the URLs in `frontend/src` (a good first step is moving them into a single `REACT_APP_API_URL` environment variable).

### Build for production

```bash
cd frontend && npm run build
cd ../backend && go build -o plan2go-server .
```

## Configuration

The backend reads settings from environment variables, and loads a `.env` file from the working directory if one exists. **Never commit `.env`.** It is already in `backend/.gitignore`.

```env
# Server
PORT=8080
SERVICE_NAME=Plan2Go
VERSION=1.0.0

# Auth
JWT_SECRETKEY=replace-with-a-long-random-string

# Database
MYSQLHOST=127.0.0.1
MYSQLPORT=3306
MYSQLUSER=root
MYSQLPASSWORD=your-db-password
MYSQLDATABASE=plan2go

# External APIs
GEMINI_API_KEY=your-gemini-key
OPENWEATHER_API_KEY=your-openweather-key
```

| Variable | Required | Description |
| --- | --- | --- |
| `PORT` / `HTTP_PORT` | No | Server port. Defaults to `8080`. `PORT` takes priority, which suits platforms like Railway. |
| `JWT_SECRETKEY` | **Yes** | Secret used to sign tokens. Always set your own. |
| `MYSQLHOST`, `MYSQLPORT`, `MYSQLUSER`, `MYSQLPASSWORD`, `MYSQLDATABASE` | **Yes** | Database connection. TLS is enabled for any host other than `localhost` or `127.0.0.1`. |
| `GEMINI_API_KEY` | **Yes** | Powers trip plan generation. |
| `OPENWEATHER_API_KEY` | For weather | OpenWeatherMap key. |

## API Reference

Base URL: `http://localhost:8080`

### Users

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/users` | Register a new user and send an OTP |
| `POST` | `/users/verify` | Verify the email OTP |
| `GET` | `/users/resend-otp` | Resend the OTP |
| `POST` | `/users/login` | Log in and receive a JWT |
| `GET` | `/users/profile` | Get the current user's profile |
| `POST` | `/users/update/profile` | Update profile details |
| `POST` | `/users/update/password` | Change password |

### Planning and Weather

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/users/plan` | Generate an AI trip plan |
| `GET` | `/users/weather?lat=&lon=` | Current weather and forecast for coordinates |

**Example plan request**

```json
{
  "division": "Sylhet",
  "district": "Sylhet",
  "days": 3,
  "budget": "normal",
  "locationType": "nature"
}
```

The response contains `locations` (each with `name`, `description`, `activities`, `rating`, `accessories`, `hotels_nearby`, `restaurants_nearby`, and `image_url`) and a `daily_itinerary` keyed by day.

### Guides

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/users/guide` | List guides |
| `POST` | `/users/guide/register` | Register the current user as a guide |

Protected endpoints expect an `Authorization: Bearer <token>` header.

## Project Structure

```
Plan2Go/
├── backend/                     # Go REST API
│   ├── main.go                  # Entry point
│   ├── cmd/Serve.go             # Wires repos, services, handlers, and server
│   ├── config/                  # Environment-based configuration
│   ├── infra/DB/                # MySQL connection
│   ├── repo/                    # Data access (users, guides, email OTP)
│   ├── rest/
│   │   ├── Server.go            # HTTP server and route registration
│   │   ├── handlers/            # user, plan, weather, guide handlers
│   │   ├── middleware/          # Logger, JWT authentication, middleware manager
│   │   └── services/            # Gemini plan generation
│   └── util/                    # JWT, OTP, password hashing, CORS router
├── frontend/                    # React application
│   ├── public/
│   └── src/
│       ├── components/          # Navbar, ProtectedRoute, PublicRoute
│       ├── context/             # AuthContext (login, register, session)
│       └── pages/               # Landing, Login, Register, VerifyOTP, Dashboard,
│                                # TourPlanning, TourGuides, TourMonitor, Maintenance
└── database/
    └── plan2go.sql              # MySQL schema
```

## Roadmap

- [ ] Nearby Essentials Map: hotels, restaurants, bus stops, train stations, hospitals, ATMs, petrol pumps, malls, and emergency services
- [ ] Weather-based packing suggestions shown in the planner using live forecasts
- [ ] Connect guide booking to the backend with booking status
- [ ] Connect Tour Monitor to real bookings
- [ ] Guide profiles with photos, availability, reviews, and ratings
- [ ] Place Explorer with popular places, hidden gems, and map visualization
- [ ] Single configurable API base URL for the frontend
- [ ] Tests and CI

## Contributing

Contributions and suggestions are welcome.

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -m "Add my feature"`
4. Push the branch: `git push origin feature/my-feature`
5. Open a pull request

## Author

**Ashraful Islam** ([@darkEye-2021831014](https://github.com/darkEye-2021831014))

---

<div align="center">

If you find Plan2Go useful, consider giving the repo a ⭐

</div>
