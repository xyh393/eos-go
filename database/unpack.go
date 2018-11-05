package database

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/eosspark/eos-go/log"
	"io"
	"io/ioutil"
	"reflect"
)

var (
	EOL                 = errors.New("rlp: end of list")
	ErrUnPointer        = errors.New("rlp: interface given to Decode must be a pointer")
	ErrElemTooLarge     = errors.New("rlp: element is larger than containing list")
	ErrValueTooLarge    = errors.New("rlp: value size exceeds available input length")
	ErrVarIntBufferSize = errors.New("rlp: invalid buffer size")
)

var TypeSize = struct {
	Bool        int
	Byte        int
	UInt8       int
	Int8        int
	UInt16      int
	Int16       int
	UInt32      int
	Int32       int
	UInt        int
	Int         int
	UInt64      int
	Int64       int
	SHA256Bytes int
}{
	Bool:        1,
	Byte:        1,
	UInt8:       1,
	Int8:        1,
	UInt16:      2,
	Int16:       2,
	UInt32:      4,
	Int32:       4,
	UInt:        4,
	Int:         4,
	UInt64:      8,
	Int64:       8,
	SHA256Bytes: 32,
}

var (
	optional           bool
	vuint32            bool
	eosArray           bool
	trxID              bool
	destaticVariantTag uint8
	rlplog             log.Logger
)

// Decoder implements the EOS unpacking, similar to FC_BUFFER
type decoder struct {
	data []byte
	pos  int
}

func init() {
	rlplog = log.New("rlp_db")
	rlplog.SetHandler(log.TerminalHandler)
}

func Decode(r io.Reader, val interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return newDecoder(data).decode(val)
}

func DecodeBytes(b []byte, val interface{}) error {
	err := newDecoder(b).decode(val)
	if err != nil {
		return err
	}
	return nil
}

func newDecoder(data []byte) *decoder {
	return &decoder{
		data: data,
		pos:  0,
	}
}

func (d *decoder) decode(v interface{}) (err error) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	if !rv.CanAddr() {
		return ErrUnPointer
	}
	t := rv.Type()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		newRV := reflect.New(t)
		rv.Set(newRV)
		rv = reflect.Indirect(newRV)
	}

	switch t.Kind() {
	case reflect.String:
		s, err := d.readString()
		if err != nil {
			return err
		}
		rv.SetString(s)
		return err
	case reflect.Bool:
		var r bool
		r, err = d.readBool()
		rv.SetBool(r)
		return
	case reflect.Int:
		var n int
		n, err = d.readInt()
		rv.SetInt(int64(n))
		return
	case reflect.Int8:
		var n int8
		n, err = d.readInt8()
		rv.SetInt(int64(n))
		return
	case reflect.Int16:
		var n int16
		n, err = d.readInt16()
		rv.SetInt(int64(n))
		return
	case reflect.Int32:
		var n int32
		n, err = d.readInt32()
		rv.SetInt(int64(n))
		return
	case reflect.Int64:
		var n int64
		n, err = d.readInt64()
		rv.SetInt(int64(n))
		return
	case reflect.Uint:
		var n uint
		n, err = d.readUint()
		rv.SetUint(uint64(n))
		return
	case reflect.Uint8:
		var n uint8
		n, err = d.readUint8()
		rv.SetUint(uint64(n))
		return
	case reflect.Uint16:
		var n uint16
		n, err = d.readUint16()
		rv.SetUint(uint64(n))
		return
	case reflect.Uint32:
		var n uint32
		n, err = d.readUint32()
		rv.SetUint(uint64(n))
		return
	case reflect.Uint64:
		var n uint64
		n, err = d.readUint64()
		rv.SetUint(n)
		return

	case reflect.Array:
		len := t.Len()

		if !eosArray {
			var l uint64
			if l, err = d.readUvarint(); err != nil {
				return
			}
			if int(l) != len {
				rlplog.Warn("the l is not equal to len of array")
			}
		}
		eosArray = false

		for i := 0; i < int(len); i++ {
			if err = d.decode(rv.Index(i).Addr().Interface()); err != nil {
				return
			}
		}
		return

	case reflect.Slice:
		var l uint64
		if l, err = d.readUvarint(); err != nil {
			return
		}
		rv.Set(reflect.MakeSlice(t, int(l), int(l)))
		for i := 0; i < int(l); i++ {
			if err = d.decode(rv.Index(i).Addr().Interface()); err != nil {
				return
			}
		}

	default:
		return errors.New("decode, unsupported type " + t.String())
	}

	return
}

func (d *decoder) readUvarint() (uint64, error) {
	l, read := binary.Uvarint(d.data[d.pos:])
	if read <= 0 {
		return l, ErrVarIntBufferSize
	}
	d.pos += read
	return l, nil
}

func (d *decoder) readByteArray() (out []byte, err error) {
	l, err := d.readUvarint()
	if err != nil {
		return nil, err
	}

	if len(d.data) < d.pos+int(l) {
		return nil, ErrValueTooLarge
	}

	out = d.data[d.pos : d.pos+int(l)]
	d.pos += int(l)

	return
}

func (d *decoder) readString() (out string, err error) {
	data, err := d.readByteArray()
	out = string(data)
	return
}

func (d *decoder) readByte() (out byte, err error) {
	if d.remaining() < TypeSize.Byte {
		err = fmt.Errorf("byte required [1] byte, remaining [%d]", d.remaining())
		return
	}

	out = d.data[d.pos]
	d.pos++
	return
}

func (d *decoder) readBool() (out bool, err error) {
	if d.remaining() < TypeSize.Bool {
		err = fmt.Errorf("rlp: bool required [%d] byte, remaining [%d]", TypeSize.Bool, d.remaining())
		return
	}

	b, err := d.readByte()
	if err != nil {
		err = fmt.Errorf("readBool, %s", err)
	}
	out = b != 0
	return

}
func (d *decoder) readUint8() (out byte, err error) {
	if d.remaining() < TypeSize.UInt8 {
		err = fmt.Errorf("rlp: byte required [1] byte, remaining [%d]", d.remaining())
		return
	}
	out = d.data[d.pos]
	d.pos++
	return
}
func (d *decoder) readUint16() (out uint16, err error) {
	if d.remaining() < TypeSize.UInt16 {
		err = fmt.Errorf("rlp: uint16 required [%d] bytes, remaining [%d]", TypeSize.UInt16, d.remaining())
		return
	}

	out = binary.BigEndian.Uint16(d.data[d.pos:])
	d.pos += TypeSize.UInt16
	return
}
func (d *decoder) readUint32() (out uint32, err error) {
	if d.remaining() < TypeSize.UInt32 {
		err = fmt.Errorf("rlp: uint32 required [%d] bytes, remaining [%d]", TypeSize.UInt32, d.remaining())
		return
	}

	out = binary.BigEndian.Uint32(d.data[d.pos:])
	d.pos += TypeSize.UInt32
	return
}
func (d *decoder) readUint() (out uint, err error) {
	if d.remaining() < TypeSize.UInt {
		err = fmt.Errorf("rlp: uint required [%d] bytes, remaining [%d]", TypeSize.UInt, d.remaining())
		return
	}

	out = uint(binary.BigEndian.Uint32(d.data[d.pos:]))
	d.pos += TypeSize.UInt
	return
}
func (d *decoder) readUint64() (out uint64, err error) {
	if d.remaining() < TypeSize.UInt64 {
		err = fmt.Errorf("rlp: uint64 required [%d] bytes, remaining [%d]", TypeSize.UInt64, d.remaining())
		return
	}

	data := d.data[d.pos : d.pos+TypeSize.UInt64]
	out = binary.BigEndian.Uint64(data)
	d.pos += TypeSize.UInt64
	return
}

func (d *decoder) readInt8() (out int8, err error) {
	n, err := d.readUint8()
	out = int8(n)
	return
}

func (d *decoder) readInt16() (out int16, err error) {
	n, err := d.readUint16()
	out = int16(n)
	return
}
func (d *decoder) readInt32() (out int32, err error) {
	n, err := d.readUint32()
	out = int32(n)
	return
}
func (d *decoder) readInt() (out int, err error) {
	n, err := d.readUint()
	out = int(n)
	return
}
func (d *decoder) readInt64() (out int64, err error) {
	n, err := d.readUint64()
	out = int64(n)
	return
}

func (d *decoder) remaining() int {
	return len(d.data) - d.pos
}