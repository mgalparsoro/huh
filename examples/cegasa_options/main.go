package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
)

type option int

const (
	AddFirmwareCandidate option = iota
	DeprecateECSFirmware
	DownloadECSFirmware
	GetAllECSFirmwares
	GetECSFirmwareDescription
	GetECSFirmwareHistory
	GetECSFirmwareUpdates
	GetFirmwareCandidates
	PromoteFirmware
	RejectFirmware
)

func (o option) String() string {
	return [...]string{
		"AddFirmwareCandidate",
		"DeprecateECSFirmware",
		"DownloadECSFirmware",
		"GetAllECSFirmwares",
		"GetECSFirmwareDescription",
		"GetECSFirmwareHistory",
		"GetECSFirmwareUpdates",
		"GetFirmwareCandidates",
		"PromoteFirmware",
		"RejectFirmware",
	}[o]
}

func makeAPICall(option option) {
	var url string
	switch option {
	case AddFirmwareCandidate:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint1"
	case DeprecateECSFirmware:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint2"
	case DownloadECSFirmware:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint3"
	case GetAllECSFirmwares:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint4"
	case GetECSFirmwareDescription:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint5"
	case GetECSFirmwareHistory:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint6"
	case GetECSFirmwareUpdates:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint7"
	case GetFirmwareCandidates:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint8"
	case PromoteFirmware:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint9"
	case RejectFirmware:
		url = "https://your-api-id.execute-api.region.amazonaws.com/your-stage/endpoint10"
	default:
		fmt.Println("Invalid option.")
		return
	}

	// Realizamos la llamada HTTP
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error realizando la llamada HTTP: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Respuesta de la API para %s: %d %s\n", option.String(), resp.StatusCode, resp.Status)
}

func main() {
	var selectedOption option

	// Crear las opciones del menú
	options := []huh.Option[option]{
		huh.NewOption("AddFirmwareCandidate", AddFirmwareCandidate),
		huh.NewOption("DeprecateECSFirmware", DeprecateECSFirmware),
		huh.NewOption("DownloadECSFirmware", DownloadECSFirmware),
		huh.NewOption("GetAllECSFirmwares", GetAllECSFirmwares),
		huh.NewOption("GetECSFirmwareDescription", GetECSFirmwareDescription),
		huh.NewOption("GetECSFirmwareHistory", GetECSFirmwareHistory),
		huh.NewOption("GetECSFirmwareUpdates", GetECSFirmwareUpdates),
		huh.NewOption("GetFirmwareCandidates", GetFirmwareCandidates),
		huh.NewOption("PromoteFirmware", PromoteFirmware),
		huh.NewOption("RejectFirmware", RejectFirmware),
	}

	// Crear el formulario y mostrar las opciones
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[option]().
				Title("Seleccione una opción para hacer la llamada a la API").
				Value(&selectedOption).
				Options(options...),
		),
	).Run()

	if err != nil {
		fmt.Println("Error en la selección de opción:", err)
		os.Exit(1)
	}

	// Hacer la llamada HTTP para la opción seleccionada
	makeAPICall(selectedOption)
}
