package handler

import "github.com/gofiber/fiber/v2"

func GetPerson(c *fiber.Ctx) error {

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})

}
func GetPersonId(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

func PostPerson(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

func PutPerson(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}

func DeletePerson(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "success"})
}
