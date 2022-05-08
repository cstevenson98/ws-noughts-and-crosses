package hub

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Hub) Login(w http.ResponseWriter, r *http.Request) {

}

type NewUser struct {
	Username string
	Password string
}

func (h *Hub) Signup(w http.ResponseWriter, r *http.Request) {
	var newUser NewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)

	h.Users = append(h.Users, User{
		Username:       newUser.Username,
		HashedPassword: string(hashedPass),
	})
}
