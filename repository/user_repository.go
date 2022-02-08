package repository

import (
	"user-service/model"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserNeo4jRepository struct {
	Drive neo4j.Driver
}

func (u *UserNeo4jRepository) findById(id string) {
}

func (u *UserNeo4jRepository) findFriendsById(id string) {
}

func (u *UserNeo4jRepository) save(user model.User) {
}

func (u *UserNeo4jRepository) deleteByMail(mail string) {
}
