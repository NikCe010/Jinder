package grpc

import (
	jobs_service "Jinder/jinder-api/jobs/pkg/handler/grpc/protos"
	"Jinder/jinder-api/jobs/pkg/service"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func Run(host string, port int, services *service.Service) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	jobs_service.RegisterResumeServiceServer(grpcServer, NewResumeService(services))
	jobs_service.RegisterVacancyServiceServer(grpcServer, NewVacancyService(services))
	grpcServer.Serve(lis)
}
