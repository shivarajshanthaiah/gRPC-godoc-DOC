package services

import (
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/models"
	pb "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/pb"
	repo_interfaces "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/repo/interfaces"
	service_interfaces "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/services/interfaces"
)

type DoctorService struct {
	repo repo_interfaces.DoctorRepoInter
}

func NewDoctorService(repo repo_interfaces.DoctorRepoInter) service_interfaces.DoctorServiceInter {
	return &DoctorService{
		repo: repo,
	}
}

func (d *DoctorService) CreateDoctorService(p *pb.DoctorModel) (*pb.DoctorResponse, error) {
	doctor := &models.Doctor{
		Name:       p.DoctorName,
		Age:        p.Age,
		Speciality: p.Speciality,
		Gender:     p.Gender,
		Email:      p.Email,
	}

	err := d.repo.CreateDoctor(doctor)
	if err != nil {
		return nil, err
	}

	return &pb.DoctorResponse{
		Status:  "Success",
		Error:   "",
		Message: "Doctor created sucessfully",
	}, nil
}

func (d *DoctorService) FindDoctorByIDService(p *pb.DoctorID) (*pb.DoctorModel, error) {
	doctor, err := d.repo.FindDoctorByID(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DoctorModel{
		DoctorName: doctor.Name,
		Age:        uint64(doctor.Age),
		Speciality: doctor.Speciality,
		Gender:     doctor.Gender,
		Email:      doctor.Email,
	}, nil
}

func (d *DoctorService) FindDoctorByNameService(p *pb.DoctorName) (*pb.DoctorModel, error) {
	doctor, err := d.repo.FindDoctorByName(p.Name)
	if err != nil {
		return nil, err
	}

	return &pb.DoctorModel{
		DoctorName: doctor.Name,
		Age:        uint64(doctor.Age),
		Speciality: doctor.Speciality,
		Gender:     doctor.Gender,
		Email:      doctor.Email,
	}, nil
}

func (d *DoctorService) FindAllDoctorsService(p *pb.NoParam) (*pb.DoctorList, error) {
	result, err := d.repo.FindAllDoctors()
	if err != nil {
		return nil, err
	}

	var doctors []*pb.DoctorModel
	for _, doctor := range result {
		doctors = append(doctors, &pb.DoctorModel{
			DoctorName: doctor.Name,
			Age:        uint64(doctor.Age),
			Speciality: doctor.Speciality,
			Gender:     doctor.Gender,
			Email:      doctor.Email,
		})
	}

	return &pb.DoctorList{
		Doctors:doctors,
	}, nil
}
