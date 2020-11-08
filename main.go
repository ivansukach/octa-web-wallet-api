package main

//func customCORSMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9080")
//		// Call the next handler, which can be another middleware in the chain, or the final handler.
//		next.ServeHTTP(w, r)
//	})
//}

//import (
//	"log"
//	"net/http"
//
//	"github.com/go-chi/chi"
//	"github.com/go-chi/chi/middleware"
//)
//
//func main() {
//	r := chi.NewRouter()
//	r.Use(middleware.Logger)
//	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("welcome"))
//	})
//	r.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("welcome from endpoint"))
//	})
//	log.Fatal(http.ListenAndServe(":3000", r))
//}
