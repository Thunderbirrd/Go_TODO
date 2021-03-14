package service

import (
	todo "github.com/Thunderbirrd/Go_TODO"
	"github.com/Thunderbirrd/Go_TODO/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userid, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userid, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewToDoItemService(repos.TodoItem, repos.TodoList),
	}
}
