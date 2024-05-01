package features

import (
	"log"

	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
)

type FeatureUsers struct {
	//*integration.FeatureTester
}

func FeatureUsersSteps(s *godog.ScenarioContext) {
	// Definimos el escenario
	s.Step(`^llamamos al endpoint para obtener usuarios$`, llamamosParaOobtenerrUsuarios)
	/*	s.Step(`^existe un libro con ID "([^"]*)"$`, existeUnLibroConID)
		s.Step(`^se envía una petición GET a "([^"]*)"$`, seEnviaUnaPeticionGETA)
		s.Step(`^se debería recibir una respuesta con código (\d+)$`, seDeberiaRecibirUnaRespuestaConCodigo)
		s.Step(`^el cuerpo de la respuesta debería ser:$`, elCuerpoDeLaRespuestaDeberiaSer)*/
}

func llamamosParaOobtenerrUsuarios() error {
	client := resty.New()
	resp, err := client.R().
		Get("http://localhost:8080/user/get-all")
	if err != nil {
		log.Fatalf("Error al hacer la petición: %v", err)
		return err
	}

	// Aquí puedes manejar la respuesta. Por ejemplo, imprimir el cuerpo de la respuesta:
	log.Println(resp)

	return nil
}
