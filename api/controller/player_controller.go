package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/june21x/project-a-api/internal_service/player_service"
	"github.com/june21x/project-a-api/repository/player_repository"
	"github.com/june21x/project-a-api/util"
)

type RegisterPlayerReq struct {
	PlayerName string `json:"playerName" binding:"required" example:"June Eleutheria"`
	Email      string `json:"email" binding:"required,email" example:"juneeleutheria@gmail.com"`
	Password   string `json:"password" binding:"required" example:"Abc123_@#"`
}

type GetPlayersQuery struct {
}

type PlayerController interface {
	RegisterPlayer(c *gin.Context)
	GetPlayers(c *gin.Context)
	GetPlayer(c *gin.Context)
}

type PlayerControllerImpl struct {
	service player_service.PlayerService
}

// @BasePath /api/v1
// Players godoc

// @Summary Register player
// @Description Register player
// @Tags Players
// @Accept json
// @Produce json
// @Param body body RegisterPlayerReq true "body"
// @Success 201 {object} player_repository.Player
// @Router /players [post]
func (p PlayerControllerImpl) RegisterPlayer(c *gin.Context) {
	var reqBodyJson RegisterPlayerReq

	if err := c.ShouldBindJSON(&reqBodyJson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	newPlayer := player_repository.Player{
		Email:      reqBodyJson.Email,
		PlayerName: reqBodyJson.PlayerName,
	}

	registeredPlayer, err := p.service.RegisterPlayer(&newPlayer, reqBodyJson.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusCreated, registeredPlayer)
}

// @Summary Get players
// @Description Get players
// @Tags Players
// @Produce json
// @Success 200 {array} player_repository.Player
// @Router /players [get]
func (p PlayerControllerImpl) GetPlayers(c *gin.Context) {
	var query GetPlayersQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, util.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	areas, err := p.service.GetPlayers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, areas)
}

// @Summary Get player
// @Description Get player
// @Tags Players
// @Produce json
// @Param uuid path string true "uuid"
// @Success 200 {object} player_repository.Player
// @Router /players/{uuid} [get]
func (p PlayerControllerImpl) GetPlayer(c *gin.Context) {
	area, err := p.service.GetPlayer(c.Param(("uuid")))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, area)
}

func PlayerControllerInit(playerService player_service.PlayerService) *PlayerControllerImpl {
	return &PlayerControllerImpl{
		service: playerService,
	}
}
