package repo

import (
	"context"
	"fmt"

	"github.com/example/todo-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLTodoRepository struct {
	db *gorm.DB
}

func NewSQLTodoRepository() *SQLTodoRepository {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=todo sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&model.Todo{}, &model.User{}); err != nil {
		panic("failed to migrate database")
	}
	return &SQLTodoRepository{db: db}
}

func (r *SQLTodoRepository) Create(ctx context.Context, todo *model.Todo) (string, error) {
	userID, ok := ctx.Value("userID").(int)
	if ok {
		todo.UserID = userID
	}
	// Implementation for creating a todo in the SQL database
	if err := r.db.Create(todo).Error; err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", todo.ID), nil
}
func (r *SQLTodoRepository) List(ctx context.Context) ([]*model.Todo, error) {
	// Implementation for listing todos from the SQL database
	userID, ok := ctx.Value("userID").(int)
	if ok {
		fmt.Println("User ID from context:", userID)
	}
	var todos []*model.Todo
	if err := r.db.Where("user_id = ?", userID).Order("id ASC").Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
func (r *SQLTodoRepository) Get(ctx context.Context, id int) (*model.Todo, error) {
	userID, ok := ctx.Value("userID").(int)
	if ok {
		fmt.Println("User ID from context:", userID)
	}
	// Implementation for getting a todo by ID from the SQL database
	var todo model.Todo
	if err := r.db.Where("id = ? AND user_id = ? AND done = ?", id, userID, false).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// func queryBuilder(filters map[string]interface{}) string {
// 	query := "1=1"
// 	for key, value := range filters {
// 		switch value.(type) {
// 		case string:
// 			query += fmt.Sprintf(" AND %s ILIKE '%%%s%%'", key, value.(string))
// 		case bool:
// 			query += fmt.Sprintf(" AND %s = %t", key, value.(bool))
// 		case int:
// 			query += fmt.Sprintf(" AND %s = %d", key, value.(int))
// 		case time.Time:
// 			query += fmt.Sprintf(" AND %s >= '%s'", key, value.(time.Time).Format("2006-01-02 15:04:05"))
// 		}
// 	}

// 	return query
// }

// func (r *SQLTodoRepository) Filter(ctx context.Context, query string) ([]*model.Todo, error) {
// 	// Implementation for filtering todos from the SQL database
// 		query += fmt.Sprintf(" AND user_id = %d", *userID)
// 	}
// 	if created != nil {
// 		query += fmt.Sprintf(" AND created_at >= '%s'", created.Format("2006-01-02 15:04:05"))
// 	}
// 	return query
// }

func (r *SQLTodoRepository) Filter(ctx context.Context, query string) ([]*model.Todo, error) {
	// Implementation for filtering todos from the SQL database
	var todos []*model.Todo
	if err := r.db.Where("title ILIKE ? AND user_id = ?", "%"+query+"%", ctx.Value("userID")).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *SQLTodoRepository) Update(ctx context.Context, todo *model.Todo) error {
	// Implementation for updating a todo in the SQL database
	return r.db.Save(todo).Error
}
func (r *SQLTodoRepository) Delete(ctx context.Context, id int) error {
	// Implementation for deleting a todo by ID from the SQL database
	return r.db.Delete(&model.Todo{}, id).Error
}
func (r *SQLTodoRepository) SetDone(ctx context.Context, id int, done bool) (*model.Todo, error) {
	// Implementation for setting a todo as done/undone in the SQL database
	var todo model.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	todo.Done = done
	if err := r.db.Save(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
func (r *SQLTodoRepository) RegisterUser(ctx context.Context, user *model.User) (int, error) {
	// Implementation for registering a user in the SQL database
	if err := r.db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *SQLTodoRepository) AuthenticateUser(ctx context.Context, username, password string) (int, error) {
	// Implementation for authenticating a user in the SQL database
	var user model.User
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}
