# set config to env
export NAME=go-url-shortener
export PORT=8080

export DB_NAME=url_shortener
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=password

export NEW_RELIC_LICENSE_KEY=4dbfe534f82770835d7b4ad8152eff5a34f4NRAL

# run go application
go run cmd/api/main.go