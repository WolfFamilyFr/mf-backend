package rivals

import (
	"strconv"

	"github.com/bchaillou003/marvel-family-backend/database"
	"github.com/bchaillou003/marvel-family-backend/models/rivals"
	"github.com/gofiber/fiber/v2"
)

func CreateAbility(c *fiber.Ctx) error {
	ability := new(rivals.Ability)
	if err := c.BodyParser(ability); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&ability)
	return c.Status(200).JSON(ability)
}

func CreateAchievement(c *fiber.Ctx) error {
	achievement := new(rivals.Achievement)
	if err := c.BodyParser(achievement); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&achievement)
	return c.Status(200).JSON(achievement)
}

func CreateCharacter(c *fiber.Ctx) error {
	character := new(rivals.Character)
	if err := c.BodyParser(character); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&character)
	return c.Status(200).JSON(character)
}

func CreateCosmetic(c *fiber.Ctx) error {
	cosmetic := new(rivals.Cosmetic)
	if err := c.BodyParser(cosmetic); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&cosmetic)
	return c.Status(200).JSON(cosmetic)
}

func CreateLore(c *fiber.Ctx) error {
	lore := new(rivals.Lore)
	if err := c.BodyParser(lore); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&lore)
	return c.Status(200).JSON(lore)
}

func getAll(c *fiber.Ctx, dest interface{}) error {
	result := database.DB.Db.Find(dest)
	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}
	return c.Status(200).JSON(dest)
}

func GetAllAbility(c *fiber.Ctx) error {
	return getAll(c, new([]rivals.Ability))
}

func GetAllAchievement(c *fiber.Ctx) error {
	return getAll(c, new([]rivals.Achievement))
}

func GetAllCharacter(c *fiber.Ctx) error {
	return getAll(c, new([]rivals.Character))
}

func GetAllCosmetic(c *fiber.Ctx) error {
	return getAll(c, new([]rivals.Cosmetic))
}

func GetAllLore(c *fiber.Ctx) error {
	return getAll(c, new([]rivals.Lore))
}

func getByID(c *fiber.Ctx, dest interface{}) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(err)
	}

	result := database.DB.Db.First(&dest, id)
	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}
	return c.Status(200).JSON(dest)
}

func GetAbilityByID(c *fiber.Ctx) error {
	return getByID(c, new(rivals.Ability))
}

func GetAchievementByID(c *fiber.Ctx) error {
	return getByID(c, new(rivals.Achievement))
}

func GetCharacterByID(c *fiber.Ctx) error {
	return getByID(c, new(rivals.Character))
}

func GetCosmeticByID(c *fiber.Ctx) error {
	return getByID(c, new(rivals.Cosmetic))
}

func GetLoreByID(c *fiber.Ctx) error {
	return getByID(c, new(rivals.Lore))
}
