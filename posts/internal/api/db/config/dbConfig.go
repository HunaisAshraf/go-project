package dbconfig

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DbConfig() *pgxpool.Pool {
	fmt.Println("postgres env is : ", os.Getenv("DB_URL"))

	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("db connected")
	}
	// defer conn.Close()

	initialiseDB(conn)

	return conn
}

func initialiseDB(conn *pgxpool.Pool) {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(15) NOT NULL
	);
	`

	_, err := conn.Exec(context.Background(), createUserTable)
	if err != nil {
		log.Fatal(err)
	}

	createPostTable := `
		CREATE TABLE IF NOT EXISTS posts(
	    id BIGSERIAL PRIMARY KEY,
	    title VARCHAR(30) NOT NULL,
	    body VARCHAR(300) NOT NULL,
	    userid VARCHAR(50),
	    FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE
	);
	`

	_, err = conn.Exec(context.Background(), createPostTable)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db initialised")

}
