package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "github.com/FelipeFernandezUSM/SistemasDistribuidos-lab3/communication"

	"google.golang.org/grpc"
)

func sendRequest(client pb.CentralServerClient, TeamId int32) (bool, error) {
	// Preparar la solicitud con cantidades aleatorias de munición
	request := &pb.Request{
		TeamId:     TeamId,
		ATQuantity: rand.Int31n(11) + 20, // Rango: 20-30
		MPQuantity: rand.Int31n(6) + 10,  // Rango: 10-15
	}

	// Enviar la solicitud al servidor central
	response, err := client.SolicitarM(context.Background(), request)
	if err != nil {
		return false, err
	}

	// Manejar la respuesta del servidor
	if response.Success {
		log.Printf("Respuesta positiva recibida para el equipo %d. Programa finalizado.", TeamId)
		// Finalizar el programa si se recibió una respuesta positiva
		// Opcional: podrías agregar aquí la lógica para realizar otras acciones si lo deseas
		return true, nil
	}
	return false, nil
}

func main() {
	// Establecer conexión con el servidor central
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no se pudo conectar: %v", err)
	}
	defer conn.Close()

	// Crear cliente para el servidor central
	client := pb.NewCentralServerClient(conn)

	// Generar un número de equipo al azar
	TeamId := rand.Int31n(1000) + 1

	// Realizar la primera consulta después de 10 segundos
	time.Sleep(10 * time.Second)
	if Success, err := sendRequest(client, TeamId); err != nil {
		log.Printf("error durante la primera consulta: %v", err)
	} else if Success {
		return
	}

	// Realizar consultas periódicas cada 3 segundos hasta recibir una respuesta positiva o negativa
	for {
		time.Sleep(3 * time.Second)
		if Success, err := sendRequest(client, TeamId); err != nil {
			log.Printf("error durante la consulta: %v", err)
			continue
		} else if Success {
			return
		}

		log.Printf("Respuesta negativa recibida para el equipo %d. Reintentando en 3 segundos...", TeamId)
		break
	}
}
