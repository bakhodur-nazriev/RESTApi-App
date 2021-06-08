package models

type Author struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	DateOfBirth string `json:"date_of_birth"`
	DateOfDeath string `json:"date_of_death"`
	Image       string `json:"imag       e"`
}
