package errors

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

var (
	_messages            atomic.Value
	_codes               = map[int]struct{}{} // register codes.
	ArchiveTypeForbidAdd = New(10001)
	ArchiveExist         = New(10002)
	ArchiveNotExist      = New(10003)
	ArchiveAlreadyDel    = New(10004)
)

func CodeString(errorNo Code) string {
	switch errorNo {
	case ArchiveTypeForbidAdd:
		return "success"
	case ArchiveExist:
		return "data error"
	case ArchiveNotExist:
		return "http request fail"
	case ArchiveAlreadyDel:
		return "param unexpected"
	default:
		return "system error"
	}
}

type argError struct {
	arg int
	msg string
}

type Code int

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

// Code return error code
func (e Code) Code() int {
	return int(e)
}

// Messsage return error message
func (e Code) Message() string {
	if cm, ok := _messages.Load().(map[int]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return e.Error()
}

func Int(i int) Code {
	return Code(i)
}

func New(e int) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e int) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return Int(e)
}

func ReturnError(e Code) Code {
	return e
}

func main() {
	a := ReturnError(ArchiveExist)
	b := ReturnError(ArchiveAlreadyDel)
	c := ReturnError(ArchiveNotExist)
	fmt.Println(a, CodeString(a))
	fmt.Println(b, CodeString(b))
	fmt.Println(c, CodeString(c))
}
