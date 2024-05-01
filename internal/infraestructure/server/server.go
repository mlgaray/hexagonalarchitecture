package server

import (
	"fmt"
	"github.com/mlgaray/hexagonalarchitecture/internal/core/ports/user"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	// We will add every new Handler here
	userHandlerPort user.UserHandlerPort
	// middlewares ports.IMiddlewares
	// paymentHandlers ports.IPaymentHandlers
}

func (s *Server) Initialize() {
	router := mux.NewRouter()

	sub := router.PathPrefix("/user").Subrouter()

	/*	uh := infraestructure.UserController{
		Service: service.NewUserServicePort(stubs.NewUserRepositoryStub()),
		//Service: service.NewUserServicePort(mongo.NewUserRepository(mongo.NewDataBaseConnection().Connect())),
	}*/

	sub.HandleFunc("/get-all", s.userHandlerPort.GetAll).Methods(http.MethodGet)

	/*router.HandleFunc("/signup", uh.SignUp).Methods(http.MethodPost)*/

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	fmt.Printf("listening at por: %s", PORT)

	handler := cors.AllowAll().Handler(router)
	/*	if os.Getenv("TESTING") == "true" {
		go func() {
			if err := http.ListenAndServe(":8080", handler); err != nil {
				log.Fatal(err)
			}
		}()
	} else {*/
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
	//}

	/*go func() {
		if err := http.ListenAndServe(":8080", handler); err != nil {
			log.Fatal(err)
		}
	}()*/
}

func NewServer(userHandlerPort user.UserHandlerPort) *Server {
	return &Server{
		userHandlerPort: userHandlerPort,
		// paymentHandlers: pHandlers
	}
}
