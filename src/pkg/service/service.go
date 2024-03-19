package service

import (
	"github.com/mauiwaui024/todo-golang"
	"github.com/mauiwaui024/todo-golang/pkg/repository"
)

type Autorization interface {
	//3.опишем метод CreateUser, принимает стракт юзера, возвращает айди созданного в базе пользователя
	CreateUser(user todo.User) (int, error)
	//принимаем пароль и логин а возвращаем токен
	GenerateToken(username string, password string) (string, error)
	//возвращает айди пользователя
	ParseToken(token string) (int, error)
}

type TodoList interface {
	//возвращает айди созданного ТуДуСписка и ошибку
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

// собирает все сервисы в одном месте
type Service struct {
	Autorization
	TodoList
	TodoItem
}

// сразу же объявляем конструктор
// сервисы будут обращаться к базе данных, поэтому в качестве аргумента передаем
// указатель на структуру репозитори
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repo.Autorization),
		TodoList:     NewTodoListService(repo.TodoList),
		TodoItem:     NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
