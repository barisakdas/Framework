package log

import (
	. "github.com/barisakdas/Framework/Log/Models"
)

type IFileLoggerService interface {
	LogRequest(request RequestModel)
	LogResponse(response ResponseModel)
}
