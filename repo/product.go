package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"_id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageUrl    string  `json:"imageUrl" db:"image_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (title, description, price, image_url)
		values ($1, $2, $3, $4)
		RETURNING id;
		`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl)
	err := row.Scan(&p.ID)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product

	query := `SELECT 
	id, 
	title, 
	description, 
	price,
	image_url
	FROM products
	WHERE id=$1;`

	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}

	return &prd, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product

	query := `SELECT
	id,
	title,
	description,
	price,
	image_url
	FROM products
	`

	err := r.db.Select(&prdList, query)
	if err != nil {
		fmt.Println("Error fetching product list:", err)
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id=$1;`

	_, err := r.db.Exec(query, id)

	return err
}

func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
		UPDATE products 
		SET title=$1, description=$2, price=$3, image_url=$4
		WHERE id=$5;
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &p, nil
}
