package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d *Delivery) CreateGroup(c *gin.Context) {
	var requestBody CreateGroupRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	res, err := d.grpcDelivery.CreateGroup(c, requestBody.Code)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (d *Delivery) AddVideo(c *gin.Context) {
	var requestBody VideoAddRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	res, err := d.grpcDelivery.AddVideo(c, requestBody.Code, string(requestBody.Content))
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (d *Delivery) ShowGroups(c *gin.Context) {
	var requestBody CreateGroupRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	res, err := d.grpcDelivery.CreateGroup(c, requestBody.Code)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (d *Delivery) RemoveVideo(c *gin.Context) {
	var requestBody RemoveVideoBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	res, err := d.grpcDelivery.RemoveVideo(c, requestBody.Id)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (d *Delivery) Login(c *gin.Context) {
	var requestBody GroupLoginRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	res, err := d.grpcDelivery.CreateGroup(c, requestBody.Code)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (d *Delivery) DeleteGroup(c *gin.Context) {
	var requestBody CreateGroupRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	res, err := d.grpcDelivery.CreateGroup(c, requestBody.Code)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (d *Delivery) ShowVideos(c *gin.Context) {
	var requestBody CreateGroupRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	res, err := d.grpcDelivery.CreateGroup(c, requestBody.Code)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
