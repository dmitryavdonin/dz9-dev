package handler

import (
	"net/http"
	"strconv"
	"time"

	"store/internal/model"

	"github.com/gin-gonic/gin"
)

// add book to store
func (h *Handler) addBookToStore(c *gin.Context) {
	var input model.StoreBook
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	// check if such a book is already presented in the store
	book, err := h.service.StoreBook.GetById(input.BookId)
	if err == nil {
		// just update in_stock for existant book
		book.InStock = input.InStock
		book.ModifiedAt = time.Now()
		h.service.StoreBook.Update(input.BookId, book)
		return
	}

	// otherwise create a new book in store
	now := time.Now()

	item := model.StoreBook{
		BookId:     input.BookId,
		InStock:    input.InStock,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	id, err := h.service.StoreBook.Create(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	book, err = h.service.StoreBook.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// update book amount in the store
func (h *Handler) updateStoreBook(c *gin.Context) {
	var input model.StoreBook
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	item := model.StoreBook{
		InStock:    input.InStock,
		ModifiedAt: time.Now(),
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	err = h.service.StoreBook.Update(id, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	book, err := h.service.StoreBook.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// get book in store by id
func (h *Handler) getStoreBookById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	book, err := h.service.StoreBook.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

// delete book from the store
func (h *Handler) deleteStoreBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	err = h.service.StoreBook.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Id:     id,
		Status: "success",
	})
}

// get all books in the store
func (h *Handler) getAllStoreBooks(c *gin.Context) {

	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var items []model.StoreBook
	items, err := h.service.StoreBook.GetAll(intLimit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(items), "data": items})
}
