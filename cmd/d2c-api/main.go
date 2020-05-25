package main

import (
	"hrple/d2c/api/internal/api/company"
	"hrple/d2c/api/internal/api/status"

	web "github.com/hrple/common/server"
)

func main() {
	company.RegisterHandler()
	status.RegisterHandler()

	logger := web.GetLogger()
	if err := web.Start(""); err != nil {
		logger.Fatalf("Could not listen: %v\n", err)
	}
}
