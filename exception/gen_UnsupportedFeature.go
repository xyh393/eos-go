// Code generated by gotemplate. DO NOT EDIT.

package exception

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/eosspark/eos-go/log"
)

// template type Exception(PARENT,CODE,WHAT)

var UnsupportedFeatureName = reflect.TypeOf(UnsupportedFeature{}).Name()

type UnsupportedFeature struct {
	_MiscException
	Elog log.Messages
}

func NewUnsupportedFeature(parent _MiscException, message log.Message) *UnsupportedFeature {
	return &UnsupportedFeature{parent, log.Messages{message}}
}

func (e UnsupportedFeature) Code() int64 {
	return 3100008
}

func (e UnsupportedFeature) Name() string {
	return UnsupportedFeatureName
}

func (e UnsupportedFeature) What() string {
	return "Feature is currently unsupported"
}

func (e *UnsupportedFeature) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e UnsupportedFeature) GetLog() log.Messages {
	return e.Elog
}

func (e UnsupportedFeature) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e UnsupportedFeature) DetailMessage() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(e.Code())))
	buffer.WriteByte(' ')
	buffer.WriteString(e.Name())
	buffer.Write([]byte{':', ' '})
	buffer.WriteString(e.What())
	buffer.WriteByte('\n')
	for _, l := range e.Elog {
		buffer.WriteByte('[')
		buffer.WriteString(l.GetMessage())
		buffer.Write([]byte{']', ' '})
		buffer.WriteString(l.GetContext().String())
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func (e UnsupportedFeature) String() string {
	return e.DetailMessage()
}

func (e UnsupportedFeature) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3100008,
		Name: UnsupportedFeatureName,
		What: "Feature is currently unsupported",
	}

	return json.Marshal(except)
}

func (e UnsupportedFeature) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*UnsupportedFeature):
		callback(&e)
		return true
	case func(UnsupportedFeature):
		callback(e)
		return true
	default:
		return false
	}
}
