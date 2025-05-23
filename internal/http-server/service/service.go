package service

import (
	"errors"
	"fmt"

	"github.com/Wladim1r/testtask/internal/http-server/repository"
	"github.com/Wladim1r/testtask/internal/lib/errs"
	"github.com/Wladim1r/testtask/internal/models"
	"github.com/Wladim1r/testtask/utils"
)

type HumanService interface {
	GetInfo(size, name, surname, patronymic string) ([]*models.Human, error)
	Delete(id string) error
	Patch(id string, human *models.Human) error
	Post(
		req models.PostRequest,
		bodyAge, bodyGender map[string]interface{},
		bodyNationality models.NationalizeResponse,
	) error
}

type humanService struct {
	repo repository.HumanRepository
}

func NewHumanService(repo repository.HumanRepository) HumanService {
	return &humanService{repo: repo}
}

func (h *humanService) GetInfo(size, name, surname, patronymic string) ([]*models.Human, error) {
	sizeInt, err := utils.IsPositive(size)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrInvalidSize, err)
	}

	humans, err := h.repo.GetInfo(uint(sizeInt), name, surname, patronymic)
	if err != nil {
		return nil, fmt.Errorf("select failed: %w", err)
	}

	return humans, nil
}

func (h *humanService) Delete(id string) error {
	idInt, err := utils.IsPositive(id)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ErrInvalidID, err)
	}

	err = h.repo.Delete(uint(idInt))
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrNotFound):
			return err
		default:
			return fmt.Errorf("delete failed: %w", err)
		}
	}

	return nil
}

func (h *humanService) Patch(id string, human *models.Human) error {
	idInt, err := utils.IsPositive(id)
	if err != nil {
		return fmt.Errorf("%w: %v", errs.ErrInvalidID, err)
	}

	err = h.repo.Patch(uint(idInt), human)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	return nil
}

func (h *humanService) Post(
	req models.PostRequest,
	bodyAge, bodyGender map[string]interface{},
	bodyNationality models.NationalizeResponse,
) error {

	age, ok := bodyAge["age"].(float64)
	if !ok {
		return errs.ErrInvalidParam
	}
	gender, ok := bodyGender["gender"].(string)
	if !ok {
		return errs.ErrInvalidParam
	}

	probabilities := map[float64]string{}
	for _, country := range bodyNationality.Country {
		probabilities[country.Probability] = country.CountryID
	}

	maxKey := 0.0
	for key := range probabilities {
		if key > maxKey {
			maxKey = key
		}
	}

	nationality := probabilities[maxKey]

	human := models.Human{
		Name:        req.Name,
		Surname:     req.Surname,
		Patronymic:  req.Patronymic,
		Age:         uint(age),
		Gender:      gender,
		Nationality: nationality,
	}

	if err := h.repo.Post(&human); err != nil {
		return fmt.Errorf("post failed %w", err)
	}

	return nil
}
