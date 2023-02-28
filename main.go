package main

import (
	"fmt"
	"net/http"

	. "github.com/MarloGrayhat/NDMSystem/queues"
)

var petQueue Pet

type Iqueue interface {
	GetQueue() []string
	SetQueue(q []string)
}

func Handler111(q Iqueue, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		if ok := r.URL.Query().Has("v"); ok { // if the v parameter is not empty
			queue := q.GetQueue()
			queue = append(queue, r.URL.Query().Get("v"))
			q.SetQueue(queue)            // add the v parameter value to the queue
			w.WriteHeader(http.StatusOK) // request status 200 (ok)
			fmt.Println(q.GetQueue())
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		break
	case http.MethodGet:
		// pop from pet queue
		queue := q.GetQueue()
		fmt.Println(len(queue))
		if len(queue) != 0 {
			el := queue[0]
			queue = queue[1:]
			q.SetQueue(queue)
			w.Write([]byte(el)) // return pet in http request body
		} else {
			if ok := r.URL.Query().Has("timeout"); ok {
				// nSec, err := strconv.ParseInt(r.URL.Query().Get("timeout"), 10, 32)
				// if err != nil {
				// 	w.WriteHeader(http.StatusBadRequest)
				// 	break
				// }
				// горутина где будет обратный отсчет
			} else {
				w.WriteHeader(http.StatusNotFound) // request status 404 (not found)
			}
		}

	default:
		fmt.Println(http.MethodHead)
		w.WriteHeader(http.StatusMethodNotAllowed)
		break
	}
}

func MW(w http.ResponseWriter, r *http.Request) {
	var handler Iqueue
	handler = &petQueue

	Handler111(handler, w, r)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/pet", MW)
	fmt.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println("ППЦ")
		panic("ППЦ")
	}
}
