package utils

import (
	"fmt"

	"github.com/Kalyug5/just-goo/model"
)

func GenerativePrompt(travelData model.TravelData) string {
	prompt := fmt.Sprintf(`Generate a personalized travel itinerary for a trip to %s from %s to %s. The budget for this trip is $%.2f. The traveler is interested in %s and wants to include %s. Provide recommendations for accommodations, popular attractions to visit, activities to do, famous landmarks to see, and local dining experiences and always make sure that the full itinerary should come under budget. Always and STRICTLY provide the response in the following JSON format ONLY:

{
  "trip_details": {
    "destination": "%s",
    "start_date": "%s",
    "end_date": "%s",
    "budget": %.2f,
    "interests": [
      %s
    ],
    "activities": [
      %s
    ]
  },
  "itinerary": [
    {
      "day": "Day 1: [Activity Title]",
      "description": "[Description of the day's activities]",
      "accommodation": "[Accommodation options with average cost per night]",
      "attractions": [
        "[Attraction 1,with average cost to roam it]",
        "[Attraction 2,with average cost to roam it]",
        "[Attraction 3,with average cost to roam it]"
      ],
      "activities": [
        "[Activity 1,with average cost to do it]",
        "[Activity 2,with average cost to do it]",
        "[Activity 3,with average cost to roam it]"
      ],
      "dining": [
        "[Dining option 1,with average cost to eat there]",
        "[Dining option 2,with average cost to eat there]",
        "[Dining option 3,with average cost to eat there]"
      ]
    },
    {
      "day": "Day 2: [Activity Title]",
      "description": "[Description of the day's activities]",
      "accommodation": "[Accommodation options with average cost per night]",
     "attractions": [
        "[Attraction 1,with average cost to roam it]",
        "[Attraction 2,with average cost to roam it]",
        "[Attraction 3,with average cost to roam it]"
      ],
      "activities": [
        "[Activity 1,with average cost to do it]",
        "[Activity 2,with average cost to do it]",
        "[Activity 3,with average cost to roam it]"
      ],
      "dining": [
        "[Dining option 1,with average cost to eat there]",
        "[Dining option 2,with average cost to eat there]",
        "[Dining option 3,with average cost to eat there]"
      ]
    },
	//generate more as per requirement
    
  ]
}`,
		travelData.Destination, travelData.TravelStartDate,travelData.TravelEndDate, travelData.Budget, travelData.Interests, travelData.Activities,
		travelData.Destination, travelData.TravelStartDate, travelData.TravelEndDate, travelData.Budget, travelData.Interests, travelData.Activities,
	)


  


	return prompt

}