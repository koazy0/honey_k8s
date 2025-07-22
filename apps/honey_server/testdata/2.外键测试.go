package main

import (
	"context"
	"honey_server/internal/service"
)

func main() {

	service.Migrations().Migrate(context.Background())
}
