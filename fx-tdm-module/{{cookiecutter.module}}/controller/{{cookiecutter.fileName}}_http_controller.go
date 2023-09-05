package controller

import (
	"{{cookiecutter.basePath}}/dto"
	"{{cookiecutter.basePath}}/model"
	"{{cookiecutter.basePath}}/service"

	"github.com/labstack/echo/v4"
	"github.com/sigmaott/gest/package/extension/i18nfx"
	base "github.com/sigmaott/gest/package/technique/database/base"
	queryBuilder "github.com/sigmaott/gest/package/technique/database/pagination"
)

type I{{cookiecutter.className}}Controller interface {
	Create()
	FindOne()
	Paginate()
	UpdateOne()
	DeleteOne()
}

func New{{cookiecutter.className}}Controller(router *echo.Group, exampleService service.IExampleService, i18nService i18nfx.II18nService) IExampleController {
	return &exampleController{
		router:         router.Group("/example"),
		exampleService: exampleService,
		i18nService:    i18nService,
	}

}

type exampleController struct {
	router         *echo.Group
	i18nService    i18nfx.II18nService
	exampleService service.I{{cookiecutter.className}}Service
}

// Create
// @Summary Create Example
// @Tags Example
// @Produce json
// @Param payload body  dto.CreateExample true "payload"
// @Success 201 {object} model.Example
// @Router /examples/{id} [post]
func (e *exampleController) Create() {
	e.router.POST("", func(c echo.Context) error {
		payload := new(dto.Create{{cookiecutter.className}})
		res, err := e.exampleService.Create(payload, "appId")
		if err != nil {
			return err
		}
		return c.JSON(201, res)
	})
}

// FindOne
// @Summary Get Example By ID
// @Tags Example
// @Produce json
// @Param id path string true "Example id"
// @Success 200 {object} model.Example
// @Router /examples/{id} [get]
func (e *exampleController) FindOne() {
	e.router.GET(":id", func(c echo.Context) error {
		id := c.Param(":id")
		res, err := e.exampleService.FindOne(struct {
			Id    string
			AppId string
		}{Id: id, AppId: ""})
		if err != nil {
			return err
		}
		return c.JSON(200, res)
	})
}

// Paginate
// @Summary Paginate Example
// @Description  Paginate Example
// @Tags Example
// @Param sort query []string false "Use only allowed properties separated by semicolon; default is ascending created_at; prefix name with hyphen/minus sign to get descending order"
// @Param page query integer false "pagination number"
// @Param perPage query integer false "per Page"
// @Accept       json
// @Produce      json
// @Success 200 {object} model.PaginateExample
// @Router /examples [get]
func (e *exampleController) Paginate() {
	e.router.GET("", func(c echo.Context) error {
		query, sort, paginate, err := queryBuilder.MongoParserQuery[model.Example](c.Request().URL.Query())
		if err != nil {
			return err
		}
		result, err := e.exampleService.Paginate(query, paginate, (*base.Sort)(&sort))
		if err != nil {
			return err
		}
		return c.JSON(200, result)

	})
}

// UpdateOne
// @Summary Update Example By ID
// @Tags Example
// @Produce json
// @Param id path string true "Example id"
// @Param payload body  dto.UpdateExample true "payload"
// @Success 200 {object} model.Example
// @Router /examples/{id} [patch]
func (e *exampleController) UpdateOne() {
	e.router.PATCH(":id", func(c echo.Context) error {
		id := c.Param(":id")
		payload := new(dto.Update{{cookiecutter.className}})
		res, err := e.exampleService.UpdateOne(struct {
			Id    string
			AppId string
		}{Id: id, AppId: ""}, payload)
		if err != nil {
			return err
		}
		return c.JSON(200, res)
	})
}

// Delete
// @Summary Delete Example By ID
// @Tags Example
// @Produce json
// @Param id path string true "Example id"
// @Success 200 {object} model.Example
// @Router /examples/{id} [delete]
func (e *exampleController) DeleteOne() {
	e.router.DELETE(":id", func(c echo.Context) error {
		id := c.Param(":id")
		res, err := e.exampleService.DeleteOne(struct {
			Id    string
			AppId string
		}{Id: id, AppId: ""})
		if err != nil {
			return err
		}
		return c.JSON(200, res)
	})
}
