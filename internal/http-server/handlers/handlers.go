package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Wladim1r/testtask/internal/http-server/service"
	"github.com/Wladim1r/testtask/internal/lib/errs"
	"github.com/Wladim1r/testtask/internal/lib/sl"
	"github.com/Wladim1r/testtask/internal/models"
	"github.com/Wladim1r/testtask/utils"
	"github.com/gin-gonic/gin"
)

type HumanHandler struct {
	serv service.HumanService
}

func NewHumanHandler(serv service.HumanService) HumanHandler {
	return HumanHandler{serv: serv}
}

// @Summary Get users information
// @Description Get information about users with filtering options
// @Tags Users
// @ID get-users-info
// @Accept json
// @Produce json
// @Param size query int false "Limit number of records" minimum(1) example(10)
// @Param name query string false "Filter by name" example("Ivan")
// @Param surname query string false "Filter by surname" example("Ivanov")
// @Param patronymic query string false "Filter by patronymic" example("Ivanovich")
// @Success 200 {array} models.Human "Slice human structs with all fields"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 404 {object} models.ErrorResponse "Record not found"
// @Failure 500 {object} models.ErrorResponse "Database error"
// @Router /api [get]
func (h *HumanHandler) GetInfo(c *gin.Context) {
	size := c.Query("size")
	name := c.Query("name")
	surname := c.Query("surname")
	patronymic := c.Query("patronymic")

	slog.Info("GetInfo request",
		"size: ", size,
		"name: ", name,
		"surname: ", surname,
		"patronymic: ", patronymic,
	)

	humans, err := h.serv.GetInfo(size, name, surname, patronymic)
	if err != nil {
		slog.Error("HumanHandler.GetInfo", sl.Err(err))

		switch {
		case errors.Is(err, errs.ErrInvalidSize):
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Invalid size value",
			})
		case errors.Is(err, errs.ErrNotFound):
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "Could not found record",
			})
		case errors.Is(err, errs.ErrDBOperation):
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Database error",
			})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Internal Server Error",
			})
		}

		return
	}

	slog.Debug("GetInfo response", "count", len(humans))

	c.JSON(http.StatusOK, gin.H{
		"humans": humans,
	})
}

// @Summary Delete user information
// @Description Permanently removes user information by ID
// @Tags Users
// @ID delete-user
// @Accept json
// @Produce json
// @Param id path int true "ID of the user to delete" minimum(1) example(3)
// @Success 200 {object} models.SuccessResponse "User deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 404 {object} models.ErrorResponse "Record not found"
// @Failure 500 {object} models.ErrorResponse "Database or Internal Server error"
// @Router /api/{id} [delete]
func (h *HumanHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	slog.Info("Delete request", "id", id)

	err := h.serv.Delete(id)
	if err != nil {
		slog.Error("HumanHandler.Delete", sl.Err(err))

		switch {
		case errors.Is(err, errs.ErrInvalidID):
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Invalid ID format",
			})
		case errors.Is(err, errs.ErrNotFound):
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "Could not found record",
			})
		case errors.Is(err, errs.ErrDBOperation):
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Database error",
			})
		default:
			slog.Error("Unexpected error in HumanHandler.Delete")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Internal server error",
			})
		}

		return
	}

	slog.Info("Delete success", "id", id)
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Info about human was successfully deleted",
	})
}

// @Summary Update user information
// @Description Change fields that were transmitted
// @Tags Users
// @ID change-user
// @Accept json
// @Produce json
// @Param id path int true "ID of the user to change" minimum(1) example(13)
// @Param request body models.Human true "New data for change existing data"
// @Success 200 {object} models.SuccessResponse "User changed successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 404 {object} models.ErrorResponse "Record not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/{id} [patch]
func (h *HumanHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	slog.Info("Patch request", "id", id)

	var human models.Human

	if err := c.ShouldBindJSON(&human); err != nil {
		slog.Error("Patch failed to parse JSON", sl.Err(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Could not parse request",
		})
		return
	}

	slog.Debug("Patch data", "id", id, "human", human)

	if err := h.serv.Patch(id, &human); err != nil {
		slog.Error("HumanHandler.Patch", sl.Err(err))

		switch {
		case errors.Is(err, errs.ErrInvalidID):
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: "Invalid ID format",
			})
		case errors.Is(err, errs.ErrNotFound):
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error: "Could not found record",
			})
		case errors.Is(err, errs.ErrDBOperation):
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Database error",
			})
		default:
			slog.Error("Unexpected error in HumanHandler.Patch")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: "Internal server error",
			})
		}

		return
	}

	slog.Info("Patch success", "id", id)
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Info about human was successfully updated",
	})
}

// @Summary Create new user
// @Description Create user full name and automatically add age, gender and nationality
// @Tags Users
// @ID post-user
// @Accept json
// @Produce json
// @Param request body models.PostRequest true "User data to create"
// @Success 201 {object} models.SuccessResponse "User created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error (API failure or database error)"
// @Router /api [post]
func (h *HumanHandler) Post(c *gin.Context) {
	var req models.PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Post failed", sl.Err(err))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: fmt.Sprintf("could not parse request %v", err),
		})
		return
	}

	slog.Info("Post request",
		"name", req.Name,
		"surname", req.Surname,
		"patronymic", req.Patronymic,
	)

	var (
		urlAge         = "https://api.agify.io/?name=" + req.Name
		urlGender      = "https://api.genderize.io/?name=" + req.Name
		urlNationality = "https://api.nationalize.io/?name=" + req.Name
	)

	slog.Debug("Extential API calls",
		"age_url", urlAge,
		"gender_url", urlGender,
		"nationality_url", urlNationality,
	)

	bodyAge, err := utils.ParseResponse(urlAge)
	if err != nil {
		slog.Error("age API failed", "url", urlAge, sl.Err(err))
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	bodyGender, err := utils.ParseResponse(urlGender)
	if err != nil {
		slog.Error("gender API failed", "url", urlGender, sl.Err(err))
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	bodyNationality, err := utils.ParseResponseNationalize(urlNationality)
	if err != nil {
		slog.Error("nationality API failed", "url", urlNationality, sl.Err(err))
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
	}

	slog.Debug("External API responses",
		"age", bodyAge,
		"gender", bodyGender,
		"nationality", bodyNationality,
	)

	err = h.serv.Post(req, bodyAge, bodyGender, bodyNationality)
	if err != nil {
		slog.Error("Post failed", sl.Err(err))

		switch {
		case errors.Is(err, errs.ErrInvalidParam):
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error: err.Error(),
			})
		case errors.Is(err, errs.ErrDBOperation):
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: err.Error(),
			})

		}

		return
	}

	slog.Info("Post success", "name", req.Name)
	c.JSON(http.StatusCreated, models.SuccessResponse{
		Message: "Info about human was successfully added",
	})
}
