package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/salsabilatts/sharing-vision-test/internal/model"
	"github.com/salsabilatts/sharing-vision-test/internal/service"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) Create(c *gin.Context) {
	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := h.service.Create(&post); err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) GetAll(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("param"))
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "limit must be a positive integer",
		})
		return
	}

	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "offset must be a non-negative integer",
		})
		return
	}

	status := c.Query("status")

	if status != "" &&
		status != "publish" &&
		status != "draft" &&
		status != "thrash" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "status must be publish, draft, or thrash",
		})
		return
	}

	posts, err := h.service.GetAll(limit, offset, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve articles",
		})
		return
	}

	c.JSON(http.StatusOK, posts)

}

func (h *PostHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("param"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid article id",
		})
		return
	}

	post, err := h.service.GetByID(uint(id))
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid article id",
		})
		return
	}

	var post model.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	updatedPost, err := h.service.Update(uint(id), &post)
	if err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedPost)
}

func (h *PostHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid article id",
		})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func handleServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrTitleRequired),
		errors.Is(err, service.ErrContentRequired),
		errors.Is(err, service.ErrCategoryRequired),
		errors.Is(err, service.ErrInvalidStatus):

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	case errors.Is(err, service.ErrPostNotFound):

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

	default:

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}
}
