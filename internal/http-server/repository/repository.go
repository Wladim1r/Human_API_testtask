package repository

import (
	"fmt"

	"github.com/Wladim1r/testtask/internal/lib/errs"
	"github.com/Wladim1r/testtask/internal/models"
	"gorm.io/gorm"
)

type HumanRepository interface {
	GetInfo(size uint, name, surname, patronymic string) ([]*models.Human, error)
	Delete(id uint) error
	Patch(id uint, human *models.Human) error
	Post(req *models.Human) error
}

type humanRepository struct {
	db *gorm.DB
}

func NewHumanRepository(db *gorm.DB) HumanRepository {
	return &humanRepository{db: db}
}

func (h *humanRepository) GetInfo(
	size uint,
	name, surname, patronymic string,
) ([]*models.Human, error) {
	var humans []*models.Human

	query := h.db.Model(models.Human{}).
		Select("id, name, surname, patronymic, age, gender, nationality")

	if name != "" {
		query = query.Where("name = ?", name)
	}
	if surname != "" {
		query = query.Where("surname = ?", surname)
	}
	if patronymic != "" {
		query = query.Where("patronymic = ?", patronymic)
	}

	if size > 0 {
		query = query.Limit(int(size))
	}

	result := query.Find(&humans)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", errs.ErrDBOperation, result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, errs.ErrNotFound
	}

	return humans, nil
}

func (h *humanRepository) Delete(id uint) error {
	err := h.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&models.Human{}, id)
		if result.Error != nil {
			return fmt.Errorf("%w: could not delete row %v", errs.ErrDBOperation, result.Error)
		}
		if result.RowsAffected == 0 {
			return errs.ErrNotFound
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (h *humanRepository) Patch(id uint, human *models.Human) error {
	return h.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(models.Human{}).Where("id = ?", id).Updates(human)

		if result.Error != nil {
			return fmt.Errorf("%w: failed update %v", errs.ErrDBOperation, result.Error)
		}

		if result.RowsAffected == 0 {
			return errs.ErrNotFound
		}

		return nil
	})
}

func (h *humanRepository) Post(human *models.Human) error {
	if err := h.db.Create(human).Error; err != nil {
		return fmt.Errorf("%w: %v", errs.ErrDBOperation, err)
	}

	return nil
}
