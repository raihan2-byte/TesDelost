package pounds

import "delos/farm"

type PoundsInput struct {
	Name         string `json:"name_pounds" binding:"required"`
	FarmCategory int    `json:"farm_category" binding:"required"`
	Farm         farm.Farm `json:"Farm"`
}

type GetIdPoundsInput struct {
	ID int `uri:"id" binding:"required"`
}

type UpdatePoundsInput struct {
	Name   string `json:"name_pounds" binding:"required"`
	FarmCategory int    `json:"farm_category" binding:"required"`
	Farm         farm.Farm `json:"Farm"`
}