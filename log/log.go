package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

func init() {

	hook := lumberjack.Logger{
		Filename: "./logs/spikeProxy.log",
		MaxSize: 128,
		MaxBackups: 30,
		MaxAge: 7,
		Compress: true,
	}

	encoderConfig:=zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "line-num",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(/*zapcore.AddSync(os.Stdout),*/zapcore.AddSync(&hook)),
		atomicLevel)

	caller := zap.AddCaller()

	development := zap.Development()

	filed := zap.Fields(zap.String("service-name","rpcx-client"))

	Logger = zap.New(core, caller, development, filed)
}
