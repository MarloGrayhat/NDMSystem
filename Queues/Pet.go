package queues

import "net/http"

type Pet struct {
	PetQueue  []string // pet queue
	UserQueue []http.ResponseWriter
}

func (p *Pet) GetQueue() []string {
	return p.PetQueue
}
func (p *Pet) SetQueue(q []string) {
	p.PetQueue = q
}
func (p *Pet) GetUser() http.ResponseWriter {
	if len(p.UserQueue) > 0 {
		user := p.UserQueue[0]
		p.UserQueue = p.UserQueue[1:]
		return user
	} else {
		return nil
	}

}
func (p *Pet) AddUsers(w http.ResponseWriter) {
	p.UserQueue = append(p.UserQueue, w)
}
