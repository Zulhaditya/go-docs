package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

// error handling
type ErrorHandler struct {
	Handler http.Handler
}

func (handler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		fmt.Println("RECOVER : ", err)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()

	handler.Handler.ServeHTTP(writer, request)
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprintf(writer, "Hello middleware")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo executed")
		fmt.Fprintf(writer, "Hello foo")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo executed")
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	// middleware untuk handle error
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
