package controller

import (
	"TokoBelanja/helper"
	"TokoBelanja/model/input"
	"TokoBelanja/model/response"
	"TokoBelanja/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	srv service.ProductService
}

func NewProductController(srv service.ProductService) *productController {
	return &productController{srv: srv}
}

func (pc *productController) Post(c *gin.Context) {
	var input *input.ProductCreateInput

	role := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productData, err2 := pc.srv.Create(role, input)
	if err2 != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productRespone := response.ProductResponse{
		ID:         productData.ID,
		Title:      productData.Title,
		Price:      productData.Price,
		Stock:      productData.Stock,
		CategoryID: productData.CategoryID,
		CreatedAt:  productData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			productRespone,
		),
	)
}

func (pc *productController) Get(c *gin.Context) {
	var products []response.ProductResponse

	productData, err := pc.srv.GetAll()
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	for _, product := range productData {
		tmpData := response.ProductResponse{
			ID:         product.ID,
			Title:      product.Title,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
			CreatedAt:  product.CreatedAt,
		}
		products = append(products, tmpData)
	}
	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			products,
		),
	)
}

func (pc *productController) Put(c *gin.Context) {
	var (
		uri       input.ProductID
		inputBody *input.ProductPutInput
	)

	role := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	err = c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productData, err := pc.srv.Put(role, uri.ID, inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productResponse := response.ProductPutResponse{
		Product: response.ProductPutResponseBody{
			ID:         productData.ID,
			Title:      productData.Title,
			Price:      productData.Price,
			Stock:      productData.Stock,
			CategoryID: productData.CategoryID,
			CreatedAt:  productData.CreatedAt,
			UpdatedAt:  productData.UpdatedAt,
		},
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			productResponse,
		),
	)
}

func (pc *productController) Delete(c *gin.Context) {
	var uri input.ProductID

	role := c.MustGet("roleUser").(string)

	err := c.ShouldBindUri(&uri)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	err2 := pc.srv.Delete(role, uri.ID)
	if err2 != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	productResponse := response.ProductDeleteResponse{
		Message: "Product has been successfully deleted",
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			productResponse,
		),
	)
}
