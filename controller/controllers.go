// En el paquete controller

package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	persona "goFIber/Persona"
)

type Persona persona.Persona

var personas []Persona = []Persona{
	{ID: 1, Nombre: "Alejandro", Edad: 22, Correo: "alejandro@example.com"},
	{ID: 2, Nombre: "Jackie", Edad: 18, Correo: "jackie@example.com"},
	{ID: 3, Nombre: "Aily", Edad: 22, Correo: "ailen@example.com"},
}

func GetPersonas(c *fiber.Ctx) error {
	return c.JSON(personas)
}

func GetPersona(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("ID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}
	var personaEncontrada *Persona
	for _, persona := range personas {
		if persona.ID == ID {
			personaEncontrada = &persona
			break
		}
	}

	if personaEncontrada != nil {
		return c.JSON(personaEncontrada)
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Persona no encontrada"})
}

func PostPersona(c *fiber.Ctx) error {
	var nuevaPersona Persona

	if err := c.BodyParser(&nuevaPersona); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos de persona inválidos"})
	}

	ultimoID := personas[len(personas)-1].ID
	nuevaPersona.ID = ultimoID + 1

	personas = append(personas, nuevaPersona)

	return c.Status(fiber.StatusCreated).JSON(nuevaPersona)
}
