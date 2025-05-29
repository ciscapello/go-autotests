package main

import (
	"context"
	"os"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/Yandex-Practicum/go-autotests/internal/fork"
)

type Iteration3BSuite struct {
	suite.Suite

	serverAddress string
	serverProcess *fork.BackgroundProcess
}

func (suite *Iteration3BSuite) SetupSuite() {
	suite.Require().NotEmpty(flagServerBinaryPath, "-binary-path non-empty flag required")

	suite.serverAddress = "http://localhost:8080"

	// Для обеспечения обратной совместимости с будущими заданиями
	envs := append(os.Environ(), []string{
		"RESTORE=false",
	}...)
	suite.serverProcess = fork.NewBackgroundProcess(context.Background(), flagServerBinaryPath,
		fork.WithEnv(envs...),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := suite.serverProcess.Start(ctx)
	if err != nil {
		suite.T().Errorf("Невозможно запустить процесс командой %s: %s. Переменные окружения: %+v", suite.serverProcess, err, envs)
		return
	}

	port := "8080"
	err = suite.serverProcess.WaitPort(ctx, "tcp", port)
	if err != nil {
		suite.T().Errorf("Не удалось дождаться пока порт %s станет доступен для запроса: %s", port, err)
		return
	}
}

func (suite *Iteration3BSuite) TearDownSuite() {
}

func (suite *Iteration3BSuite) TestCounter() {
}
