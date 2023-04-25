package authAPI

import (
	"fmt"

	"github.com/dancondo/users-api/common"
	"github.com/dancondo/users-api/domain/auth"
	"github.com/dancondo/users-api/domain/user"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Get user email and password and return a token.
// @Description  User Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200 {object} user.LoginUserResponseDto
// @Param        user body user.UserRequestDto true "the user email and password"
// @Router       /api/auth/login [post]
func LoginHandler(c *fiber.Ctx) error {
	userReq := new(user.UserRequestDto)

	if err := c.BodyParser(userReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userService := user.NewService()
	authService := auth.NewService(common.GetEnv("APP_SECRET"))

	user, err := userService.GetUserByUsername(userReq.Username)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	if user == nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, fmt.Sprintf("Username %s not found", userReq.Username))
	}

	err = userService.ValidateUserPassword(user, userReq.Password)

	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Incorrect password")
	}

	token, err := authService.GenerateJWT(user.Username)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user.ToUserResponseDto(token))
}
