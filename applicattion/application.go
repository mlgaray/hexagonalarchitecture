package applicattion

import (
	"github.com/mlgaray/hexagonalarchitecture/internal/core/services"
	"github.com/mlgaray/hexagonalarchitecture/internal/handlers"
	"github.com/mlgaray/hexagonalarchitecture/internal/infraestructure/adapters/repositories/postgresql"
	"github.com/mlgaray/hexagonalarchitecture/internal/infraestructure/server"
)

func Init() *server.Server {
	// defer wg.Done()
	userHandler := initUser()
	httpServer := server.NewServer(userHandler)
	return httpServer
}

func initUser() *handlers.UserHandler {
	// userRepository := mongo2.NewUserRepository(mongo2.NewDataBaseConnection().Connect())
	//userRepository := mysql2.NewUserRepository(mysql2.NewDataBaseConnection().Connect())
	userRepository := postgresql.NewUserRepository(postgresql.NewDataBaseConnection().Connect())
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}
