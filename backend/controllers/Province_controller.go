package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/james/app/ent"
	"github.com/james/app/ent/province"
)

// ProvinceController defines the struct for the Province controller
type ProvinceController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateProvince handles POST requests for adding province entities
// @Summary Create province
// @Description Create province
// @ID create-province
// @Accept   json
// @Produce  json
// @Param gender body ent.Province true "Province entity"
// @Success 200 {object} ent.Province
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /provinces [post]
func (ctl *ProvinceController) CreateProvince(c *gin.Context) {
	obj := ent.Province{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Province binding failed",
		})
		return
	}

	pr, err := ctl.client.Province.
		Create().
		SetProvinceName(obj.ProvinceName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pr)
}

// GetProvince handles GET requests to retrieve a province entity
// @Summary Get a province entity by ID
// @Description get province by ID
// @ID get-province
// @Produce  json
// @Param id path int true "Province ID"
// @Success 200 {object} ent.Province
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /provinces/{id} [get]
func (ctl *ProvinceController) GetProvince(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pr, err := ctl.client.Province.
		Query().
		Where(province.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pr)
}

// ListProvince handles request to get a list of Province entities
// @Summary List province entities
// @Description list province entities
// @ID list-province
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Province
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Provinces [get]
func (ctl *ProvinceController) ListProvince(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	provinces, err := ctl.client.Province.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, provinces)
}

// NewProvinceController creates and registers handles for the Province controller
func NewProvinceController(router gin.IRouter, client *ent.Client) *ProvinceController {
	prc := &ProvinceController{
		client: client,
		router: router,
	}
	prc.register()
	return prc
}

// InitProvinceController registers routes to the main engine
func (ctl *ProvinceController) register() {

	provinces := ctl.router.Group("/provinces")
	provinces.GET("", ctl.ListProvince)
	provinces.POST("", ctl.CreateProvince)
	provinces.GET(":id", ctl.GetProvince)

}
