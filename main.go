package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type permitOrder struct {
	permit_id string
	location  string
	car_id    string
}

func main() {
	database, err := sql.Open("sqlite3", "./autos.db")
	if err != nil {
		fmt.Println(err)
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go autoRegister(database)
	menu(database)
	wg.Done()
}

func autoRegister(database *sql.DB) {
	for {
		generateQueue(database)
		fmt.Println("AutoParking!")
		time.Sleep(time.Second * 28800)
	}
}

func menu(database *sql.DB) {
	var userOption string
	var run bool
	run = true
	fmt.Println("Welcome to Autoparking. Would you like to: \nMake Permits\nAdd Cars\nRegister Permit\nView Permits\nView Cars\nView Apartments\nQuit")
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
		case "vp":
			err := viewEntry(database, "view perm")
			if err != nil {
				fmt.Println(err)
			}
		case "dp":
			err := viewEntry(database, "view perm")
			if err != nil {
				fmt.Println(err)
			}
			selectPermit(database)
		case "q":
			run = false
		default:
			println("Invalid Option")
			println("Valid Options Are:\n(m) make permits\n(ac) add car\n(ap) add permit\n(vc) view cars\n(vp) view permits\n(va) view apartments\n(q) quit")
		}
	}
}

func selectPermit(database *sql.DB) {
	stdin := bufio.NewReader(os.Stdin)
	//get user input for a permit entry to delete
	var delInput int = 0
	fmt.Println("Enter the ID of the permit to be deleted")
	//	fmt.Scanln(&delInput)
	//validates the input from the user
	for {
		_, err := fmt.Fscan(stdin, &delInput)
		if err == nil {
			if delInput > 0 {
				err := deletePermit(database, strconv.Itoa(delInput))
				if err != nil {
					fmt.Println(err)
				}
			} else if delInput <= 0 {
				delInput = 0
				fmt.Println("Please enter a valid ID (Greater than 0 and listed)")
				fmt.Println(delInput)
				break
			}

		}
		fmt.Println("Please enter a valid ID (Greater than 0 and listed)")
		break
	}
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
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	//	defer rows.Close()
	for rows.Next() {
		rows.Scan(&permitID, &permitTime, &carID, &location)
		if check24hrs(permitTime, currentTime) {

			newEntry.permit_id = permitID
			newEntry.car_id = carID
			newEntry.location = location
			permitQueue = append(permitQueue, newEntry)
		}
	}
	//	defer rows.Close()
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
	//validate data in the permit queue before using it
	for _, permitOrder := range permitQueue {
		if _, err := strconv.Atoi(permitOrder.car_id); err == nil {
			runAutoparking(permitOrder.location, permitOrder.car_id, permitOrder.permit_id, database)
		} else if err != nil {
			fmt.Printf("CarID number %q is invalid\n", permitOrder.car_id)
		}
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
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	//currentTime := time.Now().Format("01-02-2006 15:04:05")
	//query data base to get vehichle info
	//Must validate Car ID
	row, _ := database.Query("SELECT email_id,address FROM email WHERE email_id=1")
	row.Next()
	row.Scan(&email_id, &email)
	row.Close()
	//Can't update after this code block

	rows, _ := database.Query("SELECT make,model,color,plate FROM cars WHERE car_id=" + carID)
	rows.Next()
	rows.Scan(&queryMake, &queryModel, &queryColor, &queryPlate)
	rows.Close()

	//check to make sure a valid entry for a car was returned
	if queryMake == "" || queryModel == "" || queryColor == "" || queryPlate == "" {
		fmt.Println("Invalid Car choice")
	} else {
		//run autoparking script with args from db
		cmd := exec.Command("python", "autoParking.py", vMake, queryMake, vModel, queryModel, vColor, queryColor, vPlate, queryPlate, vApt, apartment, vEmail, email)
		update, err := database.Prepare("UPDATE permits SET active_time=?,car_id=?,active=1 WHERE permit_id=?")
		if err != nil {
			fmt.Println(err)
		}
		//update permit if it has not just been added
		if permit_id != "-3" {
			update.Exec(currentTime, carID, permit_id)
		}
		//set the permit to active after running auto parking
		//may add some additonal checking to see if this actually works later
		//update is broken
		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
		}
	}

}
