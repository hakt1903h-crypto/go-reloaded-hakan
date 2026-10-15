package main

import (
	"strconv"
	"strings"
	"unicode"
)

//accents avec runes
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

//modifie le mot précédent selon le marqueur (Séance 4)
func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		// gestion marqueurs (up)(low)(cap)(hex)(bin)
		if mot == "(up)" || mot == "(low)" || mot == "(cap)" || mot == "(hex)" || mot == "(bin)" {
			// anti-plantage (Séance 4 ne jamais planter si le texte commence par un marqueur)
			if len(resultat) == 0 {
				continue
			}
			dernier := len(resultat) - 1

			if mot == "(up)" {
				resultat[dernier] = strings.ToUpper(resultat[dernier])
			} else if mot == "(low)" {
				resultat[dernier] = strings.ToLower(resultat[dernier])
			} else if mot == "(cap)" {
				resultat[dernier] = capitaliser(resultat[dernier])
			} else if mot == "(hex)" {
				if val, err := strconv.ParseInt(resultat[dernier], 16, 64); err == nil {
					resultat[dernier] = strconv.FormatInt(val, 10)
				}
			} else if mot == "(bin)" {
				if val, err := strconv.ParseInt(resultat[dernier], 2, 64); err == nil {
					resultat[dernier] = strconv.FormatInt(val, 10)
				}
			}
		} else {
			// Mot normal on l'ajoute à notre liste de résultats (Séance 4)
			resultat = append(resultat, mot)
		}
	}

	return resultat
}

// change "a" en "an" si le mot d'après commence par une voyelle ou un h
func corrigerArticles(mots []string) []string {
	for i := 0; i < len(mots)-1; i++ {
		m := strings.ToLower(mots[i])
		if m == "a" {
			prochain := strings.ToLower(mots[i+1])
			if len(prochain) > 0 {
				p := prochain[0]
				if p == 'a' || p == 'e' || p == 'i' || p == 'o' || p == 'u' || p == 'h' {
					if mots[i] == "A" {
						mots[i] = "An"
					} else {
						mots[i] = "an"
					}
				}
			}
		}
	}
	return mots
}

// colle les ponctuations au mot précédent
func formaterPonctuation(ligne string) string {
	// On découpe et on recolle proprement avec des conditions simples
	mots := strings.Fields(ligne)
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		// Si c est un signe de ponctuation simple on le colle au dernier mot ajoute
		if (mot == "." || mot == "," || mot == "!" || mot == "?" || mot == ":" || mot == ";") && len(resultat) > 0 {
			resultat[len(resultat)-1] += mot
		} else {
			resultat = append(resultat, mot)
		}
	}
	return strings.Join(resultat, " ")
}

// transformer gère le texte ligne par ligne (Séance 3 et 4)
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r") //(Séance 3)
		
		mots := strings.Fields(ligne)          
		mots = appliquerMarqueurs(mots)        
		mots = corrigerArticles(mots)         
		
		ligneIntermediaire := strings.Join(mots, " ")
		lignes[i] = formaterPonctuation(ligneIntermediaire)
	}

	return strings.Join(lignes, "\n")
}
