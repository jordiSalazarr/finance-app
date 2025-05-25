package zapLogger

import "go.uber.org/zap"

type ZapService struct {
	Logger *zap.Logger
}

func NewDevelopmentZapLogger() (*ZapService, error) {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	service := &ZapService{
		Logger: logger,
	}
	return service, nil
}

func fieldsToZapFields(fields ...interface{}) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key, ok := fields[i].(string)
			if ok {
				zapFields = append(zapFields, zap.Any(key, fields[i+1]))
			} else {
				zapFields = append(zapFields, zap.Any("field", fields[i]))
			}
		} else {
			zapFields = append(zapFields, zap.Any("field", fields[i]))
		}
	}
	return zapFields
}

func (z *ZapService) Info(msg string, fields ...interface{}) {
	zapFields := fieldsToZapFields(fields...)

	z.Logger.Info(msg, zapFields...)
}
func (z *ZapService) Error(msg string, fields ...interface{}) {
	zapFields := fieldsToZapFields(fields...)
	z.Logger.Error(msg, zapFields...)
}
func (z *ZapService) Debug(msg string, fields ...interface{}) {
	zapFields := fieldsToZapFields(fields...)

	z.Logger.Debug(msg, zapFields...)
}
func (z *ZapService) Warn(msg string, fields ...interface{}) {
	zapFields := fieldsToZapFields(fields...)

	z.Logger.Warn(msg, zapFields...)
}
func (z *ZapService) Close() error {
	return z.Logger.Sync()
}
