package main

import (
	"fmt"
	"os"
)

//Lit le fichier (vu e n seance 3 cours charlepeau)
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}

//ecris dans le fichier sorti(Séance 3)
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire le fichier %s", chemin)
	}
	return nil
}
