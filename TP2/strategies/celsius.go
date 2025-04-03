package strategies

import (
	"fmt"
)

// CelsiusStrategy implémente l'affichage en Celsius
type CelsiusStrategy struct {
	valeurCourante float64
}

// NewCelsiusStrategy crée une nouvelle instance de CelsiusStrategy
func NewCelsiusStrategy() *CelsiusStrategy {
	return &CelsiusStrategy{}
}

// StockerValeur stocke la valeur courante
func (c *CelsiusStrategy) StockerValeur(valeur float64) {
	c.valeurCourante = valeur
}

// AfficherTemperature affiche la température en Celsius
func (c *CelsiusStrategy) AfficherTemperature(valeur float64) string {
	return fmt.Sprintf("%.1f°C", c.valeurCourante)
}
