package product2

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type RepositoryImpl struct {
	Database *sql.DB
}

func NewRepository2(Database *sql.DB) Repository2 {
	return &RepositoryImpl{Database}
}

func (respositoryImpl *RepositoryImpl) GetByID(id int) (product *domain.Product, err error) {

	query := `
	SELECT
	id, name, qantity, code_value, is_published, expiration, price
	FROM products
	WHERE id = ?
	`
	row := respositoryImpl.Database.QueryRow(query, id)

	err = row.Scan(
		&product.Id,
		&product.Name,
		&product.Quantity,
		&product.CodeValue,
		&product.IsPublished,
		&product.Expiration,
		&product.Price,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrorNotFound
		}
		return
	}

	return product, nil
}

func (respositoryImpl *RepositoryImpl) Create(product *domain.Product) (err error) {

	statement, err := respositoryImpl.Database.Prepare(`
	INSERT INTO products (name, qantity, code_value, is_published, expiration, price)
	VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
	)
	if err != nil {
		mysqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return
		}

		switch mysqlError.Number {
		case 1062:
			err = ErrAlreadyExists
		case 1586:
			err = ErrAlreadyExists
		}
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	product.Id = int(lastId)

	return
}

func (respositoryImpl *RepositoryImpl) Update(id int, product *domain.Product) (err error) {

	statement, err := respositoryImpl.Database.Prepare(`
		UPDATE products
		SET name = ?, qantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ?
		WHERE id = ?
	`)
	if err != nil {
		return
	}

	_, err = statement.Exec(
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
		product.Id,
	)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}

	return
}

func (respositoryImpl *RepositoryImpl) Delete(id int) (err error) {
	statement, err := respositoryImpl.Database.Prepare(`
		DELETE FROM products
		WHERE id = ?
	`)
	if err != nil {
		return
	}

	_, err = statement.Exec(id)
	if err != nil {
		// TODO: Cast to MySQL error.
		return
	}

	return
}
