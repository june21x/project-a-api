package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/june21x/project-a-api/internal_service/world_map_service"
	"github.com/june21x/project-a-api/repository/world_map_repository"
	"github.com/june21x/project-a-api/util"
)

type CreateAreaReq struct {
	Coordinate world_map_repository.Coordinate `json:"coordinate" binding:"required"`
}

type GetAreasQuery struct {
	Radius *int64 `form:"radius" binding:"omitempty,gte=1"`
}

type AreaController interface {
	CreateArea(c *gin.Context)
	GetAreas(c *gin.Context)
	GetArea(c *gin.Context)
}

type AreaControllerImpl struct {
	service world_map_service.AreaService
}

// @BasePath /api/v1
// Areas godoc

// @Summary Create area
// @Description Create area
// @Tags Areas
// @Accept json
// @Produce json
// @Param body body CreateAreaReq true "body"
// @Success 201 {object} world_map_repository.Area
// @Router /world-map/areas [post]
func (a AreaControllerImpl) CreateArea(c *gin.Context) {
	var reqBodyJson CreateAreaReq

	// TODO validate non-zero and integer only

	if err := c.ShouldBindJSON(&reqBodyJson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	createdArea, err := a.service.CreateArea(&reqBodyJson.Coordinate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusCreated, createdArea)
}

// @Summary Get areas
// @Description Get areas
// @Tags Areas
// @Produce json
// @Param radius query int false "radius"
// @Success 200 {array} world_map_repository.Area
// @Failure 400 {object} util.ErrorRes
// @Router /world-map/areas [get]
func (a AreaControllerImpl) GetAreas(c *gin.Context) {
	var query GetAreasQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	areas, err := a.service.GetAreas(query.Radius)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, areas)
}

// @Summary Get area
// @Description Get area
// @Tags Areas
// @Produce json
// @Param uuid path string true "uuid"
// @Success 200 {object} world_map_repository.Area
// @Router /world-map/areas/{uuid} [get]
func (a AreaControllerImpl) GetArea(c *gin.Context) {
	area, err := a.service.GetArea(c.Param(("uuid")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, area)
}

func AreaControllerInit(areaService world_map_service.AreaService) *AreaControllerImpl {
	return &AreaControllerImpl{
		service: areaService,
	}
}
