package middleware

import (
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return fn
}

func recoverHandler(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case error:
					defaultError := err.(error)
					log.Printf("panic: %+v", err)
					http.Error(w, defaultError.Error(), http.StatusInternalServerError)
				default:
					log.Printf("panic: %+v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	}

	return fn
}

func Load(handler http.HandlerFunc) http.HandlerFunc {
	resolvedHandler := recoverHandler(loggingHandler(handler))
	return resolvedHandler
}
