package review

import (
	"Office-Booking/app/config"
	domain "Office-Booking/domain/review"
	"Office-Booking/domain/review/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	ReviewUsecase domain.ReviewUsecase
}

func NewReviewController(e *echo.Echo, Usecase domain.ReviewUsecase) {
	ReviewController := &ReviewController{
		ReviewUsecase: Usecase,
	}

	e.POST("/review", ReviewController.Create)
	e.DELETE("/admin/review/:id", ReviewController.Delete)
}
func (u *ReviewController) Create(c echo.Context) error {
	var req request.ReviewPost

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.ReviewUsecase.Create(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    401,
			"status":  false,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        200,
		"status":      true,
		"ID":          res.ID,
		"Rating":      res.Rating,
		"Description": res.Description,
	})

}

func (u *ReviewController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, e := u.ReviewUsecase.Delete(id)

	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "not found",
			"code":     404,
		})
	}

	config.DB.Unscoped().Delete(&domain.Review{}, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user with id " + strconv.Itoa(id),
		"code":     200,
	})
}
