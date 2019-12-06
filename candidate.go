package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"encoding/json"
	"github.com/gorilla/mux"
)


type Candidate struct {
	gorm.Model
	candidate_id string
	candidate_name string
	candidate_phonenumber string
	status string
}


func AllCandidates(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Candidates")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=golang dbname=startetelelogic_interview password=password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	var candidates []Candidate
	db.Find(&candidates)
	json.NewEncoder(w).Encode(candidates)
}


func NewCandidate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Candidate")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=golang dbname=startetelelogic_interview password=password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	phone := vars["phone"]
	status := vars["status"]

	db.Create(&Candidate{candidate_name: name, candidate_phonenumber: phone, status: status})

	fmt.Fprintf(w, "New Candidate Created")
}


func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Candidate")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=golang dbname=startetelelogic_interview password=password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	phone := vars["phone"]

	var candidate Candidate

	db.Where("candidate_phonenumber = ?", phone).Find(&candidate)
	db.Delete(&candidate)

	fmt.Fprintf(w, "Candidate Deleted")
}

func UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Candidate")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=golang dbname=startetelelogic_interview password=password sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	phone := vars["phone"]
	status := vars["status"]

	var candidate Candidate

	db.Where("candidate_phonenumber = ?", phone).Find(&candidate)

	candidate.status = status

	db.Save(&candidate)

	fmt.Fprintf(w, "Candidate Updated")
}