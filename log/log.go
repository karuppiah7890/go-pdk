package log

import (
	"github.com/karuppiah7890/go-pdk/bridge"
)

type Log struct {
	bridge.PdkBridge
}

func New(ch chan interface{}) Log {
	return Log{bridge.New(ch)}
}

func (r Log) Alert(args ...interface{}) error {
	_, err := r.Ask(`kong.log.alert`, args...)
	return err
}

func (r Log) Crit(args ...interface{}) error {
	_, err := r.Ask(`kong.log.crit`, args...)
	return err
}

func (r Log) Err(args ...interface{}) error {
	_, err := r.Ask(`kong.log.err`, args...)
	return err
}

func (r Log) Warn(args ...interface{}) error {
	_, err := r.Ask(`kong.log.warn`, args...)
	return err
}

func (r Log) Notice(args ...interface{}) error {
	_, err := r.Ask(`kong.log.notice`, args...)
	return err
}

func (r Log) Info(args ...interface{}) error {
	_, err := r.Ask(`kong.log.info`, args...)
	return err
}

func (r Log) Debug(args ...interface{}) error {
	_, err := r.Ask(`kong.log.debug`, args...)
	return err
}

func (r Log) Serialize() (s string, err error) {
	return r.AskString(`kong.log.serialize`)
}
