package server

import (
	"fmt"
	"os"
)

func Router() {
	// router := mux.NewRouter()

	/*uh := infraestructure.UserController{
		Service: service.NewUserServicePort(stubs.NewUserRepositoryStub()),
		//Service: service.NewUserServicePort(mongo.NewUserRepository(mongo.NewDataBaseConnection().Connect())),
	}

	router.HandleFunc("/get-all", uh.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/signup", uh.SignUp).Methods(http.MethodPost)*/

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	fmt.Printf("listening at por: %s", PORT)
	// handler := cors.AllowAll().Handler(router)
	// log.Fatal(http.ListenAndServe(":"+PORT, handler))
	// return router
}
