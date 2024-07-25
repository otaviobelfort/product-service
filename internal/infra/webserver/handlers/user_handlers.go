package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/otaviobelfort/go/product_service/internal/dto"
	"github.com/otaviobelfort/go/product_service/internal/entity"
	"github.com/otaviobelfort/go/product_service/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExpiriesIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiriesIn int) *UserHandler {
	return &UserHandler{UserDB: userDB, Jwt: jwt, JwtExpiriesIn: jwtExpiriesIn}
}

func (h *UserHandler) GetJwtToken(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//token1 := jwtauth.New("HS256", []byte("secret"), nil)
	//
	//_, tokenString2, _ := token1.Encode(map[string]interface{}{
	//	"sub": u.ID.String(),
	//	"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiriesIn)).Unix(),
	//})
	//println("tokenString2: ", tokenString2)

	_, tokenString, err := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiriesIn)).Unix(),
	})
	if err != nil {
		// Log the error or handle it accordingly
		println("Error encoding token_>>>>: ", string(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	println("tokenString2: ", tokenString)
	if tokenString == "" {
		println("tokenString: ", tokenString)
	}

	//tokenData := map[string]interface{}{
	//	"sub": u.ID.String(),
	//	"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiriesIn)).Unix(),
	//}
	//
	//_, tokenString, err := h.Jwt.Encode(tokenData)
	//if err != nil {
	//	// Log the error or handle it accordingly
	//	log.Printf("Error encoding token: %v", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}

	token := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		return
	}
}

// Create user godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dto.CreateUserRequest true "User request"
// @Success 201 {object} string
// @Failure 500 {object} string
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.UserDB.Create(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUserEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.UserDB.FindByEmail(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
