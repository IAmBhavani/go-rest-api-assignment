package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our students service
type Handler struct {
	Router  *mux.Router
	Service StudentService
	Server  *http.Server
}

// Response objecgi
type Response struct {
	Message string `json:"message"`
}

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

var users = map[string]string{
	"teacher1": "password1",
	"teacher2": "password2",
}

var mySigningKey = []byte("missionimpossible")

// NewHandler - returns a pointer to a Handler
func NewHandler(service StudentService) *Handler {
	log.Info("setting up our handler")
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	// Sets up our middleware functions
	h.Router.Use(JSONMiddleware)
	// we also want to log every incoming request
	h.Router.Use(LoggingMiddleware)
	// We want to timeout all requests that take longer than 15 seconds
	h.Router.Use(TimeoutMiddleware)
	// set up the routes
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}
	// return our wonderful handler
	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", h.AliveCheck).Methods("GET")
	h.Router.HandleFunc("/ready", h.ReadyCheck).Methods("GET")
	h.Router.HandleFunc("/api/v1/student", JWTAuth(h.PostStudent)).Methods("POST")
	h.Router.HandleFunc("/api/v1/students", JWTAuth(h.GetStudents)).Methods("GET")
	h.Router.HandleFunc("/api/v1/student/{id}", JWTAuth(h.GetStudent)).Methods("GET")
	h.Router.HandleFunc("/api/v1/student/{id}", JWTAuth(h.UpdateStudent)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/student/{id}", JWTAuth(h.DeleteStudent)).Methods("DELETE")

	h.Router.HandleFunc("/auth", Authenticate).Methods("POST")

}

func (h *Handler) AliveCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
		panic(err)
	}
}

func (h *Handler) ReadyCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.Service.ReadyCheck(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "I am Ready!"}); err != nil {
		panic(err)
	}
}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// to catch the OS Kill interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down gracefully")
	return nil
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if password, exists := users[user.ID]; exists && password == user.Password {
		tokenString, err := GenerateJWT(user.ID)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func GenerateJWT(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
