package repository

import (
	"user-service/model"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type UserNeo4jRepository struct {
	Drive neo4j.Driver
}

func (u *UserNeo4jRepository) FindById(id int) (*model.UserModel, error) {
	driver, err := neo4j.NewDriver(
		"neo4j://localhost:7687",
		neo4j.BasicAuth("neo4j", "secret", ""),
	)
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("MATCH (user:UserModel) WHERE ID(user)=$id RETURN user", map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return nil, err
	}

	neo4jUser := result.(dbtype.Node)

	user := model.UserModel{
		ID:       int(neo4jUser.Id),
		Email:    neo4jUser.Props["email"].(string),
		Username: neo4jUser.Props["username"].(string),
	}

	return &user, nil
}

func (u *UserNeo4jRepository) FindFriendsById(id int) {
}

func (u *UserNeo4jRepository) Save(user model.UserModel) {
}

func (u *UserNeo4jRepository) DeleteByEmail(email string) {
}
