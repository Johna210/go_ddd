package chapter2

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	isUserSubscriptionActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	isActive bool
}

func router(u UserHandler) {
	m := mux.NewRouter()
	m.HandleFunc("/user/{userID}/subscription/active", func(w http.ResponseWriter, r *http.Request) {
		// check auth, etc

		uID := mux.Vars(r)["userID"]
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		isActive := u.isUserSubscriptionActive(r.Context(), uID)

		b, err := json.Marshal(UserActiveResponse{isActive: isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(b)
	}).Methods(http.MethodGet)
}
