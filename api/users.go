package api

import (
	//	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	//"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CtcUser struct {
	ID           int
	UserName     string
	UserCategory string
	TeamID       int
	UserAgegroup string
}

func Users(w http.ResponseWriter, r *http.Request) {
	// Load environment variables from file.
	/*if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
	*/

	// Connect to PlanetScale database using DSN environment variable.
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

	getCtcUsers(db, w, r)
}

// getCtcUsers is the HTTP handler for GET /users.
func getCtcUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var users []CtcUser
	result := db.Find(&users)
	if result.Error != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(users)
	body, err := json.Marshal(enc)

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(body)
	}
}
