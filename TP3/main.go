package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"file-modifier/decorator"
	"file-modifier/facade"
)

func listerFichiersTxt(dossier string) ([]string, error) {
	var fichiers []string
	entries, err := os.ReadDir(dossier)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".txt") {
			fichiers = append(fichiers, entry.Name())
		}
	}
	return fichiers, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	manager := facade.NewSauvegardeManager()

	for {
		fmt.Println("\n=== Système de Gestion des Fichiers ===")
		fmt.Println("1. Lister les fichiers .txt disponibles")
		fmt.Println("2. Compresser un fichier")
		fmt.Println("3. Chiffrer un fichier")
		fmt.Println("4. Compresser et chiffrer un fichier")
		fmt.Println("5. Quitter")
		fmt.Print("\nVotre choix : ")

		scanner.Scan()
		choix := scanner.Text()

		switch choix {
		case "1":
			fichiers, err := listerFichiersTxt("input_files")
			if err != nil {
				fmt.Printf("Erreur lors de la lecture du dossier : %v\n", err)
				continue
			}
			if len(fichiers) == 0 {
				fmt.Println("Aucun fichier .txt trouvé dans le dossier input_files")
				continue
			}
			fmt.Println("\nFichiers disponibles :")
			for i, f := range fichiers {
				fmt.Printf("%d. %s\n", i+1, f)
			}

		case "2", "3", "4":
			fichiers, err := listerFichiersTxt("input_files")
			if err != nil {
				fmt.Printf("Erreur lors de la lecture du dossier : %v\n", err)
				continue
			}
			if len(fichiers) == 0 {
				fmt.Println("Aucun fichier .txt trouvé dans le dossier input_files")
				continue
			}

			fmt.Println("\nFichiers disponibles :")
			for i, f := range fichiers {
				fmt.Printf("%d. %s\n", i+1, f)
			}

			fmt.Print("\nChoisissez un fichier (numéro) : ")
			scanner.Scan()
			index := 0
			_, err = fmt.Sscanf(scanner.Text(), "%d", &index)
			if err != nil || index < 1 || index > len(fichiers) {
				fmt.Println("Choix invalide")
				continue
			}

			// Lecture du fichier sélectionné
			nomFichier := fichiers[index-1]
			contenu, err := os.ReadFile(filepath.Join("input_files", nomFichier))
			if err != nil {
				fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
				continue
			}

			// Création du fichier de base
			fichier := decorator.NewFichierTexte(nomFichier, string(contenu))

			// Application des décorateurs selon le choix
			var fichierDecore decorator.IFichier = fichier
			switch choix {
			case "2":
				fichierDecore = decorator.NewCompressionDecorator(fichier)
			case "3":
				fichierDecore = decorator.NewChiffrementDecorator(fichier)
			case "4":
				fichierCompresse := decorator.NewCompressionDecorator(fichier)
				fichierDecore = decorator.NewChiffrementDecorator(fichierCompresse)
			}

			// Sauvegarde du fichier
			err = manager.SauvegarderFichier(fichierDecore)
			if err != nil {
				fmt.Printf("Erreur lors du traitement : %v\n", err)
			} else {
				fmt.Println("Opération réussie !")
			}

		case "5":
			fmt.Println("Au revoir !")
			return

		default:
			fmt.Println("Choix invalide")
		}
	}
}
