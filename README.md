# Postgres

A simple PostgreSQL connect library for golang.

將常用連線設定結構化，方便建立DB連線

## Usage

* Connect by structured config
	```go
	func main() {
		// Initialize connection pool
		pool, err := postgres.OpenByConf(postgres.DataSource{
			Host:        "localhost",
			Port:        5432,
			User:        "postgres",
			Password:    "password",
			DBName:      "postgres",
			SSLMode:     postgres.SSLModeDisable,
			ConnTimeout: 10 * time.Second,
		})
		if err != nil {
			// handle error
		}
		defer pool.Close()

		// Acquire new session
		conn, err := pool.Conn(context.Context())
		if err != nil {
			// handle error
		}
		defer conn.Close()

		// ... query sql
	}
	```

* Connect by environment variables
	```go
	func main() {
		pool, err := postgres.OpenDefault()
		if err != nil {
			// handle error
		}

		// ... 
	}
	```

* Connect by dataSourceName string (the same as original sql.Open)
	```go
	func main() {
		pool, err := postgres.Open("host=... port=... user=... password=... dbname=... ...")
		if err != nil {
			// handle error
		}

		// ...
	}
	```
