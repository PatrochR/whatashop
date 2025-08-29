package repository

import (
	"database/sql"
	"time"

	"github.com/PatrochR/whatashop/model"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

type UserRepository interface{
	GetAll() (*[]model.User , error)
	GetById(id *model.UserID) (*model.User , error)
	Add(*model.User) error
	UpdateUsername(*model.UserID,string) error
	UpdatePassword(*model.UserID,string) error
	UpdateEmail(*model.UserID,string) error
	Delete(id *model.UserID) error
}

func (p *UserPostgres) Init() error {
	query := `create table if not exist users (
		id UUID PRIMARY KEY,
		username text not null unique,
		email text not null unique,
		password text not null,
		created_at timestamp not null,
		update_at timestamp not null,
		delete_at timestamp ,

	)`
	_, err := p.db.Exec(query)
	return err
}
func (p *UserPostgres) GetAll() (*[]model.User, error) {
	query := `select * from users`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (p *UserPostgres) GetById(id *model.UserID) (*model.User, error) {
	query := `select * form users where id = $1`
	row := p.db.QueryRow(query, id)
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (p *UserPostgres) Add(user *model.User) error {
	query := `insert into users (username, email , password , created_at , update_at , delete_at)
		values ($1,$2,$3,$4,$5,$6)
	`
	_, err := p.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	return err
}

func (p *UserPostgres) UpdateUsername(id *model.UserID, username string) error {
	query := `update users set username = $1 , update_at = $2 where id = $3`
	_, err := p.db.Exec(query, username, time.Now().UTC(), id)
	return err
}

func (p *UserPostgres) UpdatePassword(id *model.UserID, password string) error {
	query := `update users set password = $1 , update_at = $2 where id = $3`
	_, err := p.db.Exec(query, password, time.Now().UTC(), id)
	return err
}

func (p *UserPostgres) UpdateEmail(id *model.UserID, email string) error {
	query := `update users set email = $1 , update_at = $2 where id = $3`
	_, err := p.db.Exec(query, email, time.Now().UTC(), id)
	return err
}

func (p *UserPostgres) Delete(id *model.UserID) error {
	query := `update users set delete_at = $1 , update_at = $2 where id = $3`
	time := time.Now().UTC()
	_, err := p.db.Exec(query, time, time, id)
	return err
}
