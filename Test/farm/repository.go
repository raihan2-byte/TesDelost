package farm

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Farm, error)
	FindById(ID int) (Farm, error)
	FindByName(Name string) (Farm, error)
	// FindByFarmID(userID int) ([]Farm, error)
	Update(farm Farm) (Farm, error)
	Delete(farm Farm) (Farm, error)
	FindDelete(ID int) (Farm, error)
	CreateFarm(farm Farm) (Farm, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByName(Name string) (Farm, error) {
	var farm Farm

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("name = ?", Name).Find(&farm).Error
	if err != nil {
		return farm, err
	}
	return farm, nil
}


func (r *repository) FindAll() ([]Farm, error) {
	var farm []Farm

	err := r.db.Find(&farm).Error
	if err != nil {
		return farm, err
	}
	return farm, nil
}

func (r *repository) FindById(ID int) (Farm, error) {
	var farm Farm

	err := r.db.Where("id = ?", ID).Find(&farm).Error

	if err != nil {
		return farm, err
	}
	return farm, nil
}

func (r *repository) Update(farm Farm) (Farm, error) {
	err := r.db.Save(&farm).Error
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (r *repository) Delete(farm Farm) (Farm, error) {
	err := r.db.Delete(&farm).Error
	if err != nil {
		return farm, err
	}

	return farm, nil
}

func (r *repository) FindDelete(ID int) (Farm, error) {
	var farm Farm
	err := r.db.Delete(&farm).Error
	if err != nil {
		return farm, err
	}

	return farm, nil
}


func (r *repository) CreateFarm(farm Farm) (Farm, error) {
	err := r.db.Create(&farm).Error

	if err != nil {
		return farm, err
	}

	return farm, nil
}
