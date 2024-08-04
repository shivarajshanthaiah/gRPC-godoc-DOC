package service_interfaces

import (
	pb "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/pb"
)

type DoctorServiceInter interface {
	CreateDoctorService(p *pb.DoctorModel) (*pb.DoctorResponse, error)
	FindDoctorByNameService(p *pb.DoctorName) (*pb.DoctorModel, error)
	FindDoctorByIDService(p *pb.DoctorID) (*pb.DoctorModel, error)
	FindAllDoctorsService(p *pb.NoParam) (*pb.DoctorList, error)
}
