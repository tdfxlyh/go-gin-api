package caller

import (
	"fmt"
	"github.com/tdfxlyh/go-gin-api/internal/constdef"
	"github.com/tdfxlyh/go-gin-api/internal/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {
	Logger = setupLogger(constdef.LogPath)
	defer Logger.Sync()
}

func setupLogger(logPath string) *zap.Logger {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		if err := os.Mkdir(logPath, os.ModePerm); err != nil {
			fmt.Printf("failed to create directory: %s, err=%s\n", logPath, err)
			panic(err)
		}
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberJackLogger{logPath})),
		zap.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller())

	return logger
}

type lumberJackLogger struct {
	logPath string
}

func (l *lumberJackLogger) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s.log", l.logPath, utils.GetCurrentDate()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(p)
}
