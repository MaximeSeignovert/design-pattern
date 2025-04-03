package interfaces

// IStrategyAffichage définit l'interface pour les stratégies d'affichage
type IStrategyAffichage interface {
	// AfficherTemperature prend une température en entrée et retourne une chaîne formatée
	AfficherTemperature(temperature float64) string
	// StockerValeur permet à la stratégie de gérer son propre stockage
	StockerValeur(valeur float64)
}
