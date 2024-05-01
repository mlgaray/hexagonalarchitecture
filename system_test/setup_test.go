package system_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/mlgaray/hexagonalarchitecture/applicattion"
	"github.com/rs/cors"
)

var (
	serverWG sync.WaitGroup
	initWG   sync.WaitGroup
)

func TestMain(m *testing.M) {
	os.Setenv("TESTING", "true")
	// initWG.Add(1)
	// go func() {
	//	defer initWG.Done()
	server := applicattion.Init()
	go server.Initialize()

	//}()
	//initWG.Wait()
	//time.Sleep(5 * time.Second)

	code := m.Run()

	os.Exit(code)
}

func TestARejectPackageScenarios(t *testing.T) {
	client := resty.New()
	serverWG.Add(1)
	go startSever()
	serverWG.Wait()

	resp, err := client.R().
		Get("http://localhost:8080/user/get-all")

	resp2, err := client.SetBaseURL("http://localhost:3030/").R().
		EnableTrace().
		Get("user/update")
	if err != nil {
		println(err.Error())
	}

	println(resp)
	println(resp2)
}

func startSever() {
	defer serverWG.Done()
	router := mux.NewRouter()

	sub := router.PathPrefix("/user").Subrouter()

	sub.HandleFunc("/update", func(writer http.ResponseWriter, request *http.Request) {
		resource, _ := json.Marshal("dummy_update")
		_, _ = writer.Write(resource)
	}).Methods(http.MethodGet)
	handler := cors.AllowAll().Handler(router)

	server := &http.Server{Addr: fmt.Sprintf(":%d", 3030), Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Error starting server: %s", err)
		}
	}()
}
