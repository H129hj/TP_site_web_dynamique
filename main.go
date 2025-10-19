package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Produit struct {
	ID          int
	Nom         string
	Prix        float64
	Image       string
	Reduction   bool
	Description string
}

var produits = []Produit{
	{1, "PALACE Pull à Capuche Unisexe Chasseur", 145, "/static/assets/img/products/hoodie1.jpg", false, "Sweat à capuche Palace en coton biologique pour homme et femme, parfait pour un style streetwear décontracté."},
	{2, "PALACE Pull à Capuche Marine", 135, "/static/assets/img/products/hoodie2.jpg", true, "Hoodie Palace couleur marine en coton doux avec logo brodé, idéal pour toutes les saisons."},
	{3, "PALACE Pull Crew Passepoil Noir", 125, "/static/assets/img/products/hoodie3.jpg", false, "Sweat crew neck Palace noir avec passepoils contrastés, matière premium pour un look raffiné."},
	{4, "PALACE Washed Terry 1/4 Placket Hoodie Mojito", 165, "/static/assets/img/products/hoodie4.jpg", true, "Hoodie Palace couleur mojito en tissu terry washed avec fermeture boutonnée, effet vintage authentique."},
	{5, "PALACE Pantalon Bossy Jean Stone", 125, "/static/assets/img/products/pants1.jpg", false, "Jean Palace coupe large avec effet stone wash, denim de qualité pour un style streetwear."},
	{6, "PALACE Pantalon Cargo Gore-Tex Noir", 110, "/static/assets/img/products/pants2.jpg", true, "Pantalon cargo Palace noir avec technologie Gore-Tex, imperméable et résistant pour activités outdoor."},
}

func nomUnique(nom string) string {
	ext := filepath.Ext(nom)
	sansExt := strings.TrimSuffix(nom, ext)
	return fmt.Sprintf("%s_%d%s", sansExt, time.Now().Unix(), ext)
}

func sauvegarderImage(r *http.Request, champ string) (string, error) {
	fichier, handler, err := r.FormFile(champ)
	if err != nil {
		return "", err
	}
	defer fichier.Close()

	dossier := "static/assets/img/products"
	if _, err := os.Stat(dossier); os.IsNotExist(err) {
		os.MkdirAll(dossier, 0755)
	}

	nomFichier := nomUnique(handler.Filename)
	chemin := filepath.Join(dossier, nomFichier)

	destination, err := os.Create(chemin)
	if err != nil {
		return "", err
	}
	defer destination.Close()

	_, err = io.Copy(destination, fichier)
	if err != nil {
		return "", err
	}

	return "/static/assets/img/products/" + nomFichier, nil
}

func main() {
	templates, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println("Erreur templates:", err)
		os.Exit(1)
	}

	repertoire, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(repertoire + "/static"))

	http.HandleFunc("/static/assets/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		http.ServeFile(w, r, repertoire+"/static/assets/css/style.css")
	})

	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.HandleFunc("/", pageAccueil(templates))
	http.HandleFunc("/produit", pageProduit(templates))
	http.HandleFunc("/ajouter", pageAjouter(templates))
	http.HandleFunc("/ajouterForm", traiterAjout)
	http.HandleFunc("/supprimer", traiterSuppression)

	fmt.Println("Serveur démarré sur http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func pageAccueil(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/" {
			http.NotFound(w, r)
			return
		}
		templates.ExecuteTemplate(w, "index", produits)
	}
}

func pageProduit(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)

		var produit *Produit
		for _, p := range produits {
			if p.ID == id {
				produit = &p
				break
			}
		}

		if produit != nil {
			templates.ExecuteTemplate(w, "details", produit)
		} else {
			http.NotFound(w, r)
		}
	}
}

func pageAjouter(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "add", nil)
	}
}

func traiterAjout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	nom := strings.TrimSpace(r.FormValue("nom"))
	if nom == "" {
		http.Error(w, "Nom obligatoire", http.StatusBadRequest)
		return
	}

	prixStr := strings.TrimSpace(r.FormValue("prix"))
	prix, err := strconv.ParseFloat(prixStr, 64)
	if err != nil || prix <= 0 {
		http.Error(w, "Prix invalide", http.StatusBadRequest)
		return
	}

	description := strings.TrimSpace(r.FormValue("description"))
	if description == "" {
		http.Error(w, "Description obligatoire", http.StatusBadRequest)
		return
	}

	reduction := r.FormValue("reduction") == "on"

	imagePath := "/static/assets/img/products/hoodie5.jpg"
	if fichier, _, err := r.FormFile("image"); err == nil {
		fichier.Close()
		if chemin, err := sauvegarderImage(r, "image"); err == nil {
			imagePath = chemin
		}
	}

	nouveau := Produit{
		ID:          len(produits) + 1,
		Nom:         nom,
		Prix:        prix,
		Image:       imagePath,
		Reduction:   reduction,
		Description: description,
	}

	produits = append(produits, nouveau)
	http.Redirect(w, r, "/produit?id="+strconv.Itoa(nouveau.ID), http.StatusSeeOther)
}

func traiterSuppression(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	for i, p := range produits {
		if p.ID == id {
			produits = append(produits[:i], produits[i+1:]...)
			break
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
