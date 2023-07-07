package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResp struct {
	Token string `json:"token,omitempty"`
}

type MarketName struct {
	Name string `json:"name,omitempty"`
}

type CreateMktResp struct {
	Name string `json:"name,omitempty"`
}

const SOUND_LAB_SERVER = "http://localhost:3000"

func Login(credentials Credentials) LoginResp {
	url := SOUND_LAB_SERVER + "/user/login"
	credentialsJson, err := json.Marshal(credentials)
	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(credentialsJson)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	loginCreateResp := LoginResp{}
	_ = json.Unmarshal(responseData, &loginCreateResp)

	return loginCreateResp
}

func CreateMarketplace(name MarketName, token string) (CreateMktResp, error) {
	url := SOUND_LAB_SERVER + "/marketplace/create"
	nameJson, err := json.Marshal(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("token: %+v\n", token)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(nameJson)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		errorMessage := string(body)
		fmt.Println("Error:", errorMessage)
		return CreateMktResp{}, errors.New(errorMessage)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	CreateMktResp := CreateMktResp{}
	_ = json.Unmarshal(responseData, &CreateMktResp)

	return CreateMktResp, nil
}
