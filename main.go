package main

import (
	"fmt"
	"main/entity"
	"main/service"
	"time"
)

func main() {

	userSvc := service.NewUserService()
	newUser := &entity.User{
		Id:        1,
		Username:  "david123",
		Email:     "david123@gmail.com",
		Password:  "Passdav!d",
		Age:       17,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	cv := &entity.Cv{
		PekerjaanSekarang: "Pengangguran serba bisa",
		Pengalaman:        2,
	}

	if user, cv, err := userSvc.Register(newUser, cv); err != nil {
		fmt.Printf("Error when register user: %+v \n", err)
	} else {
		fmt.Printf("Success register user: %+v \n", user)
		fmt.Println("----------------------------------")
		fmt.Printf("cv: %+v \n", cv)
		fmt.Println("----------------------------------")

		userSvc.Perkenalan()
	}
}
