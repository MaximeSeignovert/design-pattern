package strategies

import (
	"fmt"
)

// FahrenheitStrategy implémente l'affichage en Fahrenheit
type FahrenheitStrategy struct {
	valeurCourante float64
}

// NewFahrenheitStrategy crée une nouvelle instance de FahrenheitStrategy
func NewFahrenheitStrategy() *FahrenheitStrategy {
	return &FahrenheitStrategy{}
}

// StockerValeur stocke la valeur courante
func (f *FahrenheitStrategy) StockerValeur(valeur float64) {
	f.valeurCourante = valeur
}

// AfficherTemperature affiche la température en Fahrenheit
func (f *FahrenheitStrategy) AfficherTemperature(valeur float64) string {
	fahrenheit := (f.valeurCourante * 9 / 5) + 32
	return fmt.Sprintf("%.1f°F", fahrenheit)
}
