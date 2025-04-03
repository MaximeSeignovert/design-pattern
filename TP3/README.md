# TP 3 - Patterns Structurels en Go
## Facade & Decorator

### Description du Projet
Ce projet implémente un système de sauvegarde de fichiers utilisant les patterns Facade et Decorator en Go. Il permet de sauvegarder, compresser et chiffrer des fichiers de manière flexible et extensible.

### Architecture

```
...
```

### Design Patterns Implémentés

#### 1. Pattern Facade
Le pattern Facade est implémenté dans le package `internal/facade` à travers la classe `SauvegardeManager`. Cette classe simplifie l'interface pour les opérations complexes de sauvegarde.

#### 2. Pattern Decorator
Le pattern Decorator est implémenté dans le package `internal/decorator` avec :
- `IFichier` : Interface de base
- `FichierTexte` : Implémentation concrète
- `CompressionDecorator` : Ajoute la compression
- `ChiffrementDecorator` : Ajoute le chiffrement

### Installation et Utilisation

```bash
# Cloner le projet
git clone [URL_DU_PROJET]

# Exécuter le projet ou lancer le .exe directement
go run cmd/main.go
```

### Exemple d'Utilisation

```go
// Création d'un fichier texte
fichier := decorator.NewFichierTexte("mon_fichier.txt", "Contenu du fichier")

// Application des décorateurs
fichierCompresse := decorator.NewCompressionDecorator(fichier)
fichierChiffre := decorator.NewChiffrementDecorator(fichierCompresse)

// Utilisation de la facade
manager := facade.NewSauvegardeManager()
manager.SauvegarderFichier(fichierChiffre)
```

### Tests
Pour exécuter les tests :
```bash
go test ./...
```

### Documentation
La documentation détaillée de chaque composant est disponible dans les fichiers source correspondants. 