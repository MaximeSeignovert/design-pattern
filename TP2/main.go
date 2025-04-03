package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"station-meteo/interfaces"
	"station-meteo/models"
	"station-meteo/observers"
	"station-meteo/strategies"
	"strconv"
	"strings"
)

func main() {
	// Création de la station météo
	station := models.NewStationMeteo()

	// Création des observateurs
	ecranSalon := observers.NewEcranConnecte("Salon")
	ecranCuisine := observers.NewEcranConnecte("Cuisine")
	appMobile := observers.NewApplicationMobile("iPhone")

	// Création des stratégies d'affichage
	strategieCelsius := strategies.NewCelsiusStrategy()
	strategieFahrenheit := strategies.NewFahrenheitStrategy()
	strategieLissage := strategies.NewLissageStrategy(3) // Lissage sur 3 valeurs

	// Configuration des observateurs avec leurs stratégies
	ecranSalon.SetStrategieAffichage(strategieCelsius)
	ecranCuisine.SetStrategieAffichage(strategieFahrenheit)
	appMobile.SetStrategieAffichage(strategieLissage)

	// Enregistrement des observateurs avec gestion des erreurs
	if err := station.AjouterObservateur(ecranSalon); err != nil {
		log.Printf("Erreur lors de l'ajout de l'écran salon: %v", err)
	}
	if err := station.AjouterObservateur(ecranCuisine); err != nil {
		log.Printf("Erreur lors de l'ajout de l'écran cuisine: %v", err)
	}
	if err := station.AjouterObservateur(appMobile); err != nil {
		log.Printf("Erreur lors de l'ajout de l'application mobile: %v", err)
	}

	fmt.Println("=== Station Météo Interactive ===")
	fmt.Println("Commandes disponibles :")
	fmt.Println("- t <température> : Définir une nouvelle température (ex: t 20)")
	fmt.Println("- s <observateur> <stratégie> : Changer la stratégie d'un observateur (ex: s Salon fahrenheit)")
	fmt.Println("- r <observateur> : Supprimer un observateur (ex: r Cuisine)")
	fmt.Println("- q : Quitter")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}
		commande := scanner.Text()
		args := strings.Fields(commande)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "t":
			if len(args) != 2 {
				fmt.Println("Usage: t <température>")
				continue
			}
			temp, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				fmt.Println("Erreur: La température doit être un nombre")
				continue
			}
			station.SetTemperature(temp)

		case "s":
			if len(args) != 3 {
				fmt.Println("Usage: s <observateur> <stratégie>")
				continue
			}
			observateur := args[1]
			strategie := args[2]

			var obs interface {
				SetStrategieAffichage(interfaces.IStrategyAffichage)
			}

			switch observateur {
			case "Salon":
				obs = ecranSalon
			case "Cuisine":
				obs = ecranCuisine
			case "iPhone":
				obs = appMobile
			default:
				fmt.Println("Observateur non trouvé")
				continue
			}

			var strat interfaces.IStrategyAffichage
			switch strategie {
			case "celsius":
				strat = strategieCelsius
			case "fahrenheit":
				strat = strategieFahrenheit
			case "lissage":
				strat = strategieLissage
			default:
				fmt.Println("Stratégie non trouvée")
				continue
			}

			obs.SetStrategieAffichage(strat)
			fmt.Printf("Stratégie %s appliquée à %s\n", strategie, observateur)

		case "r":
			if len(args) != 2 {
				fmt.Println("Usage: r <observateur>")
				continue
			}
			observateur := args[1]
			var obs interfaces.IObservateur

			switch observateur {
			case "Salon":
				obs = ecranSalon
			case "Cuisine":
				obs = ecranCuisine
			case "iPhone":
				obs = appMobile
			default:
				fmt.Println("Observateur non trouvé")
				continue
			}

			if err := station.SupprimerObservateur(obs); err != nil {
				fmt.Printf("Erreur: %v\n", err)
				continue
			}
			fmt.Printf("Observateur %s supprimé\n", observateur)

		case "q":
			fmt.Println("Au revoir!")
			return

		default:
			fmt.Println("Commande non reconnue")
		}
	}
}
