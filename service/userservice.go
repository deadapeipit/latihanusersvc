package service

import (
	"fmt"
	"main/entity"
	"time"
	"unicode"
	"unicode/utf8"
)

type UserServiceIface interface {
	Register(user *entity.User, cv *entity.Cv) (users *entity.User, cvs *entity.Cv, err error)
	Perkenalan()
}

type UserSvc struct {
	User entity.User
	Cv   entity.Cv
}

func NewUserService() UserServiceIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User, cv *entity.Cv) (users *entity.User, cvs *entity.Cv, err error) {
	u.User = *user
	u.Cv = *cv
	fmt.Println("Data berhasil dimasukan.")
	return user, cv, nil
}

func (u *UserSvc) Perkenalan() {
	fmt.Printf("Nama user adalah %s. \n", u.User.Username)
	u.HitungUmur()
	fmt.Printf("Pekerjaan %s adalah %s, telah bekerja selama %d tahun. \n", u.User.Username, u.Cv.PekerjaanSekarang, u.Cv.Pengalaman)
	u.TebakPassword()
}

func (u *UserSvc) HitungUmur() {
	tahunSekarang := time.Now().Year()
	fmt.Printf("%s lahir pada tahun %d. \n", u.User.Username, tahunSekarang-u.User.Age)
}

func (u *UserSvc) TebakPassword() {
	jumlahKarakter := utf8.RuneCountInString(u.User.Password)
	hasUpper := false
	hasSpecialCharacter := false
	adaKapital := "Tidak terdapat"
	jumlahKapital := 0
	adaKarakterSpesial := "Tidak terdapat"
	jumlahKarakterSpesial := 0
	for _, r := range u.User.Password {
		if unicode.IsUpper(r) {
			hasUpper = true
			jumlahKapital += 1
		}
		if !unicode.IsLetter(r) {
			hasSpecialCharacter = true
			jumlahKarakterSpesial += 1
		}
	}
	if hasUpper {
		adaKapital = fmt.Sprintf("Terdapat %d", jumlahKapital)
	}
	if hasSpecialCharacter {
		adaKarakterSpesial = fmt.Sprintf("Terdapat %d", jumlahKarakterSpesial)
	}
	fmt.Printf("Password %s terdiri dari %d karakter, %s huruf kapital, %s karakter spesial. \n", u.User.Username, jumlahKarakter, adaKapital, adaKarakterSpesial)
}
