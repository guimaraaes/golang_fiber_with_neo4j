package handler

import (
	url "net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/guimaraaes/golang_fiber_with_neo4j/model"
	"github.com/guimaraaes/golang_fiber_with_neo4j/repository"
)

// GetPerson godoc
// @Tags person
// @Summary Show all persons
// @Produce  json
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person [get]
func GetPerson(c *fiber.Ctx) error {
	var person model.Person
	p, err := repository.FindR(person, nil)
	if err != "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(p)

}

// GetPersonId godoc
// @Tags person
// @Summary Get person by id
// @Produce json
// @Param name path string true "name Person"
// @Param born path int true "year person born"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person/{name}/{born} [get]
func GetPersonId(c *fiber.Ctx) error {
	n := c.Params("name")
	name, _ := url.QueryUnescape(n)
	born, _ := strconv.ParseInt(c.Params("born"), 10, 64)

	var person model.Person
	// , map[string]interface{}{"name": name, "born": born}
	p, err := repository.FindR(person, map[string]interface{}{"name": name, "born": born})
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err})
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// PostPerson godoc
// @Tags person
// @Summary Create a person
// @Produce  json
// @Param person body model.Person true "Person model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person/ [post]
func PostPerson(c *fiber.Ctx) error {
	var person model.Person
	if err := c.BodyParser(&person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	p, err := repository.CreateR(person)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err, "person": p})
	}
	return c.Status(fiber.StatusOK).JSON(p)
}

// PostPersonWithRel godoc
// @Tags person
// @Summary Create a person with relationship
// @Produce  json
// @Param relationship body model.PersonRelationship true "Relationship model"
// @Success 200 "OK"
// @Failure 400 "Bad Request"
// @Router /person_with_relationship [post]
func PostPersonWithRelationship(c *fiber.Ctx) error {

	var person model.PersonRelationship
	if err := c.BodyParser(&person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	p, err := repository.CreateWithRElR(person.Person, person.Relationship, person.Movie)
	if err != "" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": err, "movie": p})
	}
	return c.Status(fiber.StatusOK).JSON(p)
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
