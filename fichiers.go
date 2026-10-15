package main

import (
	"fmt"
	"os"
)

// lireFichier lit le fichier d'entrée (Séance 3)[cite: 3]
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}

// ecrireFichier écrit le résultat dans le fichier de sortie (Séance 3)[cite: 3]
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire le fichier %s", chemin)
	}
	return nil
}
