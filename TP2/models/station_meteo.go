package models

import (
	"errors"
	"station-meteo/interfaces"
)

// StationMeteo représente la station météo
type StationMeteo struct {
	observateurs []interfaces.IObservateur
	temperature  float64
}

// NewStationMeteo crée une nouvelle instance de StationMeteo
func NewStationMeteo() *StationMeteo {
	return &StationMeteo{
		observateurs: make([]interfaces.IObservateur, 0),
	}
}

// SetTemperature définit la température actuelle et notifie les observateurs
func (s *StationMeteo) SetTemperature(temperature float64) {
	s.temperature = temperature
	s.notifierObservateurs()
}

// AjouterObservateur ajoute un observateur à la liste
func (s *StationMeteo) AjouterObservateur(observateur interfaces.IObservateur) error {
	if observateur == nil {
		return errors.New("l'observateur ne peut pas être nil")
	}
	s.observateurs = append(s.observateurs, observateur)
	return nil
}

// SupprimerObservateur supprime un observateur de la liste
func (s *StationMeteo) SupprimerObservateur(observateur interfaces.IObservateur) error {
	if observateur == nil {
		return errors.New("l'observateur ne peut pas être nil")
	}

	for i, obs := range s.observateurs {
		if obs == observateur {
			s.observateurs = append(s.observateurs[:i], s.observateurs[i+1:]...)
			return nil
		}
	}
	return errors.New("observateur non trouvé")
}

// notifierObservateurs notifie tous les observateurs enregistrés
func (s *StationMeteo) notifierObservateurs() {
	for _, observateur := range s.observateurs {
		observateur.MettreAJour(s.temperature)
	}
}
