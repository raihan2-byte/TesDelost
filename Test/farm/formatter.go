package farm

import "time"

type GetAllFarm struct {
	ID uint `json:"id"`
	Name      string `json:"name_farm"`
	CreatedAt time.Time`json:"created_at"`
	UpdatedAt time.Time`json:"updated_at"`
}

func FormatterGetAllFarm (farm Farm) GetAllFarm {
	newFarm := GetAllFarm{}
	newFarm.ID = farm.ID
	newFarm.Name = farm.Name
	newFarm.CreatedAt = farm.CreatedAt
	newFarm.UpdatedAt = farm.UpdatedAt

	return newFarm
}

func FormatterGetFarms(farms []Farm) []GetAllFarm {
	newFarmGetFormatter := []GetAllFarm{}

	for _, newFarm := range farms {
		newFarmFormatter := FormatterGetAllFarm(newFarm)
		newFarmGetFormatter = append(newFarmGetFormatter, newFarmFormatter)
	}

	return newFarmGetFormatter
}