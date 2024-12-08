package api

import (
	"encoding/json"
	"log"
	"net/http"
	"cs/db"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var c credentials

	json.NewDecoder(r.Body).Decode(&c)

	query := "SELECT password FROM users WHERE username = ?"
	rows, err := db.DB.Query(query, c.Username)
	if err != nil {
		log.Fatalf("Erreur lors de l'exécution de la requête : %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		var passwordDB string
		rows.Scan(&passwordDB)

		if c.Password == passwordDB {
			w.Write([]byte("Connexion réussie"))
		} else {
			w.Write([]byte("Mot de passe incorrect"))
		}
	} else {
		query = "INSERT INTO users (username, password) VALUES (?, ?)"
		_, err = db.DB.Exec(query, c.Username, c.Password)
		if err != nil {
			log.Fatalf("Erreur lors de l'exécution de la requête : %v", err)
		}

		w.Write([]byte("Utilisateur crée"))
	}
}
