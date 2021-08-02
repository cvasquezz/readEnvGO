package external

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Wea struct {
	Default bool
	User    interface{}
}

type User struct {
	Nombre string
}

func (w *Wea) ReadEnv(u *User) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	fmt.Printf("os.Getenv(\"MESSAGE\"): %v\n", os.Getenv("MESSAGE"))
	if w.Default {
		fmt.Println(w.User)
	} else {

		fmt.Printf("User: %v\n", u)
	}
}
