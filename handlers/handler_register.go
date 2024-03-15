package handlers

import (
	"forum/db"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer les données du formulaire
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Vérifier si l'utilisateur existe déjà
	_, err := db.FindUser(username, password)
	if err == nil {
		http.Error(w, "L'utilisateur existe déjà", http.StatusConflict)
		return
	}

	// Ajouter l'utilisateur à la base de données
	err = db.AddUser(username, email, password)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout de l'utilisateur à la base de données", http.StatusInternalServerError)
		return
	}

	// Traiter les données d'inscription ici
	// (par exemple, valider les entrées et enregistrer l'utilisateur dans la base de données)

	// Redirection vers une page de confirmation d'inscription
	http.Redirect(w, r, "/registration-successful", http.StatusSeeOther)
}
