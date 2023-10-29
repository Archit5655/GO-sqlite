package routes

import (
	"errors"
	models "fiber/Models"
	"fiber/database"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Products) Product {
	return Product{Id:productModel.ID,Name: productModel.Name,SerialNumber: productModel.SerialNumber}


}
func CreteProduct(c *fiber.Ctx) error{

var  product models.Products
if err:=c.BodyParser(&product);err!=nil{
	return c.Status(400).JSON(err.Error())
}
database.Database.DB.Create(&product)
respnseProduct:=CreateResponseProduct(product)
return c.Status(200).JSON(respnseProduct)
}


func  GetProducts(c *fiber.Ctx)error  {
	products:=[] models.Products{}
	database.Database.DB.Find(&products)
	responseProducts:=[] Product{}
	for _,product:= range products{
		responseProduct:=CreateResponseProduct(product)
		responseProducts=append(responseProducts,responseProduct)

	}

	return c.Status(200).JSON(responseProducts)

	
}
func findProduct(id int, product *models.Products) error {
	database.Database.DB.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("user does not exist")

	}
	return nil

}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Products
	if err != nil {
		return c.Status(400).JSON("PleaseEnsure that id in an Integar")

	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	reponseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(reponseProduct)

}


func UpdateProduct( c *fiber.Ctx) error{

	id, err := c.ParamsInt("id")
	
	var product models.Products
	if err != nil {
		return c.Status(400).JSON("PleaseEnsure that id in an Integar")

	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
		type UpdateProduct struct{
			Name string `json:"name"`
			SerialNumber string `json:"serial_number"`
		}
		var updateData UpdateProduct
		if err:=c.BodyParser(&updateData);err!=nil{
			return c.Status(500).JSON(err.Error())

		}
		product.Name=updateData.Name
		product.SerialNumber=updateData.SerialNumber
		database.Database.DB.Save(&product)
		reponseProduct:=CreateResponseProduct(product)
		return c.Status(200).JSON(reponseProduct)


}
func DeleteProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var product models.Products
	if err != nil {
		return c.Status(400).JSON("PleaseEnsure that id in an Integar")

	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err:=database.Database.DB.Delete(&product).Error;err!=nil{
		return c.Status(400).JSON(err.Error())

	}
	return c.Status(200).JSON("Successfully deleted the user")
	
	
}