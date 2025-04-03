package decorator

import (
	"fmt"
	"os"
	"path/filepath"
	"file-modifier/pkg"
)

// ChiffrementDecorator implémente le pattern Decorator pour le chiffrement
type ChiffrementDecorator struct {
	fichier IFichier
}

// NewChiffrementDecorator crée une nouvelle instance de ChiffrementDecorator
func NewChiffrementDecorator(fichier IFichier) *ChiffrementDecorator {
	return &ChiffrementDecorator{
		fichier: fichier,
	}
}

// Enregistrer implémente la méthode de l'interface IFichier avec chiffrement
func (c *ChiffrementDecorator) Enregistrer() error {
	// Lecture du contenu
	contenu := []byte(c.fichier.GetContenu())

	// Chiffrement du contenu
	contenuChiffre, err := pkg.Encrypt(contenu)
	if err != nil {
		return fmt.Errorf("erreur lors du chiffrement : %v", err)
	}

	// Création du dossier de sortie si nécessaire
	outputDir := filepath.Join("output_files", "encrypted")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("erreur lors de la création du dossier : %v", err)
	}

	// Écriture du fichier chiffré
	outputPath := filepath.Join(outputDir, c.GetNom())
	if err := os.WriteFile(outputPath, contenuChiffre, 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier : %v", err)
	}

	return nil
}

// GetNom retourne le nom du fichier avec l'extension de chiffrement
func (c *ChiffrementDecorator) GetNom() string {
	return c.fichier.GetNom() + ".enc"
}

// GetContenu retourne le contenu chiffré
func (c *ChiffrementDecorator) GetContenu() string {
	contenu := []byte(c.fichier.GetContenu())
	contenuChiffre, err := pkg.Encrypt(contenu)
	if err != nil {
		return fmt.Sprintf("Erreur de chiffrement: %v", err)
	}
	return string(contenuChiffre)
}
