package repo

import (
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/models"
	"gorm.io/gorm"
)

type DoctorRepository struct {
	DB *gorm.DB
}

func NewDoctorRepo(db *gorm.DB) *DoctorRepository {
	return &DoctorRepository{DB: db}
}

func (d *DoctorRepository) CreateDoctor(Doc *models.Doctor) error {
	if err := d.DB.Create(Doc).Error; err != nil {
		return err
	}
	return nil
}

func (d *DoctorRepository) FindDoctorByID(id uint) (*models.Doctor, error) {
	var Doc *models.Doctor
	err := d.DB.First(&Doc, id).Error
	return Doc, err
}

func (d *DoctorRepository) FindDoctorByName(name string) (*models.Doctor, error) {
	var Doc *models.Doctor
	err := d.DB.Where("doctor_name = ?", name).First(&Doc).Error
	return Doc, err
}

func (d *DoctorRepository) FindAllDoctors() ([]*models.Doctor, error) {
	var Doc []*models.Doctor
	err := d.DB.Find(&Doc).Error
	return Doc, err
}
