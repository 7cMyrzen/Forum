package db

import (
	"crypto/sha256"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connexion à la base de données
var db *sql.DB

func dbConnect() error {
	// Informations de connexion à la base de données
	username := "root"
	password := "**Tax1p9"
	hostname := "127.0.0.1:3306"
	dbname := "forum_go"

	// Chaîne de connexion à la base de données
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)

	// Connexion à la base de données MySQL
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Vérifier si la connexion à la base de données est réussie
	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Connexion à la base de données réussie!")
	return nil
}

// Fonction pour ajouter un utilisateur à la base de données pour registerHandler
func AddUser(username, email, password string) error {
	if db == nil {
		if err := dbConnect(); err != nil {
			return err
		}
	}

	insert, err := db.Prepare("INSERT INTO user(username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer insert.Close()

	_, err = insert.Exec(username, email, hashPassword(password))
	return err
}

// Fonction pour chercher un utilisateur et mot de passe dans la base de données pour loginHandler qui renvoie l'ID de l'utilisateur ou 0 si l'utilisateur n'existe pas
func FindUser(username, password string) (int, error) {
	if db == nil {
		if err := dbConnect(); err != nil {
			return 0, err
		}
	}

	var id int
	err := db.QueryRow("SELECT id FROM user WHERE (username = ? OR email = ?) AND password = ?", username, username, hashPassword(password)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Foction pour se connecter avec soit l'emal ou le nom d'utilisateur et le mot de passe qiu renvoie l'ID de l'utilisateur ou 0 si l'utilisateur n'existe pas
func AuthenticateUser(username, password string) (int, error) {
	if db == nil {
		if err := dbConnect(); err != nil {
			return 0, err
		}
	}

	var id int
	err := db.QueryRow("SELECT id FROM user WHERE (username = ? OR email = ?) AND password = ?", username, username, hashPassword(password)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func deleteUser(id int) error {
	if db == nil {
		if err := dbConnect(); err != nil {
			return err
		}
	}

	delete, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return err
	}
	defer delete.Close()

	_, err = delete.Exec(id)
	return err
}

func showUsers() error {
	if db == nil {
		if err := dbConnect(); err != nil {
			return err
		}
	}

	type User struct {
		ID          int
		Username    string
		Email       string
		Description string
		Level       int
	}

	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		return err
	}
	defer results.Close()

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.Username, &user.Email, &user.Description, &user.Level)
		if err != nil {
			return err
		}
		fmt.Println(user)
	}

	return nil
}

// Fonction pour hacher le mot de passe
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
