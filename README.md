# Weather API

REST API γραμμένο σε Go (Gin) για διαχείριση πόλεων και δεδομένων καιρού. Υποστηρίζει εγγραφή/login χρηστών με JWT, CRUD για πόλεις και τρέχοντα καιρό, και ένα background job που τραβάει live δεδομένα καιρού από το [Open-Meteo](https://open-meteo.com/) κάθε λεπτό, ενημερώνοντας τον τρέχοντα καιρό κάθε πόλης και κρατώντας ιστορικό (θερμοκρασία, υγρασία, ταχύτητα ανέμου).

## Τεχνολογίες

- Go 1.24 + [Gin](https://github.com/gin-gonic/gin)
- MySQL 8.4
- [golang-migrate](https://github.com/golang-migrate/migrate) (migrations τρέχουν αυτόματα στο startup)
- JWT auth ([golang-jwt](https://github.com/golang-jwt/jwt))
- [robfig/cron](https://github.com/robfig/cron) για το scheduled job
- Docker / Docker Compose (MySQL + phpMyAdmin)

## Προαπαιτούμενα

- Go 1.24+
- Docker & Docker Compose (για MySQL — ή τοπικό MySQL 8.x αν προτιμάς)
- (προαιρετικά) [air](https://github.com/air-verse/air) για hot reload

## Γρήγορο ξεκίνημα

```bash
git clone <repo-url>
cd weather-api

# 1. Environment variables
cp .env.example .env

# 2. Βάση δεδομένων (MySQL + phpMyAdmin)
docker compose up -d

# 3. Dependencies
go mod download

# 4. Εκκίνηση του API
go run ./cmd/api
```

Το API θα ακούει στο `http://localhost:8085` (ή στο `SERVER_PORT` που έχεις ορίσει στο `.env`). Τα migrations (`migrations/sql/*.sql`) τρέχουν **αυτόματα** στο startup — δεν χρειάζεται manual βήμα.

Για hot reload κατά την ανάπτυξη:

```bash
go install github.com/air-verse/air@latest
air
```

### Έλεγχος ότι όλα δουλεύουν

- Λογαριασμός στο MySQL μέσω phpMyAdmin: [http://localhost:8089](http://localhost:8089) (χρήστης/κωδικός όπως στο `.env`)
- Logs της εφαρμογής: στο terminal και στο `logs/app.log`
- Health check γρήγορα με ένα request χωρίς auth:
  ```bash
  curl -i http://localhost:8085/signup -X POST -H "Content-Type: application/json" -d '{}'
  ```
  Αν πάρεις `400 Bad Request` (αντί για connection error), το server τρέχει κανονικά.

## Environment Variables (`.env`)

| Μεταβλητή | Περιγραφή | Default |
|---|---|---|
| `MYSQL_HOST` | Host της MySQL | `localhost` |
| `MYSQL_PORT` | Port της MySQL (το docker-compose εκθέτει `8088`) | `3306` |
| `MYSQL_DATABASE` | Όνομα βάσης | `project_go` |
| `MYSQL_USER` | Χρήστης MySQL | – |
| `MYSQL_PASSWORD` | Κωδικός χρήστη | – |
| `MYSQL_ROOT_PASSWORD` | Root password (μόνο για docker-compose) | – |
| `SERVER_PORT` | Port στο οποίο ακούει το API | `8085` |
| `JWT_SECRET` | Secret για την υπογραφή JWT tokens | `your-secret-key` |
| `ENVIRONMENT` | `development` / `production` | `development` |
| `DB_MAX_OPEN_CONNS` | Max open DB connections | `10` |
| `DB_MAX_IDLE_CONNS` | Max idle DB connections | `5` |

Οι default τιμές του `.env.example` είναι ήδη συμβατές με το `docker-compose.yml`, οπότε αρκεί το `cp .env.example .env` για local development.

## API Endpoints

### Auth (χωρίς token)

| Method | Path | Body |
|---|---|---|
| POST | `/signup` | `{ "username", "email", "password" }` |
| POST | `/login` | `{ "email", "password" }` → επιστρέφει `token` |

### Cities (χρειάζονται `Authorization: Bearer <token>`)

| Method | Path | Body |
|---|---|---|
| GET | `/cities` | – |
| GET | `/cities/:id` | – |
| POST | `/cities` | `{ "name", "latitude", "longitude" }` |
| PUT | `/cities/:id` | `{ "name", "latitude", "longitude" }` |
| DELETE | `/cities/:id` | – |

### Weather (χρειάζονται `Authorization: Bearer <token>`)

| Method | Path | Body |
|---|---|---|
| GET | `/weather/:city_id` | – |
| POST | `/weather` | `{ "city_id", "condition", "temperature_c" }` |
| PUT | `/weather/:city_id` | `{ "city_id", "condition", "temperature_c" }` |
| DELETE | `/weather/:city_id` | – |

## Δοκιμή του API τοπικά

### Επιλογή Α — `.http` αρχεία (πιο εύκολο)

Στο [api-test/](api-test/) υπάρχουν έτοιμα requests (`register.http`, `login.http`, `city_crud.http`, `get_weather_by_city.http`). Άνοιξέ τα με το [REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) στο VS Code (ή το ενσωματωμένο HTTP Client του JetBrains/GoLand) και πάτα "Send Request" πάνω από κάθε block.

### Επιλογή Β — curl

```bash
# 1. Signup
curl -s -X POST http://localhost:8085/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"password123"}'

# 2. Login — πάρε το token από την απάντηση
TOKEN=$(curl -s -X POST http://localhost:8085/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}' | jq -r .token)

# 3. Δημιουργία πόλης
curl -s -X POST http://localhost:8085/cities \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"Athens","latitude":37.9838,"longitude":23.7275}'

# 4. Λίστα πόλεων
curl -s http://localhost:8085/cities -H "Authorization: Bearer $TOKEN"

# 5. Καταχώρηση τρέχοντος καιρού για την πόλη (π.χ. city_id=1)
curl -s -X POST http://localhost:8085/weather \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"city_id":1,"condition":"Sunny","temperature_c":28.5}'

# 6. Ανάκτηση καιρού για πόλη
curl -s http://localhost:8085/weather/1 -H "Authorization: Bearer $TOKEN"
```

### Background job (ιστορικό καιρού)

Μόλις υπάρχει τουλάχιστον μία πόλη στη βάση, ένα cron job τρέχει **κάθε λεπτό** και:

1. Τραβάει live δεδομένα από το Open-Meteo API για κάθε πόλη (βάσει `latitude`/`longitude`)
2. Ενημερώνει τον τρέχοντα καιρό της πόλης (πίνακας `weathers`)
3. Αποθηκεύει μία εγγραφή ιστορικού στον πίνακα `weather_histories`

Μπορείς να δεις τα αποτελέσματα είτε μέσω `GET /weather/:city_id`, είτε απευθείας στη βάση μέσω phpMyAdmin ([http://localhost:8089](http://localhost:8089)), κοιτώντας τους πίνακες `weathers` και `weather_histories`.

## Δομή project

```
cmd/api/           entrypoint (main.go)
config/            φόρτωση env vars, logger setup
db/                σύνδεση με MySQL
migrations/        SQL migrations (τρέχουν αυτόματα)
models/            domain structs
repositories/      πρόσβαση στη βάση
services/          business logic
handlers/          HTTP handlers (Gin)
routes/            ορισμός routes
middlewares/       JWT auth middleware
jobs/ + schedulers/ background cron job για ανάκτηση καιρού
clients/           HTTP client για το Open-Meteo API
api-test/          έτοιμα .http requests για δοκιμή
```

## Εκτέλεση ολόκληρου του stack με Docker

Το `docker-compose.yml` περιέχει ήδη (σχολιασμένο) ένα `app` service. Αφαίρεσε τα σχόλια από αυτό το block και τρέξε:

```bash
docker compose up -d --build
```

## Αντιμετώπιση προβλημάτων

- **`Could not run migrations` / "Dirty database version"**: κάποιο migration απέτυχε στη μέση. Χρησιμοποίησε το `scripts/fix-dirty-migration.sh <version>` για να καθαρίσεις το dirty flag (χρειάζεται τρέχον `docker compose` MySQL container).
- **`401 Unauthorized`**: λείπει ή είναι λάθος το header `Authorization: Bearer <token>` — κάνε πρώτα `POST /login` για να πάρεις token.
- **`connection refused` στη MySQL**: βεβαιώσου ότι το `docker compose up -d` έχει τρέξει και ότι το `MYSQL_PORT` στο `.env` ταιριάζει με αυτό που εκθέτει το `docker-compose.yml` (`8088`).
