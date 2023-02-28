package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	. "github.com/MarloGrayhat/NDMSystem/Queues"
)

var petQueue Pet
var roleQueue Role

type Iqueue interface {
	GetQueue() []string
	SetQueue(q []string)
	GetUser() http.ResponseWriter
	AddUsers(w http.ResponseWriter)
}

func Handler(q Iqueue, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		if ok := r.URL.Query().Has("v"); ok { // if the v parameter is not empty
			queue := q.GetQueue()
			value := r.URL.Query().Get("v")
			user := q.GetUser() // получение превого в очереди польлзователя, ожидающего ответа (совсем не проверяет, есть ли соединение или нет)
			if user != nil {
				user.Write([]byte(value))
			} else {
				queue = append(queue, value)
				q.SetQueue(queue)            // add the v parameter value to the queue
				w.WriteHeader(http.StatusOK) // request status 200 (ok)
				fmt.Println(q.GetQueue())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	case http.MethodGet:
		queue := q.GetQueue()
		fmt.Println(len(queue))
		if len(queue) != 0 {
			// pop from queue
			el := queue[0]
			queue = queue[1:]
			q.SetQueue(queue)
			w.Write([]byte(el)) // return pet in http request body
		} else {
			if ok := r.URL.Query().Has("timeout"); ok {
				nSec, err := strconv.ParseInt(r.URL.Query().Get("timeout"), 10, 32)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					break
				}
				q.AddUsers(w)                                 // добавляет соединение в очередь ожидающих ответа
				time.Sleep(time.Second * time.Duration(nSec)) // чтобы соединение не закрылось.
				w.WriteHeader(http.StatusNotFound)

			} else {
				w.WriteHeader(http.StatusNotFound) // request status 404 (not found)
			}
		}

	default:
		fmt.Println(http.MethodHead)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func MWPet(w http.ResponseWriter, r *http.Request) {
	var handler Iqueue
	handler = &petQueue

	Handler(handler, w, r)
}
func MWRole(w http.ResponseWriter, r *http.Request) {
	var handler Iqueue
	handler = &petQueue

	Handler(handler, w, r)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/pet", MWPet)
	mux.HandleFunc("/role", MWRole)
	fmt.Println("Запуск веб-сервера на http://127.0.0.1")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println("ППЦ")
		panic("ППЦ")
	}
}
