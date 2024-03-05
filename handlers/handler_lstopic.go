package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

// Handler for the route /lstopic
func LsTopic(w http.ResponseWriter, r *http.Request) {
	lsTopicFilePath := "templates/lstopic.html"
	lsTopicFile, err := template.ParseFiles(lsTopicFilePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier HTML : %s", err), http.StatusInternalServerError)
		return
	}

	err = lsTopicFile.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du modèle HTML : %s", err), http.StatusInternalServerError)
		return
	}
}
