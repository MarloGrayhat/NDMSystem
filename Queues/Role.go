package queues

import "net/http"

type Role struct {
	RoleQueue []string // pet queue
	UserQueue []http.ResponseWriter
}

func (p *Role) GetQueue() []string {
	return p.RoleQueue
}
func (p *Role) SetQueue(q []string) {
	p.RoleQueue = q
}
func (p *Role) GetUser() http.ResponseWriter {
	if len(p.UserQueue) > 0 {
		user := p.UserQueue[0]
		p.UserQueue = p.UserQueue[1:]
		return user
	} else {
		return nil
	}

}
func (p *Role) AddUsers(w http.ResponseWriter) {
	p.UserQueue = append(p.UserQueue, w)
}
