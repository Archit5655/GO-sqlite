package routes

import (
	// "crypto/internal/e/dwards25519/field"
	"errors"
	models "fiber/Models"
	"fiber/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// {
// 	id:1,
// 	user:{
// 		id:23,
// 		first_name:"Marley",
// 		last_name:"bOg"
// 	},
// 	products:{
// 		id:24,
// 		name:"MAckvoob",
// 		serialNuimber:"132"
// 	}
// }

type Order struct {
	ID         uint      `josn:"id"`
	User       User      `josn:"user"`
	Product    Product   `josn:"product"`
	Created_At time.Time `josn:"order_date"`
}

func CreateResponseOrder(order models.Orders, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, Created_At: order.Created_At}

}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Orders
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	var product models.Products
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Create(&order)
	responseUser := CreateResponseUser(user)
	responseporduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseporduct)
	return c.Status(200).JSON(responseOrder)

}
func GetOrders(c *fiber.Ctx) error {
	orders := []models.Orders{}
	database.Database.DB.Find(&orders)
	resposeOrders := []Order{}
	for _, order := range orders {
		var user models.User
		var product models.Products
		database.Database.DB.Find(&user, "id=?", order.UserRefer)
		database.Database.DB.Find(&product, "id=?", order.ProductRefer)
		reponseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		resposeOrders = append(resposeOrders, reponseOrder)

	}
	return c.Status(200).JSON(resposeOrders)

}

func findOrder(id int, order *models.Orders) error {
	database.Database.DB.Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("Order does not exists")

	}
	return nil

}

func GetOrderById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Orders
	if err != nil {

		return c.Status(400).JSON("PleaseEnsure that id in an Integar")

	}
	if err:=findOrder(id,&order);err!=nil{
	
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	var product models.Products
	database.Database.DB.First(&user,order.UserRefer)
	database.Database.DB.First(&product,order.ProductRefer)
	reponseUser:=CreateResponseUser(user)
	reponseProduct:=CreateResponseProduct(product)
	responseOrder:=CreateResponseOrder(order,reponseUser,reponseProduct)
	

	return c.Status(200).JSON(responseOrder)
	




}
