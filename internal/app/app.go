package app

import (
	"github.com/antongoncharik/crypto-knight-api/internal/api/http/route"
)

func Run() {
	route.Init().Run(":8080")

	// db, err := sqlx.Connect("postgres", "user=postgres dbname=yourdatabase sslmode=disable password=yourpassword host=localhost")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully Connected")
	// }

	// type User struct {
	// 	Name  string `db:"username"`
	// 	Email string `db:"email"`
	// }

	// place := User{}

	// rows, _ := db.Queryx("SELECT username, email FROM users")

	// for rows.Next() {
	// 	err := rows.StructScan(&place)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	log.Printf("%#v\n", place)
	// }
}
