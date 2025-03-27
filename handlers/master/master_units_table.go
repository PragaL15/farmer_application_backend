package Masterhandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Unit struct {
	ID       int64  `json:"id"`
	UnitName string `json:"unit_name"`
}

func GetAllUnits(db *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query(context.Background(), "SELECT * FROM admin_schema.get_all_units()")
		if err != nil {
			log.Println("Error fetching units:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch units"})
		}
		defer rows.Close()
		var units []Unit
		for rows.Next() {
			var u Unit
			if err := rows.Scan(&u.ID, &u.UnitName); err != nil {
				log.Println("Error scanning row:", err)
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse data"})
			}
			units = append(units, u)
		}
		return c.JSON(units)
	}
}

func GetUnitByID(db *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var u Unit
		err := db.QueryRow(context.Background(), "SELECT * FROM admin_schema.get_unit_by_id($1)", id).Scan(&u.ID, &u.UnitName)
		if err != nil {
			log.Println("Error fetching unit:", err)
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Unit not found"})
		}
		return c.JSON(u)
	}
}

func CreateUnit(db *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var u Unit
		if err := c.BodyParser(&u); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}
		_, err := db.Exec(context.Background(), "SELECT admin_schema.insert_unit($1)", u.UnitName)
		if err != nil {
			log.Println("Error inserting unit:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert unit"})
		}
		return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Unit created successfully"})
	}
}

func UpdateUnit(db *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var u Unit
		if err := c.BodyParser(&u); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}
		_, err := db.Exec(context.Background(), "SELECT admin_schema.update_unit($1, $2)", id, u.UnitName)
		if err != nil {
			log.Println("Error updating unit:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update unit"})
		}
		return c.JSON(fiber.Map{"message": "Unit updated successfully"})
	}
}
