package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id    primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Email string `json:"email"`
}


type User struct {
	Name string `json:"name"`
	Email string `json:"email" bson:"_id"`
	Password string `json:"password"`
}

func (t *Todo) IsEmpty() bool {
	return t.Title == "" || t.Description == ""
}


type TravelData struct {
	Id primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Destination string `json:"destination"`
    TravelStartDate string `json:"start_date"`
    TravelEndDate string `json:"end_date"`
	Budget      float64 `json:"budget"`
	Interests   []string `json:"interests"`
	Activities  []string `json:"activities"`
	Email string `json:"email"`
}

func (t *TravelData) TravelInfo() bool{
	return t.Destination != "" && t.TravelStartDate != "" && t.Budget != 0 && t.TravelEndDate != ""
}


type TripDetails struct {
	Destination string   `json:"destination"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Budget      float64  `json:"budget"`
	Interests   []string `json:"interests"`
	Activities  []string `json:"activities"`
}

type Itinerary struct {
	Day          string   `json:"day"`
	Description  string   `json:"description"`
	Accommodation string  `json:"accommodation"`
	Attractions  []string `json:"attractions"`
	Activities   []string `json:"activities"`
	Dining       []string `json:"dining"`
}

type Response struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TripDetails *TripDetails `json:"trip_details"`
	Itinerary   *[]Itinerary  `json:"itinerary"`
	Email string `json:"email"`
	
}

