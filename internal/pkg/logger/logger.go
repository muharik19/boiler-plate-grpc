package logger

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/muharik19/boiler-plate-grpc/configs"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/elasticsearch"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/utils"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
	global "github.com/muharik19/boiler-plate-grpc/pkg/utils"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
)

type LogActivity struct {
	Log       string    `json:"log"`
	CreatedAt time.Time `json:"createdAt"`
}

func SetIdentifierId(ctx context.Context) context.Context {
	nodeNumber := global.Getenv("NODE_NUMBER")
	n, _ := strconv.ParseInt(*nodeNumber, 10, 64)
	node, err := utils.NewSnowflakeNode(n)
	if err != nil {
		logger.Errf("NewSnowflakeNode Err: %v", err)
	}

	ctx = context.WithValue(ctx, configs.IdentifierId, node.GenerateID().String())

	return ctx
}

func ActivityLogger(ctx context.Context, layer, function, url, method string, request, response any) {
	var log string
	if url != "" {
		if request != nil && response != nil {
			log = fmt.Sprintf("[%v:log][identifier-id]:%v,[layer:%v],[function:%v],[url]:%v,[method]:%v,[request]:%v,[response]:%v", *global.Getenv("APP_NAME"), ctx.Value(configs.IdentifierId).(string), layer, function, url, method, request, response)
			logger.Info(log)
		} else if request != nil && response == nil {
			log = fmt.Sprintf("[%v:log][identifier-id]:%v,[layer:%v],[function:%v],[url]:%v,[method]:%v,[request]:%v", *global.Getenv("APP_NAME"), ctx.Value(configs.IdentifierId).(string), layer, function, url, method, request)
			logger.Info(log)
		} else if request == nil && response != nil {
			log = fmt.Sprintf("[%v:log][identifier-id]:%v,[layer:%v],[function:%v],[url]:%v,[method]:%v,[response]:%v", *global.Getenv("APP_NAME"), ctx.Value(configs.IdentifierId).(string), layer, function, url, method, response)
			logger.Info(log)
		} else if request == nil && response == nil {
			log = fmt.Sprintf("[%v:log][identifier-id]:%v,[layer:%v],[function:%v],[url]:%v,[method]:%v", *global.Getenv("APP_NAME"), ctx.Value(configs.IdentifierId).(string), layer, function, url, method)
			logger.Info(log)
		}
	} else {
		log = fmt.Sprintf("[%v:log][identifier-id]:%v,[layer:%v],[function:%v]", *global.Getenv("APP_NAME"), ctx.Value(configs.IdentifierId).(string), layer, function)
		logger.Info(log)
	}

	logActivity := LogActivity{
		Log:       log,
		CreatedAt: time.Now().UTC().Local(),
	}
	elasticsearch.Insert(ctx, INDEX_LOG_ACTIVITY, logActivity)
}

func ErrorLogger(ctx context.Context, layer, function, url, method string, request, response any, err error) {
	log := fmt.Sprintf("[%v:log][identifier-id]:%v,[layer:%v],[function:%v],[url]:%v,[method]:%v,[request]:%v,[response]:%v,[error]:%v", *global.Getenv("APP_NAME"), ctx.Value(configs.IdentifierId).(string), layer, function, url, method, request, response, err.Error())
	logger.Info(log)
	logActivity := LogActivity{
		Log:       log,
		CreatedAt: time.Now().UTC().Local(),
	}
	elasticsearch.Insert(ctx, INDEX_LOG_ERROR, logActivity)
}
