package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"
	"github.com/jmoiron/sqlx"
)

// type User struct {
// 	ID          int    `json:"id" db:"id"`
// 	FirstName   string `json:"first_name" db:"first_name"`
// 	LastName    string `json:"last_name" db:"last_name"`
// 	Email       string `json:"email" db:"email"`
// 	Password    string `json:"password" db:"password"`
// 	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
// }

type UserRepo interface {
	user.UserRepo //embedding the user service
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
		VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
		RETURNING id;
	`

	var id int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&id)
	}
	rows.Close()

	user.ID = id

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
    // SQL query to find the user
    query := `
        SELECT id, first_name, last_name, email, password, is_shop_owner
        FROM users
        WHERE email = $1 AND password = $2
        LIMIT 1
    `

    var user domain.User
    // Get one row
    err := r.db.Get(&user, query, email, pass)
    if err != nil {
        if err == sql.ErrNoRows || err.Error() == "sql: no rows in result set" {
            return nil, nil // user not found
        }
        return nil, err // some other error
    }

    return &user, nil
}
