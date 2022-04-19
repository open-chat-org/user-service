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
	session := u.Drive.NewSession(neo4j.SessionConfig{})
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
		Avatar:   neo4jUser.Props["avatar"].(string),
		Profile:  neo4jUser.Props["profile"].(string),
	}

	return &user, nil
}

func (u *UserNeo4jRepository) DeleteById(id int) error {
	session := u.Drive.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("MATCH (user:UserModel) WHERE ID(user)=$id DELETE user", map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserNeo4jRepository) UpdateAvatarById(id int, avatar string) error {
	session := u.Drive.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("MATCH (user:UserModel) WHERE ID(user)=$id SET user.avatar=$avatar", map[string]interface{}{"id": id, "avatar": avatar})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *UserNeo4jRepository) UpdateProfileById(id int, avatar string) error {
	session := u.Drive.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run("MATCH (user:UserModel) WHERE ID(user)=$id SET user.profile=$profile", map[string]interface{}{"id": id, "profile": avatar})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return err
	}

	return nil
}
