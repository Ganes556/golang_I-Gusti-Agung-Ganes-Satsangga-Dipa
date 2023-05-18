package item

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/usecases/item"
	"github.com/labstack/echo/v4"
)

type ItemHandler interface {
	GetAll(c echo.Context) error
	Create(c echo.Context) error
	Get(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetByCategory(c echo.Context) error
}

type itemHandler struct {
	itemUC item.ItemUsecase
}

func NewItemHandler(itemUC item.ItemUsecase) ItemHandler {
	return &itemHandler{itemUC: itemUC}
}