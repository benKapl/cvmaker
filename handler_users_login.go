package main

//func (cfg *apiConfig) handlerUsersLogin(w http.ResponseWriter, r *http.Request) {
//	type parameters struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//
//	type response struct {
//		User
//	}
//
//	decoder := json.NewDecoder(r.Body)
//	params := parameters{}
//	err := decoder.Decode(&params)
//	if err != nil {
//		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
//		return
//	}
//
//	user, err := cfg.db.GetUser(r.Context(), params.Email)
//	if err != nil {
//		respondWithError(w, http.StatusUnauthorized, "Incorrect Email or Password", err)
//		return
//	}
//
//	err = goauth.CheckPasswordHash(params.Password, user.Password)
//	if err != nil {
//		respondWithError(w, http.StatusUnauthorized, "Incorrect Email or Password", err)
//		return
//	}
//
//	respondWithJSON(w, http.StatusCreated, response{
//		User: User{
//			ID:        int(user.ID),
//			CreatedAt: user.CreatedAt,
//			UpdatedAt: user.UpdatedAt,
//			Email:     user.Email,
//		},
//	})
//}
//
