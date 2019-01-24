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

var ProducerNotInScheduleName = reflect.TypeOf(ProducerNotInSchedule{}).Name()

type ProducerNotInSchedule struct {
	_ProducerException
	Elog log.Messages
}

func NewProducerNotInSchedule(parent _ProducerException, message log.Message) *ProducerNotInSchedule {
	return &ProducerNotInSchedule{parent, log.Messages{message}}
}

func (e ProducerNotInSchedule) Code() int64 {
	return 3170006
}

func (e ProducerNotInSchedule) Name() string {
	return ProducerNotInScheduleName
}

func (e ProducerNotInSchedule) What() string {
	return "The producer is not part of current schedule"
}

func (e *ProducerNotInSchedule) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e ProducerNotInSchedule) GetLog() log.Messages {
	return e.Elog
}

func (e ProducerNotInSchedule) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); len(msg) > 0 {
			return msg
		}
	}
	return e.String()
}

func (e ProducerNotInSchedule) DetailMessage() string {
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

func (e ProducerNotInSchedule) String() string {
	return e.DetailMessage()
}

func (e ProducerNotInSchedule) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3170006,
		Name: ProducerNotInScheduleName,
		What: "The producer is not part of current schedule",
	}

	return json.Marshal(except)
}

func (e ProducerNotInSchedule) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*ProducerNotInSchedule):
		callback(&e)
		return true
	case func(ProducerNotInSchedule):
		callback(e)
		return true
	default:
		return false
	}
}
