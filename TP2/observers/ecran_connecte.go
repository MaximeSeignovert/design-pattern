package observers

import (
	"fmt"
	"station-meteo/interfaces"
)

// EcranConnecte représente un écran connecté qui affiche la température
type EcranConnecte struct {
	nom       string
	strategie interfaces.IStrategyAffichage
}

// NewEcranConnecte crée une nouvelle instance d'EcranConnecte
func NewEcranConnecte(nom string) *EcranConnecte {
	return &EcranConnecte{
		nom: nom,
	}
}

// SetStrategieAffichage définit la stratégie d'affichage à utiliser
func (e *EcranConnecte) SetStrategieAffichage(strategie interfaces.IStrategyAffichage) {
	e.strategie = strategie
}

// MettreAJour met à jour l'affichage avec la nouvelle valeur
func (e *EcranConnecte) MettreAJour(valeur float64) {
	if e.strategie == nil {
		fmt.Printf("[%s] Valeur brute: %.1f\n", e.nom, valeur)
		return
	}
	e.strategie.StockerValeur(valeur)
	affichage := e.strategie.AfficherTemperature(valeur)
	fmt.Printf("[%s] %s\n", e.nom, affichage)
}
