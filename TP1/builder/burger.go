package burger

// TypePain représente les types de pain disponibles
type TypePain string

const (
	PainBlanc   TypePain = "blanc"
	PainComplet TypePain = "complet"
)

// TypeViande représente les types de viande disponibles
type TypeViande string

const (
	ViandeBoeuf      TypeViande = "boeuf"
	ViandePoulet     TypeViande = "poulet"
	ViandeVegetarien TypeViande = "végétarien"
)

// Accompagnement représente un accompagnement pour le burger
type Accompagnement string

const (
	Fromage Accompagnement = "fromage"
	Tomate  Accompagnement = "tomate"
	Salade  Accompagnement = "salade"
	Sauce   Accompagnement = "sauce"
	Oignon  Accompagnement = "oignon"
	Bacon   Accompagnement = "bacon"
)

// Burger représente un burger personnalisé
type Burger struct {
	Pain            TypePain
	Viande          TypeViande
	Accompagnements []Accompagnement
	Prix            float64
}

// String retourne une représentation textuelle du burger
func (b *Burger) String() string {
	result := "Burger avec pain " + string(b.Pain) + ", viande " + string(b.Viande)

	if len(b.Accompagnements) > 0 {
		result += " et accompagnements: "
		for i, acc := range b.Accompagnements {
			if i > 0 {
				result += ", "
			}
			result += string(acc)
		}
	}

	return result
}

// BurgerBuilder est l'interface pour construire un burger
type BurgerBuilder interface {
	SetPain(pain TypePain) BurgerBuilder
	SetViande(viande TypeViande) BurgerBuilder
	AjouterAccompagnement(accompagnement Accompagnement) BurgerBuilder
	FixerPrix(prix float64) BurgerBuilder
	Build() *Burger
}

// ConcreteBuilder implémente l'interface BurgerBuilder
type ConcreteBuilder struct {
	burger *Burger
}

// NouveauBurgerBuilder crée un nouveau builder pour construire un burger
func NouveauBurgerBuilder() BurgerBuilder {
	return &ConcreteBuilder{
		burger: &Burger{
			Accompagnements: []Accompagnement{},
			Prix:            0.0,
		},
	}
}

// SetPain définit le type de pain du burger
func (b *ConcreteBuilder) SetPain(pain TypePain) BurgerBuilder {
	b.burger.Pain = pain
	return b
}

// SetViande définit le type de viande du burger
func (b *ConcreteBuilder) SetViande(viande TypeViande) BurgerBuilder {
	b.burger.Viande = viande
	return b
}

// AjouterAccompagnement ajoute un accompagnement au burger
func (b *ConcreteBuilder) AjouterAccompagnement(accompagnement Accompagnement) BurgerBuilder {
	b.burger.Accompagnements = append(b.burger.Accompagnements, accompagnement)
	return b
}

// FixerPrix définit le prix du burger
func (b *ConcreteBuilder) FixerPrix(prix float64) BurgerBuilder {
	b.burger.Prix = prix
	return b
}

// Build construit et retourne le burger final
func (b *ConcreteBuilder) Build() *Burger {
	return b.burger
}
