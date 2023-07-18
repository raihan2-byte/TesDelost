package pounds

import (
	"time"
)

type FormatPounds struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name_pounds"`
	FarmID int `json:"farm_category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func FormatterPounds(pounds Pounds) FormatPounds {
	newPounds := FormatPounds{}
	newPounds.ID = pounds.ID
	newPounds.Name = pounds.Name
	newPounds.FarmID = pounds.FarmID
	newPounds.CreatedAt = pounds.CreatedAt
	newPounds.UpdatedAt = pounds.UpdatedAt

	return newPounds
}

type FormatGetPounds struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name_pounds"`
	FarmID int `json:"farm_category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Farms FormatFarms `json:"farm"`
}

type FormatFarms struct {
	ID uint `json:"id"`
	Name string `json:"name_farm"`
}

func FormatterGetPound(pounds1 Pounds) FormatGetPounds {
	newGetPondsFormat := FormatGetPounds{}
	newGetPondsFormat.ID = pounds1.ID
	newGetPondsFormat.Name = pounds1.Name
	newGetPondsFormat.FarmID = pounds1.FarmID
	newGetPondsFormat.CreatedAt = pounds1.CreatedAt
	newGetPondsFormat.UpdatedAt = pounds1.UpdatedAt

	farm := pounds1.Farm

	farmFormatter := FormatFarms{}
	farmFormatter.ID = farm.ID
	farmFormatter.Name = farm.Name

	newGetPondsFormat.Farms = farmFormatter

	return newGetPondsFormat
}

func FormatterGetPounds(pounds []Pounds) []FormatGetPounds {
	newPoundsGetFormatter := []FormatGetPounds{}

	for _, newPounds := range pounds {
		newPoundsFormatter := FormatterGetPound(newPounds)
		newPoundsGetFormatter = append(newPoundsGetFormatter, newPoundsFormatter)
	}

	return newPoundsGetFormatter
}