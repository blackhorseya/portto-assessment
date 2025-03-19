package handler

import (
	"net/http"
	_ "portto/docs" // import the generated swagger docs
	"portto/entity"
	"portto/pkg/contextx"
	"portto/pkg/httpx"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	defaultScore = 1
)

type handlerImpl struct {
	coinRepo entity.CoinRepository
}

func RegisterRoutes(coinRepo entity.CoinRepository) httpx.InitRouterFn {
	instance := &handlerImpl{
		coinRepo: coinRepo,
	}

	return func(engine *gin.Engine) {
		api := engine.Group("/api")
		{
			api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

			api.GET("/readiness", instance.Readiness)
			api.GET("/liveness", instance.Liveness)

			v1 := api.Group("/v1")
			{
				coins := v1.Group("/coins")
				{
					coins.POST("", instance.CreateCoin)
					coins.GET("/:id", instance.GetCoinById)
					coins.PATCH("/:id", instance.UpdateCoinById)
					coins.DELETE("/:id", instance.DeleteCoinById)
					coins.POST("/:id/poke", instance.PokeCoin)
				}
			}
		}
	}
}

// Readiness
// @Summary Readiness check
// @Description Check if the service is ready to accept requests
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "status"
// @Router /readiness [get]
func (i *handlerImpl) Readiness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

// Liveness
// @Summary Liveness check
// @Description Check if the service is alive
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "status"
// @Router /liveness [get]
func (i *handlerImpl) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}

type createCoinInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CreateCoin
// @Summary Create a new coin
// @Description Create a new coin with the given name and description
// @Tags Coins
// @Accept json
// @Produce json
// @Param coin body createCoinInput true "Coin"
// @Success 201 {object} entity.Coin "Created coin"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/coins [post]
func (i *handlerImpl) CreateCoin(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	var input createCoinInput
	if err := c.ShouldBindJSON(&input); err != nil {
		ctx.Error("invalid input", "error", err, "input", input)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coin := &entity.Coin{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := i.coinRepo.Create(ctx, coin); err != nil {
		ctx.Error("failed to create coin", "error", err, "coin", coin)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, coin)
}

// GetCoinById
// @Summary Get a coin by ID
// @Description Get a coin by its ID
// @Tags Coins
// @Accept json
// @Produce json
// @Param id path int true "Coin ID"
// @Success 200 {object} entity.Coin "Coin"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/coins/{id} [get]
func (i *handlerImpl) GetCoinById(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctx.Error("invalid id", "error", err, "id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	coin, err := i.coinRepo.GetByID(ctx, uint(id))
	if err != nil {
		ctx.Error("failed to get coin", "error", err, "id", id)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coin)
}

type updateCoinInput struct {
	Description string `json:"description" binding:"required"`
}

// UpdateCoinById
// @Summary Update a coin by ID
// @Description Update a coin's description by its ID
// @Tags Coins
// @Accept json
// @Produce json
// @Param id path int true "Coin ID"
// @Param coin body updateCoinInput true "Coin"
// @Success 200 {object} entity.Coin "Updated coin"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/coins/{id} [patch]
func (i *handlerImpl) UpdateCoinById(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctx.Error("invalid id", "error", err, "id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input updateCoinInput
	if err = c.ShouldBindJSON(&input); err != nil {
		ctx.Error("invalid input", "error", err, "input", input)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = i.coinRepo.UpdateDescription(ctx, uint(id), input.Description); err != nil {
		ctx.Error("failed to update coin", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	coin, err := i.coinRepo.GetByID(ctx, uint(id))
	if err != nil {
		ctx.Error("failed to get coin", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coin)
}

// DeleteCoinById
// @Summary Delete a coin by ID
// @Description Delete a coin by its ID
// @Tags Coins
// @Accept json
// @Produce json
// @Param id path int true "Coin ID"
// @Success 204 {object} map[string]string "No content"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/coins/{id} [delete]
func (i *handlerImpl) DeleteCoinById(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctx.Error("invalid id", "error", err, "id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = i.coinRepo.Delete(ctx, uint(id)); err != nil {
		ctx.Error("failed to delete coin", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// PokeCoin
// @Summary Poke a coin by ID
// @Description Poke a coin by its ID
// @Tags Coins
// @Accept json
// @Produce json
// @Param id path int true "Coin ID"
// @Success 200 {object} entity.Coin "Poked coin"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/coins/{id}/poke [post]
func (i *handlerImpl) PokeCoin(c *gin.Context) {
	ctx := contextx.WithContext(c.Request.Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctx.Error("invalid id", "error", err, "id", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err = i.coinRepo.Poke(ctx, uint(id), defaultScore); err != nil {
		ctx.Error("failed to poke coin", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	coin, err := i.coinRepo.GetByID(ctx, uint(id))
	if err != nil {
		ctx.Error("failed to get coin", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coin)
}
