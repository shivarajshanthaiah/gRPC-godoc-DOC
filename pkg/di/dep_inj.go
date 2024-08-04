package di

import (
	"github.com/shivaraj-shanthaiah/godoc/docservice/config"
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/db"
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/handler"
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/repo"
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/server"
	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/services"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	doctorRepo := repo.NewDoctorRepo(db)
	doctorsvc := services.NewDoctorService(doctorRepo)
	doctorHandler := handler.NewDoctorHandler(doctorsvc)
	server.NewGrpcServer(doctorHandler)
}
