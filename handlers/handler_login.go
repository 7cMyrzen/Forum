package handlers

import (
	"fmt"
	"forum/db"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer les données du formulaire
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Chercher l'id de l'utilisateur dans la base de données
	id, err := db.FindUser(email, password)
	if err != nil {
		http.Error(w, "Erreur lors de la recherche de l'utilisateur dans la base de données", http.StatusInternalServerError)
		return
	}

	// Afficher l'id dans la console
	fmt.Println("ID de l'utilisateur:", id)

	// Redirection vers une page d'accueil après connexion réussie
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
