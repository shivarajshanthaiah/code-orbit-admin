package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_admin/config"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/problem"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/clients/user"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/db"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/handler"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/repo"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/server"
	"github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/service"
)

func Init() {
	cnfg := config.LoadConfig()

	db := db.ConnectDB(cnfg)

	userClient, err := user.ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("failed to connect to user client")
	}

	problemClient, err := problem.ClientDial(*cnfg)
	if err != nil {
		log.Fatal("failed to connect to problem client")
	}

	adminRepo := repo.NewAdminRepository(db)

	adminService := service.NewAdminService(adminRepo, userClient, problemClient)

	adminHandler := handler.NewAdminHandler(adminService)

	err = server.NewGrpcAdminServer(cnfg.GrpcPort, adminHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
