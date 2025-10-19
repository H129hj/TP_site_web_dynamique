# StreetShop - Site Web Dynamique

Site web de vente de vêtements streetwear développé en Go.

## Structure du projet

```
TP_site_web_dynamique
├── main.go                    # Serveur web principal
├── templates/                 # Templates HTML
│   ├── index.html            # Page d'accueil
│   ├── details.html          # Page produit
│   ├── add.html              # Formulaire ajout
│   └── layout.html           # Template de base
└── static/                   # Fichiers statiques
    └── assets/
        ├── css/
        │   └── style.css     # Styles CSS
        └── img/
            ├── logo/
            └── products/
```

## Fonctionnalités

- Affichage des produits en grille
- Page de détails pour chaque produit
- Ajout de nouveaux produits avec image
- Suppression de produits créés
- Formulaire de validation des données

## Installation et utilisation

1. Installer Go sur votre machine
2. Ouvrir un terminal dans le dossier du projet
3. Exécuter la commande :
   ```
   go run main.go
   ```
4. Ouvrir le navigateur sur http://localhost:8000

## Les Pages disponibles

- `/` - Page d'accueil avec tous les produits
- `/produit?id=X` - Détails d'un produit spécifique
- `/ajouter` - Formulaire pour ajouter un produit
- `/ajouterForm` - Traitement du formulaire d'ajout
- `/supprimer` - Suppression d'un produit

## Technologies utilisées

- Go (langage de programmation)
- HTML/CSS (interface utilisateur)
- Templates Go (génération de pages)
- Serveur HTTP intégré

## Structure des données

Chaque produit contient :
- ID unique
- Nom du produit
- Prix
- Chemin vers l'image
- Statut de réduction
- Description

## Validation des données

Le formulaire d'ajout vérifie :
- Nom obligatoire
- Prix numérique et positif
- Description obligatoire
- Image optionnelle (image par défaut si non fournie)