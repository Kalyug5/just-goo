package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kalyug5/just-goo/api"
	"github.com/Kalyug5/just-goo/model"
	"github.com/Kalyug5/just-goo/utils"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




var collection *mongo.Collection
var dbName="Todos"
var collectionName="my_todos"

var collectionUser *mongo.Collection
var collectionNameUser="my_users"

var collectionTrip *mongo.Collection
var collectionNameTrip="my_trips"



func Home(c *fiber.Ctx) error {
	
	return c.Status(200).JSON(fiber.Map{"msg": "Hi i am Priyanshu Mishra, and welcome to my backend "})
	
}

func init(){
	if os.Getenv("DB_URI") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	DB_URI := os.Getenv("DB_URI")
	if DB_URI == "" {
		log.Fatal("DB_URI environment variable not set")
	}
	
	clientOption :=options.Client().ApplyURI(DB_URI)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB Connected successfully")	
	collection = client.Database(dbName).Collection(collectionName)
	collectionUser=client.Database(dbName).Collection(collectionNameUser)
	collectionTrip=client.Database(dbName).Collection(collectionNameTrip)
}

func insertOneTrip(trip model.Response,email string){
	trip.Email=email
	_,err := collectionTrip.InsertOne(context.Background(),trip)
	if err!=nil{
		log.Fatal(err)
	}

}

func getOneTrip(tripid string,email string) model.Response {
	var trip model.Response
	id,err := primitive.ObjectIDFromHex(tripid)
	if err!=nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id":id,"email":email}
	if err := collectionTrip.FindOne(context.Background(),filter).Decode(&trip); err!=nil{
		log.Fatal(err)
	}
	return trip

}

func getAllTrip(email string) []model.Response {
	var trips []model.Response
	filter := bson.M{"email":email}
	curr,err := collectionTrip.Find(context.Background(),filter)
	if err!=nil{
		return trips
	}
	for curr.Next(context.Background()) {
		var trip model.Response
		if err:= curr.Decode(&trip); err !=nil{
			log.Fatal(err)
		}
		trips = append(trips,trip)
	}
	// Close the cursor once finished
	defer curr.Close(context.Background())

	return trips
}

func deleteOneTrip(id string) {
	tripid,_ :=primitive.ObjectIDFromHex(id)

	filters := bson.M{"_id":tripid}
	_,err := collectionTrip.DeleteOne(context.Background(),filters)
	if err != nil {
		log.Fatal(err)
	}
}

func insertOneTodo(todo model.Todo) {
	_,err := collection.InsertOne(context.TODO(),todo)
	if err != nil {
		log.Fatal(err)
	}

}

func updateOneTodo(todoid string){
	id,err := primitive.ObjectIDFromHex(todoid)
	if err != nil {
		log.Fatal(err)
	}

	filters := bson.M{"_id":id}
	_,err = collection.UpdateOne(context.Background(),filters,bson.M{"$set":bson.M{"completed":true}})
	if err != nil {
		log.Fatal(err)
	}
}

func getOneTodo(todoid string,email string) model.Todo {
	id,err := primitive.ObjectIDFromHex(todoid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id":id,"email":email}
	var todo model.Todo
	err = collection.FindOne(context.Background(),filter).Decode(&todo)
	if err != nil {
		return todo
	}

	return todo


}

func getAllTodo(email string) []model.Todo {
	var todos []model.Todo
	filter := bson.M{"email":email}
	curr,err := collection.Find(context.Background(),filter)
	if err != nil {
		return todos
	}

	for curr.Next(context.Background()){
		var todo model.Todo
		if err := curr.Decode(&todo); err!=nil{
			log.Fatal(err)
		}
		todos = append(todos,todo)


	}

	defer curr.Close(context.Background())

	return todos
}

func deleteTodo(todoid string) {
	id,err := primitive.ObjectIDFromHex(todoid)
	if err != nil {
		log.Fatal(err)
	}

	_,error := collection.DeleteOne(context.Background(),bson.M{"_id":id})

	if error != nil {
		log.Fatal(error)
	}

}

func deleteTodos(){
	_,error := collection.DeleteMany(context.Background(),bson.M{})
	if error != nil {
		log.Fatal(error)
	}
}

func createUser(user model.User){
	_,err := collectionUser.InsertOne(context.Background(),user)
	if err != nil {
		log.Fatal(err)
	}
}

func loginUser(email string) model.User {
	var user model.User
	err := collectionUser.FindOne(context.Background(),bson.M{"_id":email}).Decode(&user)
	if err != nil {
		return user
	}

	return user
}


//controllers
func CreateTodo(c *fiber.Ctx) error {
	var todo model.Todo

	err := c.BodyParser(&todo)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	if todo.IsEmpty() {
		return c.Status(400).JSON(fiber.Map{"error":"todo is empty"})
	}

	insertOneTodo(todo)

	return c.Status(201).JSON(fiber.Map{"message":"Todo is created successfully","status":201})

}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	
	updateOneTodo(id)

	

return c.Status(200).JSON(fiber.Map{"success":"Todo is updated"})
}


func GetOneTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var email model.Todo
	if err := c.BodyParser(&email); err != nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	todo := getOneTodo(id,email.Email)

	return c.Status(200).JSON(fiber.Map{"data":todo,"status":200})
}


func GetTodos(c *fiber.Ctx) error {
	
	var input struct{
		Email string `json:"email"`
	}
	if err := c.BodyParser(&input); err != nil{
		return c.Status(400).JSON(fiber.Map{"error":err.Error()})
	}
	
	todos := getAllTodo(input.Email)

	return c.Status(200).JSON(fiber.Map{"data":todos,"status":200})
}


func DeleteOneTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	deleteTodo(id)
	return c.Status(200).JSON(fiber.Map{"success":"Todo is deleted"})
}

func DeleteTodos(c *fiber.Ctx) error {
	deleteTodos()
	return c.Status(200).JSON(fiber.Map{"success":"All Todos is deleted"})
}

func CreateTravelIternery(c *fiber.Ctx) error{
	var travelDetail model.TravelData
	var travel model.Response
	err := c.BodyParser(&travelDetail)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message":"Invalid request","status":400})
	}

	if !travelDetail.TravelInfo() {
		return c.Status(400).JSON(fiber.Map{"error":"Travel Detail is not sufficient","status":400})
	}

	model := api.Api()

	resp,error := model.GenerateContent(context.Background(),genai.Text(utils.GenerativePrompt(travelDetail)))
	if error != nil {
		return c.Status(400).JSON(fiber.Map{"error":error,"status":400})
	}
	
	jsonData,_ := json.Marshal(resp)

	
	
	var generatedResp api.ContentResponse
	_ = json.Unmarshal(jsonData,&generatedResp)
	for _,cad := range *generatedResp.Candidates{
		if cad.Content!=nil{
			for _,part := range cad.Content.Parts {
				err := json.Unmarshal([]byte(part), &travel)
				if err != nil {
					return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse itinerary", "status": http.StatusInternalServerError})
				}
			}
		}
	}

	insertOneTrip(travel,travelDetail.Email)


	return c.Status(201).JSON(fiber.Map{"response":travel,"status":201})

}


func GetTrip(c *fiber.Ctx) error {
	var travel model.Response
	var data struct{
		Email string `json:"email"`
		Id string `json:"id"`
	}
	if err:=c.BodyParser(&data); err!=nil{
		return c.Status(400).JSON(fiber.Map{"error":"Invalid request","status":400})
	}

	travel=getOneTrip(data.Id,data.Email)
	return c.Status(200).JSON(fiber.Map{"response":travel,"status":200})
}

func GetAllTrip(c *fiber.Ctx) error {
	var travels []model.Response
	var data struct{
		Email string `json:"email"`
		}
		if err:=c.BodyParser(&data); err!=nil{
			return c.Status(400).JSON(fiber.Map{"error":"Invalid request","status":400})
			}
			travels=getAllTrip(data.Email)
			return c.Status(200).JSON(fiber.Map{"response":travels,"status":200})

}

func DeleteOneTrip(c *fiber.Ctx) error {
	id := c.Params("id")
	deleteOneTrip(id)

	return c.Status(200).JSON(fiber.Map{
		"response": "Trip deleted successfully",
		"status": 200,
	})
}


func Register(c *fiber.Ctx) error {
	var user model.User

	if err:=c.BodyParser(&user); err!=nil{
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error":"Invalid request","status":400})
	}

	password,_ := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	user.Password=string(password)
	createUser(user)
	

	return c.Status(201).JSON(fiber.Map{
		"data":"User registered successfully",
		"status":201,
		"user":user,
	})
}


func Login(c *fiber.Ctx) error {

	if os.Getenv("SECRET_KEY") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}


	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable not set")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	var user model.User
	if err:=c.BodyParser(&user); err!=nil{
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error":"Invalid request","status":400})
	}
	// Check if user exists
	userFound:=loginUser(user.Email)
	
	if userFound.Email==""{
		return c.Status(400).JSON(fiber.Map{
			"error":"Invalid email",
			"status":400,
		})
	}

	if err:=bcrypt.CompareHashAndPassword([]byte(userFound.Password),[]byte(user.Password)); err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid credentials","status":fiber.StatusUnauthorized})
	}

	
	claim := &jwt.RegisteredClaims{
		ExpiresAt:jwt.NewNumericDate(time.Now().Add(time.Hour*1)),
		Issuer:userFound.Email,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

	token,err:=claims.SignedString([]byte(secretKey))
	
	if err!=nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Auth failed","status":fiber.StatusUnauthorized})
	}


	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		Path:     "/",
		Domain:   "travelhat.onrender.com",
		HTTPOnly: false,

        SameSite: "None",             
	}
	c.Cookie(&cookie)


	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data":"Login successful","status":200,"token":token})

}

func User(c *fiber.Ctx) error {
	if os.Getenv("SECRET_KEY") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}


	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable not set")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	cookie := c.Cookies("token")

	

	token,err := jwt.ParseWithClaims(cookie,&jwt.RegisteredClaims{},func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey),nil
	})

	if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"UnAuthroized","status":fiber.StatusUnauthorized})
	}

	claims,_ :=token.Claims.GetIssuer()

	return c.Status(200).JSON(fiber.Map{
		
		"status":200,
		"data":claims,
	})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1*time.Hour),
		Path:     "/",
		Domain:   "travelhat.onrender.com",
		HTTPOnly: false,
        SameSite: "None",   
	}
		c.Cookie(&cookie)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data":"Logout successful","status":200})

}
