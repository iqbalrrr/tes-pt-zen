package helpers

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDatabase() {
	var err error
	connStr := "user=yourusername password=yourpassword dbname=online_store sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	log.Println("Connected to the database successfully")

	// Buat tabel products
	createProductTable := `
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        description TEXT,
        price DECIMAL(10, 2),
        stock INTEGER
    );`
	_, err = db.Exec(createProductTable)
	if err != nil {
		log.Fatal("Failed to create products table:", err)
	}

	// Buat tabel orders
	createOrderTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        customer_name VARCHAR(100),
        product_id INTEGER REFERENCES products(id),
        quantity INTEGER,
        status VARCHAR(20)
    );`
	_, err = db.Exec(createOrderTable)
	if err != nil {
		log.Fatal("Failed to create orders table:", err)
	}

	log.Println("Tables created or already exist")
}
