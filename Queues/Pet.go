package queues

import (
	"net/http"
	"time"
)

type Pet struct {
	PetQueue  []string // pet queue
	UserQueue []struct {
		time.Time
		http.ResponseWriter
	}
}

func (p *Pet) GetQueue() []string {
	return p.PetQueue
}
func (p *Pet) SetQueue(q []string) {
	p.PetQueue = q
}

func (p *Pet) GetUserQueue() []struct {
	time.Time
	http.ResponseWriter
} {
	return p.UserQueue
}

func (p *Pet) SetUserToQueue(sec int, w http.ResponseWriter) {
	t := time.Now()
	p.UserQueue = append(p.UserQueue, struct {
		time.Time
		http.ResponseWriter
	}{
		t,
		w,
	})
}
