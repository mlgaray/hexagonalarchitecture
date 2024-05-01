package postgresql

import (
	"database/sql"
	models2 "github.com/mlgaray/hexagonalarchitecture/internal/core/models"
	"log"
)

type UserSqlRepository struct {
	db       *sql.DB
	users    []models2.User
	Incident []models2.Incident
}

func (s *UserSqlRepository) GetAll() ([]models2.User, error) {
	rows, err := s.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models2.User
		// Lee los resultados en la estructura User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		// Agrega el usuario al array
		s.users = append(s.users, user)
	}

	// Imprime los resultados
	/*	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}*/

	// Maneja cualquier error que ocurra durante el proceso de iteración
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return s.users, nil
}

func (s *UserSqlRepository) SignUp(user models2.User) error {
	return nil
}

func NewUserRepository(client *sql.DB) *UserSqlRepository {
	return &UserSqlRepository{
		db: client,
	}
}
