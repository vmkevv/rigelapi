package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(db *ent.Client, config config.Config) func(*fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type res struct {
		Teacher Teacher `json:"teacher"`
		JWT     string  `json:"jwt"`
	}
	return func(c *fiber.Ctx) error {
		var reqData req
		err := c.BodyParser(&reqData)
		if err != nil {
			return err
		}
		teacher, err := db.Teacher.Query().Where(teacher.EmailEQ(reqData.Email)).First(c.Context())
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return NewClientErr(fiber.StatusBadRequest, "Credenciales incorrectas.")
			}
			return err
		}
		err = bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(reqData.Password))
		if err != nil {
			return NewClientErr(fiber.StatusBadRequest, "Credenciales incorrectas.")
		}
		claims := AppClaims{
			teacher.ID,
			jwt.RegisteredClaims{},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(config.App.JWTKey))
		if err != nil {
			return err
		}
		return c.JSON(res{
			Teacher: Teacher{
				teacher.ID,
				teacher.Name,
				teacher.LastName,
				teacher.Email,
			},
			JWT: tokenStr,
		})
	}
}
