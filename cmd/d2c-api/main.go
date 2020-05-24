package main

import (
	"hrple/d2c/api/internal/handlers"

	web "github.com/hrple/common/server"
)

func main() {
	handlers.RegisterCompanyHandler()

	logger := web.GetLogger()
	if err := web.Start(""); err != nil {
		logger.Fatalf("Could not listen: %v\n", err)
	}
}
