package service

import (
	"context"
	"{{cookiecutter.basePath}}/dto"
	"{{cookiecutter.basePath}}/model"
	"{{cookiecutter.basePath}}/repository"

	"github.com/google/uuid"
	base "github.com/sigmaott/gest/package/technique/database/base"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type I{{cookiecutter.className}}Service interface {
	Create(payload *dto.Create{{cookiecutter.className}}, appId string) (*model.{{cookiecutter.className}}, error)
	FindOne(params struct {
		Id    string
		AppId string
	}) (*model.{{cookiecutter.className}}, error)
	Paginate(query any, paginate *base.Paginate, sort *base.Sort) (results *base.PaginateResponse[model.{{cookiecutter.className}}], err error)
	UpdateOne(params struct {
		Id    string
		AppId string
	}, payload *dto.Update{{cookiecutter.className}}) (*model.{{cookiecutter.className}}, error)
	DeleteOne(params struct {
		Id    string
		AppId string
	}) (*model.{{cookiecutter.className}}, error)
}
type {{cookiecutter.fieldPrefix}}Service struct {
	{{cookiecutter.fieldPrefix}}Repository repository.I{{cookiecutter.className}}Repository
}

func (s *{{cookiecutter.fieldPrefix}}Service) Create(payload *dto.Create{{cookiecutter.className}}, appId string) (*model.{{cookiecutter.className}}, error) {

	modelCreated := &model.{{cookiecutter.className}}{
		Id: uuid.New().String(),
	}
	res, err := s.{{cookiecutter.fieldPrefix}}Repository.CreateOne(context.TODO(), modelCreated)
	return res, err
}

func (s *{{cookiecutter.fieldPrefix}}Service) FindOne(params struct {
	Id    string
	AppId string
}) (*model.{{cookiecutter.className}}, error) {
	res, err := s.{{cookiecutter.fieldPrefix}}Repository.FindOne(context.TODO(), bson.M{
		"id": params.Id,
	})
	if err != nil {
		return res, nil
	}
	if res.AppId != params.AppId {
		return nil, mongo.ErrNilDocument
	}
	return res, err
}

// UpdateOne implements IExampleService.
func (s *{{cookiecutter.fieldPrefix}}Service) UpdateOne(params struct {
	Id    string
	AppId string
}, payload *dto.Update{{cookiecutter.className}}) (*model.{{cookiecutter.className}}, error) {
	res, err := s.{{cookiecutter.fieldPrefix}}Repository.FindOne(context.TODO(), bson.M{
		"id": params.Id,
	})
	if err != nil {
		return res, nil
	}
	if res.AppId != params.AppId {
		return nil, mongo.ErrNilDocument
	}

	res, err = s.{{cookiecutter.fieldPrefix}}Repository.UpdateOne(context.TODO(), bson.M{
		"id": params.Id,
	}, payload)
	return res, err
}

func (s *{{cookiecutter.fieldPrefix}}Service) DeleteOne(params struct {
	Id    string
	AppId string
}) (*model.{{cookiecutter.className}}, error) {
	res, err := s.{{cookiecutter.fieldPrefix}}Repository.FindOne(context.TODO(), bson.M{
		"id": params.Id,
	})
	if err != nil {
		return res, nil
	}
	if res.AppId != params.AppId {
		return nil, mongo.ErrNilDocument
	}

	res, err = s.{{cookiecutter.fieldPrefix}}Repository.DeleteOne(context.TODO(), bson.M{
		"id": params.Id,
	})
	return res, err
}

// Paginate implements IExampleService.
func (s *{{cookiecutter.fieldPrefix}}Service) Paginate(query any, paginate *base.Paginate, sort *base.Sort) (results *base.PaginateResponse[model.{{cookiecutter.className}}], err error) {
	return s.{{cookiecutter.fieldPrefix}}Repository.Paginate(context.TODO(), query, paginate, sort)
}

func New{{cookiecutter.className}}Service(exampleRepository repository.I{{cookiecutter.className}}Repository) I{{cookiecutter.className}}Service {

	return &{{cookiecutter.fieldPrefix}}Service{
		{{cookiecutter.fieldPrefix}}Repository: {{cookiecutter.fieldPrefix}}Repository,
	}

}
