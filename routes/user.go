package routes

import (
	"errors"
	models "fiber/Models"
	"fiber/database"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	// THis is the model of the user this is a serializer

	ID         uint   `json:"id"`
	First_Name string `json:"firstname"`

	Last_Name string `json:"lastname"`
}

func CreateResponseUser(usermodel models.User) User {
	return User{ID: usermodel.Id, First_Name: usermodel.First_Name, Last_Name: usermodel.Last_Name}

}

func CretaeUser(c *fiber.Ctx) error{
	var user models.User
	if err:=c.BodyParser(&user); err!=nil{
		return c.Status(400).JSON(err.Error())


	}  

	database.Database.DB.Create(&user)
	responseUser:=CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)




}
func GetUsers(c *fiber.Ctx)error{
users:=[] models.User{}
database.Database.DB.Find(&users)
responseUsers:=[] User{}
for _,user:=range users{
	responseUser:=CreateResponseUser(user)
	responseUsers=append(responseUsers,responseUser)

}
return c.Status(200).JSON(responseUsers)

}
func findUser(id int,user *models.User) error{
	database.Database.DB.Find(&user,"id=?",id)
	if user.Id==0 {
		return errors.New("user does not exist")

		
	}
return  nil



}
func Getuser(c *fiber.Ctx) error{
	id,err:=c.ParamsInt("id")
	var user models.User
	if err!=nil {
		return c.Status(400).JSON("PleaseEnsure that id in an Integar")
		
	}
	if err:=findUser(id,&user);err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	responseUser:=CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)



}