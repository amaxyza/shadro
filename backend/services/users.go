package services

import (
	"fmt"
	"log"
	"os"

	"context"

	"github.com/amaxyza/shadro/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var pool *pgxpool.Pool

func Connect() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file.")
		return err
	}

	pool, err = pgxpool.New(context.Background(), os.Getenv("PSQL_EXTERNAL_CONNECT"))
	if err != nil {
		fmt.Errorf("error creating db pool")
		return err
	}

	return nil
}

func ClosePool() {
	pool.Close()
}

func hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func AddUser(name string, password string) (*models.User, error) {
	password_hashed, err := hash(password)
	if err != nil {
		fmt.Println("Error encrypting password of new user.")
		return nil, err
	}

	tag, err := pool.Exec(
		context.Background(),
		`INSERT INTO "user" (name, password) VALUES (@username, @userpassword)`,
		pgx.NamedArgs{
			"username":     name,
			"userpassword": password_hashed,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("unable to add row")
	}

	fmt.Println("Rows affected: ", tag.RowsAffected())

	return &models.User{Name: name, Password_Hash: password_hashed}, nil
}

func DeleteUserByName(name string) error {

	tag, err := pool.Exec(
		context.Background(),
		`DELETE FROM "user" WHERE name = @username`,
		pgx.NamedArgs{"username": name},
	)

	if err != nil {
		return err
	}

	fmt.Println("Deleted ", tag.RowsAffected(), "rows. (Should be only 1)")
	return nil
}

func DeleteUserByID(id int) error {
	tag, err := pool.Exec(
		context.Background(),
		`DELETE FROM "user" WHERE id = @userid`,
		pgx.NamedArgs{"username": id},
	)

	if err != nil {
		return err
	}

	fmt.Println("Deleted ", tag.RowsAffected(), "rows. (Should be only 1)")
	return nil
}

func ValidateUser(name string, password string) (bool, error) {
	var hashed string

	pool.QueryRow(
		context.Background(),
		`SELECT password FROM "user" WHERE name = @username`,
		pgx.NamedArgs{"username": name},
	).Scan(&hashed)
	fmt.Println("Hashed password in db: ", hashed)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User

	row := pool.QueryRow(
		context.Background(),
		`SELECT * FROM "user" WHERE id = @userid`,
		pgx.NamedArgs{"userid": id},
	)

	if err := row.Scan(&user.ID, &user.Name, &user.Password_Hash); err == pgx.ErrNoRows {
		return nil, err
	}

	return &user, nil
}

func GetUserByName(name string) (*models.User, error) {
	var user models.User

	row := pool.QueryRow(
		context.Background(),
		`SELECT * FROM "user" WHERE id = @username`,
		pgx.NamedArgs{"username": name},
	)

	if err := row.Scan(&user.ID, &user.Name, &user.Password_Hash); err == pgx.ErrNoRows {
		return nil, err
	}

	return &user, nil
}

func GetUserList() []models.PublicUser {
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM "user"`,
	)
	if err != nil {
		log.Fatal(err)
		return []models.PublicUser{}
	}

	var users []models.PublicUser

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name, nil)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, models.PublicUser{
			ID:   id,
			Name: name,
		})
	}

	return users
}
