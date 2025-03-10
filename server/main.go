// main.go
// Entry point for the server

// Reference:
// https://thenewstack.io/make-a-restful-json-api-go/

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"encoding/json"
)

type Organization struct {
	OrgID int `json:"orgID"`
	OrgName string `json:"orgName"`
	NumberOfXWorkers int `json:"numberOfXWorkers"`
}

type Organizations []Organization

type User struct {
	UserID int
	Username string
	FirstName string
	LastName string
	Title string
	Image int
	SecurityGroup int
}

type Users []User

type XWorker struct {
	XWorkerID int
	XWorkerName string
	Title string
	Image int
	PermissionGroup int
}

type XWorkers []XWorker

type Role struct {
	RoleID int
	RoleName string
}

type Roles []Role

type UserMessage struct {
	UserMessageID int	
	Sender int
	Recipient int
	Content string
}

type UserMessages []UserMessage

type Job struct {
	JobID int
	XWorkerID int
	DateAssigned time.Time
	DateUpdated time.Time
	DateScheduled time.Time
	IsRecurring bool
	NextScheduledRun time.Time
	DateCompleted time.Time
	IsPast bool
	IsInProgress bool
	PercentDone int
}

type Jobs []Job

type Task struct {
	TaskID int	
	JobID int
	TaskType string
	SequenceNumber int
}

type Tasks []Task

type Action struct {
 	ActionID int
 	ActionType string
 	PermissionNeeded int
}

type Actions []Action

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Root)

	router.HandleFunc("/organizations", OrganizationsRoot)
	router.HandleFunc("/organizations/{orgID}", OrganizationByID)

	router.HandleFunc("/users", UsersRoot)
	router.HandleFunc("/users/{userID}", UserByID)

	router.HandleFunc("/xworkers", XWorkersRoot)
	router.HandleFunc("/xworkers/{xworkerID}", XWorkerByID)

	router.HandleFunc("/roles", RolesRoot)
	router.HandleFunc("/roles/{roleID}", RoleByID)

	router.HandleFunc("/usermessages", UserMessagesRoot)
	router.HandleFunc("/usermessages/{usermessageID}", UserMessageByID)

	router.HandleFunc("/jobs", JobsRoot)
	router.HandleFunc("/jobs/{jobID}", JobByID)

	router.HandleFunc("/tasks", TasksRoot)
	router.HandleFunc("/tasks/{taskID}", TaskByID)

	router.HandleFunc("/actions", ActionsRoot)
	router.HandleFunc("/actions/{actionID}", ActionByID)

	fmt.Println("XANTHA Server started on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Root(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "XANTHA API Root")
}

func OrganizationsRoot(w http.ResponseWriter, r *http.Request){
	organizations := Organizations{
		Organization{OrgName: "IBM"},
		Organization{OrgName: "Microsoft"},
	}
	json.NewEncoder(w).Encode(organizations)
	//fmt.Fprintln(w, "Organizations Root")
}

func OrganizationByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	orgID := vars["orgID"]
	fmt.Fprintln(w, "Organization: ", orgID)
}

func UsersRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Users Root")
}

func UserByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userID := vars["userID"]
	fmt.Fprintln(w, "User: ", userID)
}

func XWorkersRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "XWorkers Root")
}

func XWorkerByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	xworkerID := vars["xworkerID"]
	fmt.Fprintln(w, "XWorker: ", xworkerID)
}

func RolesRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Roles Root")
}

func RoleByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	roleID := vars["roleID"]
	fmt.Fprintln(w, "Role: ", roleID)
}

func UserMessagesRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "UserMessages Root")
}

func UserMessageByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	usermessageID := vars["usermessageID"]
	fmt.Fprintln(w, "UserMessage: ", usermessageID)
}

func JobsRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Jobs Root")
}

func JobByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	jobID := vars["jobID"]
	fmt.Fprintln(w, "Job: ", jobID)
}

func TasksRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Tasks Root")
}

func TaskByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID := vars["taskID"]
	fmt.Fprintln(w, "Task: ", taskID)
}

func ActionsRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Actions Root")
}

func ActionByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	actionID := vars["actionID"]
	fmt.Fprintln(w, "Action: ", actionID)
}
