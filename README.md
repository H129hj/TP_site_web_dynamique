# 🛍️ StreetShop - Site Web Dynamique

Ce projet a été réalisé dans le cadre du TP **"Rendre un site web dynamique"**.
Il s'agit d'un site web en **Go (Golang)** permettant de gérer un catalogue de produits streetwear (hoodies et pantalons).

## Fonctionnalités
- 🏠 Page d'accueil avec liste des produits en grille 3x2
- 🔍 Détails d'un produit sélectionné
- ➕ Formulaire d'ajout de nouveaux produits avec upload d'images
- 🗑️ Suppression de produits
- 📱 Design responsif moderne
- ✅ Validation des données côté serveur
- 🎨 Interface utilisateur cohérente avec cartes grises et boutons rouges arrondis

## Structure du projet
```
V2/
├── main.go              # Code principal Go
├── templates/           # Templates HTML
│   ├── index.html      # Page d'accueil (grille 3x2)
│   ├── details.html    # Page de détails produit
│   └── add.html        # Formulaire d'ajout
├── static/             # Fichiers statiques
│   ├── assets/
│   │   ├── img/
│   │   │   ├── logo/logo.png
│   │   │   └── products/ (toutes les images des produits)
│   │   └── style.css   # Styles CSS modernes
│   └── assets/img/products/ (dossier d'upload automatique)
└── README.md
```

## Lancer le projet
1. Installer Go (https://go.dev)
2. Dans le terminal :
   ```bash
   go run main.go
   ```
3. Ouvrir http://localhost:8080

## Auteur
Projet réalisé par [Votre Nom] – StreetShop 2025.