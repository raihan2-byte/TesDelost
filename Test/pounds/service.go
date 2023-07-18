package pounds

import (
	"delos/farm"
	"errors"
)

type Service interface {
	CreatePounds(input PoundsInput) (Pounds, error)
	GetPounds(input int) ([]Pounds, error)
	UpdatePounds(inputID GetIdPoundsInput, input UpdatePoundsInput) (Pounds, error)
	DeletePounds(ID int) (Pounds, error)
	GetPoundsById(input int) (Pounds, error)
}

type service struct {
	repository Repository
	repositoryFarms farm.Repository
}

func NewService(repository Repository, repositoryFarms farm.Repository) *service {
	return &service{repository, repositoryFarms}
}

func (s *service) CreatePounds(input PoundsInput) (Pounds, error) {
	inputPounds := Pounds{}

	inputPounds.Name = input.Name
	inputPounds.FarmID = input.FarmCategory

	cekPounds1 := input.Name

	cekPounds, err := s.repository.FindByName(cekPounds1)
	if err != nil {
		return cekPounds, err
	}

	if cekPounds.ID != 0 {
		return cekPounds, errors.New("pounds with that name already exists")
	}

	_, err = s.repositoryFarms.FindById(input.FarmCategory)
	if err != nil {
		return Pounds{}, err
	}

	newPounds, err := s.repository.CreatePounds(inputPounds)
	if err != nil {
		return newPounds, errors.New("farm with that id doesn't exist")
	}
	return newPounds, nil
}

func (s *service) DeletePounds(ID int) (Pounds, error) {
	pounds, err := s.repository.FindById(ID)
	if err != nil {
		return pounds, errors.New("pounds with that id doesn't exist")
	}
	poundsDel, err := s.repository.Delete(pounds)

	if err != nil {
		return poundsDel, err
	}
	return poundsDel, nil
}

func (s *service) GetPoundsById(input int) (Pounds, error) {
	pounds, err := s.repository.FindById(int(input))
	if err != nil {
		return pounds, err
	} 	
	if pounds.ID == 0 {
		return pounds, errors.New("pounds with that id doesn't exist")
	}

	return pounds, nil
}

func (s *service) GetPounds(input int) ([]Pounds, error) {
	pounds, err := s.repository.FindAll()
	if err != nil {
		return pounds, err
	}

	if len(pounds) == 0 {
		return pounds, errors.New("data Not Found")
	}

	return pounds, nil
}

func (s *service) UpdatePounds(inputID GetIdPoundsInput, input UpdatePoundsInput) (Pounds, error) {
	pounds, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return pounds, err
	}

	pounds.Name = input.Name
	
	cek, err := s.repository.FindByName(pounds.Name)
	if err != nil {
		return cek, err
	}
	if cek.ID != 0 && cek.ID != pounds.ID {
		return cek, errors.New("pounds with that name already exists")
	}

	if input.FarmCategory != 0 {
		farm, err := s.repositoryFarms.FindById(input.FarmCategory)
		if err != nil {
			return Pounds{}, errors.New("farm with that id not found")
		}
		pounds.Farm = farm
	}

	updatedPounds, err := s.repository.Update(pounds)
	if err != nil {
		return updatedPounds, err
	}
	return updatedPounds, nil
}
