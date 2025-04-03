package interfaces

// IObservateur définit l'interface pour les observateurs de la station météo
type IObservateur interface {
	// MettreAJour est appelée lorsque la station météo a une nouvelle valeur
	MettreAJour(valeur float64)
}
