package main

import (
	"database/sql"
	"fmt"
	"time"
)

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
	insert.Exec(carID, currentTime, location, 1)
	if err != nil {
		return nil
	} else {

		runAutoparking(location, carID, "-3", database)
		fmt.Println("Registering car " + carID + " at " + location)
		return nil
	}
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

func deletePermit(database *sql.DB, permit_id string) error {
	update, err := database.Prepare("DELETE FROM permits WHERE permit_id=?")
	if err != nil {
		return err
	}
	update.Exec(permit_id)
	return nil
}

func replacePermit(carID string, location string, database *sql.DB) {
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	insert, err := database.Prepare("INSERT INTO permits(car_id,active_time,location,active) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Registering car " + carID + " at " + location)
	insert.Exec(carID, currentTime, location, 1)
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
			defer rows.Close()
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
			defer rows.Close()
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
		defer rows.Close()
		if err != nil {
			return err
		}
		for rows.Next() {
			rows.Scan(&apt_id, &apartmentName)
			fmt.Println("[" + apt_id + " | " + apartmentName + "]")
		}
	} else if option == "view perm" {
		var permit_id string
		var car_id string
		var active_time string
		var location string
		var active string
		rows, err := database.Query("SELECT CAST(permit_id AS VARCHAR),car_id,CAST(active_time AS VARCHAR),location,active FROM permits")
		defer rows.Close()
		if err != nil {
			return err
		}
		for rows.Next() {
			rows.Scan(&permit_id, &car_id, &active_time, &location, &active)
			switch active {
			case "1":
				active = "Active"
			case "0":
				active = "Inactive"
			}
			fmt.Println("[ " + permit_id + " | " + car_id + " | " + active_time + " | " + location + " | " + active + " ]")
		}
	}
	return nil
}
