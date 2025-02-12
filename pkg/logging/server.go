// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package logging

import (
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"strings"

	"github.com/onosproject/onos-lib-go/api/logging"
	"github.com/onosproject/onos-lib-go/pkg/logging/service"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewService returns a new device Service
func NewService() (service.Service, error) {
	return &Service{}, nil
}

// Service is an implementation of C1 service.
type Service struct {
	service.Service
}

// Register registers the logging Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{}
	logging.RegisterLoggerServer(r, server)
}

// Server implements the logging gRPC service
type Server struct {
}

func splitLoggerName(name string) []string {
	names := strings.Split(name, nameSep)
	return names
}

// GetLevel implements GetLevel rpc function to get a logger level
func (s *Server) GetLevel(ctx context.Context, req *logging.GetLevelRequest) (*logging.GetLevelResponse, error) {

	name := req.GetLoggerName()
	if name == "" {
		return &logging.GetLevelResponse{}, errors.NewInvalid("precondition for get level request is failed")
	}

	names := splitLoggerName(name)
	logger := GetLogger(names...)
	level := logger.GetLevel()

	var loggerLevel logging.Level
	switch level {
	case InfoLevel:
		loggerLevel = logging.Level_INFO
	case DebugLevel:
		loggerLevel = logging.Level_DEBUG
	case WarnLevel:
		loggerLevel = logging.Level_WARN
	case ErrorLevel:
		loggerLevel = logging.Level_ERROR
	case PanicLevel:
		loggerLevel = logging.Level_PANIC
	case DPanicLevel:
		loggerLevel = logging.Level_DPANIC
	case FatalLevel:
		loggerLevel = logging.Level_FATAL

	}

	return &logging.GetLevelResponse{
		Level: loggerLevel,
	}, nil

}

// SetLevel implements SetLevel rpc function to set a logger level
func (s *Server) SetLevel(ctx context.Context, req *logging.SetLevelRequest) (*logging.SetLevelResponse, error) {
	name := req.GetLoggerName()
	level := req.GetLevel()
	if name == "" {
		return &logging.SetLevelResponse{
			ResponseStatus: logging.ResponseStatus_PRECONDITION_FAILED,
		}, errors.NewInvalid("precondition for set level request is failed")
	}

	names := splitLoggerName(name)
	logger := GetLogger(names...)

	switch level {
	case logging.Level_INFO:
		logger.SetLevel(InfoLevel)
	case logging.Level_DEBUG:
		logger.SetLevel(DebugLevel)
	case logging.Level_WARN:
		logger.SetLevel(WarnLevel)
	case logging.Level_ERROR:
		logger.SetLevel(ErrorLevel)
	case logging.Level_PANIC:
		logger.SetLevel(PanicLevel)
	case logging.Level_DPANIC:
		logger.SetLevel(DPanicLevel)
	case logging.Level_FATAL:
		logger.SetLevel(FatalLevel)

	default:
		return &logging.SetLevelResponse{
			ResponseStatus: logging.ResponseStatus_PRECONDITION_FAILED,
		}, errors.NewNotSupported("the requested level is not supported")

	}

	return &logging.SetLevelResponse{
		ResponseStatus: logging.ResponseStatus_OK,
	}, nil
}
