package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unnagirirao/Unique/unique/pkg/rest/server/models"
	"github.com/unnagirirao/Unique/unique/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type UniqueController struct {
	uniqueService *services.UniqueService
}

func NewUniqueController() (*UniqueController, error) {
	uniqueService, err := services.NewUniqueService()
	if err != nil {
		return nil, err
	}
	return &UniqueController{
		uniqueService: uniqueService,
	}, nil
}

func (uniqueController *UniqueController) CreateUnique(context *gin.Context) {
	// validate input
	var input models.Unique
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger unique creation
	if _, err := uniqueController.uniqueService.CreateUnique(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Unique created successfully"})
}

func (uniqueController *UniqueController) UpdateUnique(context *gin.Context) {
	// validate input
	var input models.Unique
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger unique update
	if _, err := uniqueController.uniqueService.UpdateUnique(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Unique updated successfully"})
}

func (uniqueController *UniqueController) FetchUnique(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger unique fetching
	unique, err := uniqueController.uniqueService.GetUnique(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, unique)
}

func (uniqueController *UniqueController) DeleteUnique(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger unique deletion
	if err := uniqueController.uniqueService.DeleteUnique(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Unique deleted successfully",
	})
}

func (uniqueController *UniqueController) ListUniques(context *gin.Context) {
	// trigger all uniques fetching
	uniques, err := uniqueController.uniqueService.ListUniques()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, uniques)
}

func (*UniqueController) PatchUnique(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*UniqueController) OptionsUnique(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*UniqueController) HeadUnique(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
