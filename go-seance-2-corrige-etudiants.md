# Séance 2 — Corrigé des exercices

**Application CLI & Parsing en Go — Bachelor 1 — Séance 2**

Pour chaque exercice : la solution complète, une explication pas à pas, et les erreurs les plus fréquentes.

## Mise en place du projet

```bash
mkdir seance2
cd seance2
go mod init seance2
mkdir conditions chaines listes fonctions crier
```

Chaque programme se lance depuis le dossier `seance2`, par exemple `go run ./conditions`.

---

# Exercice 1 — Les conditions

## `conditions/main.go`

```go
package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Tu es majeur.")
	} else {
		fmt.Println("Tu es mineur.")
	}

	note := 14

	if note >= 16 {
		fmt.Println("Très bien")
	} else if note >= 10 {
		fmt.Println("Admis")
	} else {
		fmt.Println("Non admis")
	}

	motDePasse := "go2026"

	if motDePasse == "go2026" {
		fmt.Println("Accès autorisé")
	} else {
		fmt.Println("Accès refusé")
	}
}
```

## Résultat

```
Tu es mineur.
Admis
Accès autorisé
```

## 1.4 — Les notes testées

| Note | Résultat | Pourquoi |
|---|---|---|
| 18 | Très bien | 18 ≥ 16, la première condition est vraie |
| 16 | Très bien | 16 ≥ 16 : avec `>=`, la limite est comprise |
| 10 | Admis | 10 n'est pas ≥ 16, mais 10 ≥ 10 |
| 9 | Non admis | Aucune condition n'est vraie, on arrive au `else` |

## Ce qu'il faut comprendre

**L'ordre des conditions compte.** Go s'arrête à la première vraie. Si on testait `note >= 10` en premier, une note de 18 afficherait « Admis », parce que 18 ≥ 10 est vrai.

**`>=` et `>` ne donnent pas le même résultat à la limite.** Avec `note > 16`, une note de 16 ne serait plus « Très bien ». C'est pour cela qu'on teste toujours les valeurs pile à la limite.

## Erreurs courantes

**`if motDePasse = "go2026"`** avec un seul `=`. Go refuse : un seul `=` range une valeur, il ne compare pas. Il faut `==`.

**`else` sur une nouvelle ligne**. En Go, `else` doit être sur la même ligne que l'accolade fermante : `} else {`.

**Les conditions dans le mauvais ordre**. Tout le monde est « Admis », même avec 18.

---

# Exercice 2 — Manipuler du texte

## `chaines/main.go`

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	mot := "Bonjour"

	fmt.Println(strings.ToUpper(mot))
	fmt.Println(strings.ToLower(mot))

	fmt.Println(strings.Contains(mot, "jour"))
	fmt.Println(strings.Contains(mot, "soir"))

	phrase := "  trop   d'espaces   ici  "
	fmt.Println(strings.TrimSpace(phrase))

	fmt.Println(strings.ReplaceAll("le chat et le chien", "le", "un"))

	message := "réunion urgent demain"
	if strings.Contains(message, "urgent") {
		fmt.Println("Message urgent !")
	}
}
```

## Résultat

```
BONJOUR
bonjour
true
false
trop   d'espaces   ici
un chat et un chien
Message urgent !
```

## 2.3 — Pourquoi les espaces du milieu restent

`TrimSpace` ne retire que les espaces **au début et à la fin** du texte. Ceux du milieu ne sont pas touchés. Pour les supprimer, il faudra découper la phrase en mots puis la recoller : c'est l'exercice 3.

## Ce qu'il faut comprendre

**`strings.Contains` renvoie `true` ou `false`.** On peut l'afficher, ou l'utiliser directement comme condition dans un `if`.

**Attention à `ReplaceAll`.** Il remplace **tous** les endroits où le texte apparaît, même au milieu d'un mot. `strings.ReplaceAll("le lecteur", "le", "un")` donne `"un uncteur"`. Il faut bien choisir ce qu'on remplace.

## Erreurs courantes

**`"strings" imported and not used`**. La boîte à outils est importée avant d'être utilisée. Ce n'est pas grave : l'erreur disparaît dès que vous écrivez le premier `strings.…`.

**`undefined: strings`**. L'inverse : vous utilisez `strings` sans l'avoir importé.

**`strings.toUpper`** avec une minuscule. Comme en séance 1, la majuscule compte.

---

# Exercice 3 — Les listes et les boucles

## `listes/main.go`

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fruits := []string{"pomme", "banane", "cerise"}

	fmt.Println(fruits[0])
	fmt.Println(len(fruits))

	fruits = append(fruits, "kiwi")
	fmt.Println(fruits)

	for i, fruit := range fruits {
		fmt.Println(i+1, "-", fruit)
	}

	var enMajuscules []string
	for _, fruit := range fruits {
		enMajuscules = append(enMajuscules, strings.ToUpper(fruit))
	}
	fmt.Println(enMajuscules)

	mots := strings.Fields("le   chat   dort")
	fmt.Println(len(mots), "mots")
	fmt.Println(strings.Join(mots, " "))
}
```

## Résultat

```
pomme
3
[pomme banane cerise kiwi]
1 - pomme
2 - banane
3 - cerise
4 - kiwi
[POMME BANANE CERISE KIWI]
3 mots
le chat dort
```

## Ce qu'il faut comprendre

**Construire une nouvelle liste.** On part d'une liste vide, `var enMajuscules []string`, et on y ajoute un élément à chaque tour de boucle. C'est **exactement** la technique qu'on utilisera pour le projet `textfix` : parcourir des mots, et construire une nouvelle liste de mots transformés.

**`strings.Fields` puis `strings.Join`.** Découper puis recoller fait disparaître les espaces en trop. C'est la solution au problème de l'exercice 2.3.

## Une variante également correcte

```go
for i := 0; i < len(fruits); i++ {
	fmt.Println(i+1, "-", fruits[i])
}
```

Une boucle qui compte de 0 jusqu'à la taille de la liste. Elle fonctionne aussi. `range` est plus court et évite les erreurs de comptage.

## Erreurs courantes

**`append(fruits, "kiwi")` tout seul**. Go refuse :

```
append(fruits, "kiwi") (value of type []string) is not used
```

Il faut récupérer le résultat : `fruits = append(fruits, "kiwi")`.

**`fruits[3]` avec trois éléments**. Le dernier est à l'index 2, puisqu'on compte à partir de 0. Le programme plante avec `index out of range`.

**`i declared and not used`**. Vous avez écrit `for i, fruit := range` sans vous servir de `i`. Remplacez-le par `_`.

---

# Exercice 4 — Les fonctions

## `fonctions/main.go`

```go
package main

import (
	"fmt"
	"strings"
)

func direBonjour(prenom string) {
	fmt.Println("Bonjour", prenom, "!")
}

func doubler(nombre int) int {
	return nombre * 2
}

func crier(texte string) string {
	return strings.ToUpper(texte) + " !"
}

func estLong(mot string) bool {
	return len(mot) > 6
}

func main() {
	direBonjour("Léa")
	direBonjour("Sam")

	fmt.Println(doubler(21))

	resultat := crier("attention")
	fmt.Println(resultat)

	fmt.Println(estLong("chat"))
	fmt.Println(estLong("ordinateur"))

	mot := "ordinateur"
	if estLong(mot) {
		fmt.Println(mot, ": mot long")
	} else {
		fmt.Println(mot, ": mot court")
	}
}
```

## Résultat

```
Bonjour Léa !
Bonjour Sam !
42
ATTENTION !
false
true
ordinateur : mot long
```

## Ce qu'il faut comprendre

**Afficher ou renvoyer.** `direBonjour` **affiche** directement. `crier` **renvoie** un texte, et c'est `main` qui décide de l'afficher. La seconde façon est plus souple : le résultat peut être affiché, rangé, ou réutilisé ailleurs.

**`return` termine la fonction.** Tout ce qui serait écrit après un `return` ne s'exécuterait jamais.

**Une petite limite de `len`.** `len` compte des octets, pas des lettres : un mot avec des accents paraît un peu plus long qu'il ne l'est. On reviendra là-dessus en séance 4.

## Erreurs courantes

**`missing return`**. La fonction annonce un résultat, par exemple `int`, mais n'a pas de `return`.

**`too many return values`**. L'inverse : un `return` avec une valeur dans une fonction qui n'annonce aucun résultat.

**Fonction écrite à l'intérieur de `main`**. Go refuse. Les fonctions s'écrivent les unes à côté des autres, jamais les unes dans les autres.

**`doubler("21")`**. Le paramètre attend un nombre, pas un texte :

```
cannot use "21" (untyped string constant) as int value in argument to doubler
```

---

# Exercice 5 — Les arguments du programme

## `crier/main.go`

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Erreur : il faut au moins un mot.")
		fmt.Println("Exemple : go run ./crier bonjour le monde")
		os.Exit(1)
	}

	mots := os.Args[1:]
	phrase := strings.Join(mots, " ")
	fmt.Println(strings.ToUpper(phrase) + " !")
}
```

## Les tests

```bash
go run ./crier attention au départ
```

```
ATTENTION AU DÉPART !
```

```bash
go run ./crier
```

```
Erreur : il faut au moins un mot.
Exemple : go run ./crier bonjour le monde
exit status 1
```

## Ce qu'il faut comprendre

**`len(os.Args) < 2`** — sans argument, `os.Args` contient seulement le nom du programme : sa taille est 1. Avec au moins un mot, elle vaut 2 ou plus.

**`exit status 1`** — cette dernière ligne n'est pas affichée par votre programme, mais par `go run`. Elle confirme que le programme s'est arrêté avec le code 1, comme demandé.

**Les accents passent en majuscules.** `ToUpper` transforme `é` en `É` sans difficulté.

## Erreurs courantes

**Le programme plante sans argument**. Le `if` de vérification a été placé **après** la lecture de `os.Args[1]`. Il doit venir avant.

**Le nom du programme apparaît dans la phrase**. Vous avez utilisé `os.Args` au lieu de `os.Args[1:]`.

**`"os" imported and not used`** ou **`undefined: os`**. L'import de `os` a été oublié, ou ajouté trop tôt.

---

# Récapitulatif des points difficiles

**`==` pour comparer, `=` pour ranger.**

**Tester les valeurs pile à la limite.** C'est là que se cachent les erreurs de `>` et `>=`.

**Un outil comme `ToUpper` fabrique un nouveau texte.** Il faut ranger le résultat.

**`liste = append(liste, x)`.** Toujours récupérer le résultat d'`append`.

**Les index commencent à 0.** Le dernier élément est à `len(liste) - 1`.

**Vérifier le nombre d'arguments avant de les lire.**
