package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	// Создаем клиент Docker
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	// Определяем параметры для контейнера
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Маппинг портов: хост (80) -> контейнер (80)
	portBindings := nat.PortMap{
		"80/tcp": []nat.PortBinding{
			{
				HostIP:   "0.0.0.0", // либо прописать конкретный IP
				HostPort: "80",      // Порт на хосте
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "nginx:latest", // Образ Nginx
		ExposedPorts: map[nat.Port]struct{}{
			"80/tcp": {}, // Открываем порт 80 контейнера
		},
	}, &container.HostConfig{
		PortBindings: portBindings,
	}, nil, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	// Запускаем контейнер
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Контейнер запущен! ID: %s\n", resp.ID)

}
