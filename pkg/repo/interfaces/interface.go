package repo_interfaces

import "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/models"

type DoctorRepoInter interface {
	CreateDoctor(Doctor *models.Doctor) error
	FindDoctorByID(id uint) (*models.Doctor, error)
	FindDoctorByName(name string) (*models.Doctor, error)
	FindAllDoctors() ([]*models.Doctor, error)
}
