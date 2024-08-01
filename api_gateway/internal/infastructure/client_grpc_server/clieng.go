package clientgrpcserver

import (
	"api_gate_way/internal/config"
	"fmt"
	tv1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/TV"
	alarm1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/smart_alarm"
	speaker1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/device/speaker"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient interface {
	TvService() tv1.TVServiceClient
	SpeakerService() speaker1.SpeakerServiceClient
	SmartAlarm() alarm1.SmartAlarmServiceClient
	UserService() user1.UserServiceClient
	Close() error
}

type serviceClient struct {
	connection []*grpc.ClientConn
	tv         tv1.TVServiceClient
	speaker    speaker1.SpeakerServiceClient
	smartAlarm alarm1.SmartAlarmServiceClient
	user       user1.UserServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connTV, err := grpc.NewClient(fmt.Sprintf("%s",
		cfg.DevicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	connUser, err := grpc.NewClient(fmt.Sprintf("%s",
		cfg.UserPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &serviceClient{
		tv:         tv1.NewTVServiceClient(connTV),
		speaker:    speaker1.NewSpeakerServiceClient(connTV),
		smartAlarm: alarm1.NewSmartAlarmServiceClient(connTV),
		user:       user1.NewUserServiceClient(connUser),
		connection: []*grpc.ClientConn{connTV},
	}, nil
}

func (s *serviceClient) TvService() tv1.TVServiceClient {
	return s.tv
}
func (s *serviceClient) SpeakerService() speaker1.SpeakerServiceClient {
	return s.speaker
}
func (s *serviceClient) SmartAlarm() alarm1.SmartAlarmServiceClient {
	return s.smartAlarm
}
func (s *serviceClient) UserService() user1.UserServiceClient {
	return s.user
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
