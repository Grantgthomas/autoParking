package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type permitOrder struct {
	permit_id string
	location  string
	car_id    string
}

func main() {
	var userOption string
	var run bool
	run = true
	fmt.Println("Welcome to Autoparking. Would you like to: \nMake Permits\nAdd Cars\nRegister Permit\nView Cars\nView Apartments\nQuit")

	database, err := sql.Open("sqlite3", "./autos.db")
	if err != nil {
		fmt.Println(err)
	}
	generateQueue(database)
	for run {
		fmt.Println("Please enter an option")
		fmt.Scanln(&userOption)
		switch userOption {
		case "m":
			generateQueue(database)
		case "ac":
			err := addCar(database)
			if err != nil {
				fmt.Println(err)
			}
		case "ap":
			addPermit(database)
		case "vc":
			err := viewEntry(database, "view cars")
			if err != nil {
				fmt.Println(err)
			}
		case "va":
			err := viewEntry(database, "view apartments")
			if err != nil {
				fmt.Println(err)
			}
		case "q":
			run = false
		default:
			println("Invalid Option")
			println("Valid Options Are:\n(m) make permits\n(ac) add car\n(ap) add permit\n(vc) view cars\n(va) view apartments\n(q) quit")
		}
	}

}
func viewEntry(database *sql.DB, option string) error {
	if option == "view cars" {
		var ownerName string
		var carID string
		var make string
		var model string
		var color string
		var plate string
		fmt.Println("Type an owner name to view their cars or type 'all' to view all cars")
		fmt.Scanln(&ownerName)
		if ownerName == "all" {
			rows, err := database.Query("SELECT CAST(car_id AS varchar),make,model,color,plate FROM cars")
			if err != nil {
				return err
			}
			for rows.Next() {
				rows.Scan(&carID, &make, &model, &color, &plate)
				fmt.Println("[" + carID + " | " + carID + " | " + make + " | " + model + " | " + color + " | " + plate + "]")
			}
		} else {
			query := `SELECT car_id,make,model,color,plate FROM cars WHERE ownerFname=$1;`
			rows, err := database.Query(query, ownerName)
			if err != nil {
				return err
			}
			for rows.Next() {
				rows.Scan(&carID, &make, &model, &color, &plate)
				fmt.Println("[" + carID + " | " + carID + " | " + make + " | " + model + " | " + color + " | " + plate + "]")
			}
		}

	} else if option == "view apartments" {
		var apartmentName string
		var apt_id string
		rows, err := database.Query("SELECT CAST(apt_id AS VARCHAR),apt_name FROM apartments WHERE apt_name NOT NULL")
		if err != nil {
			return err
		}
		for rows.Next() {
			rows.Scan(&apt_id, &apartmentName)
			fmt.Println("[" + apt_id + " | " + apartmentName + "]")
		}
	}
	return nil
}
func addPermit(database *sql.DB) error {
	var carID string
	var location string
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	viewEntry(database, "view cars")
	fmt.Println("Enter the ID of the car:")
	fmt.Scanln(&carID)
	fmt.Println("Enter the ID for the location of the permit:")
	fmt.Scanln(&location)
	query := `SELECT apt_name FROM apartments WHERE apt_id=$1`
	apt_name := database.QueryRow(query, location)
	apt_name.Scan(&location)
	insert, err := database.Prepare("INSERT INTO permits(car_id,active_time,location,active) VALUES(?,?,?,?)")
	fmt.Println("Registering car " + carID + " at " + location)
	insert.Exec(carID, currentTime, location, 1)
	if err != nil {
		return nil
	}

	return nil

}

func addCar(database *sql.DB) error {
	var make string
	var model string
	var color string
	var plate string
	var name string

	fmt.Println("Enter the make of your car:")
	fmt.Scanln(&make)
	fmt.Println("Enter the model of your car:")
	fmt.Scanln(&model)
	fmt.Println("Enter the color")
	fmt.Scanln(&color)
	fmt.Println("Enter .your car's plate:")
	fmt.Scanln(&plate)
	fmt.Println("Enter your first name:")
	fmt.Scanln(&name)
	validate := []string{make, model, color, plate, name}
	for n := range validate {
		if validate[n] == "" {
			fmt.Println(validate)
			fmt.Println("Option was null")
			return nil
		}
	}
	insert, err := database.Prepare("INSERT INTO cars (make,model,color,plate,ownerFname) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	insert.Exec(make, model, color, plate, name)
	return nil
}
func generateQueue(database *sql.DB) {
	permitQueue := make([]permitOrder, 0)
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	var permitTime string
	var permitID string
	var carID string
	var location string
	var newEntry permitOrder
	//check active permits to validate they are active
	//look to see if 24 hrs has passed
	rows, err := database.Query("SELECT permit_id,CAST(active_time AS varchar),car_id,location FROM permits WHERE active=1")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&permitID, &permitTime, &carID, &location)
		if check24hrs(permitTime, currentTime) {

			newEntry.permit_id = permitID
			newEntry.car_id = carID
			newEntry.location = location
			permitQueue = append(permitQueue, newEntry)
		}
	}
	if len(permitQueue) > 0 {
		executeQueue(permitQueue, database)
	}
}
func check24hrs(evalTime string, current string) bool {
	//check if 24 hours has passed
	//convert datetime values from SQL database to numerical
	evalMonth, _ := strconv.Atoi(evalTime[0:2])
	evalDay, _ := strconv.Atoi(evalTime[3:5])
	evalYear, _ := strconv.Atoi(evalTime[6:10])
	evalHour, _ := strconv.Atoi(evalTime[11:13])
	evalMinute, _ := strconv.Atoi(evalTime[14:16])
	evalSec, _ := strconv.Atoi(evalTime[17:])
	currentMonth, _ := strconv.Atoi(current[0:2])
	currentDay, _ := strconv.Atoi(current[3:5])
	currentYear, _ := strconv.Atoi(current[6:10])
	currentHour, _ := strconv.Atoi(current[11:13])
	currentMinute, _ := strconv.Atoi(current[14:16])
	currentSec, _ := strconv.Atoi(current[17:])
	//prepare values for comparison
	var evalTimePassed float64
	var currentTimePassed float64
	evalTimePassed = float64(86400 * evalDay)
	evalTimePassed = float64(3600*evalHour) + evalTimePassed + float64(60*evalMinute) + float64(evalSec)
	currentTimePassed = float64(86400 * currentDay)
	currentTimePassed = float64(3600*currentHour) + currentTimePassed + float64(60*currentMinute) + float64(currentSec)
	//any major values are not equal 24 hours has definetly passed
	if evalMonth != currentMonth || evalYear != currentYear {
		return true
	} else if float64(currentTimePassed-evalTimePassed) > float64(86400) {
		return true
	}
	return false
}

func executeQueue(permitQueue []permitOrder, database *sql.DB) {
	for _, order := range permitQueue {
		runAutoparking(order.location, order.car_id, order.permit_id, database)
	}
}

func runAutoparking(apartment string, carID string, permit_id string, database *sql.DB) {
	vMake, vModel, vColor, vPlate, vApt, vEmail := "--ma", "--mo", "--co", "--pl", "--apt", "--e"
	var queryMake string
	var queryModel string
	var queryColor string
	var queryPlate string
	var email string
	var email_id string
	//query data base to get vehichle info
	row, _ := database.Query("SELECT email_id,address FROM email WHERE email_id=1")
	row.Next()
	row.Scan(&email_id, &email)
	rows, _ := database.Query("SELECT make,model,color,plate FROM cars WHERE car_id=" + carID)
	rows.Next()
	rows.Scan(&queryMake, &queryModel, &queryColor, &queryPlate)

	//check to make sure a valid entry for a car was returned
	if queryMake == "" || queryModel == "" || queryColor == "" || queryPlate == "" {
		fmt.Println("Invalid Car choice")
	} else {
		//run autoparking script with args from db
		cmd := exec.Command("python", "autoParking.py", vMake, queryMake, vModel, queryModel, vColor, queryColor, vPlate, queryPlate, vApt, apartment, vEmail, email)
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
		}
		//set the permit to active after running auto parking
		//may add some additonal checking to see if this actually works later
		database.Exec("UPDATE permits SET active=1 WHERE permit_id=" + permit_id)
	}
}
