# Power 4 - Projet Web

Un jeu de Power 4 (Puissance 4) interactif développé en Go avec une interface web moderne.

## 📋 Description

Ce projet est une implémentation web du célèbre jeu Power 4, permettant à deux joueurs de s'affronter avec trois niveaux de difficulté différents. Le jeu propose une interface élégante avec un thème sombre et des effets visuels modernes.

## 🎮 Fonctionnalités

- **Trois niveaux de difficulté** :
  - Facile : Grille 6x7 (classique)
  - Normal : Grille 6x9
  - Difficile : Grille 7x8

- **Système de score** : Suivi des points entre les parties
- **Rejouer** : Possibilité de recommencer une partie en gardant les scores
- **Changement de difficulté** : Modifier le niveau en cours de jeu


## 🚀 Installation et lancement

### En ligne
- Ouvrez votre navigateur à l'adresse :
```
https://power4.prettyflacko.fr/
```


### Installation Locale

## Prérequis
- Go 1.22.2 ou supérieur


1. Clonez le dépôt :
```bash
git clone https://github.com/ANome1/projet_Power4.git
cd projet_Power4
```

2. Lancez le serveur :
```bash
go run server.go
```

3. Ouvrez votre navigateur à l'adresse :
```
http://localhost:8080
```

## 🎯 Comment jouer

1. **Entrez les noms des joueurs** sur la page d'accueil
2. **Choisissez une difficulté** (Easy, Normal ou Hard)
3. **Jouez à tour de rôle** en cliquant sur les cases
4. **Alignez 4 jetons** horizontalement, verticalement ou en diagonale pour gagner
5. **Rejouez** ou revenez à l'accueil après une victoire

## 📊 Logique du jeu

Le jeu implémente la logique complète de Power 4 :
- Placement des jetons par gravité (toujours dans la case vide la plus basse)
- Détection de victoire sur 4 directions :
  - Horizontale
  - Verticale
  - Diagonale descendante
  - Diagonale montante
- Alternance automatique des joueurs
- Gestion des scores cumulés

## 👤 Auteur

**Pretty Flacko**
- GitHub: [@ANome1](https://github.com/ANome1)
