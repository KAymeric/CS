package api

import (
	"cs/db"
	"encoding/json"
	"log"
	"net/http"
)

type client struct {
	Id      int    `json:"id"`
	Name     string `json:"name"`
	Adress string `json:"adress"`
	Siret   string `json:"siret"`
}

func Clients(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		create(w, r)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	var c client

	json.NewDecoder(r.Body).Decode(&c)

	query := "INSERT INTO clients (nom, adresse, siret) VALUES (?, ?, ?)"
	_, err := db.DB.Exec(query, c.Name, c.Adress, c.Siret)
	if err != nil {
		w.Write([]byte("Erreur lors de la création du client"))
		log.Fatalf("Erreur lors de l'exécution de la requête : %v", err)
		return
	}

	w.Write([]byte("Client crée"))
}

func list(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM clients")
	if err != nil {
		w.Write([]byte("Erreur lors de la récupération des clients"))
		log.Fatalf("Erreur lors de l'exécution de la requête : %v", err)
		return
	}
	defer rows.Close()

	var clients []client
	for rows.Next() {
		var c client
		rows.Scan(&c.Id, &c.Name, &c.Adress, &c.Siret)
		clients = append(clients, c)
	}

	json.NewEncoder(w).Encode(clients)
}
