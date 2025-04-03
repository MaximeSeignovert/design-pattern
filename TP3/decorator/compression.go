package decorator

import (
	"fmt"
	"os"
	"path/filepath"
	"file-modifier/pkg"
)

// CompressionDecorator implémente le pattern Decorator pour la compression
type CompressionDecorator struct {
	fichier IFichier
}

// NewCompressionDecorator crée une nouvelle instance de CompressionDecorator
func NewCompressionDecorator(fichier IFichier) *CompressionDecorator {
	return &CompressionDecorator{
		fichier: fichier,
	}
}

// Enregistrer implémente la méthode de l'interface IFichier avec compression
func (c *CompressionDecorator) Enregistrer() error {
	// Lecture du contenu
	contenu := []byte(c.fichier.GetContenu())

	// Compression du contenu
	contenuCompresse, err := pkg.Compress(contenu)
	if err != nil {
		return fmt.Errorf("erreur lors de la compression : %v", err)
	}

	// Création du dossier de sortie si nécessaire
	outputDir := filepath.Join("output_files", "compressed")
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("erreur lors de la création du dossier : %v", err)
	}

	// Écriture du fichier compressé
	outputPath := filepath.Join(outputDir, c.GetNom())
	if err := os.WriteFile(outputPath, contenuCompresse, 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier : %v", err)
	}

	return nil
}

// GetNom retourne le nom du fichier avec l'extension de compression
func (c *CompressionDecorator) GetNom() string {
	return c.fichier.GetNom() + ".gz"
}

// GetContenu retourne le contenu compressé
func (c *CompressionDecorator) GetContenu() string {
	contenu := []byte(c.fichier.GetContenu())
	contenuCompresse, err := pkg.Compress(contenu)
	if err != nil {
		return fmt.Sprintf("Erreur de compression: %v", err)
	}
	return string(contenuCompresse)
}
