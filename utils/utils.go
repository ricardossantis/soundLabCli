package utils

import (
	"bytes"
	"encoding/json"
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

const SOUND_LAB_SERVER = "http://localhost:5000"

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
	responseData, _ := ioutil.ReadAll(resp.Body)

	loginCreateResp := LoginResp{}
	_ = json.Unmarshal(responseData, &loginCreateResp)

	return loginCreateResp
}
