package main

import (
	"context"
	
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "../communication"
)

type server struct {
	pb.UnimplementedCentralServerServer
	inventory struct {
		AT int
		MP int
	}
}

func (s *server) SolicitarM(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	// Check if there is enough ammunition in the inventory
	if s.inventory.AT >= int(in.ATQuantity) && s.inventory.MP >= int(in.MPQuantity) {
		// Deduct ammunition from the inventory
		s.inventory.AT -= int(in.ATQuantity)
		s.inventory.MP -= int(in.MPQuantity)
		log.Printf("Recepción de solicitud desde equipo %d, %d AT y %d MP -- APROBADA --", in.TeamId, in.ATQuantity, in.MPQuantity)
		return &pb.Response{Success: true}, nil
	}
	// Not enough ammunition in inventory
	log.Printf("Recepción de solicitud desde equipo %d, %d AT y %d MP -- DENEGADA --", in.TeamId, in.ATQuantity, in.MPQuantity)
	return &pb.Response{Success: false}, nil
}

func (s *server) generateAmmunition() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {

		<-ticker.C

		newAT := 10
		newMP := 5

		if s.inventory.AT+newAT <= 50 {
			s.inventory.AT += newAT
		} else {
			s.inventory.AT = 50
		}
		if s.inventory.MP+newMP <= 20 {
			s.inventory.MP += newMP
		} else {
			s.inventory.MP = 20
		}
		log.Printf("Current inventory: AT=%d, MP=%d", s.inventory.AT, s.inventory.MP)
	}
}

func main() {
	s := grpc.NewServer()
	centralServer := &server{}
	pb.RegisterCentralServerServer(s, centralServer)

	go centralServer.generateAmmunition()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on port 8080...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
