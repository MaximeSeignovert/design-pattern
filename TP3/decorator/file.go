package decorator

import (
	"os"
)

// IFichier définit l'interface de base pour tous les types de fichiers
type IFichier interface {
	// Enregistrer sauvegarde le fichier
	Enregistrer() error
	// GetNom retourne le nom du fichier
	GetNom() string
	// GetContenu retourne le contenu du fichier
	GetContenu() string
}

// FichierTexte implémente l'interface IFichier
type FichierTexte struct {
	nom     string
	contenu string
}

// NewFichierTexte crée une nouvelle instance de FichierTexte
func NewFichierTexte(nom, contenu string) *FichierTexte {
	return &FichierTexte{
		nom:     nom,
		contenu: contenu,
	}
}

// Enregistrer implémente la méthode de l'interface IFichier
func (f *FichierTexte) Enregistrer() error {
	// Création du fichier
	file, err := os.Create(f.nom)
	if err != nil {
		return err
	}
	defer file.Close()

	// Écriture du contenu
	_, err = file.WriteString(f.contenu)
	if err != nil {
		return err
	}

	return nil
}

// GetNom retourne le nom du fichier
func (f *FichierTexte) GetNom() string {
	return f.nom
}

// GetContenu retourne le contenu du fichier
func (f *FichierTexte) GetContenu() string {
	return f.contenu
}
