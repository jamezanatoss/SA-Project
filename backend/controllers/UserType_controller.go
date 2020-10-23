package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/james/app/ent"
	"github.com/james/app/ent/usertype"
)

// UserTypeController defines the struct for the UserType controller
type UserTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateUserType handles POST requests for adding usertype entities
// @Summary Create usertype
// @Description Create usertype
// @ID create-usertype
// @Accept   json
// @Produce  json
// @Param gender body ent.UserType true "UserType entity"
// @Success 200 {object} ent.UserType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes [post]
func (ctl *UserTypeController) CreateUserType(c *gin.Context) {
	obj := ent.UserType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "UserType binding failed",
		})
		return
	}

	ut, err := ctl.client.UserType.
		Create().
		SetUserTypeName(obj.UserTypeName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ut)
}

// GetUserType handles GET requests to retrieve a UserType entity
// @Summary Get a usertype entity by ID
// @Description get usertype by ID
// @ID get-usertype
// @Produce  json
// @Param id path int true "UserType ID"
// @Success 200 {object} ent.UserType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes/{id} [get]
func (ctl *UserTypeController) GetUserType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ut, err := ctl.client.UserType.
		Query().
		Where(usertype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ut)
}

// ListUserType handles request to get a list of UserType entities
// @Summary List usertype entities
// @Description list usertype entities
// @ID list-usertype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.UserType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /usertypes [get]
func (ctl *UserTypeController) ListUserType(c *gin.Context) {
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

	usertypes, err := ctl.client.UserType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, usertypes)
}

// NewUserTypeController creates and registers handles for the UserType controller
func NewUserTypeController(router gin.IRouter, client *ent.Client) *UserTypeController {
	utc := &UserTypeController{
		client: client,
		router: router,
	}
	utc.register()
	return utc
}

// InitUserTypeController registers routes to the main engine
func (ctl *UserTypeController) register() {
	usertypes := ctl.router.Group("/usertypes")

	usertypes.GET("", ctl.ListUserType)
	usertypes.POST("", ctl.CreateUserType)
	usertypes.GET(":id", ctl.GetUserType)

}
