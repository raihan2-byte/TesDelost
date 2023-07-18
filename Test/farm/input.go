package farm

type FarmInput struct {
	Name string `json:"name_farm" binding:"required"`
}

type GetIdFarmInput struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateFarmInput struct {
	Name string `json:"name_farm" binding:"required"`
	Farm Farm
}