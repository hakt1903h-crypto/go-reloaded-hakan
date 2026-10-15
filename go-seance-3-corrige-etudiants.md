# Séance 3 — Corrigé des exercices

**Application CLI & Parsing en Go — Bachelor 1 — Séance 3**

Pour chaque exercice : la solution complète, une explication pas à pas, et les erreurs les plus fréquentes.

---

# Exercice 1 — Transformer un texte en nombre

## Mise en place

```bash
mkdir seance3
cd seance3
go mod init seance3
mkdir nombre
```

## `nombre/main.go`

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Erreur : il faut un nombre.")
		fmt.Println("Exemple : go run ./nombre 21")
		os.Exit(1)
	}

	texte := os.Args[1]

	nombre, err := strconv.Atoi(texte)
	if err != nil {
		fmt.Println("Erreur :", texte, "n'est pas un nombre entier.")
		os.Exit(1)
	}

	fmt.Println("Le double de", nombre, "est", nombre*2)
}
```

## Les trois tests

```
go run ./nombre 21      →  Le double de 21 est 42
go run ./nombre douze   →  Erreur : douze n'est pas un nombre entier.
go run ./nombre         →  Erreur : il faut un nombre.
```

## Ce qu'il faut comprendre

**Pourquoi transformer ?** Tout ce qu'on tape dans le terminal arrive sous forme de **texte**, même `21`. Pour calculer `21 * 2`, il faut d'abord transformer le texte `"21"` en nombre : c'est le rôle de `strconv.Atoi`.

**Pourquoi ça peut échouer ?** `"douze"` est un texte qui ne représente pas un nombre en chiffres. `Atoi` ne peut pas le transformer : il renvoie une erreur.

**Le `if err != nil` vient juste après.** On vérifie l'erreur **avant** d'utiliser `nombre`. Si on calculait d'abord, on utiliserait une valeur fausse.

## Erreurs courantes

**`assignment mismatch: 1 variable but strconv.Atoi returns 2 values`**. Vous avez écrit `nombre := strconv.Atoi(texte)`. La fonction renvoie deux choses : il faut deux variables, `nombre, err :=`.

**`declared and not used`** sur `err`. L'erreur est récupérée mais jamais vérifiée. Go l'interdit, justement pour vous obliger à y penser.

**Utiliser `_` à la place de `err`**. `nombre, _ := strconv.Atoi(texte)` est accepté par Go, mais c'est une très mauvaise idée : avec `douze`, le programme affiche « Le double de 0 est 0 » sans rien signaler.

---

# Exercice 2 — Lire un fichier

## `fichiers.go`

```go
package main

import (
	"fmt"
	"os"
)

// lireFichier renvoie le contenu du fichier, ou une erreur.
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}
```

## `main.go`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Erreur : il faut deux noms de fichiers.")
		os.Exit(1)
	}

	entree := os.Args[1]

	texte, err := lireFichier(entree)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}

	fmt.Println(texte)
}
```

## Les tests

```bash
go run . entree.txt sortie.txt
```

affiche le contenu de `entree.txt`.

```bash
go run . absent.txt sortie.txt
```

```
Erreur : impossible de lire le fichier absent.txt
```

## Ce qu'il faut comprendre

**`len(os.Args) != 3`** — on attend deux noms de fichiers. Avec le nom du programme, `os.Args` doit donc contenir **trois** éléments.

**La fonction ne fait qu'une chose.** `lireFichier` lit, et c'est tout. Elle n'affiche rien et n'arrête pas le programme : elle renvoie une erreur, et c'est `main` qui décide quoi en faire.

**`%s` dans `fmt.Errorf`** — c'est un emplacement, remplacé par la valeur qui suit. `fmt.Errorf("impossible de lire le fichier %s", "notes.txt")` donne le message `impossible de lire le fichier notes.txt`.

## Erreurs courantes

**`undefined: lireFichier`**. Deux causes possibles : le programme a été lancé avec `go run main.go`, qui n'inclut que ce fichier, ou `fichiers.go` ne commence pas par `package main`. Lancez avec `go run .`.

**`cannot use contenu (variable of type []byte) as string value`**. La conversion `string(contenu)` a été oubliée dans le `return`.

---

# Exercice 3 — Écrire un fichier

## `fichiers.go` complet

```go
package main

import (
	"fmt"
	"os"
)

// lireFichier renvoie le contenu du fichier, ou une erreur.
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}

// ecrireFichier écrit le texte dans le fichier, et renvoie une erreur si ça échoue.
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire le fichier %s", chemin)
	}
	return nil
}
```

## `main.go`

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Erreur : il faut deux noms de fichiers.")
		os.Exit(1)
	}

	entree := os.Args[1]
	sortie := os.Args[2]

	if entree == sortie {
		fmt.Println("Erreur : le fichier de sortie doit être différent du fichier d'entrée.")
		os.Exit(1)
	}

	texte, err := lireFichier(entree)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}

	err = ecrireFichier(sortie, texte)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}

	fmt.Println("Terminé :", sortie, "a été créé.")
}
```

## Ce qu'il faut comprendre

**`err = ecrireFichier(…)`** — avec un seul `=`. La variable `err` existe déjà depuis la lecture : on réutilise la même boîte.

**La vérification `entree == sortie` vient avant la lecture.** Inutile de lire un fichier si on sait déjà qu'on va refuser de continuer.

**Pourquoi cette vérification est indispensable.** `os.WriteFile` efface le fichier de sortie avant d'écrire. Si c'était le fichier d'entrée, le texte d'origine serait perdu.

## Erreurs courantes

**`no new variables on left side of :=`**. Vous avez écrit `err := ecrireFichier(…)`, alors que `err` existe déjà. Utilisez `=`.

**`644` au lieu de `0644`**. Le zéro du début est important : sans lui, Go comprend un tout autre nombre, et le fichier créé a des permissions bizarres.

**Le fichier de sortie n'apparaît pas dans VS Code**. Il a peut-être été créé dans un autre dossier : celui où se trouvait le terminal. Vérifiez avec `pwd` et `ls`.

---

# Exercice 4 — Messages clairs et nettoyage du texte

## `main.go`

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

## `fichiers.go`

Inchangé par rapport à l'exercice 3.

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
```

## Un test

`entree.txt` :

```
bonjour    le   monde

   deux   lignes
```

`sortie.txt` :

```
bonjour le monde

deux lignes
```

## Ce qu'il faut comprendre

**Le message de réussite reste sur la sortie normale.** `fmt.Println` pour ce qui s'est bien passé, `fmt.Fprintln(os.Stderr, …)` pour les erreurs.

**Les lignes vides sont conservées.** Sur une ligne vide, `strings.Fields` renvoie une liste vide, et `strings.Join` renvoie un texte vide : la ligne reste, simplement vide.

**Trois fichiers, trois rôles.** `main.go` organise le déroulé, `fichiers.go` s'occupe des fichiers, `regles.go` transforme le texte. Quand on cherche quelque chose, on sait dans quel fichier regarder.

## Erreurs courantes

**`ligne = strings.Join(mots, " ")`** au lieu de `lignes[i] = …`. Le programme tourne sans erreur, mais le fichier de sortie est identique à l'entrée. La variable `ligne` n'est qu'une copie : la modifier ne change pas la liste.

**Toutes les lignes collées en une seule**. Le texte a été découpé en mots directement, sans passer par les lignes. `strings.Fields` considère le retour à la ligne comme un espace.

**`undefined: strings`** dans `regles.go`. Chaque fichier a ses propres imports : `regles.go` doit importer `strings` lui-même.

---

# Exercice 5 — Synthèse

## Le tableau complété

| Commande | Message obtenu |
|---|---|
| `go run .` | Erreur : il faut deux noms de fichiers. + utilisation |
| `go run . entree.txt` | Erreur : il faut deux noms de fichiers. + utilisation |
| `go run . absent.txt sortie.txt` | Erreur : impossible de lire le fichier absent.txt |
| `go run . entree.txt entree.txt` | Erreur : le fichier de sortie doit être différent du fichier d'entrée. |
| `go run . entree.txt dossier_absent/sortie.txt` | Erreur : impossible d'écrire le fichier dossier_absent/sortie.txt |
| `go run . entree.txt sortie.txt` | Terminé : sortie.txt a été créé. |

Après chaque erreur, `go run` ajoute la ligne `exit status 1` : c'est normal, elle confirme que le programme s'est arrêté avec un code d'erreur.

## 5.2 — Améliorer un message

Le message `impossible de lire le fichier absent.txt` est clair, mais il ne dit pas **pourquoi**. Le fichier n'existe pas ? Il est protégé ?

Pour ceux qui veulent aller plus loin, on peut distinguer le cas le plus fréquent, le fichier absent :

```go
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("le fichier %s n'existe pas", chemin)
		}
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}
```

`errors.Is(err, os.ErrNotExist)` vérifie si l'erreur signifie « ce fichier n'existe pas ». Il faut ajouter `"errors"` aux imports. Cette amélioration est facultative.

---

# Exercice 6 — CampusLink : la commande `import-salles`

## L'organisation du projet

```
campuslink/
├── go.mod
├── main.go
├── fichiers.go
├── salles.go
└── donnees/
    ├── salles.csv
    └── salles-erreurs.csv
```

## `fichiers.go`

Copié depuis `textfix`, sans modification.

```go
package main

import (
	"fmt"
	"os"
)

// lireFichier renvoie le contenu du fichier, ou une erreur.
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}

// ecrireFichier écrit le texte dans le fichier, et renvoie une erreur si ça échoue.
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'écrire le fichier %s", chemin)
	}
	return nil
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
```

## `main.go`

```go
package main

import (
	"errors"
	"fmt"
	"os"
)

const usage = `Utilisation :
  campuslink aide
  campuslink import-salles <fichier_salles>`

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
	} else {
		err = fmt.Errorf("commande inconnue : %s\n%s", commande, usage)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}

// commandeImportSalles lit le fichier de salles et affiche les salles valides.
func commandeImportSalles(args []string) error {
	if len(args) != 1 {
		return errors.New("la commande import-salles attend un nom de fichier\n" + usage)
	}

	texte, err := lireFichier(args[0])
	if err != nil {
		return err
	}

	salles, problemes := lireSalles(texte)

	for _, s := range salles {
		afficherSalle(s)
	}
	for _, p := range problemes {
		fmt.Fprintln(os.Stderr, "Ligne ignorée :", p)
	}

	fmt.Println(len(salles), "salle(s) importée(s),", len(problemes), "ligne(s) ignorée(s).")
	return nil
}

// afficherSalle affiche une salle sur une ligne.
func afficherSalle(s Salle) {
	fmt.Println(s.ID, "-", s.Nom, "- bâtiment", s.Batiment, "-", s.Capacite, "places")
}
```

## Les tests

```bash
go run . import-salles donnees/salles.csv
```

```
S101 - Salle Turing - bâtiment A - 30 places
S102 - Salle Lovelace - bâtiment A - 24 places
S201 - Amphi Hopper - bâtiment B - 120 places
S202 - Salle Hamilton - bâtiment B - 18 places
4 salle(s) importée(s), 0 ligne(s) ignorée(s).
```

```bash
go run . import-salles donnees/salles-erreurs.csv
```

```
S101 - Salle Turing - bâtiment A - 30 places
S201 - Amphi Hopper - bâtiment B - 120 places
Ligne ignorée : ligne 3 : 4 champs attendus, 3 trouvés
Ligne ignorée : ligne 4 : la capacité "trente" n'est pas un nombre
Ligne ignorée : ligne 5 : l'identifiant est vide
Ligne ignorée : ligne 6 : la capacité doit être positive, 0 trouvé
2 salle(s) importée(s), 4 ligne(s) ignorée(s).
```

La ligne 7 du fichier est vide : elle est sautée sans être signalée.

| Commande | Résultat |
|---|---|
| `go run .` | Erreur : aucune commande indiquée. + utilisation |
| `go run . truc` | Erreur : commande inconnue : truc + utilisation |
| `go run . import-salles` | Erreur : la commande import-salles attend un nom de fichier + utilisation |
| `go run . import-salles absent.csv` | Erreur : impossible de lire le fichier absent.csv |

## Ce qu'il faut comprendre

**Le même schéma que `textfix`.** Lire le fichier, le découper en lignes, traiter chaque ligne, rassembler le résultat. Seul le traitement d'une ligne change : on ne corrige plus des mots, on remplit une fiche.

**Deux fonctions, deux niveaux.** `lireSalle` ne connaît qu'une ligne : elle ne sait pas quel est son numéro. `lireSalles` connaît toutes les lignes : c'est elle qui ajoute le numéro au message. Chaque fonction ajoute l'information qu'elle est seule à connaître.

**`lireSalles` ne s'arrête jamais sur une erreur.** Une ligne invalide est rangée dans `problemes`, puis `continue` passe à la ligne suivante. C'est pour cela que la fonction renvoie deux listes et pas d'`error` : il n'y a aucune situation où tout l'import échoue.

**`salles.go` n'affiche rien.** Il fabrique des salles et des messages. C'est `main.go` qui décide de les afficher, et où : les salles sur la sortie normale, les problèmes sur la sortie d'erreur.

**`strings.TrimSpace` sur chaque champ.** `S101, Salle Turing, A, 30` avec des espaces après les virgules est ainsi accepté. Elle retire aussi le `\r` des fichiers Windows en fin de ligne.

**Les crochets dans `os.Args[2:]`.** Si l'utilisateur tape seulement `go run . import-salles`, `os.Args[2:]` est une liste vide, sans plantage. C'est `commandeImportSalles` qui vérifie qu'elle contient bien un fichier.

## Erreurs courantes

**L'en-tête est signalé comme invalide** : `la capacité "capacite" n'est pas un nombre`. La condition `i == 0` a été oubliée.

**Les numéros de ligne sont décalés d'un cran.** Le message utilise `i` au lieu de `i+1`.

**`unknown field Capacité in struct literal of type Salle`.** Le nom de la case a été écrit avec un accent. Les noms des cases s'écrivent exactement comme dans la structure.

**`s.capacite undefined`.** Les cases commencent par une majuscule : `s.Capacite`.

**`cannot use champs[3] (variable of type string) as int value in struct literal`.** La capacité a été rangée sans conversion. Il faut passer par `strconv.Atoi`.

**L'import s'arrête à la première erreur.** Un `return` a remplacé le `continue` dans `lireSalles` : toutes les lignes suivantes sont perdues.

---

# Récapitulatif des points difficiles

**Une fonction qui peut échouer renvoie deux choses.** Il faut deux variables pour les récupérer.

**`if err != nil` juste après.** On vérifie l'erreur avant d'utiliser le résultat.

**Jamais `_` à la place de `err`.** Le programme continuerait avec une valeur fausse, sans prévenir.

**`go run .` avec le point.** Pour inclure tous les fichiers du projet.

**`lignes[i] = …` pour modifier une liste dans une boucle.**

**Ne jamais écrire dans le fichier qu'on lit.** `os.WriteFile` l'effacerait.

**Une ligne CSV se découpe avec `strings.Split(ligne, ",")`.** Tous les champs sont du texte.

**Une structure regroupe les informations d'une salle.** Les noms des cases s'écrivent exactement comme dans la structure.

**Une ligne invalide est ignorée et signalée, l'import continue.** `continue`, pas `return`.
