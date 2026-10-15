# Séance 4 — Corrigé des exercices

**Application CLI & Parsing en Go — Bachelor 1 — Séance 4**

Pour chaque exercice : la solution complète, une explication pas à pas, et les erreurs les plus fréquentes.

---

# Exercice 1 — Découper en lignes et en mots

## Mise en place

```bash
mkdir seance4
cd seance4
go mod init seance4
mkdir decouper
```

## `decouper/main.go`

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	texte := "le chat dort\nle chien   joue"

	lignes := strings.Split(texte, "\n")
	fmt.Println("Nombre de lignes :", len(lignes))

	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		fmt.Println("Ligne", i+1, ":", len(mots), "mots")

		for _, mot := range mots {
			fmt.Println("  -", mot)
		}
	}
}
```

## Résultat

```
Nombre de lignes : 2
Ligne 1 : 3 mots
  - le
  - chat
  - dort
Ligne 2 : 3 mots
  - le
  - chien
  - joue
```

## Ce qu'il faut comprendre

**`"\n"` dans le texte** — c'est un retour à la ligne écrit directement dans le code. Le texte contient donc bien deux lignes.

**La boucle intérieure tourne entièrement pour chaque ligne.** Pour la ligne 1, elle passe sur ses 3 mots ; puis la boucle extérieure passe à la ligne 2, et la boucle intérieure recommence.

**`_` dans la boucle intérieure** — on n'a pas besoin du numéro du mot, seulement du mot.

## Erreurs courantes

**`strings.Split(ligne, " ")` pour les mots**. La ligne 2 donne alors 5 éléments au lieu de 3 : les espaces en trop produisent des mots vides.

**`i declared and not used`** dans la boucle intérieure. Utilisez `_` quand vous n'avez pas besoin de l'index.

---

# Exercice 2 — Le marqueur `(up)`

## 2.1 — Le tableau sur papier

Texte : `bonjour (up) le monde`

| Mot lu | Ce qu'on fait | `resultat` après |
|---|---|---|
| `bonjour` | ajouté | `bonjour` |
| `(up)` | modifie le dernier mot | `BONJOUR` |
| `le` | ajouté | `BONJOUR le` |
| `monde` | ajouté | `BONJOUR le monde` |

## `regles.go`

```go
package main

import "strings"

// transformer nettoie le texte ligne par ligne :
// les espaces en trop sont supprimés.
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}

// appliquerMarqueurs construit une nouvelle liste de mots :
// un mot normal est ajouté, un marqueur modifie le dernier mot ajouté.
func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if mot == "(up)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}
```

## Le test temporaire dans `main.go`

En tout début de `func main()` :

```go
	fmt.Println(appliquerMarqueurs(strings.Fields("il est (up) parti")))
	fmt.Println(appliquerMarqueurs(strings.Fields("bonjour (up) le monde")))
	fmt.Println(appliquerMarqueurs(strings.Fields("il fait beau (up)")))
```

avec `"strings"` ajouté aux imports de `main.go`.

```
[il EST parti]
[BONJOUR le monde]
[il fait BEAU]
```

Les crochets indiquent que `Println` affiche une **liste**.

## Ce qu'il faut comprendre

**Le marqueur n'est jamais ajouté à `resultat`.** Il est traité dans le `if`, et l'`append` se trouve dans le `else`. C'est ce qui le fait disparaître du texte final.

**Pourquoi construire une nouvelle liste ?** On pourrait être tenté de modifier directement la liste `mots`. Mais le mot à transformer est « le dernier mot **gardé** », pas « le mot d'avant dans le texte ». Avec deux marqueurs à la suite, ces deux idées ne désignent plus le même mot.

**`il fait beau (up)`** — un marqueur en fin de ligne fonctionne : quand on le lit, `beau` est déjà dans `resultat`.

## Erreurs courantes

**Le marqueur apparaît dans le résultat**. L'`append` a été placé avant le `if`, et s'applique donc à tous les mots, y compris aux marqueurs.

**`resultat[len(resultat)]`** au lieu de `len(resultat) - 1`. Le programme plante : l'index `len(resultat)` est juste après le dernier élément.

---

# Exercice 3 — `(low)` et `(cap)`

## La fonction `capitaliser`

```go
// capitaliser met une majuscule au début du mot et le reste en minuscules.
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}
```

`regles.go` doit maintenant importer `strings` **et** `unicode`. Le code complet est donné à l'exercice 4.

## La partie de `appliquerMarqueurs` qui change

```go
	for _, mot := range mots {
		if mot == "(up)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		} else if mot == "(low)" {
			dernier := len(resultat) - 1
			resultat[dernier] = strings.ToLower(resultat[dernier])
		} else if mot == "(cap)" {
			dernier := len(resultat) - 1
			resultat[dernier] = capitaliser(resultat[dernier])
		} else {
			resultat = append(resultat, mot)
		}
	}
```

## Les tests

```
paris (cap)                              →  Paris
PARIS (cap)                              →  Paris
élise (cap)                              →  Élise
bonjour (up) TOUT (low) le monde (cap)   →  BONJOUR tout le Monde
```

## Ce qu'il faut comprendre

**`strings.ToLower` d'abord.** Sans lui, `PARIS (cap)` resterait `PARIS` : on mettrait en majuscule une lettre déjà majuscule, et le reste ne changerait pas.

**Les runes protègent les accents.** `[]rune("élise")` donne 5 lettres, dont la première est `é` entière. Sans cette conversion, on couperait le `é` en deux.

**`string(lettres)`** recolle la liste de lettres en un texte.

## Erreurs courantes

**`undefined: unicode`**. L'import manque dans `regles.go`.

**`strings.ToUpper(mot[:1]) + mot[1:]`**. Cette astuce fonctionne sur `paris`, mais donne un caractère illisible sur `élise`. C'est exactement le piège des octets.

**`(cap)` met tout le mot en majuscules**. `strings.ToUpper` a été utilisé au lieu de `capitaliser`.

---

# Exercice 4 — Ne jamais planter

## 4.1 — Le plantage

```
panic: runtime error: index out of range [-1]
```

Avec `(up) bonjour`, le marqueur arrive alors que `resultat` est vide. `len(resultat) - 1` vaut −1, un index qui n'existe pas.

## `regles.go` complet

```go
package main

import (
	"strings"
	"unicode"
)

// transformer applique les règles ligne par ligne.
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		mots = appliquerMarqueurs(mots)
		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}

// appliquerMarqueurs construit une nouvelle liste de mots :
// un mot normal est ajouté, un marqueur modifie le dernier mot ajouté.
func appliquerMarqueurs(mots []string) []string {
	var resultat []string

	for _, mot := range mots {
		if mot == "(up)" || mot == "(low)" || mot == "(cap)" {
			if len(resultat) == 0 {
				continue
			}

			dernier := len(resultat) - 1

			if mot == "(up)" {
				resultat[dernier] = strings.ToUpper(resultat[dernier])
			} else if mot == "(low)" {
				resultat[dernier] = strings.ToLower(resultat[dernier])
			} else {
				resultat[dernier] = capitaliser(resultat[dernier])
			}
		} else {
			resultat = append(resultat, mot)
		}
	}

	return resultat
}

// capitaliser met une majuscule au début du mot et le reste en minuscules.
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}
```

La ligne `mots = appliquerMarqueurs(mots)` dans `transformer` correspond à l'exercice 5.

## Ce qu'il faut comprendre

**Une seule vérification pour les trois marqueurs.** Le premier `if` regroupe les trois avec `||`. La vérification de la liste vide et le calcul de `dernier` sont écrits une seule fois. Le second `if` choisit ensuite la transformation.

**`continue` ignore le marqueur.** On passe directement au mot suivant, sans rien modifier et sans rien ajouter.

**Le dernier `else` du second `if`** — à cet endroit, le mot ne peut être que `(cap)`, puisqu'on sait déjà que c'est l'un des trois marqueurs et que ce n'est ni `(up)` ni `(low)`.

## 4.3 — Les situations du tableau

| Texte | Résultat |
|---|---|
| `(up) en début de ligne` | `en début de ligne` |
| `(up)` seul | ligne vide |
| `un (truc) inconnu` | `un (truc) inconnu` |
| ligne vide | ligne vide |
| `des   espaces   en   trop` | `des espaces en trop` |

## 4.4 — D'autres idées de textes mal écrits

Deux marqueurs à la suite, `bonjour (up) (low)` : le mot est mis en majuscules puis en minuscules, et donne `bonjour`. Aucun plantage.

Un marqueur collé à un mot, `bonjour(up)` : ce n'est pas un marqueur séparé, donc le mot reste tel quel. C'est une **limite** du programme, qu'on documentera en séance 5.

## Erreurs courantes

**La vérification placée après le calcul de `dernier`**. Le calcul en lui-même ne plante pas, mais la ligne suivante, oui. La vérification doit venir **avant** d'utiliser `dernier`.

**`return` au lieu de `continue`**. Le programme ne plante plus, mais tous les mots situés après le marqueur disparaissent : `return` arrête toute la fonction.

**`&&` au lieu de `||`**. Un mot ne peut pas être `(up)` **et** `(low)` à la fois : la condition n'est jamais vraie, et tous les marqueurs restent dans le texte.

---

# Exercice 5 — Brancher dans `textfix`

## `main.go`

Identique à la fin de la séance 3. La ligne de test et l'import de `strings` ont été retirés.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Erreur : il faut deux noms de fichiers.")
		fmt.Fprintln(os.Stderr, "Utilisation : textfix <fichier_entree> <fichier_sortie>")
		os.Exit(1)
	}

	entree := os.Args[1]
	sortie := os.Args[2]

	if entree == sortie {
		fmt.Fprintln(os.Stderr, "Erreur : le fichier de sortie doit être différent du fichier d'entrée.")
		os.Exit(1)
	}

	texte, err := lireFichier(entree)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}

	resultat := transformer(texte)

	err = ecrireFichier(sortie, resultat)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}

	fmt.Println("Terminé :", sortie, "a été créé.")
}
```

## `regles.go`

Celui de l'exercice 4, complet.

## `fichiers.go`

Inchangé depuis la séance 3.

## Le résultat sur `essai.txt`

```
BONJOUR tout le Monde
CE TEXTE crie un peu trop
Élise est arrivée
en début de ligne
un (truc) inconnu

des espaces en trop
```

## Ce qu'il faut comprendre

**Une seule ligne ajoutée dans `transformer`, et rien dans `main.go`.** C'est l'avantage d'avoir séparé le code en fichiers dès la séance 3 : la nouvelle règle s'est glissée au bon endroit sans toucher au reste.

**`CE TEXTE CRIE (low)`** — seul `CRIE` change. Un marqueur n'agit que sur **un** mot, celui qui le précède directement.

## Erreurs courantes

**`"strings" imported and not used`** dans `main.go`. L'import ajouté pour le test de l'exercice 2 est resté.

**`appliquerMarqueurs(mots)` sans `mots =`**. La fonction renvoie une nouvelle liste : sans récupérer le résultat, rien ne change dans le fichier de sortie.

---

# Bonus — `(hex)` et `(bin)`

## Le cas à ajouter dans `appliquerMarqueurs`

Juste avant le dernier `else` :

```go
		} else if mot == "(hex)" || mot == "(bin)" {
			if len(resultat) == 0 {
				continue
			}

			dernier := len(resultat) - 1
			resultat[dernier] = convertir(resultat[dernier], mot)
		} else {
```

## La fonction `convertir`

```go
// convertir transforme un nombre hexadécimal ou binaire en nombre décimal.
// Si le mot n'est pas un nombre valide, il est renvoyé sans changement.
func convertir(mot string, marqueur string) string {
	base := 16
	if marqueur == "(bin)" {
		base = 2
	}

	valeur, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(valeur, 10)
}
```

`regles.go` doit importer `strconv`.

## Les tests

```
il y a 1E (hex) fichiers et 10 (bin) dossiers   →  il y a 30 fichiers et 2 dossiers
la valeur zz (hex) reste                         →  la valeur zz reste
```

## Ce qu'il faut comprendre

**Le 64 de `ParseInt`** indique la taille maximale du nombre. Mettez toujours 64.

**Un mot invalide reste inchangé.** Si `err` n'est pas `nil`, on renvoie le mot d'origine : un seul mot invalide ne doit pas faire échouer tout le fichier.

---

# Exercice 6 — CampusLink : `controle`, `stats` et `export`

## `main.go`

```go
package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const usage = `Utilisation :
  campuslink aide
  campuslink import-salles <fichier_salles>
  campuslink controle <fichier_salles>
  campuslink stats <fichier_salles>
  campuslink export <fichier_salles> <fichier_sortie>`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Erreur : aucune commande indiquée.")
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	commande := os.Args[1]
	var err error

	if commande == "aide" {
		fmt.Println(usage)
	} else if commande == "import-salles" {
		err = commandeImportSalles(os.Args[2:])
	} else if commande == "controle" {
		err = commandeControle(os.Args[2:])
	} else if commande == "stats" {
		err = commandeStats(os.Args[2:])
	} else if commande == "export" {
		err = commandeExport(os.Args[2:])
	} else {
		err = fmt.Errorf("commande inconnue : %s\n%s", commande, usage)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}

// chargerSalles lit le fichier, renvoie les salles valides,
// et affiche les lignes ignorées sur la sortie d'erreur.
func chargerSalles(chemin string) ([]Salle, error) {
	texte, err := lireFichier(chemin)
	if err != nil {
		return nil, err
	}

	salles, problemes := lireSalles(texte)
	for _, p := range problemes {
		fmt.Fprintln(os.Stderr, "Ligne ignorée :", p)
	}

	return salles, nil
}

// commandeImportSalles lit le fichier de salles et affiche les salles valides.
func commandeImportSalles(args []string) error {
	if len(args) != 1 {
		return errors.New("la commande import-salles attend un nom de fichier\n" + usage)
	}

	salles, err := chargerSalles(args[0])
	if err != nil {
		return err
	}

	for _, s := range salles {
		afficherSalle(s)
	}
	fmt.Println(len(salles), "salle(s) importée(s).")
	return nil
}

// commandeControle vérifie qu'aucun identifiant de salle n'apparaît deux fois.
func commandeControle(args []string) error {
	if len(args) != 1 {
		return errors.New("la commande controle attend un nom de fichier\n" + usage)
	}

	salles, err := chargerSalles(args[0])
	if err != nil {
		return err
	}

	doublons := trouverDoublons(salles)
	for _, d := range doublons {
		fmt.Println("Problème :", d)
	}

	if len(doublons) > 0 {
		return fmt.Errorf("%d problème(s) trouvé(s)", len(doublons))
	}

	fmt.Println("Contrôle terminé : aucun problème trouvé.")
	return nil
}

// commandeStats affiche des statistiques simples sur les salles.
func commandeStats(args []string) error {
	if len(args) != 1 {
		return errors.New("la commande stats attend un nom de fichier\n" + usage)
	}

	salles, err := chargerSalles(args[0])
	if err != nil {
		return err
	}

	total := 0
	parBatiment := map[string]int{}

	for _, s := range salles {
		total = total + s.Capacite
		parBatiment[s.Batiment] = parBatiment[s.Batiment] + 1
	}

	var batiments []string
	for batiment := range parBatiment {
		batiments = append(batiments, batiment)
	}
	sort.Strings(batiments)

	fmt.Println("Nombre de salles :", len(salles))
	fmt.Println("Capacité totale :", total, "places")
	fmt.Println("Salles par bâtiment :")
	for _, batiment := range batiments {
		fmt.Println("  Bâtiment", batiment, ":", parBatiment[batiment])
	}
	return nil
}

// commandeExport écrit les salles valides dans un nouveau fichier CSV.
func commandeExport(args []string) error {
	if len(args) != 2 {
		return errors.New("la commande export attend deux noms de fichiers\n" + usage)
	}

	entree := args[0]
	sortie := args[1]
	if entree == sortie {
		return errors.New("le fichier de sortie doit être différent du fichier d'entrée")
	}

	salles, err := chargerSalles(entree)
	if err != nil {
		return err
	}

	lignes := []string{"id,nom,batiment,capacite"}
	for _, s := range salles {
		ligne := s.ID + "," + s.Nom + "," + s.Batiment + "," + strconv.Itoa(s.Capacite)
		lignes = append(lignes, ligne)
	}

	err = ecrireFichier(sortie, strings.Join(lignes, "\n")+"\n")
	if err != nil {
		return err
	}

	fmt.Println(len(salles), "salle(s) exportée(s) dans", sortie)
	return nil
}

// afficherSalle affiche une salle sur une ligne.
func afficherSalle(s Salle) {
	fmt.Println(s.ID, "-", s.Nom, "- bâtiment", s.Batiment, "-", s.Capacite, "places")
}
```

## `salles.go`

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Salle représente une salle du campus.
type Salle struct {
	ID       string
	Nom      string
	Batiment string
	Capacite int
}

// lireSalles transforme le contenu d'un fichier CSV en liste de salles.
// La première ligne est l'en-tête, elle est ignorée, comme les lignes vides.
// Chaque ligne invalide est signalée dans la liste des problèmes.
func lireSalles(texte string) ([]Salle, []string) {
	var salles []Salle
	var problemes []string

	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimSpace(ligne)
		if i == 0 || ligne == "" {
			continue
		}

		salle, err := lireSalle(ligne)
		if err != nil {
			problemes = append(problemes, fmt.Sprintf("ligne %d : %v", i+1, err))
			continue
		}

		salles = append(salles, salle)
	}

	return salles, problemes
}

// lireSalle transforme une ligne CSV en salle, ou renvoie une erreur.
func lireSalle(ligne string) (Salle, error) {
	champs := strings.Split(ligne, ",")
	if len(champs) != 4 {
		return Salle{}, fmt.Errorf("4 champs attendus, %d trouvés", len(champs))
	}

	for i, champ := range champs {
		champs[i] = strings.TrimSpace(champ)
	}

	if champs[0] == "" {
		return Salle{}, errors.New("l'identifiant est vide")
	}

	capacite, err := strconv.Atoi(champs[3])
	if err != nil {
		return Salle{}, fmt.Errorf("la capacité %q n'est pas un nombre", champs[3])
	}
	if capacite <= 0 {
		return Salle{}, fmt.Errorf("la capacité doit être positive, %d trouvé", capacite)
	}

	salle := Salle{
		ID:       champs[0],
		Nom:      champs[1],
		Batiment: champs[2],
		Capacite: capacite,
	}
	return salle, nil
}

// trouverDoublons renvoie un message pour chaque identifiant déjà rencontré.
func trouverDoublons(salles []Salle) []string {
	var problemes []string
	vus := map[string]bool{}

	for _, s := range salles {
		if vus[s.ID] {
			problemes = append(problemes, "l'identifiant "+s.ID+" apparaît plusieurs fois")
		}
		vus[s.ID] = true
	}

	return problemes
}
```

## `fichiers.go`

Inchangé depuis la séance 3. La fonction `ecrireFichier` sert maintenant à la commande `export`.

## Les tests

```bash
go run . stats donnees/salles.csv
```

```
Nombre de salles : 4
Capacité totale : 192 places
Salles par bâtiment :
  Bâtiment A : 2
  Bâtiment B : 2
```

```bash
go run . controle donnees/salles.csv
```

```
Contrôle terminé : aucun problème trouvé.
```

```bash
go run . controle donnees/salles-doublons.csv
```

```
Problème : l'identifiant S101 apparaît plusieurs fois
Erreur : 1 problème(s) trouvé(s)
exit status 1
```

```bash
go run . export donnees/salles-erreurs.csv donnees/propre.csv
```

Les quatre lignes ignorées s'affichent sur la sortie d'erreur, puis :

```
2 salle(s) exportée(s) dans donnees/propre.csv
```

`donnees/propre.csv` contient :

```
id,nom,batiment,capacite
S101,Salle Turing,A,30
S201,Amphi Hopper,B,120
```

## Ce qu'il faut comprendre

**`chargerSalles` évite quatre copies du même code.** Si demain il faut changer la façon de signaler une ligne ignorée, on modifie un seul endroit.

**`parBatiment[s.Batiment] + 1` fonctionne dès la première salle.** Une clé absente vaut 0 : la map crée le compteur toute seule.

**Le tri des clés rend l'affichage stable.** Sans `sort.Strings`, deux lancements sur le même fichier pourraient afficher B avant A. Un résultat qui change d'un lancement à l'autre est impossible à tester.

**Dans `trouverDoublons`, le test vient avant `vus[s.ID] = true`.** On regarde d'abord si l'identifiant a déjà été vu, puis on le note. Dans l'autre ordre, chaque salle se verrait elle-même comme un doublon.

**`controle` renvoie une erreur quand il trouve un problème.** C'est `main` qui l'affiche et arrête le programme avec le code 1, comme pour toutes les autres erreurs.

**`export` réutilise l'en-tête d'origine.** Le fichier produit peut être relu par `import-salles` : c'est un CSV propre, sans lignes invalides.

**`+ "\n"` à la fin de l'export.** `strings.Join` met un retour à la ligne **entre** les lignes, pas après la dernière. On l'ajoute pour que le fichier se termine proprement.

## Erreurs courantes

**`panic: assignment to entry in nil map`.** La map a été déclarée avec `var parBatiment map[string]int`. Sans `{}`, elle n'existe pas encore et on ne peut rien y ranger. Écrivez `map[string]int{}`.

**Les bâtiments changent d'ordre d'un lancement à l'autre.** La map a été parcourue directement, sans trier ses clés.

**`invalid operation: … (mismatched types untyped string and int)`.** La capacité a été collée au texte sans `strconv.Itoa`.

**`controle` affiche le doublon mais le code de sortie reste 0.** La commande renvoie `nil` au lieu d'une erreur.

**`undefined: sort`.** L'import `"sort"` manque dans `main.go`.

---

# Récapitulatif des points difficiles

**On construit une nouvelle liste.** Un mot normal est ajouté, un marqueur modifie le dernier mot et n'est pas ajouté.

**Le dernier élément est à `len(resultat) - 1`.**

**Vérifier que la liste n'est pas vide avant de modifier son dernier mot.** Sinon : `index out of range [-1]`.

**`continue` passe au mot suivant, `return` arrête toute la fonction.**

**`[]rune` pour toucher aux lettres une par une.** Une lettre accentuée occupe plusieurs octets.

**`liste = fonction(liste)`.** Toujours récupérer la liste renvoyée.

**Deux fois le même code, c'est une fonction.** `chargerSalles` sert aux quatre commandes.

**Une clé absente d'une map vaut 0, ou `false`.** Pas besoin de créer le compteur avant.

**L'ordre d'une map n'est pas fixe.** On trie les clés pour un affichage stable.

**Un contrôle qui trouve un problème renvoie une erreur.** Le programme se termine alors avec le code 1.
