package controller

import (
	"fmt"
	"log"
	"rest-api/config"
	"rest-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllItems(ctx *fiber.Ctx) error{
	db := config.GetDb()
	coll:=db.Client.Database("rest-api").Collection("item")
	cursor, err := coll.Find(db.Ctx,bson.M{})
	if(err!=nil){
		log.Fatal(err)
	}
	return ctx.Status(200).JSON(cursor)
}

func PostItem(ctx *fiber.Ctx) error{
	db:=config.GetDb()
	coll:=db.Client.Database("rest-api").Collection("item")

	data:= new(models.Item)
	if err := ctx.BodyParser(data); err != nil {
		fmt.Println(err)
		return ctx.SendStatus(400)
	}
	res,err:=coll.InsertOne(db.Ctx,data)
	if(err!=nil){
		fmt.Println(err)
		return ctx.SendStatus(400)
	}

	return ctx.Status(200).JSON(res)
}