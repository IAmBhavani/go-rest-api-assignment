package main

import (
	"go-rest-api-assignment/internal/database"
	"go-rest-api-assignment/internal/student"
	transportHTTP "go-rest-api-assignment/internal/transport/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Run - sets up our application
func Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting Up Our APP")

	var err error
	store, err := database.NewDatabase()
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}
	//err = store.MigrateDB()
	//if err != nil {
	//	log.Error("failed to setup database")
	//	return err
	//}

	studentService := student.NewService(store)
	handler := transportHTTP.NewHandler(studentService)

	if err := handler.Serve(); err != nil {
		log.Error("failed to gracefully serve our application")
		return err
	}

	return nil
}

func main() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)
	// var mySigningKey = []byte("missionimpossible")

	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["authorized"] = true
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// claims["iat"] = time.Now().Unix()

	// tokenString, err := token.SignedString(mySigningKey)
	// if err != nil {
	// 	fmt.Println("Error generating token:", err)
	// 	return
	// }

	// fmt.Println("Generated Token:", tokenString)

	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}

}
