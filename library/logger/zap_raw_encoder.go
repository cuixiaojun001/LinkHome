package log

import (
	"sync"
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var bufferPool = buffer.NewPool()

var _zapRawPool = sync.Pool{New: func() interface{} {
	return &zapRawEncoder{}
}}

func getZapRawEncoder() *zapRawEncoder {
	return _zapRawPool.Get().(*zapRawEncoder)
}

func putZapRawEncoder(enc *zapRawEncoder) {
	enc.EncoderConfig = nil
	enc.buf = nil
	_zapRawPool.Put(enc)
}

type zapRawEncoder struct {
	*zapcore.EncoderConfig
	buf *buffer.Buffer
}

// NewzapRawEncoder creates a key=value encoder
func NewZapRawEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return &zapRawEncoder{
		EncoderConfig: &cfg,
		buf:           bufferPool.Get(),
	}
}

// 下面为Encoder的ObjectEncoder接口，该Encoder不需要，都没有内容，后续可以参考JsonEncoder的事项
func (enc *zapRawEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error   { return nil }
func (enc *zapRawEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error { return nil }
func (enc *zapRawEncoder) AddBinary(key string, value []byte)                            {}
func (enc *zapRawEncoder) AddByteString(key string, value []byte)                        {}
func (enc *zapRawEncoder) AddBool(key string, value bool)                                {}
func (enc *zapRawEncoder) AddComplex128(key string, value complex128)                    {}
func (enc *zapRawEncoder) AddComplex64(key string, value complex64)                      {}
func (enc *zapRawEncoder) AddDuration(key string, value time.Duration)                   {}
func (enc *zapRawEncoder) AddFloat64(key string, value float64)                          {}
func (enc *zapRawEncoder) AddFloat32(key string, value float32)                          {}
func (enc *zapRawEncoder) AddInt(key string, value int)                                  {}
func (enc *zapRawEncoder) AddInt64(key string, value int64)                              {}
func (enc *zapRawEncoder) AddInt32(key string, value int32)                              {}
func (enc *zapRawEncoder) AddInt16(key string, value int16)                              {}
func (enc *zapRawEncoder) AddInt8(key string, value int8)                                {}
func (enc *zapRawEncoder) AddString(key, value string)                                   {}
func (enc *zapRawEncoder) AddTime(key string, value time.Time)                           {}
func (enc *zapRawEncoder) AddUint(key string, value uint)                                {}
func (enc *zapRawEncoder) AddUint64(key string, value uint64)                            {}
func (enc *zapRawEncoder) AddUint32(key string, value uint32)                            {}
func (enc *zapRawEncoder) AddUint16(key string, value uint16)                            {}
func (enc *zapRawEncoder) AddUint8(key string, value uint8)                              {}
func (enc *zapRawEncoder) AddUintptr(key string, value uintptr)                          {}
func (enc *zapRawEncoder) AddReflected(key string, value interface{}) error              { return nil }
func (enc *zapRawEncoder) OpenNamespace(key string)                                      {}

func (enc *zapRawEncoder) Clone() zapcore.Encoder {
	clone := enc.clone()
	clone.buf.Write(enc.buf.Bytes())
	return clone
}

func (enc *zapRawEncoder) clone() *zapRawEncoder {
	clone := getZapRawEncoder()
	clone.EncoderConfig = enc.EncoderConfig
	clone.buf = bufferPool.Get()
	return clone
}

func (enc *zapRawEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	final := enc.clone()
	final.buf.AppendString(ent.Message)
	if final.LineEnding != "" {
		final.buf.AppendString(final.LineEnding)
	} else {
		final.buf.AppendString(zapcore.DefaultLineEnding)
	}
	ret := final.buf
	putZapRawEncoder(final)
	return ret, nil
}
