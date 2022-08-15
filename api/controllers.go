package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/guionardo/go-api-example/domain"
	"gorm.io/gorm"
)

type (
	FeiraController struct {
		Service domain.FeiraService
	}
	ErrorMessage struct {
		Message string `json:"message"`
	}
)

// GetFeira godoc
// @Summary  Get Feira
// @Schemes
// @Description  Reads and feira by id.
// @Tags         read
// @Accept       json
// @Produce      json
// @Success      200	{object}   	domain.Feira
// @Param 			registro path string true "Feira Registro"
// @NotFound     404	{object}   	ErrorMessage
// @Failure      500	{object}	ErrorMessage
// @Router       /feiras/{registro} [get]
func (c *FeiraController) GetFeira(ctx *gin.Context) {
	registro := ctx.Param("registro")

	if len(registro) == 0 {
		errorResponse(ctx, fmt.Errorf("missing registro"), 404)
		return
	}
	feira, err := c.Service.FindByRegistro(registro)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse(ctx, fmt.Errorf("not found"), 404)
		} else {
			errorResponse(ctx, err)
		}
		return
	}
	if feira == nil {
		errorResponse(ctx, fmt.Errorf("not found"), 404)
	} else {
		ctx.JSON(200, feira)
	}
}

// GetFeiras godoc
// @Summary  Get all feiras
// @Schemes
// @Description  Reads and returns all the tasks.
// @Tags         read
// @Accept       json
// @Produce      json
// @Param distrito query string false "Distrito"
// @Param regiao5 query string false "Região 5"
// @Param nome_feira query string false "Nome Feira"
// @Param bairro query string false "Bairro"
// @Success      200    {array}   []domain.Feira
// @Failure      500  	{object}  ErrorMessage
// @Router       /feiras/ [get]
func (c *FeiraController) GetFeiras(ctx *gin.Context) {
	var registro, regiao5, nomeFeira, bairro string
	unpackVars(getQueryVars(ctx, "distrito", "regiao5", "nome_feira", "bairro"), &registro, &regiao5, &nomeFeira, &bairro)

	feiras, err := c.Service.Query(registro, regiao5, nomeFeira, bairro)
	if err != nil {
		errorResponse(ctx, err)
	} else {
		if len(feiras) == 0 {
			errorResponse(ctx, fmt.Errorf("not found"), 404)
		} else {
			ctx.JSON(200, feiras)
		}
	}
}

// CreateFeira godoc
// @Summary  Creates feira
// @Schemes
// @Description  Creates a new feira.
// @Tags         write
// @Accept       json
// @Produce      json
// @Param		feira body domain.Feira true "New Feira"
// @Success      201 {object} domain.Feira
// @Failure      500  	{object}  ErrorMessage
// @Response 	409  {object}  ErrorMessage
// @Router       /feiras/ [post]
func (c *FeiraController) CreateFeira(ctx *gin.Context) {
	var feira domain.Feira
	if err := ctx.BindJSON(&feira); err != nil {
		errorResponse(ctx, err, 400)
		return
	}
	if err := c.Service.Create(&feira); err != nil {
		if strings.Contains(err.Error(),"UNIQUE"){
			errorResponse(ctx, fmt.Errorf("registro %s already exists",feira.Registro), 409)
		} else {
			errorResponse(ctx, err)
		}		
	} else {
		ctx.JSON(201, feira)

	}
}

// GetFeira godoc
// @Summary  Delete feira
// @Schemes
// @Description  Deletes a feira by registro
// @Tags         write
// @Accept       json
// @Produce      json
// @Success      202	{object}   	domain.Feira
// @Param 		 registro path string true "Feira Registro"
// @NotFound     404	{object}   	ErrorMessage
// @Failure      500	{object}	ErrorMessage
// @Router       /feiras/{registro} [delete]
func (c *FeiraController) DeleteFeira(ctx *gin.Context) {
	registro := ctx.Param("registro")

	if len(registro) == 0 {
		errorResponse(ctx, fmt.Errorf("missing registro"), 404)
		return
	}

	if err := c.Service.DeleteByRegistro(registro); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse(ctx, fmt.Errorf("not found"), 404)
		} else {
			errorResponse(ctx, err)
		}
	} else {
		ctx.JSON(202, "Deleted")

	}
}

// CreateFeira godoc
// @Summary  Updates feira
// @Schemes
// @Description  Updates feira by registro
// @Tags         write
// @Accept       json
// @Produce      json
// @Param		feira body domain.Feira true "Feira"
// @Success      202
// @Failure      500  	{object}  ErrorMessage
// @Response	400 	{object}  ErrorMessage
// @Router       /feiras/ [put]
func (c *FeiraController) UpdateFeira(ctx *gin.Context) {
	var feira domain.Feira
	if err := ctx.BindJSON(&feira); err != nil {
		errorResponse(ctx, err, 400)
		return
	}

	if err := c.Service.Update(&feira); err != nil {
		errorResponse(ctx, err, 500)
	} else {
		ctx.JSON(202, feira)
	}

}
