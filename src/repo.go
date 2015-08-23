package main

import (
	"fmt"
)

type Repository interface {
	Init()
	CreateUser(user UserInDb) UserInDb
	DeleteUser(id int) error
	FindUser(id int) UserInDb
	GetUsers() UsersInDb
	CreateUserByEmail(email string) UserInDb
}

// Not thread-safe
type InMemoryRepository struct {
	Repository
	currentId int
	users     UsersInDb
}

// Give us some seed data
func (repo *InMemoryRepository) Init() {
	repo.CreateUser(UserInDb{FullName: "Hubert J. Farnsworth", Phone:"123-45-66", Address:"New New York"})
	repo.CreateUser(UserInDb{FullName: "John A. Zoidberg", Phone:"236-22-62", Address:"New New York"})
	repo.CreateUser(UserInDb{FullName: "Zapp Brannigan", Phone:"821-31-44", Address:"New New York"})
}

func (repo *InMemoryRepository) FindUser(id int) UserInDb {
	for _, t := range repo.users {
		if t.Id == id {
			return t
		}
	}
	// return empty User if not found
	return UserInDb{}
}

func (repo *InMemoryRepository) GetUsers() UsersInDb {
	return repo.users
}

func (repo *InMemoryRepository) CreateUser(user UserInDb) UserInDb {
	repo.currentId += 1
	user.Id = repo.currentId
	repo.users = append(repo.users, user)
	return user
}

func (repo *InMemoryRepository) CreateUserByEmail(email string) UserInDb {
	repo.currentId += 1
	newUser := UserInDb{Id:repo.currentId, Email:email}
	repo.users = append(repo.users, newUser)
	return newUser
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