# Station Météo - Patterns Comportementaux

## Description
Ce projet implémente un mini système de notification météo en utilisant les patterns de conception comportementaux Observer et Strategy en Go.

## Objectifs
- Implémenter le pattern Observer pour la notification des changements de température
- Implémenter le pattern Strategy pour différents algorithmes d'affichage

## Architecture

### Pattern Observer
- `StationMeteo` : Le sujet qui notifie les observateurs des changements de température
- `IObservateur` : Interface définissant le contrat pour les observateurs
- Implémentations d'observateurs :
  - `EcranConnecte` : Affiche la température sur un écran connecté
  - `ApplicationMobile` : Affiche la température sur une application mobile

### Pattern Strategy
- `IStrategyAffichage` : Interface définissant le contrat pour les différentes stratégies d'affichage
- Implémentations disponibles :
  - `CelsiusStrategy` : Affichage en degrés Celsius (°C)
  - `FahrenheitStrategy` : Affichage en degrés Fahrenheit (°F)
  - `LissageStrategy` : Affichage avec lissage des données sur plusieurs mesures

## Structure du Projet
```
.
├── interfaces/
│   ├── iobservateur.go
│   └── istrategy_affichage.go
├── models/
│   ├── station_meteo.go
│   └── station_meteo_test.go
├── observers/
│   ├── ecran_connecte.go
│   └── application_mobile.go
├── strategies/
│   ├── celsius.go
│   ├── fahrenheit.go
│   ├── lissage.go
│   └── strategies_test.go
├── main.go
├── README.md
└── ROADMAP.md
```

## Installation et Utilisation

### Prérequis
- Go 1.16 ou supérieur

### Installation
```bash
git clone [URL_DU_REPO]
cd station-meteo
go mod download
```

### Exécution
```bash
# Exécuter le projet ou lancer le .exe directement
go run main.go

```