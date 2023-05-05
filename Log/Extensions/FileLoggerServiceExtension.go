package log

import (
	interfaces "github.com/barisakdas/Framework/Log/Interfaces"
	services "github.com/barisakdas/Framework/Log/Services"
)

func NewFileLoggerService(logPath string) interfaces.IFileLoggerService {
	return &services.FileLoggerService{LogDir: logPath}
}
