package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	DefaultRedisUrl      = "localhost:6379"
	DefaultRedisPassword = ""
)

func main() {

	// dbURL := env.EnvString("DsATABASE_URL", DefaultRedisUrl)
	// redisPassword := env.EnvString("REsDIS_PASSWORD", DefaultRedisPassword)
	// rconn := redisConnect(dbURL, redisPassword)
	// defer rconn.Close()

	// userRepo := redisdb.NewRedisUserRepository(rconn)
	// userService := user.NewUserService(userRepo)
	// userHandler := user.NewUserHandler(userService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ok", Ok).Methods("GET")
	// router.HandleFunc("/auth", userHandler.GetToken).Methods("GET")

	http.Handle("/", accessControl(router))

	errs := make(chan error, 2)
	go func() {
		logrus.Info("Listening server mode on port :3000")
		errs <- http.ListenAndServe(":3000", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logrus.Errorf("terminated %s", <-errs)

}
func  Ok(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Test-Header", "Works")
	w.WriteHeader(http.StatusOK)
	
}

// func redisConnect(url string, password string) *redis.Client {

// 	logrus.WithField("connection", url).Info("Connecting to Redis DB")
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     url,
// 		Password: password, // no password set
// 		DB:       0,        // use default DB
// 	})
// 	err := client.Ping().Err()

// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// 	return client

// }

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
