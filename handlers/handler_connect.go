package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// Connect est le gestionnaire pour la route /connect

func ConnectHandler(w http.ResponseWriter, r *http.Request) {
	connectFilePath := "templates/connect.html"
	connectFile, err := template.ParseFiles(connectFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = connectFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
