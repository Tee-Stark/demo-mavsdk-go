package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/mavlink/MAVSDK-Go/Sources/action"
	"github.com/mavlink/MAVSDK-Go/Sources/core"
	"github.com/mavlink/MAVSDK-Go/Sources/geofence"
	"github.com/mavlink/MAVSDK-Go/Sources/telemetry"
	"google.golang.org/grpc"
)

//Drone creates a drone object to interact with drone related plugins
type Drone struct {
	port         string
	mavsdkServer string
	action       action.ServiceImpl
	core         core.ServiceImpl
	telemetry    telemetry.ServiceImpl
	geofence     geofence.ServiceImpl
}

//Connect Starts a mavsdk server and create a connection to it
func (s *Drone) Connect() {
	//start mavsdk server
	// s.startMAVSDKServer()
	grpcConnection := s.connectToMAVSDKServer()
	s.InitPlugins(grpcConnection)

}

//InitPlugins initializes all the plugins
func (s *Drone) InitPlugins(cc *grpc.ClientConn) {

	s.telemetry = telemetry.ServiceImpl{
		Client: telemetry.NewTelemetryServiceClient(cc),
	}
	s.core = core.ServiceImpl{
		Client: core.NewCoreServiceClient(cc),
	}
	s.action = action.ServiceImpl{
		Client: action.NewActionServiceClient(cc),
	}
	s.action = action.ServiceImpl{
		Client: action.NewActionServiceClient(cc),
	}
	s.geofence = geofence.ServiceImpl{
		Client: geofence.NewGeofenceServiceClient(cc),
	}
}

func (s *Drone) startMAVSDKServer() {
	cmd := exec.Command("C:\\Users\\ykhedar\\Downloads\\mavsdk_server_win32", "-p", "50051")
	print("Command [%s]", cmd)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		print("Error starting mavsdkserver %s", err)
	}
}

func (s *Drone) connectToMAVSDKServer() *grpc.ClientConn {
	dialoption := grpc.WithInsecure()

	serverAddr := s.mavsdkServer + ":" + s.port
	cc, err := grpc.Dial(serverAddr, dialoption)
	if err != nil {
		fmt.Printf("Error while dialing %v", err)
	}
	grpc.ConnectionTimeout(5)
	return cc
}

func main() {
	drone := &Drone{port: "50051", mavsdkServer: "192.168.0.119"}
	drone.Connect()
	// drone.core.Init()
	// x := drone.core.ListRunningPlugins()
	// fmt.Printf("%v\n", x)
	//ch := drone.telemetry.Position()
	//for x := range ch {
	//	fmt.Printf("%v\n", x)
	//}
	drone.action.Arm(context.Background())
	drone.action.Takeoff(context.Background())
	drone.action.Land(context.Background())
	// drone.telemetry.SetRatePosition(60)

	// drone.action.GotoLocation(54, 56, 0.5, 30)
	drone.core.ConnectionState(context.Background())
	// fmt.Printf("channel value %v \n", <-ch1)
	lat := 47.3977508
	lon := 8.5456074
	p1 := &geofence.Point{
		LatitudeDeg:  lat - 0.0001,
		LongitudeDeg: lon - 0.0001,
	}
	p2 := &geofence.Point{
		LatitudeDeg:  lat + 0.0001,
		LongitudeDeg: lon - 0.0001,
	}
	p3 := &geofence.Point{
		LatitudeDeg:  lat + 0.0001,
		LongitudeDeg: lon + 0.0001,
	}
	p4 := &geofence.Point{
		LatitudeDeg:  lat - 0.0001,
		LongitudeDeg: lon + 0.0001,
	}

	polygon := &geofence.Polygon{
		Points:    []*geofence.Point{p1, p2, p3, p4},
		FenceType: geofence.Polygon_FENCE_TYPE_EXCLUSION}
	drone.geofence.UploadGeofence(context.Background(), []*geofence.Polygon{polygon})

}
