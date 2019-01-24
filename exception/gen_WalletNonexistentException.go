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

var WalletNonexistentExceptionName = reflect.TypeOf(WalletNonexistentException{}).Name()

type WalletNonexistentException struct {
	_WalletException
	Elog log.Messages
}

func NewWalletNonexistentException(parent _WalletException, message log.Message) *WalletNonexistentException {
	return &WalletNonexistentException{parent, log.Messages{message}}
}

func (e WalletNonexistentException) Code() int64 {
	return 3120002
}

func (e WalletNonexistentException) Name() string {
	return WalletNonexistentExceptionName
}

func (e WalletNonexistentException) What() string {
	return "Nonexistent wallet"
}

func (e *WalletNonexistentException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e WalletNonexistentException) GetLog() log.Messages {
	return e.Elog
}

func (e WalletNonexistentException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e WalletNonexistentException) DetailMessage() string {
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

func (e WalletNonexistentException) String() string {
	return e.DetailMessage()
}

func (e WalletNonexistentException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3120002,
		Name: WalletNonexistentExceptionName,
		What: "Nonexistent wallet",
	}

	return json.Marshal(except)
}

func (e WalletNonexistentException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*WalletNonexistentException):
		callback(&e)
		return true
	case func(WalletNonexistentException):
		callback(e)
		return true
	default:
		return false
	}
}
