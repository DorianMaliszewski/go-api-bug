package handlers

import (
	"bug-tracker/backend/daos"
	"bug-tracker/backend/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetupBugRoutes(r fiber.Router) {
	r.Get("/", findAll)
	r.Post("/", create)
	r.Get("/:bug_id", findOne)
	r.Delete("/:bug_id", delete)
	r.Put("/:bug_id", update)
	r.Put("/:bug_id/status/:new_status", updateStatus)
}

func findAll(c *fiber.Ctx) error {
	bugs := daos.FindAllBugs()

	return c.JSON(&bugs)
}

func findOne(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("bug_id"))
	if err != nil {
		log.Fatal("Error when converting id")
		return c.SendString("Cant convert bug id")
	}
	b := daos.FindBug(int32(id))
	if b != nil {
		return c.JSON(&b)
	} else {
		return c.SendString("No bug found")
	}
}

func create(c *fiber.Ctx) error {
	b := &models.Bug{}
	if err := c.BodyParser(b); err != nil {
		return err
	}
	b.ID = 0
	return c.JSON(daos.CreateBug(b))
}

func update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("bug_id"))
	if err != nil {
		log.Fatal("Error when converting id")
		return c.SendString("Cant convert bug id")
	}
	b := &models.Bug{}

	if err := c.BodyParser(b); err != nil {
		return err
	}
	return c.JSON(daos.UpdateBug(int32(id), b))
}

func delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("bug_id"))
	if err != nil {
		log.Fatal("Error when converting id")
		return c.SendString("Cant convert bug id")
	}
	return c.JSON(daos.DeleteBug(int32(id)))
}

func updateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("bug_id"))
	if err != nil {
		log.Fatal("Error when converting id")
		return c.SendString("Cant convert bug id")
	}
	return c.JSON(daos.UpdateBugStatus(int32(id), c.Params("new_status")))
}
