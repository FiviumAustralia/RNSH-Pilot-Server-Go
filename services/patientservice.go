package services

import (
	"fmt"

	"github.com/FiviumAustralia/RNSH-Pilot-Server-Go/models"
)

type PatientService interface {
	GetAllPatients() []models.Patient
	GetPatient(id int) models.Patient
	GetEhrId(mrn string) string
}

var currentPatientService PatientService

func GetPatientService() *PatientService {
	// TODO  fancy xml config or something to get different services
	// for now just get openEHR
	if currentPatientService == nil {
		fmt.Println("initing service")
		currentPatientService = OpenEHRPatientService{"https://ehrscape.code-4-health.org/rest/v1/", "rnsh.mrn", "rnshpilot_c4h", "lIsombRI"}
	}

	return &currentPatientService
}
