package item

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/dto"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/helpers"
	"github.com/labstack/echo/v4"
)

func (i *itemHandler) GetAll(c echo.Context) error {

	keyword := c.QueryParam("keyword")

	if keyword != "" {
		items, err := i.itemUC.GetByKeyword(keyword)
		
		if err != nil {
			return helpers.ResponseError(c, getStatus(err), err)
		}

		return helpers.ResponseSuccess(c, getStatus(err), "get item by keyword success", items)
	}


	items, err := i.itemUC.GetAll()

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccess(c, getStatus(err), "get all items success", items)
}

func (i *itemHandler) Create(c echo.Context) error {
	req := dto.ItemCreateDTO{}

	c.Bind(&req)
	
	if err := isRequestValid(&req); err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	err := i.itemUC.Create(req)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}
	return helpers.ResponseSuccessWithoutData(c, getStatus(err), "create item success")
}


func (i *itemHandler) Get(c echo.Context) error {

	id := c.Param("id")
	
	if err := isReuqestUuid(id); err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}
	
	item, err := i.itemUC.GetById(id)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccess(c, getStatus(err), "get item by id success", item)
}


func (i *itemHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := isReuqestUuid(id); err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	err := i.itemUC.Delete(id)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccessWithoutData(c, getStatus(err), "delete item success")
}

func (i *itemHandler) Update(c echo.Context) error {
	id := c.Param("id")

	if err := isReuqestUuid(id); err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	req := dto.ItemUpdateDTO{}

	err := c.Bind(&req)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	err = i.itemUC.Update(id, req)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccessWithoutData(c, getStatus(err), "update item success")
}

func (i *itemHandler) GetByCategory(c echo.Context) error {
	category := c.Param("category_id")

	categoryIdInt, err := helpers.CheckId(category, "category_id")

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	items, err := i.itemUC.GetByCategory(categoryIdInt)

	if err != nil {
		return helpers.ResponseError(c, getStatus(err), err)
	}

	return helpers.ResponseSuccess(c, getStatus(err), "get item by category success", items)
}

