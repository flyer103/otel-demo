// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alibabacloudlogserviceexporter

import (
	"errors"
	"net"
	"os"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	slsutil "github.com/aliyun/aliyun-log-go-sdk/util"
	"go.uber.org/zap"
)

// LogServiceClient log Service's client wrapper
type LogServiceClient interface {
	// SendLogs send message to LogService
	SendLogs(logs []*sls.Log) error
}

type logServiceClientImpl struct {
	clientInstance *producer.Producer
	project        string
	logstore       string
	topic          string
	source         string
	logger         *zap.Logger
}

func getIPAddress() (ipAddress string, err error) {
	as, err := net.InterfaceAddrs()
	for _, a := range as {
		if in, ok := a.(*net.IPNet); ok && !in.IP.IsLoopback() {
			if in.IP.To4() != nil {
				ipAddress = in.IP.String()
			}
		}
	}
	return ipAddress, err
}

// NewLogServiceClient Create Log Service client
func NewLogServiceClient(config *Config, logger *zap.Logger) (LogServiceClient, error) {
	if config == nil || config.Endpoint == "" || config.Project == "" || config.Logstore == "" {
		return nil, errors.New("missing logservice params: Endpoint, Project, Logstore")
	}

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = config.Endpoint
	producerConfig.AccessKeyID = config.AccessKeyID
	producerConfig.AccessKeySecret = config.AccessKeySecret
	if config.ECSRamRole != "" || config.TokenFilePath != "" {
		tokenUpdateFunc, _ := slsutil.NewTokenUpdateFunc(config.ECSRamRole, config.TokenFilePath)
		producerConfig.UpdateStsToken = tokenUpdateFunc
		producerConfig.StsTokenShutDown = make(chan struct{})
	}

	c := &logServiceClientImpl{
		project:        config.Project,
		logstore:       config.Logstore,
		clientInstance: producer.InitProducer(producerConfig),
		logger:         logger,
	}
	c.clientInstance.Start()
	// do not return error if get hostname or ip address fail
	c.topic, _ = os.Hostname()
	c.source, _ = getIPAddress()
	logger.Info("Create LogService client success", zap.String("project", config.Project), zap.String("logstore", config.Logstore))
	return c, nil
}

// SendLogs send message to LogService
func (c *logServiceClientImpl) SendLogs(logs []*sls.Log) error {
	return c.clientInstance.SendLogListWithCallBack(c.project, c.logstore, c.topic, c.source, logs, c)
}

// Success is impl of producer.CallBack
func (c *logServiceClientImpl) Success(*producer.Result) {}

// Fail is impl of producer.CallBack
func (c *logServiceClientImpl) Fail(result *producer.Result) {
	c.logger.Warn("Send to LogService failed",
		zap.String("project", c.project),
		zap.String("store", c.logstore),
		zap.String("code", result.GetErrorCode()),
		zap.String("error_message", result.GetErrorMessage()),
		zap.String("request_id", result.GetRequestId()))
}
