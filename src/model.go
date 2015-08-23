package main

type User struct {
	Id       int `json:"id"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type Users []User

type SignUpEmail struct {
	Email string `json:"email"`
}

type UserInDb struct {
	Id       int
	FullName string
	Phone    string
	Address  string
	Email    string
}

type UsersInDb []UserInDb

func User2UserInDb(user User) UserInDb {
	userInDb := UserInDb{Id:user.Id, FullName:user.FullName, Phone:user.Phone, Address:user.Address}
	return userInDb
}

func UserInDb2User(userInDb UserInDb) User {
	user := User{Id:userInDb.Id, FullName:userInDb.FullName, Phone:userInDb.Phone, Address:userInDb.Address}
	return user
}