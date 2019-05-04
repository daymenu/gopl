package models

import (
	"crypto/md5"
	"fmt"
)

type User struct {
	Id        int64  `json:"id" db:"id"`
	UserName  string `json:"username" db:"userName"`
	Password  string `json:"password" db:"password"`
	Status    int    `json:"" db:"status"`
	CreatedAt int64  `json:"createAt" db:"createAt"`
	UpdatedAt int64  `json:"updateAt" db:"updateAt"`
	DeletedAt int64  `json:"deleteAt" db:"deleteAt"`
}

const UserTableName = "user"

func GetUserById(id int) (User, error) {
	var user User
	db, err := New()
	if err != nil {
		return user, err
	}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id=?", UserTableName)
	r := db.QueryRow(sql, id)
	r.Scan(
		&user.Id,
		&user.UserName,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	return user, nil
}

func GetUsers(status int) (users []User, err error) {
	db, err := New()
	if err != nil {
		return users, err
	}
	sql := fmt.Sprintf("SELECT * FROM %s where status=?", UserTableName)
	rows, err := db.Query(sql, status)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		rows.Scan(
			&users[i].Id,
			&users[i].UserName,
			&users[i].Password,
			&users[i].Status,
			&users[i].CreatedAt,
			&users[i].UpdatedAt,
			&users[i].DeletedAt,
		)
	}
	return users, nil
}

func AddUser(user *User) (id int64, err error) {
	db, err := New()
	if err != nil {
		return id, err
	}
	sql := fmt.Sprintf(
		"INSERT INTO %s (userName,password,status,createdAt,updatedAt) VALUES(?,?,?,?,?)",
		UserTableName)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return id, err
	}
	stmt.Close()
	r, err := stmt.Exec(
		user.UserName,
		Password(user.Password),
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return r.LastInsertId()
}

func Password(password string) string {
	pwdByte := md5.Sum([]byte(password))
	return fmt.Sprintf("%x", pwdByte)
}

func UpdateUser(id int64, user *User) (int64, error) {
	if id < 1 {
		return 0, fmt.Errorf("请传入要修改的ID")
	}
	db, err := New()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("UPDATE %s SET userName=?, status=?, updatedAt=?, deletedAt=? WHERE id=?", UserTableName)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	r, err := stmt.Exec(
		user.UserName,
		user.Status,
		user.UpdatedAt,
		user.DeletedAt,
		id,
	)

	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

func DeleteUser(id int64) (int64, error) {
	if id < 1 {
		return 0, fmt.Errorf("请传入要删除的ID")
	}
	db, err := New()
	if err != nil {
		return 0, err
	}
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=?", UserTableName)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	r, err := stmt.Exec(id)

	if err != nil {
		return 0, err
	}
	return r.RowsAffected()

}
