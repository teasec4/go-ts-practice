package repository

import (
	"database/sql"
	"justforfun/internal/model"

	"github.com/google/uuid"
)

type SQLiteUserRepostiroyImpl struct{
	DB *sql.DB
}

func NewSQLiteUserRepositoryImpl(db *sql.DB) *SQLiteUserRepostiroyImpl{
	return &SQLiteUserRepostiroyImpl{DB: db}
}

func (r *SQLiteUserRepostiroyImpl) CreateUser(user model.User) error{
	_, err := r.DB.Exec("INSERT INTO users (id, name) VALUES (?, ?)", user.Id.String(), user.Name)
	return err
}

func (r *SQLiteUserRepostiroyImpl) GetAllUsers()([]model.User, error){
	rows, err := r.DB.Query("SELECT id, name FROM users")
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	
	var users []model.User
	for rows.Next(){
		var idStr string
        var name string
        if err := rows.Scan(&idStr, &name); err != nil {
            return nil, err
        }
        id, _ := uuid.Parse(idStr)
        users = append(users, model.User{Id: id, Name: name})
	}
	
	return users, nil
}

func (r *SQLiteUserRepostiroyImpl) FindUserByID(id string) (*model.User, error) {
    row := r.DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

    var idStr string
    var name string
    if err := row.Scan(&idStr, &name); err != nil {
        return nil, err
    }
    uid, _ := uuid.Parse(idStr)
    return &model.User{Id: uid, Name: name}, nil
}