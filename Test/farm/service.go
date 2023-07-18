package farm

import "errors"

type Service interface {
	CreateFarm(input FarmInput) (Farm, error)
	GetFarms(input int) ([]Farm, error)
	UpdateFarms(inputID GetIdFarmInput, input UpdateFarmInput) (Farm, error)
	DeleteFarm(ID int) (Farm, error)
	// FindDeleteFarm(ID int) (Farm, error)
	GetFarmById(input int) (Farm, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateFarm(input FarmInput) (Farm, error) {
	inputFarm := Farm{}

	cekFarm1 := input.Name

	inputFarm.Name = input.Name

	cekFarm, err := s.repository.FindByName(cekFarm1)
	if err != nil {
		return cekFarm, err
	}

	if cekFarm.ID != 0 {
		return cekFarm, errors.New("farm with that name already exists")
	}

	newFarm, err := s.repository.CreateFarm(inputFarm)
	if err != nil {
		return newFarm, err
	}
	return newFarm, nil
}

func (s *service) DeleteFarm(ID int) (Farm, error) {
	farm, err := s.repository.FindById(ID)
	if err != nil {
		return farm, err
	}
	farmDel, err := s.repository.Delete(farm)

	if err != nil {
		return farmDel, err
	}
	return farmDel, nil
}

func (s *service) GetFarmById(input int) (Farm, error) {
	farm, err := s.repository.FindById(int(input))

	if err != nil {
		return farm, err
	}

	if farm.ID == 0 {
		return farm, errors.New("farm not found with that id")
	}

	return farm, nil
}

func (s *service) GetFarms(input int) ([]Farm, error) {

	farm, err := s.repository.FindAll()
	if err != nil {
		return farm, err
	}

	if len(farm) == 0 {
		return nil, errors.New("farms not found")
	}

	return farm, nil
}

func (s *service) UpdateFarms(inputID GetIdFarmInput, input UpdateFarmInput) (Farm, error) {
	farm, err := s.repository.FindById((inputID.ID))
	if err != nil {
		return farm, err
	}

	farm.Name = input.Name
	cek, err := s.repository.FindByName(farm.Name)
	if err != nil {
			return cek, err
	}
	if cek.Name == farm.Name {
		return cek, errors.New("farm with that name already exists")
	}

	updatedFarm, err := s.repository.Update(farm)
	if err != nil {
		return updatedFarm, err
	}
	return updatedFarm, nil
}
