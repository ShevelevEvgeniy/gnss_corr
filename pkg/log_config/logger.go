package log_config

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

type EscapeSeqJSONEncoder struct {
	zapcore.Encoder
}

func newLogger() *zap.Logger {
	encoder := &EscapeSeqJSONEncoder{
		Encoder: zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:      "ts",
			LevelKey:     "level",
			CallerKey:    "caller",
			MessageKey:   "msg",
			EncodeTime:   zapcore.TimeEncoderOfLayout(time.RFC3339),
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
		}),
	}

	core := zapcore.NewCore(
		encoder,
		os.Stdout,
		zapcore.InfoLevel,
	)

	return zap.New(core, zap.AddCaller())
}

func (enc *EscapeSeqJSONEncoder) Clone() zapcore.Encoder {
	return &EscapeSeqJSONEncoder{
		Encoder: enc.Encoder.Clone(),
	}
}

func (enc *EscapeSeqJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	output := bytes.Replace(b.Bytes(), []byte("\\n"), []byte("\n"), -1)

	outputStr := string(output)
	outputStr = strings.ReplaceAll(outputStr, "\n\\", "\n    ")
	outputStr = strings.ReplaceAll(outputStr, "\\", "")

	newb := buffer.NewPool().Get()
	newb.Write([]byte(outputStr))

	return newb, nil
}
