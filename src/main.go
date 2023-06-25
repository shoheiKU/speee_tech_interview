package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func fibonacci(num, n_1, n_2 int64) int64 {
	if num == 1 {
		return n_1
	}
	if n_1 > math.MaxInt64/2 {
		panic(errors.New("num becomes too big"))
	}
	return fibonacci(num-1, n_1+n_2, n_1)
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			v := []byte(fmt.Sprintf(`{
	"status": 500,
	"message": "internal server error: %v"
}`, e))
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write(v)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	n, err := strconv.ParseInt(r.FormValue("n"), 10, 64)
	if err != nil {
		v := []byte(`{
	"status": 400,
	"message": "Bad request (a correct example:/fib?n=5)"
}`)
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	n = fibonacci(n, 1, 0)
	// ret := []byte(fmt.Sprintf("{\n\t\"result\": %d\n}", n))
	ret := []byte(fmt.Sprintf(`{
	"result": %d
}`, n))
	_, err = w.Write(ret)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/fib", fibonacciHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
