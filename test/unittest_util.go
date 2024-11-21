package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"hub-service/cmd/tool/migration"
	"hub-service/config"
	"hub-service/container"
	"hub-service/logger"
	"hub-service/middleware"
	"hub-service/repository"
	"hub-service/session"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// PrepareForControllerTest func prepares the controllers for testing.
func PrepareForControllerTest(isSecurity bool) (*echo.Echo, container.Container) {
	e := echo.New()

	conf := createConfig(isSecurity)
	logger := initTestLogger()
	container := initContainer(conf, logger)

	migration.CreateDatabase(container)
	migration.InitMasterData(container)

	middleware.InitLoggerMiddleware(e, container)

	middleware.InitSessionMiddleware(e, container)
	return e, container
}

// PrepareForServiceTest func prepares the services for testing.
func PrepareForServiceTest() container.Container {
	conf := createConfig(false)
	logger := initTestLogger()
	container := initContainer(conf, logger)

	migration.CreateDatabase(container)
	migration.InitMasterData(container)

	return container
}

func createConfig(isSecurity bool) *config.Config {
	conf := &config.Config{}
	conf.Database.Dialect = "sqlite3"
	conf.Database.Host = "file::memory:?cache=shared"
	conf.Database.Migration = true
	conf.Extension.MasterGenerator = true
	conf.Extension.SecurityEnabled = isSecurity
	conf.Log.RequestLogFormat = "${remote_ip} ${User_name} ${uri} ${method} ${status}"
	return conf
}

func initContainer(conf *config.Config, logger logger.Logger) container.Container {
	rep := repository.NewRepository(logger, conf)
	sess := session.NewSession(logger, conf)
	messages := map[string]string{
		"ValidationErrMessageBookTitle": "Please enter the title with 3 to 50 characters.",
		"ValidationErrMessageBookISBN":  "Please enter the ISBN with 10 to 20 characters."}
	container := container.NewContainer(rep, sess, conf, messages, logger, "test")
	return container
}

func initTestLogger() logger.Logger {
	myConfig := createLoggerConfig()
	zap, err := myConfig.Build()
	if err != nil {
		fmt.Printf("Error")
	}
	sugar := zap.Sugar()
	// set package varriable logger.
	logger := logger.NewLogger(sugar)
	logger.GetZapLogger().Infof("Success to read zap logger configuration")
	_ = zap.Sync()
	return logger
}

func createLoggerConfig() zap.Config {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	return zap.Config{
		Level:       level,
		Encoding:    "console",
		Development: true,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// ConvertToString func converts model to string.
func ConvertToString(model interface{}) string {
	bytes, _ := json.Marshal(model)
	return string(bytes)
}

// NewJSONRequest func creates a new request using JSON format.
func NewJSONRequest(method string, target string, param interface{}) *http.Request {
	req := httptest.NewRequest(method, target, strings.NewReader(ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

// GetCookie func gets a cookie from a HTTP request.
func GetCookie(rec *httptest.ResponseRecorder, cookieName string) string {
	parser := &http.Request{Header: http.Header{"Cookie": rec.Header()["Set-Cookie"]}}
	if cookie, err := parser.Cookie(cookieName); cookie != nil && err == nil {
		return cookie.Value
	}
	return ""
}
