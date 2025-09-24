package repo

import (
	"database/sql"
	"errors"

	"github.com/arthu/shop-api-go/internal/db"
	"github.com/arthu/shop-api-go/internal/models"
)

func CreateUser(u *models.User) (int64, error) {
	s := `INSERT INTO users (username, email, password, firstName, lastName, address, phoneNumber, role) VALUES (?,?,?,?,?,?,?,?)`
	res, err := db.DB().Exec(s, u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.Address, u.PhoneNumber, u.Role)
	if err != nil { return 0, err }
	return res.LastInsertId()
}

func GetUserByID(id int64) (*models.User, error) {
	row := db.DB().QueryRow(`SELECT id, username, email, password, firstName, lastName, address, phoneNumber, role FROM users WHERE id = ?`, id)
	u := models.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.Address, &u.PhoneNumber, &u.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) { return nil, nil }
		return nil, err
	}
	return &u, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	row := db.DB().QueryRow(`SELECT id, username, email, password, firstName, lastName, address, phoneNumber, role FROM users WHERE email = ?`, email)
	u := models.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.Address, &u.PhoneNumber, &u.Role); err != nil {
		if errors.Is(err, sql.ErrNoRows) { return nil, nil }
		return nil, err
	}
	return &u, nil
}

func ListUsers() ([]models.User, error) {
	rows, err := db.DB().Query(`SELECT id, username, email, password, firstName, lastName, address, phoneNumber, role FROM users`)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.User
	for rows.Next() {
		u := models.User{}
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.Address, &u.PhoneNumber, &u.Role); err != nil { return nil, err }
		out = append(out, u)
	}
	return out, nil
}

func UpdateUser(id int64, u *models.User) (bool, error) {
	s := `UPDATE users SET username=?, email=?, firstName=?, lastName=?, address=?, phoneNumber=?, role=? WHERE id=?`
	res, err := db.DB().Exec(s, u.Username, u.Email, u.FirstName, u.LastName, u.Address, u.PhoneNumber, u.Role, id)
	if err != nil { return false, err }
	affected, _ := res.RowsAffected()
	return affected > 0, nil
}

func DeleteUser(id int64) (bool, error) {
	res, err := db.DB().Exec(`DELETE FROM users WHERE id=?`, id)
	if err != nil { return false, err }
	affected, _ := res.RowsAffected()
	return affected > 0, nil
}
