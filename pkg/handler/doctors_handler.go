package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/pb"
	service_interfaces "github.com/shivaraj-shanthaiah/godoc/docservice/pkg/services/interfaces"
)

type DoctorsHandler struct {
	pb.DoctorServicesServer
	svc service_interfaces.DoctorServiceInter
}

func NewDoctorHandler(svc service_interfaces.DoctorServiceInter) *DoctorsHandler {
	return &DoctorsHandler{
		svc: svc,
	}
}

func (d *DoctorsHandler) CreateDoctor(ctx context.Context, p *pb.DoctorModel) (*pb.DoctorResponse, error) {
	result, err := d.svc.CreateDoctorService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DoctorsHandler) FetchDoctorByID(ctx context.Context, p *pb.DoctorID) (*pb.DoctorModel, error) {
	result, err := d.svc.FindDoctorByIDService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DoctorsHandler) FetchDoctorByName(ctx context.Context, p *pb.DoctorName) (*pb.DoctorModel, error) {
	result, err := d.svc.FindDoctorByNameService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DoctorsHandler) FetchAllDoctors(ctx context.Context, p *pb.NoParam) (*pb.DoctorList, error) {
	result, err := d.svc.FindAllDoctorsService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}
