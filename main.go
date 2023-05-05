package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	m "./models"
	s "./services/user.services"

	"io/ioutil"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := s.Read()

	if err != nil {
		fmt.Print("Se ha presntado un error en la consulta de usuarios")
	}

	if len(users) == 0 {
		fmt.Print("La consulta no retorno datos")
	} else {
		fmt.Print("La lectura finalizo con exito")
	}
	json.NewEncoder(w).Encode(users)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API HelloGO")
}

func createUsers(w http.ResponseWriter, r *http.Request) {

	var user m.User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid user")
	}

	json.Unmarshal(reqBody, &user)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err1 := s.Create(user)

	if err1 != nil {
		fmt.Print("No pudo crearse el usuario")
	} else {
		fmt.Print("Se creo con exito el usuario")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func updateUsers(w http.ResponseWriter, r *http.Request) {

	var user m.User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Update a valid user")
	}

	json.Unmarshal(reqBody, &user)

	err1 := s.Update(user, user.ID.Hex())

	if err1 != nil {
		fmt.Print("No pudo actualizarse el usuario")
	} else {
		fmt.Print("Se actualizo el usuario con exito")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	var user m.User
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "delete a valid user")
	}

	json.Unmarshal(reqBody, &user)

	err1 := s.Delete(user.ID.Hex())

	if err1 != nil {
		fmt.Print("No pudo eliminarse el usuario")
	} else {
		fmt.Print("Se elimino el usuario con exito")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUsers).Methods("POST")
	router.HandleFunc("/users", updateUsers).Methods("PUT")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
	//fmt.Println("conexion a MongoDB")
}
