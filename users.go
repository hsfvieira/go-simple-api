package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"regexp"
	"strconv"
	"errors"
)

type user struct {
    Nome	string	`json:"nome"`
	Idade	int		`json:"idade"`
}

const db = "./database.json"

func handleUsers(w http.ResponseWriter, req *http.Request) {

	methods := map[string]func(*http.Request) ([]byte, error) {
		"GET": getUsers,
		"POST": postUser,
		"DELETE": removeUser,
		"PUT": updateUser,
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")

	content, err := methods[req.Method](req)
	if err != nil {
		w.WriteHeader(400)
	} else {
		w.Write(content)
	}
}

func getUsers(req *http.Request) ([]byte, error) {

	r := regexp.MustCompile("/users/([0-9]+)")

	if r.MatchString(req.URL.Path) {
		return getUserByIndex(req)
	}

	file, _ := ioutil.ReadFile(db)

	return file, nil
}

func getUserByIndex(req *http.Request) ([]byte, error) {
	var users []user
	file, _ := ioutil.ReadFile(db)
	r := regexp.MustCompile("/users/([0-9]+)")
	
	index, _ := strconv.Atoi(string(r.FindSubmatch([]byte(req.URL.Path))[1]))

	json.Unmarshal(file, &users)

	if len(users) <= index {
		return nil, errors.New("Usuário não encontrado")
	}

	stringUsers, _ := json.Marshal(users[index])

	return []byte(stringUsers), nil
	
}

func postUser(req *http.Request) ([]byte, error) {

	value, _ := ioutil.ReadAll(req.Body)
	file, _ := ioutil.ReadFile(db)

	var newUser user
	var users []user

	json.Unmarshal(file, &users)
	json.Unmarshal(value, &newUser)

	bytesUsers, _ := json.Marshal(append(users, newUser))

	ioutil.WriteFile(db, bytesUsers, 0664)

	return nil, nil
}

func removeUser(req *http.Request) ([]byte, error) {
	file, _ := ioutil.ReadFile(db)
	r := regexp.MustCompile("/users/([0-9]+)")
	
	index, _ := strconv.Atoi(string(r.FindSubmatch([]byte(req.URL.Path))[1]))

	var users []user

	json.Unmarshal(file, &users)

	newUsers, _ := json.Marshal(append(users[:index], users[index+1:]...))

	ioutil.WriteFile(db, newUsers, 0664)

	return nil, nil
}

func updateUser(req *http.Request) ([]byte, error) {
	file, _ := ioutil.ReadFile(db)
	var users []user
	var newUser user

	r := regexp.MustCompile("/users/([0-9]+)")
	index, _ := strconv.Atoi(string(r.FindSubmatch([]byte(req.URL.Path))[1]))
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(file, &users)
	json.Unmarshal(body, &newUser)

	users[index] = newUser

	bytesNewUsers, _ := json.Marshal(users)

	ioutil.WriteFile(db, bytesNewUsers, 0664)

	return nil, nil
	
}