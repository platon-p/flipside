## Flipside

Содержимое `.env`:
```sh
POSTGRES_USER=admin
POSTGRES_DB=flipside
POSTGRES_PASSWORD=yourpassword

JWT_SIGN_KEY=yoursignkey
JWT_EXPIRES_IN=15m
REFRESH_TOKEN_EXPIRES_IN=24h
```
> A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
