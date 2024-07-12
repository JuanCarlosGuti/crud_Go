package user

import (
	"context"
	"database/sql"
	"github.com/JuanCarlosGuti/Go_users.git/internal/domain"
	"log"
)

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id uint64) (*domain.User, error)
		Update(ctx context.Context, user *domain.User) error
		Update2(ctx context.Context, id uint64, firstName *string, lastName *string, email *string) error
	}

	repository struct {
		db  *sql.DB
		log *log.Logger
	}
)

func NewRepository(db *sql.DB, l *log.Logger) Repository {
	return &repository{
		db:  db,
		log: l,
	}
}
func (r *repository) Create(ctx context.Context, user *domain.User) error {

	sqlQ := "INSERT INTO users(first_name, last_name, email) VALUES(?, ?, ?)"
	res, err := r.db.Exec(sqlQ, user.FirstName, user.LastName, user.Email)
	id, err := res.LastInsertId()
	if err != nil {
		r.log.Printf("Error inserting user into database: %v", err)
		return err

	}
	user.ID = uint64(id)
	r.log.Println("User inserted successfully into database with id:", id)

	return nil

}
func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	sqlQ := "SELECT id, first_name, last_name, email FROM users"
	rows, err := r.db.Query(sqlQ)
	if err != nil {
		r.log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			r.log.Printf("Error getting all users: %v", err)
			return nil, err
		}
		users = append(users, u)
	}
	r.log.Println("Users retrieved successfully into database: ", len(users))
	return users, nil
}
func (r *repository) Get(ctx context.Context, id uint64) (*domain.User, error) {
	sqlQ := "SELECT id, first_name, last_name, email FROM users WHERE id=?"
	row := r.db.QueryRow(sqlQ, id)

	var u domain.User
	if err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			r.log.Printf("No user found with id %v", id)
			return nil, ErrorNotFound{id}
		}

	}

	return &u, nil

}

func (r *repository) Update(ctx context.Context, user *domain.User) error {

	sqlQ := "UPDATE users SET first_name=?, last_name=?,email=? WHERE id=?"
	if _, err := r.db.Exec(sqlQ, user.FirstName, user.LastName, user.Email, user.ID); err != nil {
		if err == sql.ErrNoRows {
			r.log.Printf("No user found with id %v", user.ID)
			return nil
		}
		r.log.Printf("Error updating user by id %v: %v", user.ID, err)
		return err
	}

	return nil
}

func (r *repository) Update2(ctx context.Context, id uint64, firstName *string, lastName *string, email *string) error {

	var sqlQ = ""
	if firstName != nil && *firstName != "" {
		sqlQ = "UPDATE users SET first_name=? WHERE id=?"
		if _, err := r.db.Exec(sqlQ, *firstName, id); err != nil {
			r.log.Printf("Error updating user by id %v: %v", id, err)
			return nil
		}

	}
	if lastName != nil && *lastName != "" {
		sqlQ = "UPDATE users SET last_name=? WHERE id=?"
		if _, err := r.db.Exec(sqlQ, *lastName, id); err != nil {
			r.log.Printf("Error updating user by id %v: %v", id, err)
			return nil
		}
	}
	if email != nil && *email != "" {
		sqlQ = "UPDATE users SET email=? WHERE id=?"
		if _, err := r.db.Exec(sqlQ, *email, id); err != nil {
			r.log.Printf("Error updating user by id %v: %v", id, err)
			return nil
		}
	}

	//chat GPT
	// Obtener el usuario actual de la base de datos
	//user, err := r.Get(ctx, id)
	//if err != nil {
	//	r.log.Printf("Error getting user by id %v: %v", id, err)
	//	return err
	//}
	//if user == nil {
	//	r.log.Printf("No user found with id %v", id)
	//	return nil
	//}
	//
	//// Construir la consulta SQL de actualización dinámicamente
	//updates := []string{}
	//args := []interface{}{}
	//
	//if firstName != nil && user.FirstName != *firstName {
	//	updates = append(updates, "first_name=?")
	//	args = append(args, *firstName)
	//	user.FirstName = *firstName
	//}
	//if lastName != nil && user.LastName != *lastName {
	//	updates = append(updates, "last_name=?")
	//	args = append(args, *lastName)
	//	user.LastName = *lastName
	//}
	//if email != nil && user.Email != *email {
	//	updates = append(updates, "email=?")
	//	args = append(args, *email)
	//	user.Email = *email
	//}
	//
	//// Si no hay actualizaciones, retornar sin hacer nada
	//if len(updates) == 0 {
	//	return nil
	//}
	//
	//// Añadir el ID del usuario al final de los argumentos
	//args = append(args, user.ID)
	//
	//// Construir la consulta SQL completa
	//sqlQ := fmt.Sprintf("UPDATE users SET %s WHERE id=?", strings.Join(updates, ", "))
	//
	//// Ejecutar la consulta SQL de actualización
	//if _, err := r.db.Exec(sqlQ, args...); err != nil {
	//	r.log.Printf("Error updating user by id %v: %v", user.ID, err)
	//	return err
	//}

	return nil
}

//func (r *repository) Update(ctx context.Context, user *domain.User) error {
//	r.log.Println("repository.Update", user)
//	users := r.db.Users
//	for i, u := range users {
//		if u.ID == user.ID {
//			users[i] = *user
//			r.db.Users = users
//			r.log.Println("repository.Update", users[i])
//		}
//	}
//	return nil
//}
