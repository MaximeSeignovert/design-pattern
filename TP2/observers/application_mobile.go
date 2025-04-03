package observers

import (
	"fmt"
	"station-meteo/interfaces"
)

// ApplicationMobile représente une application mobile qui affiche la température
type ApplicationMobile struct {
	nom string
	strategie interfaces.IStrategyAffichage
}

// NewApplicationMobile crée une nouvelle instance d'ApplicationMobile
func NewApplicationMobile(nom string) *ApplicationMobile {
	return &ApplicationMobile{
		nom: nom,
	}
}

// SetStrategieAffichage définit la stratégie d'affichage à utiliser
func (a *ApplicationMobile) SetStrategieAffichage(strategie interfaces.IStrategyAffichage) {
	a.strategie = strategie
}

// MettreAJour met à jour l'affichage avec la nouvelle valeur
func (a *ApplicationMobile) MettreAJour(valeur float64) {
	if a.strategie == nil {
		fmt.Printf("📱 [%s] Valeur brute: %.1f\n", a.nom, valeur)
		return
	}
	a.strategie.StockerValeur(valeur)
	affichage := a.strategie.AfficherTemperature(valeur)
	fmt.Printf("📱 [%s] %s\n", a.nom, affichage)
}
