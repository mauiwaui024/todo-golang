package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mauiwaui024/todo-golang"
)

// если у нас есть значение типа интерфейс - мы не знаем ничего о том что это
// мы знаем только то, что оно умеет делать
type Autorization interface {
	CreateUser(user todo.User) (int, error)
	//если пользователь есть то генерируем метод в который записываем айди пользователя
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

// собирает все сервисы в одном месте
type Repository struct {
	Autorization
	TodoList
	TodoItem
}

// сразу же объявляем конструктор
// 8. в файле репозитория инициализируем наш репозиторий в конструкторе и в service.go тоже самое сделаем
// только для этого меняем
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		TodoList:     NewTodoListPostgres(db),
		TodoItem:     NewTodoItemPostgres(db),
	}
}
