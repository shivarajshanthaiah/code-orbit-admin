package handler

import (
	pb "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/proto"
	inter "github.com/shivaraj-shanthaiah/code_orbit_admin/pkg/service/interfaces"
)

type AdminHandler struct {
	SVC inter.AdminServiceInter
	pb.AdminServiceServer
}

func NewAdminHandler(svc inter.AdminServiceInter) *AdminHandler {
	return &AdminHandler{
		SVC: svc,
	}
}
