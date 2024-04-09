package f_logger

import "testing"

func TestNewZeroLogger(t *testing.T) {
	logger, err := NewZeroLogger(ZeroLoggerOptions{})
	if err != nil {
		t.Error(err)
	}
	logger.Err(nil).Msg("test")
}
