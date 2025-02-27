package Masterhandlers

import (
	"context"
	"log"
	"strconv"
"time"
	"github.com/gofiber/fiber/v2"
	"github.com/PragaL15/go_newBackend/go_backend/db"
	"github.com/go-playground/validator/v10"
)
func InsertMasterVehicle(c *fiber.Ctx) error {
	type Request struct {
		InsuranceID          int    `json:"insurance_id"`
		VehicleName          string `json:"vehicle_name" validate:"required,max=255"`
		VehicleManufactureYear string `json:"vehicle_manufacture_year" validate:"max=4"`
		VehicleWarranty      string `json:"vehicle_warranty" validate:"max=255"`
		VehicleMake          int    `json:"vehicle_make" validate:"required"`
		VehicleModel         int    `json:"vehicle_model" validate:"required"`
		VehicleRegistrationNo string `json:"vehicle_registration_no" validate:"required,max=50"`
		VehicleEngineType    int    `json:"vehicle_engine_type" validate:"required"`
		VehiclePurchaseDate  string `json:"vehicle_purchase_date" validate:"required"`
		VehicleColor         string `json:"vehicle_color" validate:"required,max=50"`
		Col1                 string `json:"col1"`
		Col2                 string `json:"col2"`
		Col3                 string `json:"col3"`
		VehicleInsuranceID   int    `json:"vehicle_insurance_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	_, err := db.Pool.Exec(context.Background(), `
		CALL insert_master_vehicle($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
	`, req.InsuranceID, req.VehicleName, req.VehicleManufactureYear, req.VehicleWarranty, 
		req.VehicleMake, req.VehicleModel, req.VehicleRegistrationNo, req.VehicleEngineType, 
		req.VehiclePurchaseDate, req.VehicleColor, req.Col1, req.Col2, req.Col3, req.VehicleInsuranceID)

	if err != nil {
		log.Printf("Failed to insert vehicle: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert vehicle"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Vehicle added successfully"})
}
func UpdateMasterVehicle(c *fiber.Ctx) error {
	type Request struct {
		VehicleID            int    `json:"vehicle_id" validate:"required,min=1"`
		InsuranceID          int    `json:"insurance_id"`
		VehicleName          string `json:"vehicle_name" validate:"required,max=255"`
		VehicleManufactureYear string `json:"vehicle_manufacture_year" validate:"max=4"`
		VehicleWarranty      string `json:"vehicle_warranty" validate:"max=255"`
		VehicleMake          int    `json:"vehicle_make" validate:"required"`
		VehicleModel         int    `json:"vehicle_model" validate:"required"`
		VehicleRegistrationNo string `json:"vehicle_registration_no" validate:"required,max=50"`
		VehicleEngineType    int    `json:"vehicle_engine_type" validate:"required"`
		VehiclePurchaseDate  string `json:"vehicle_purchase_date" validate:"required"`
		VehicleColor         string `json:"vehicle_color" validate:"required,max=50"`
		Col1                 string `json:"col1"`
		Col2                 string `json:"col2"`
		Col3                 string `json:"col3"`
		VehicleInsuranceID   int    `json:"vehicle_insurance_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	_, err := db.Pool.Exec(context.Background(), `
		CALL update_master_vehicle($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
	`, req.VehicleID, req.InsuranceID, req.VehicleName, req.VehicleManufactureYear, req.VehicleWarranty, 
		req.VehicleMake, req.VehicleModel, req.VehicleRegistrationNo, req.VehicleEngineType, 
		req.VehiclePurchaseDate, req.VehicleColor, req.Col1, req.Col2, req.Col3, req.VehicleInsuranceID)

	if err != nil {
		log.Printf("Failed to update vehicle: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update vehicle"})
	}

	return c.JSON(fiber.Map{"message": "Vehicle updated successfully"})
}
func DeleteMasterVehicle(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}
	_, err = db.Pool.Exec(context.Background(), `
		CALL delete_master_vehicle($1);
	`, idInt)

	if err != nil {
		log.Printf("Failed to delete vehicle: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete vehicle"})
	}

	return c.JSON(fiber.Map{"message": "Vehicle deleted successfully"})
}

func GetVehicles(c *fiber.Ctx) error {
	rows, err := db.Pool.Query(context.Background(), "SELECT * FROM get_master_vehicles()")
	if err != nil {
		log.Printf("Failed to fetch vehicle records: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch vehicles"})
	}
	defer rows.Close()

	var vehicles []map[string]interface{}

	for rows.Next() {
		var vehicleID, insuranceID, vehicleMake, vehicleModel, vehicleEngineType, vehicleInsuranceID *int
		var vehicleMakeName, vehicleModelName, vehicleEngineTypeName *string
		var vehicleName, vehicleManufactureYear, vehicleWarranty, vehicleRegistrationNo, vehicleColor*string
		var vehiclePurchaseDate *time.Time

		if err := rows.Scan(
			&vehicleID, &insuranceID, &vehicleName, &vehicleManufactureYear, &vehicleWarranty,
			&vehicleMake, &vehicleMakeName, &vehicleModel, &vehicleModelName, &vehicleRegistrationNo,
			&vehicleEngineType, &vehicleEngineTypeName, &vehiclePurchaseDate, &vehicleColor,
		 &vehicleInsuranceID,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error processing data"})
		}

		vehicles = append(vehicles, map[string]interface{}{
			"vehicle_id":              vehicleID,
			"insurance_id":            insuranceID,
			"vehicle_name":            vehicleName,
			"vehicle_manufacture_year": vehicleManufactureYear,
			"vehicle_warranty":        vehicleWarranty,
			"vehicle_make":            vehicleMake,
			"vehicle_make_name":       vehicleMakeName,
			"vehicle_model":           vehicleModel,
			"vehicle_model_name":      vehicleModelName,
			"vehicle_registration_no": vehicleRegistrationNo,
			"vehicle_engine_type":     vehicleEngineType,
			"vehicle_engine_type_name": vehicleEngineTypeName,
			"vehicle_purchase_date":   vehiclePurchaseDate,
			"vehicle_color":           vehicleColor,
			"vehicle_insurance_id":    vehicleInsuranceID,
		})
	}

	return c.JSON(vehicles)
}
