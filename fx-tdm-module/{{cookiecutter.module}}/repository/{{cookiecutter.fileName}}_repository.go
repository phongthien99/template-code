package repository

import (
	"{{cookiecutter.basePath}}/model"

	base "github.com/sigmaott/gest/package/technique/database/base"
	mongoRepository "github.com/sigmaott/gest/package/technique/database/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type I{{cookiecutter.className}}Repository interface {
	base.IRepository[model.{{cookiecutter.className}}]
}
type {{cookiecutter.fieldPrefix}}Repository struct {
	mongoRepository.BaseMongoRepository[model.{{cookiecutter.className}}]
}

func New{{cookiecutter.className}}Repository(db *mongo.Database) I{{cookiecutter.className}}Repository {
	return &{{cookiecutter.fieldPrefix}}Repository{
		BaseMongoRepository: mongoRepository.BaseMongoRepository[model.{{cookiecutter.className}}]{
			Collection: db.Collection("{{cookiecutter.fileName}}"),
		},
	}
}
