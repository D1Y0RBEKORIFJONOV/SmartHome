package clientgrpcserver

import (
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"user_service_smart_home/internal/config"
)

type ServiceClient interface {
	UserServiceClient() user1.UserServiceClient
	Close() error
}

type serviceClient struct {
	connection  []*grpc.ClientConn
	userService user1.UserServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connSoldiersService, err := grpc.NewClient(fmt.Sprintf("%s", cfg.RPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &serviceClient{
		userService: user1.NewUserServiceClient(connSoldiersService),
		connection:  []*grpc.ClientConn{connSoldiersService},
	}, nil
}

func (s *serviceClient) UserServiceClient() user1.UserServiceClient {
	return s.userService
}

func (s *serviceClient) Close() error {
	var err error
	for _, conn := range s.connection {
		if cerr := conn.Close(); cerr != nil {
			log.Println("Error while closing gRPC connection:", cerr)
			err = cerr
		}
	}
	return err
}
