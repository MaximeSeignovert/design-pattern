package menu

import (
	"burger-system/builder"
	"fmt"
)

// TypeBoisson représente les types de boissons disponibles
type TypeBoisson string

const (
	BoissonSoda TypeBoisson = "soda"
	BoissonEau  TypeBoisson = "eau"
	BoissonJus  TypeBoisson = "jus"
)

// TailleFrites représente les tailles de frites disponibles
type TailleFrites string

const (
	FritesPetite  TailleFrites = "petite"
	FritesMoyenne TailleFrites = "moyenne"
	FritesGrande  TailleFrites = "grande"
)

// Menu représente un menu complet (burger + frites + boisson)
type Menu struct {
	Nom     string
	Burger  *burger.Burger
	Frites  TailleFrites
	Boisson TypeBoisson
	Prix    float64
}

// String retourne une représentation textuelle du menu
func (m *Menu) String() string {
	return fmt.Sprintf("Menu %s (%.2f€):\n- %s\n- Frites %s\n- Boisson %s",
		m.Nom, m.Prix, m.Burger.String(), string(m.Frites), string(m.Boisson))
}

// MenuFactory est l'interface pour créer différents types de menus
type MenuFactory interface {
	CreerMenu() *Menu
}

// MenuEnfantFactory crée des menus pour enfants
type MenuEnfantFactory struct{}

// CreerMenu crée un menu enfant prédéfini
func (f *MenuEnfantFactory) CreerMenu() *Menu {
	burgerBuilder := burger.NouveauBurgerBuilder()
	burgerEnfant := burgerBuilder.
		SetPain(burger.PainBlanc).
		SetViande(burger.ViandePoulet).
		AjouterAccompagnement(burger.Fromage).
		AjouterAccompagnement(burger.Tomate).
		FixerPrix(4.50).
		Build()

	return &Menu{
		Nom:     "Enfant",
		Burger:  burgerEnfant,
		Frites:  FritesPetite,
		Boisson: BoissonJus,
		Prix:    7.90,
	}
}

// MenuStandardFactory crée des menus standards
type MenuStandardFactory struct{}

// CreerMenu crée un menu standard prédéfini
func (f *MenuStandardFactory) CreerMenu() *Menu {
	burgerBuilder := burger.NouveauBurgerBuilder()

	// Création du burger standard
	burgerStandard := burgerBuilder.
		SetPain(burger.PainBlanc).
		SetViande(burger.ViandeBoeuf).
		AjouterAccompagnement(burger.Fromage).
		AjouterAccompagnement(burger.Tomate).
		AjouterAccompagnement(burger.Salade).
		AjouterAccompagnement(burger.Sauce).
		FixerPrix(6.50).
		Build()

	return &Menu{
		Nom:     "Standard",
		Burger:  burgerStandard,
		Frites:  FritesMoyenne,
		Boisson: BoissonSoda,
		Prix:    10.90,
	}
}

// MenuXLFactory crée des menus XL
type MenuXLFactory struct{}

// CreerMenu crée un menu XL prédéfini
func (f *MenuXLFactory) CreerMenu() *Menu {
	burgerBuilder := burger.NouveauBurgerBuilder()

	// Création du burger XL
	burgerXL := burgerBuilder.
		SetPain(burger.PainComplet).
		SetViande(burger.ViandeBoeuf).
		AjouterAccompagnement(burger.Fromage).
		AjouterAccompagnement(burger.Tomate).
		AjouterAccompagnement(burger.Salade).
		AjouterAccompagnement(burger.Sauce).
		AjouterAccompagnement(burger.Oignon).
		AjouterAccompagnement(burger.Bacon).
		FixerPrix(8.50).
		Build()

	return &Menu{
		Nom:     "XL",
		Burger:  burgerXL,
		Frites:  FritesGrande,
		Boisson: BoissonSoda,
		Prix:    13.90,
	}
}

// ObtenirFactory retourne une factory de menu en fonction du type demandé
func ObtenirFactory(typeMenu string) MenuFactory {
	switch typeMenu {
	case "enfant":
		return &MenuEnfantFactory{}
	case "standard":
		return &MenuStandardFactory{}
	case "xl":
		return &MenuXLFactory{}
	default:
		// Par défaut, on renvoie le menu standard
		return &MenuStandardFactory{}
	}
}
