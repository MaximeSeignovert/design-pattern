package strategies

import (
	"fmt"
)

// LissageStrategy implémente l'affichage avec lissage des températures
type LissageStrategy struct {
	historique []float64
	tailleMax  int
}

// NewLissageStrategy crée une nouvelle instance de LissageStrategy
func NewLissageStrategy(tailleMax int) *LissageStrategy {
	return &LissageStrategy{
		historique: make([]float64, 0),
		tailleMax:  tailleMax,
	}
}

// StockerValeur ajoute une valeur à l'historique
func (l *LissageStrategy) StockerValeur(valeur float64) {
	l.historique = append(l.historique, valeur)
	if len(l.historique) > l.tailleMax {
		l.historique = l.historique[1:]
	}
}

// calculerMoyenne calcule la moyenne des valeurs de l'historique
func (l *LissageStrategy) calculerMoyenne() float64 {
	if len(l.historique) == 0 {
		return 0
	}

	var somme float64
	for _, val := range l.historique {
		somme += val
	}
	return somme / float64(len(l.historique))
}

// AfficherTemperature affiche la température lissée
func (l *LissageStrategy) AfficherTemperature(valeur float64) string {
	moyenne := l.calculerMoyenne()
	return fmt.Sprintf("%.1f°C (lissé sur %d mesures, moyenne: %.1f°C)",
		valeur,
		len(l.historique),
		moyenne)
}
