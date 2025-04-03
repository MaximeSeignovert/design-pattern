package main

import (
	"bufio"
	"burger-system/builder"
	"burger-system/factory"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("===== Système de Commande de Burgers Personnalisés =====")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nQue souhaitez-vous faire ?")
		fmt.Println("1. Commander un menu prédéfini")
		fmt.Println("2. Créer un burger personnalisé")
		fmt.Println("3. Quitter")
		fmt.Print("Votre choix : ")

		scanner.Scan()
		choix := scanner.Text()

		switch choix {
		case "1":
			commanderMenu(scanner)
		case "2":
			creerBurgerPersonnalise(scanner)
		case "3":
			fmt.Println("Merci d'avoir utilisé notre système de commande. À bientôt !")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}

func commanderMenu(scanner *bufio.Scanner) {
	fmt.Println("\n=== Commander un menu ===")
	fmt.Println("Quel type de menu souhaitez-vous ?")
	fmt.Println("1. Menu Enfant")
	fmt.Println("2. Menu Standard")
	fmt.Println("3. Menu XL")
	fmt.Print("Votre choix : ")

	scanner.Scan()
	choix := scanner.Text()

	var typeMenu string

	switch choix {
	case "1":
		typeMenu = "enfant"
	case "2":
		typeMenu = "standard"
	case "3":
		typeMenu = "xl"
	default:
		fmt.Println("Option invalide, menu standard sélectionné par défaut.")
		typeMenu = "standard"
	}

	menuFactory := menu.ObtenirFactory(typeMenu)
	menuCommande := menuFactory.CreerMenu()

	fmt.Println("\nVous avez commandé :")
	fmt.Println(menuCommande)
}

func creerBurgerPersonnalise(scanner *bufio.Scanner) {
	fmt.Println("\n=== Créer un burger personnalisé ===")

	burgerBuilder := burger.NouveauBurgerBuilder()

	// Choisir le pain
	fmt.Println("Choisissez le type de pain :")
	fmt.Println("1. Pain blanc")
	fmt.Println("2. Pain complet")
	fmt.Print("Votre choix : ")

	scanner.Scan()
	choixPain := scanner.Text()

	switch choixPain {
	case "1":
		burgerBuilder.SetPain(burger.PainBlanc)
	case "2":
		burgerBuilder.SetPain(burger.PainComplet)
	default:
		fmt.Println("Option invalide, pain blanc sélectionné par défaut.")
		burgerBuilder.SetPain(burger.PainBlanc)
	}

	// Choisir la viande
	fmt.Println("\nChoisissez le type de viande :")
	fmt.Println("1. Bœuf")
	fmt.Println("2. Poulet")
	fmt.Println("3. Végétarien")
	fmt.Print("Votre choix : ")

	scanner.Scan()
	choixViande := scanner.Text()

	switch choixViande {
	case "1":
		burgerBuilder.SetViande(burger.ViandeBoeuf)
	case "2":
		burgerBuilder.SetViande(burger.ViandePoulet)
	case "3":
		burgerBuilder.SetViande(burger.ViandeVegetarien)
	default:
		fmt.Println("Option invalide, bœuf sélectionné par défaut.")
		burgerBuilder.SetViande(burger.ViandeBoeuf)
	}

	// Ajouter les accompagnements
	fmt.Println("\nAjoutez des accompagnements :")
	fmt.Println("1. Fromage")
	fmt.Println("2. Tomate")
	fmt.Println("3. Salade")
	fmt.Println("4. Sauce")
	fmt.Println("5. Oignon")
	fmt.Println("6. Bacon")
	fmt.Println("7. Terminer la sélection d'accompagnements")

	accompagnements := []burger.Accompagnement{
		burger.Fromage,
		burger.Tomate,
		burger.Salade,
		burger.Sauce,
		burger.Oignon,
		burger.Bacon,
	}

	for {
		fmt.Print("Votre choix (1-7) : ")
		scanner.Scan()
		choixAcc := scanner.Text()

		if choixAcc == "7" {
			break
		}

		choixAccNum, err := strconv.Atoi(choixAcc)
		if err != nil || choixAccNum < 1 || choixAccNum > 6 {
			fmt.Println("Option invalide, veuillez réessayer.")
			continue
		}

		burgerBuilder.AjouterAccompagnement(accompagnements[choixAccNum-1])
		fmt.Printf("Ajout de %s\n", accompagnements[choixAccNum-1])
	}

	// Fixer le prix du burger (ici, un prix fictif basé sur le nombre d'accompagnements)
	prix := 5.0 // Prix de base
	burger := burgerBuilder.Build()
	prix += float64(len(burger.Accompagnements)) * 0.5 // +0.50€ par accompagnement

	// Construire le burger final
	burger = burgerBuilder.FixerPrix(prix).Build()

	fmt.Println("\nVotre burger personnalisé :")
	fmt.Println(burger)
	fmt.Printf("Prix : %.2f€\n", burger.Prix)

	// Proposer d'ajouter des frites et une boisson pour en faire un menu
	fmt.Println("\nSouhaitez-vous ajouter des frites et une boisson pour en faire un menu ? (oui/non)")
	scanner.Scan()
	choixMenu := strings.ToLower(scanner.Text())

	if choixMenu == "oui" || choixMenu == "o" {
		fmt.Println("Choisissez la taille des frites :")
		fmt.Println("1. Petite")
		fmt.Println("2. Moyenne")
		fmt.Println("3. Grande")

		scanner.Scan()
		choixFrites := scanner.Text()

		var frites menu.TailleFrites
		switch choixFrites {
		case "1":
			frites = menu.FritesPetite
		case "2":
			frites = menu.FritesMoyenne
		case "3":
			frites = menu.FritesGrande
		default:
			fmt.Println("Option invalide, taille moyenne sélectionnée par défaut.")
			frites = menu.FritesMoyenne
		}

		fmt.Println("\nChoisissez votre boisson :")
		fmt.Println("1. Soda")
		fmt.Println("2. Eau")
		fmt.Println("3. Jus")

		scanner.Scan()
		choixBoisson := scanner.Text()

		var boisson menu.TypeBoisson
		switch choixBoisson {
		case "1":
			boisson = menu.BoissonSoda
		case "2":
			boisson = menu.BoissonEau
		case "3":
			boisson = menu.BoissonJus
		default:
			fmt.Println("Option invalide, soda sélectionné par défaut.")
			boisson = menu.BoissonSoda
		}

		// Prix du menu = prix du burger + supplément pour les frites et la boisson
		var supplementMenu float64
		switch frites {
		case menu.FritesPetite:
			supplementMenu += 2.0
		case menu.FritesMoyenne:
			supplementMenu += 2.5
		case menu.FritesGrande:
			supplementMenu += 3.0
		}

		switch boisson {
		case menu.BoissonSoda, menu.BoissonJus:
			supplementMenu += 2.0
		case menu.BoissonEau:
			supplementMenu += 1.5
		}

		menuPersonnalise := &menu.Menu{
			Nom:     "Personnalisé",
			Burger:  burger,
			Frites:  frites,
			Boisson: boisson,
			Prix:    burger.Prix + supplementMenu,
		}

		fmt.Println("\nVotre menu personnalisé :")
		fmt.Println(menuPersonnalise)
	}
}
