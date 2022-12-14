package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/dpto"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

func ProvsHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		depID := c.Params("depid")
		provs, err := db.Provincia.Query().
			Where(provincia.HasDepartamentoWith(dpto.ID(depID))).
			Order(ent.Asc(provincia.FieldName)).All(c.Context())
		if err != nil {
			return err
		}
		provsRes := make([]Provincia, len(provs))
		for i, prov := range provs {
			provsRes[i] = Provincia{prov.ID, prov.Name}
		}
		return c.JSON(provsRes)
	}
}
