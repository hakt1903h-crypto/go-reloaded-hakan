# Séance 1 — Corrigé des exercices

**Application CLI & Parsing en Go — Bachelor 1 — Séance 1**

Pour chaque exercice : la solution, une explication pas à pas, et les erreurs les plus fréquentes.

---

# Exercice 1 — Le terminal

## Les commandes, dans l'ordre

```bash
pwd
ls
cd Documents
mkdir go-cours
cd go-cours
cd ..
cd go-c        ← puis appuyez sur Tab : le terminal complète en go-cours
go version
```

## Ce que vous devez voir

`pwd` affiche le dossier courant, par exemple :

```
C:\Users\lea\Documents\go-cours
```

`go version` affiche un numéro de version :

```
go version go1.22.2 windows/amd64
```

Le numéro exact peut être différent : ce qui compte, c'est qu'il s'affiche.

## Ce qu'il faut comprendre

**Le terminal est toujours « quelque part »**. Chaque commande s'exécute dans le dossier indiqué par l'invite. C'est la source de la plupart des erreurs des premières séances : une commande correcte, lancée dans le mauvais dossier.

**`cd ..`** — les deux points désignent toujours le dossier parent, celui qui contient le dossier actuel.

## Erreurs courantes

**`cd Documents` ne fonctionne pas**. Le dossier Documents n'est peut-être pas à l'endroit où vous êtes, ou il porte un autre nom sur votre ordinateur. Utilisez `ls` pour voir les vrais noms des dossiers présents.

**Un nom de dossier avec un espace**. `cd Mes Documents` échoue : le terminal croit recevoir deux mots. Mettez le nom entre guillemets : `cd "Mes Documents"`. Mieux encore, évitez les espaces dans les noms de dossiers que vous créez.

**`go version` n'est pas reconnu**. Go n'est pas installé, ou VS Code était déjà ouvert pendant l'installation. Fermez complètement VS Code et rouvrez-le.

---

# Exercice 2 — Le premier programme

## Mise en place

```bash
mkdir bonjour
cd bonjour
go mod init bonjour
```

## `main.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("Bonjour !")
	fmt.Println("Je m'appelle Léa.")
	fmt.Println("J'apprends le Go.")
}
```

## Lancement

```bash
go run .
```

```
Bonjour !
Je m'appelle Léa.
J'apprends le Go.
```

## Ce qu'il faut comprendre

**Les instructions s'exécutent dans l'ordre**, de haut en bas. Inversez deux lignes, l'affichage s'inverse aussi.

**Chaque `Println` produit une ligne**. Le « ln » à la fin veut dire « line » : un retour à la ligne est ajouté automatiquement.

**L'apostrophe de `J'apprends` ne pose aucun problème**, parce que le texte est délimité par des guillemets doubles `"`.

## 2.3 — Le guillemet manquant

```go
	fmt.Println("Bonjour !)
```

```
./main.go:6:25: newline in string
```

Go a atteint la fin de la ligne sans trouver le guillemet fermant. Le message indique la ligne 6 : c'est bien là que se trouve l'oubli.

## 2.4 — La minuscule

```go
	fmt.println("Bonjour !")
```

```
./main.go:6:6: undefined: fmt.println
```

`undefined` veut dire « inconnu ». La boîte à outils `fmt` contient `Println`, avec une majuscule, mais pas `println`. Pour Go, ce sont deux mots différents.

## 2.5 — L'import supprimé

```
./main.go:4:2: undefined: fmt
```

Sans la ligne `import "fmt"`, Go ne sait pas ce qu'est `fmt`.

## Erreurs courantes

**`go.mod file not found`**. Le `go mod init` n'a pas été fait, ou le terminal n'est pas dans le dossier `bonjour`. Tapez `pwd` pour vérifier.

**Le fichier s'appelle `main.go.txt`**. Certains éditeurs ajoutent une extension cachée. Dans VS Code, créez le fichier depuis l'explorateur de gauche pour éviter ce problème.

**Les accolades mal placées**. En Go, l'accolade ouvrante `{` doit être sur la **même ligne** que `func main()`. La mettre à la ligne suivante provoque une erreur.

---

# Exercice 3 — Les variables

## Mise en place

```bash
cd ..
mkdir identite
cd identite
go mod init identite
```

## `main.go`

```go
package main

import "fmt"

func main() {
	prenom := "Léa"
	ville := "Paris"
	age := 19

	fmt.Println("Prénom :", prenom)
	fmt.Println("Ville :", ville)
	fmt.Println("Âge :", age)

	age = age + 1
	fmt.Println("L'année prochaine, j'aurai", age, "ans.")

	phrase := prenom + " habite à " + ville + "."
	fmt.Println(phrase)
}
```

## Résultat

```
Prénom : Léa
Ville : Paris
Âge : 19
L'année prochaine, j'aurai 20 ans.
Léa habite à Paris.
```

## Ce qu'il faut comprendre

**`age = age + 1`** — Go calcule d'abord ce qui est à droite du `=` : la valeur actuelle de `age`, plus 1. Puis il range le résultat dans la boîte `age`. L'ancienne valeur est remplacée.

**La virgule dans `Println`** — elle sépare plusieurs éléments à afficher, et Go ajoute un espace entre chacun. C'est pourquoi on n'écrit pas d'espace après `"Prénom :"`.

**Le `+` entre deux textes** — il les colle sans rien ajouter. D'où les espaces écrits à l'intérieur des guillemets : `" habite à "`.

## Une variante également correcte

```go
	fmt.Println(prenom, "habite à", ville+".")
```

Le résultat est presque identique. Les deux façons sont justes : choisissez celle que vous trouvez la plus lisible.

## Erreurs courantes

**`no new variables on left side of :=`**. Vous avez écrit `age := age + 1`. La boîte `age` existe déjà : on utilise `=`, sans les deux-points.

**`undefined: age`**. L'inverse : vous avez écrit `age = 19` alors que la boîte n'existait pas encore. La première fois, c'est `:=`.

**`declared and not used`**. Une variable créée mais jamais utilisée. Go refuse, pour éviter le code inutile. Utilisez-la ou supprimez-la.

**`Léahabite àParis.`**. Les espaces ont été oubliés dans les guillemets.

**Mélanger texte et nombre avec `+`**. `"Âge : " + age` provoque une erreur, parce qu'on ne peut pas coller un texte et un nombre. Utilisez la virgule dans `Println` : `fmt.Println("Âge :", age)`.

---

# Récapitulatif des points difficiles

**Toujours vérifier dans quel dossier se trouve le terminal.** `pwd` avant de chercher plus loin.

**Lire le numéro de ligne dans un message d'erreur.** C'est là qu'il faut regarder.

**Les majuscules comptent.** `Println` et `println` sont deux mots différents.

**`:=` la première fois, `=` ensuite.**

**Le texte entre guillemets, les noms de variables sans.**
