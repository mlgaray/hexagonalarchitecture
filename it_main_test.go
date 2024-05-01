package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/mlgaray/hexagonalarchitecture/applicattion"
	"github.com/mlgaray/hexagonalarchitecture/godogs/features"
)

/*func TestMain(m *testing.M) {
	os.Setenv("TESTING", "true")
	//initWG.Add(1)
	//go func() {
	//	defer initWG.Done()
	var wg sync.WaitGroup
	wg.Add(1)
	server := applicattion.Init()

	go func() {
		go server.Initialize()
		wg.Done()
	}()
	//}()
	//initWG.Wait()
	//time.Sleep(5 * time.Second)

	code := m.Run()

	os.Exit(code)
}*/

func TestFeatures(t *testing.T) {
	status := godog.TestSuite{
		Name:                 "integration",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  initializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"godogs/features"},
		},
	}.Run()

	os.Exit(status)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	// Iniciamos el servidor mock
	/*server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hola mundo!"}`))
	}))*/

	// Creamos un nuevo escenario para cada caso de prueba

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {

	//wg.Done()
	//}()
	ctx.BeforeSuite(func() {
		server := applicattion.Init()
		go server.Initialize()
		/*server := applicattion.Init()

		// Crear un canal para esperar a que el servidor se inicialice
		done := make(chan bool)

		go func() {
			server.Initialize()
			// Cuando el servidor se haya inicializado, enviar un valor por el canal
			done <- true
		}()

		// Esperar a que el servidor se haya inicializado antes de continuar
		<-done*/
	})
	/*
		go ctx.BeforeSuite(func() {
			server := applicattion.Init()
			go server.Initialize()
			//wg.Wait()
			/*server := applicattion.Init()
			go server.Initialize()*/
	//time.Sleep(10 * time.Second)
	//})

	ctx.AfterSuite(func() {
		// Después de ejecutar la suite de pruebas, hacemos alguna limpieza adicional
	})

	// Registramos los pasos que utilizaremos en nuestros casos de prueba
}

func initializeScenario(s *godog.ScenarioContext) {
	// Aquí agregaríamos nuestros escenarios y steps definidos en archivos .feature y .go

	features.FeatureUsersSteps(s)
}

/*func TestFeatures(t *testing.T) {
	//logging.Log.Level = logrus.WarnLevel

	suite := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: func(s *godog.TestSuiteContext) { suiteContext(s) },
		ScenarioInitializer:  nil,
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{"features"},
			Tags:      "~@wip",
			Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
		},
	}

	if suite.Run() != 0 {
		t.Fail()
	}
}*/

/*func suiteContext(suite *godog.TestSuiteContext) {
	var serverManager mockserver.Manager
	suite.BeforeSuite(
		func() {
			prepareMockServer()
			serverManager = mockserver.Run(":8088", mockResources)
			beforeSuite(serverManager, suite)
		},
	)
	suite.AfterSuite(func() { _ = serverManager.Shutdown(context.Background()) })
}*/
