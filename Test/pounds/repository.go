package pounds

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Pounds, error)
	FindById(ID int) (Pounds, error)
	FindByName(Name string) (Pounds, error)
	Update(pounds Pounds) (Pounds, error)
	Delete(pounds Pounds) (Pounds, error)
	CreatePounds(pounds Pounds) (Pounds, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByName(Name string) (Pounds, error) {
	var pounds Pounds

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("name = ?", Name).Find(&pounds).Error
	if err != nil {
		return pounds, err
	}
	return pounds, nil
}

func (r *repository) FindAll() ([]Pounds, error) {
	var pounds []Pounds

	err := r.db.Preload("Farm").Find(&pounds).Error
	if err != nil {
		return pounds, err
	}
	return pounds, nil
}

func (r *repository) FindById(ID int) (Pounds, error) {
	var pounds Pounds

	err := r.db.Preload("Farm").Where("id = ?", ID).First(&pounds).Error

	if err != nil {
		return pounds, err
	}
	return pounds, nil
}

func (r *repository) Update(pounds Pounds) (Pounds, error) {
	err := r.db.Save(&pounds).Error
	if err != nil {
		return pounds, err
	}

	return pounds, nil
}

func (r *repository) Delete(pounds Pounds) (Pounds, error) {
	err := r.db.Delete(&pounds).Error
	if err != nil {
		return pounds, err
	}

	return pounds, nil
}

func (r *repository) CreatePounds(pounds Pounds) (Pounds, error) {
	err := r.db.Create(&pounds).Error
	if err != nil {
		return pounds, err
	}

	return pounds, nil
}

// func (r *repository) FindByUserId(userID int) ([]Farm, error) {
// 	var farm []Farm

// 	err := r.db.Joins("User", r.db.Where(&user.User{ID: userID})).Find(&farm).Error
// 	fmt.Println("eror", farm)
// 	if err != nil {
// 		return farm, err
// 	}
// 	return farm, nil
// }
