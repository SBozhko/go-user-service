package main

import (
	"fmt"
)

type Repository interface {
	Init()
	CreateUser(user User) User
	DeleteUser(id int) error
	FindUser(id int) User
	GetUsers() Users
}

type InMemoryRepository struct {
	Repository
	currentId int
	users     Users
}

// Give us some seed data
func (repo *InMemoryRepository) Init() {
	repo.CreateUser(User{FullName: "Hubert J. Farnsworth", Phone:"123-45-66", Address:"New New York"})
	repo.CreateUser(User{FullName: "John A. Zoidberg", Phone:"236-22-62", Address:"New New York"})
	repo.CreateUser(User{FullName: "Zapp Brannigan", Phone:"821-31-44", Address:"New New York"})
}

func (repo *InMemoryRepository) FindUser(id int) User {
	for _, t := range repo.users {
		if t.Id == id {
			return t
		}
	}
	// return empty User if not found
	return User{}
}

func (repo *InMemoryRepository) GetUsers() Users {
	return repo.users
}

//this is bad, I don't think it passes race conditions
func (repo *InMemoryRepository) CreateUser(user User) User {
	repo.currentId += 1
	user.Id = repo.currentId
	repo.users = append(repo.users, user)
	return user
}

func (repo *InMemoryRepository) DeleteUser(id int) error {
	for i, t := range repo.users {
		if t.Id == id {
			repo.users = append(repo.users[:i], repo.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find User with id of %d to delete", id)
}