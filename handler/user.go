package handler

import (
	"goTempl/model"
	"goTempl/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "jeff@jeff.com",
	}
	return render(c, user.Show(u))
}
