package service

import (
	todo "github.com/Thunderbirrd/Go_TODO"
	"github.com/Thunderbirrd/Go_TODO/pkg/repository"
)

type ToDoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
	
}

func NewToDoItemService(repo repository.TodoItem, listRepo repository.TodoList) *ToDoItemService {
	return &ToDoItemService{repo: repo, listRepo: listRepo}
}

func (s * ToDoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil{
		// list does not exists or does not belong to user
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *ToDoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *ToDoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *ToDoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *ToDoItemService) Update(userid, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userid, itemId, input)
}
