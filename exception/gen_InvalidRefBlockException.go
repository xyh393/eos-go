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

var InvalidRefBlockExceptionName = reflect.TypeOf(InvalidRefBlockException{}).Name()

type InvalidRefBlockException struct {
	_TransactionException
	Elog log.Messages
}

func NewInvalidRefBlockException(parent _TransactionException, message log.Message) *InvalidRefBlockException {
	return &InvalidRefBlockException{parent, log.Messages{message}}
}

func (e InvalidRefBlockException) Code() int64 {
	return 3040007
}

func (e InvalidRefBlockException) Name() string {
	return InvalidRefBlockExceptionName
}

func (e InvalidRefBlockException) What() string {
	return "Invalid Reference Block"
}

func (e *InvalidRefBlockException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e InvalidRefBlockException) GetLog() log.Messages {
	return e.Elog
}

func (e InvalidRefBlockException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e InvalidRefBlockException) DetailMessage() string {
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

func (e InvalidRefBlockException) String() string {
	return e.DetailMessage()
}

func (e InvalidRefBlockException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3040007,
		Name: InvalidRefBlockExceptionName,
		What: "Invalid Reference Block",
	}

	return json.Marshal(except)
}

func (e InvalidRefBlockException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*InvalidRefBlockException):
		callback(&e)
		return true
	case func(InvalidRefBlockException):
		callback(e)
		return true
	default:
		return false
	}
}
