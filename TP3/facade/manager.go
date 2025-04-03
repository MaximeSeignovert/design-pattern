package facade

import (
	"file-modifier/decorator"
)

// SauvegardeManager implémente le pattern Facade pour simplifier les opérations de sauvegarde
type SauvegardeManager struct {
	// TODO: Ajouter des champs pour la configuration si nécessaire
}

// NewSauvegardeManager crée une nouvelle instance de SauvegardeManager
func NewSauvegardeManager() *SauvegardeManager {
	return &SauvegardeManager{}
}

// SauvegarderFichier est la méthode principale qui simplifie le processus de sauvegarde
func (s *SauvegardeManager) SauvegarderFichier(fichier decorator.IFichier) error {
	// La facade encapsule la complexité des opérations
	// Ici, nous utilisons simplement la méthode Enregistrer du fichier
	// qui a déjà été décorée avec les fonctionnalités nécessaires
	return fichier.Enregistrer()
}
