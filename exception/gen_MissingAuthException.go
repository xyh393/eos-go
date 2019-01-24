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

var MissingAuthExceptionName = reflect.TypeOf(MissingAuthException{}).Name()

type MissingAuthException struct {
	_AuthorizationException
	Elog log.Messages
}

func NewMissingAuthException(parent _AuthorizationException, message log.Message) *MissingAuthException {
	return &MissingAuthException{parent, log.Messages{message}}
}

func (e MissingAuthException) Code() int64 {
	return 3090004
}

func (e MissingAuthException) Name() string {
	return MissingAuthExceptionName
}

func (e MissingAuthException) What() string {
	return "Missing required authority"
}

func (e *MissingAuthException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e MissingAuthException) GetLog() log.Messages {
	return e.Elog
}

func (e MissingAuthException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e MissingAuthException) DetailMessage() string {
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

func (e MissingAuthException) String() string {
	return e.DetailMessage()
}

func (e MissingAuthException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3090004,
		Name: MissingAuthExceptionName,
		What: "Missing required authority",
	}

	return json.Marshal(except)
}

func (e MissingAuthException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*MissingAuthException):
		callback(&e)
		return true
	case func(MissingAuthException):
		callback(e)
		return true
	default:
		return false
	}
}
