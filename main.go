package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"net/http"
	"strings"
	"strconv"
	"github.com/gorilla/mux"
)

type book struct {
	UserID		string `json:"UserID"`
	BookID      string `json:"BookID"`
	dateOfIssue string 
	currentIssuee string 
	rating 		float64 `json:"rating"`
	noOfIssues 	float64	
	Title       string `json:"Title"`
	noOfRaters	float64
}

type bookAvailablity struct {
	Description		string `json:"Description"`
}

type allEvents []book

var events = allEvents{
	{
		UserID:		"1234",	
		BookID:       "1",
		Title:       "book 1",
		dateOfIssue:       "NA",
		currentIssuee:       "NA",
		rating:       0.0,
		noOfIssues:       0,
		noOfRaters: 0,
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome User!")

}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newEvent book
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter the data in the specified format ")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	if strings.Compare(newEvent.UserID,"1234") == 0{
		newEvent.currentIssuee = "NA"
		newEvent.noOfIssues = 0
		newEvent.noOfRaters = 0
		newEvent.rating = 0.0	
		events = append(events, newEvent)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newEvent)
	}
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var check bookAvailablity

	for _, singleEvent := range events {
		if singleEvent.BookID == eventID  {
			if strings.Compare(singleEvent.currentIssuee,"NA")==0 {
				check.Description = "Available!"
			} else{
				check.Description = "Issued to - "+singleEvent.currentIssuee
			}
		}
	}
	json.NewEncoder(w).Encode(check)
}

func rateBook(w http.ResponseWriter, r *http.Request) {
	
	eventID := mux.Vars(r)["id"]

	var newRating bookAvailablity

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter the data in the specified format ")
	}
			
	json.Unmarshal(reqBody, &newRating)

	for i, singleEvent := range events {
		if singleEvent.BookID == eventID  {

			s, _ := strconv.ParseFloat(newRating.Description, 64)
			singleEvent.rating = (singleEvent.rating*singleEvent.noOfRaters + s)/(singleEvent.noOfRaters+1)
			singleEvent.noOfRaters++
			events = append(events[:i], singleEvent)
			r1 := fmt.Sprintf("%f", singleEvent.rating)
			newRating.Description = "New Rating : "+r1
			json.NewEncoder(w).Encode(newRating)
		}
	}
	
}

func issueBook(w http.ResponseWriter, r *http.Request) {
	
	eventID := mux.Vars(r)["id"]

	var newRating bookAvailablity

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter the data in the specified format ")
	}
			
	json.Unmarshal(reqBody, &newRating)

	for i, singleEvent := range events {
		if singleEvent.BookID == eventID  {
			if strings.Compare("NA",singleEvent.currentIssuee) == 0{
				dt := time.Now()
				singleEvent.currentIssuee = newRating.Description
				singleEvent.noOfIssues++
				singleEvent.dateOfIssue = dt.String()

				events = append(events[:i], singleEvent)
				newRating.Description = "Book Issued !"
				json.NewEncoder(w).Encode(newRating)
				} else{
					newRating.Description = "Book Unavailable !"
					json.NewEncoder(w).Encode(newRating)
				}
			
		}
	}
	
}

func returnBook(w http.ResponseWriter, r *http.Request) {
	
	eventID := mux.Vars(r)["id"]

	var newRating bookAvailablity

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter the data in the specified format ")
	}
			
	json.Unmarshal(reqBody, &newRating)

	for i, singleEvent := range events {
		if singleEvent.BookID == eventID  {
			if strings.Compare(singleEvent.currentIssuee,newRating.Description) == 0 {
				singleEvent.currentIssuee = "NA"
				singleEvent.dateOfIssue = "NA"
				events = append(events[:i], singleEvent)
				newRating.Description = "Book Returned !"
				json.NewEncoder(w).Encode(newRating)
			} else{
				newRating.Description = "Unauthorized to return !"
				json.NewEncoder(w).Encode(newRating)	
			}
		}
	}
	
}

func getRatings(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var check bookAvailablity

	for _, singleEvent := range events {
		if singleEvent.BookID == eventID  {
			s := fmt.Sprintf("%f", singleEvent.rating)
			if strings.Compare(singleEvent.currentIssuee,"NA")==0 {
				check.Description = "Available! Rated as "+s
			} else{
				check.Description = "Unavailable now! Rated as "+s
			}
		}
	}
	json.NewEncoder(w).Encode(check)
}

func getPopular(w http.ResponseWriter, r *http.Request) {

	max := 0.0
	var popularEvent book
	for _, singleEvent := range events {
		if singleEvent.rating > max  {
			max = singleEvent.rating
			popularEvent = singleEvent
		}
	}
	json.NewEncoder(w).Encode(popularEvent)
}

func getIssued(w http.ResponseWriter, r *http.Request) {

	max := 0.0
	var issuedEvent book
	for _, singleEvent := range events {
		if singleEvent.noOfIssues > max  {
			max = singleEvent.noOfIssues
			issuedEvent = singleEvent
		}
	}
	json.NewEncoder(w).Encode(issuedEvent)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.BookID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/addBook", addBook).Methods("POST")
	router.HandleFunc("/books", getAllEvents).Methods("GET")
	router.HandleFunc("/books/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/status/{id}", getStatus).Methods("GET")
	router.HandleFunc("/ratings/{id}", getRatings).Methods("GET")
	router.HandleFunc("/issue/{id}", issueBook).Methods("POST")
	router.HandleFunc("/return/{id}", returnBook).Methods("POST")
	router.HandleFunc("/mostPopular", getPopular).Methods("GET")
	router.HandleFunc("/mostIssued", getIssued).Methods("GET")
	router.HandleFunc("/rate/{id}", rateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}