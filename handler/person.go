package handler

import "github.com/gofiber/fiber/v2"

// GetPerson godoc
// @Tags person
// @Summary Show all persons
// @Produce  json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person [get]
func GetPerson(c *fiber.Ctx) error {

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})

}

// GetPersonId godoc
// @Tags person
// @Summary Get person by id
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person/id} [get]
func GetPersonId(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

// PostPerson godoc
// @Tags person
// @Summary Create a person
// @Produce  json
// @Param id path int true "Person ID"
// @Param person body model.Person true "Person model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person/ [post]
func PostPerson(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

// PutPerson godoc
// @Tags person
// @Summary Edit a person
// @Produce  json
// @Param id path int true "Person ID"
// @Param person body model.Person true "Person model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person/{id} [put]
func PutPerson(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

// DeletePerson godoc
// @Tags person
// @Summary Delete a person
// @Produce  json
// @Param id path int true "Person ID"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Failure 401 "Unauthorized"
// @Router /person/{id} [delete]
func DeletePerson(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}
