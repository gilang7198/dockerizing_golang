package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func returnVersion(w http.ResponseWriter, r *http.Request) {
	var versions Versions
	var arrVersions []Versions
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM version")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&versions.ID, &versions.NameApps, &versions.Version); err != nil {
			log.Fatal(err.Error())
		} else {
			arrVersions = append(arrVersions, versions)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrVersions

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func inserUser(w http.ResponseWriter, r *http.Request) {
// 	// var users Users
// 	// var arr_user []Users
// 	var response Response

// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}

// 	userID := r.FormValue("user_id")
// 	userFullname := r.FormValue("user_fullname")
// 	userAddress := r.FormValue("user_address")
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	_, err = db.Exec("INSERT INTO users (user_id, user_fullname, user_address, username, password) VALUES (?,?,?,?,?)",
// 		userID,
// 		userFullname,
// 		userAddress,
// 		username,
// 		password,
// 	)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	response.Status = 200
// 	response.Message = "Success Add User"
// 	log.Print("Insert data to database")

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func updateUsers(w http.ResponseWriter, r *http.Request) {
// 	var response Response

// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}

// 	userID := r.FormValue("user_id")
// 	userFullname := r.FormValue("user_fullname")
// 	userAddress := r.FormValue("user_address")
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	_, err = db.Exec("UPDATE users SET user_fullname = ?, user_address = ?, username = ?, password = ? WHERE user_id = ?",
// 		userFullname,
// 		userAddress,
// 		username,
// 		password,
// 		userID,
// 	)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	response.Status = 200
// 	response.Message = "Success Update Data"
// 	log.Print("Update data to database")

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func deleteUsers(w http.ResponseWriter, r *http.Request) {
// 	var response Response

// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}

// 	userID := r.FormValue("user_id")

// 	_, err = db.Exec("DELETE FROM users WHERE user_id = ?",
// 		userID,
// 	)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	response.Status = 200
// 	response.Message = "Success delete data"
// 	log.Print("Delete data database")

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/version", returnVersion).Methods("GET")
	// router.HandleFunc("/users", inserUser).Methods("POST")
	// router.HandleFunc("/users", updateUsers).Methods("PUT")
	// router.HandleFunc("/users", deleteUsers).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 9091")
	log.Fatal(http.ListenAndServe(":9091", router))
}
