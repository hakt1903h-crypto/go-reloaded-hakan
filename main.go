package main

import (
	"fmt"
	"os"
)

func main() {
	// verif des argument (Séance 2 & 3)
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Erreur : il manque des arguments.")
		fmt.Fprintln(os.Stderr, "Utilisation : go run . <entree.txt> <sortie.txt>")
		os.Exit(1)
	}

	cheminEntree := os.Args[1]
	cheminSortie := os.Args[2]

	// interdire decraser le fichier d'entree (Séance 3)
	if cheminEntree == cheminSortie {
		fmt.Fprintln(os.Stderr, "Erreur : le fichier d'entrée et de sortie ne peuvent pas être identiques.")
		os.Exit(1)
	}

	//lecture fichier (Séance 3)
	texteBrut, err := lireFichier(cheminEntree)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}

	// transfo texte (Séance 4)
	texteModifie := transformer(texteBrut)

	// ecriture fichier sortie (Séance 3
	err = ecrireFichier(cheminSortie, texteModifie)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}

	fmt.Println("Terminé :", cheminSortie, "a été créé.")
}
