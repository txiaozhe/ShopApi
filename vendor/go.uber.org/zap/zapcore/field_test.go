// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zapcore_test

import (
	"errors"
	"fmt"
	"io"
	"math"
	"testing"
	"time"

	"go.uber.org/zap"

	richErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "go.uber.org/zap/zapcore"
)

type users int

func (u users) String() string {
	return fmt.Sprintf("%d users", int(u))
}

func (u users) Error() string {
	return fmt.Sprintf("%d too many users", int(u))
}

func (u users) Format(s fmt.State, verb rune) {
	// Implement fmt.Formatter, but don't add any information beyond the basic
	// Error method.
	if verb == 'v' && s.Flag('+') {
		io.WriteString(s, u.Error())
	}
}

func (u users) MarshalLogObject(enc ObjectEncoder) error {
	if int(u) < 0 {
		return errors.New("too few users")
	}
	enc.AddInt("users", int(u))
	return nil
}

func (u users) MarshalLogArray(enc ArrayEncoder) error {
	if int(u) < 0 {
		return errors.New("too few users")
	}
	for i := 0; i < int(u); i++ {
		enc.AppendString("user")
	}
	return nil
}

func TestUnknownFieldType(t *testing.T) {
	unknown := Field{Key: "k", String: "foo"}
	assert.Equal(t, UnknownType, unknown.Type, "Expected zero value of FieldType to be UnknownType.")
	assert.Panics(t, func() {
		unknown.AddTo(NewMapObjectEncoder())
	}, "Expected using a field with unknown type to panic.")
}

func TestFieldAddingError(t *testing.T) {
	tests := []struct {
		t    FieldType
		want interface{}
	}{
		{ArrayMarshalerType, []interface{}(nil)},
		{ObjectMarshalerType, map[string]interface{}{}},
	}
	for _, tt := range tests {
		f := Field{Key: "k", Interface: users(-1), Type: tt.t}
		enc := NewMapObjectEncoder()
		assert.NotPanics(t, func() { f.AddTo(enc) }, "Unexpected panic when adding fields returns an error.")
		assert.Equal(t, tt.want, enc.Fields["k"], "On error, expected zero value in field.Key.")
		assert.Equal(t, "too few users", enc.Fields["kError"], "Expected error message in log context.")
	}
}

func TestFields(t *testing.T) {
	tests := []struct {
		t     FieldType
		i     int64
		s     string
		iface interface{}
		want  interface{}
	}{
		{t: ArrayMarshalerType, iface: users(2), want: []interface{}{"user", "user"}},
		{t: ObjectMarshalerType, iface: users(2), want: map[string]interface{}{"users": 2}},
		{t: BinaryType, iface: []byte("foo"), want: []byte("foo")},
		{t: BoolType, i: 0, want: false},
		{t: ByteStringType, iface: []byte("foo"), want: []byte("foo")},
		{t: Complex128Type, iface: 1 + 2i, want: 1 + 2i},
		{t: Complex64Type, iface: complex64(1 + 2i), want: complex64(1 + 2i)},
		{t: DurationType, i: 1000, want: time.Microsecond},
		{t: Float64Type, i: int64(math.Float64bits(3.14)), want: 3.14},
		{t: Float32Type, i: int64(math.Float32bits(3.14)), want: float32(3.14)},
		{t: Int64Type, i: 42, want: int64(42)},
		{t: Int32Type, i: 42, want: int32(42)},
		{t: Int16Type, i: 42, want: int16(42)},
		{t: Int8Type, i: 42, want: int8(42)},
		{t: StringType, s: "foo", want: "foo"},
		{t: TimeType, i: 1000, iface: time.UTC, want: time.Unix(0, 1000).In(time.UTC)},
		{t: Uint64Type, i: 42, want: uint64(42)},
		{t: Uint32Type, i: 42, want: uint32(42)},
		{t: Uint16Type, i: 42, want: uint16(42)},
		{t: Uint8Type, i: 42, want: uint8(42)},
		{t: UintptrType, i: 42, want: uintptr(42)},
		{t: ReflectType, iface: users(2), want: users(2)},
		{t: NamespaceType, want: map[string]interface{}{}},
		{t: StringerType, iface: users(2), want: "2 users"},
		{t: ErrorType, iface: users(2), want: "2 too many users"},
		{t: SkipType, want: interface{}(nil)},
	}

	for _, tt := range tests {
		enc := NewMapObjectEncoder()
		f := Field{Key: "k", Type: tt.t, Integer: tt.i, Interface: tt.iface, String: tt.s}
		f.AddTo(enc)
		assert.Equal(t, tt.want, enc.Fields["k"], "Unexpected output from field %+v.", f)

		delete(enc.Fields, "k")
		assert.Equal(t, 0, len(enc.Fields), "Unexpected extra fields present.")

		assert.True(t, f.Equals(f), "Field does not equal itself")
	}
}

func TestRichErrorSupport(t *testing.T) {
	f := Field{
		Type:      ErrorType,
		Interface: richErrors.WithMessage(richErrors.New("egad"), "failed"),
		Key:       "k",
	}
	enc := NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, "failed: egad", enc.Fields["k"], "Unexpected basic error message.")

	serialized := enc.Fields["kVerbose"]
	// Don't assert the exact format used by a third-party package, but ensure
	// that some critical elements are present.
	assert.Regexp(t, `egad`, serialized, "Expected original error message to be present.")
	assert.Regexp(t, `failed`, serialized, "Expected error annotation to be present.")
	assert.Regexp(t, `TestRichErrorSupport`, serialized, "Expected calling function to be present in stacktrace.")
}

func TestEquals(t *testing.T) {
	tests := []struct {
		a, b Field
		want bool
	}{
		{
			a:    zap.Int16("a", 1),
			b:    zap.Int32("a", 1),
			want: false,
		},
		{
			a:    zap.String("k", "a"),
			b:    zap.String("k", "a"),
			want: true,
		},
		{
			a:    zap.String("k", "a"),
			b:    zap.String("k2", "a"),
			want: false,
		},
		{
			a:    zap.String("k", "a"),
			b:    zap.String("k", "b"),
			want: false,
		},
		{
			a:    zap.Time("k", time.Unix(1000, 1000)),
			b:    zap.Time("k", time.Unix(1000, 1000)),
			want: true,
		},
		{
			a:    zap.Time("k", time.Unix(1000, 1000).In(time.UTC)),
			b:    zap.Time("k", time.Unix(1000, 1000).In(time.FixedZone("TEST", -8))),
			want: false,
		},
		{
			a:    zap.Time("k", time.Unix(1000, 1000)),
			b:    zap.Time("k", time.Unix(1000, 2000)),
			want: false,
		},
		{
			a:    zap.Binary("k", []byte{1, 2}),
			b:    zap.Binary("k", []byte{1, 2}),
			want: true,
		},
		{
			a:    zap.Binary("k", []byte{1, 2}),
			b:    zap.Binary("k", []byte{1, 3}),
			want: false,
		},
		{
			a:    zap.ByteString("k", []byte("abc")),
			b:    zap.ByteString("k", []byte("abc")),
			want: true,
		},
		{
			a:    zap.ByteString("k", []byte("abc")),
			b:    zap.ByteString("k", []byte("abd")),
			want: false,
		},
		{
			a:    zap.Ints("k", []int{1, 2}),
			b:    zap.Ints("k", []int{1, 2}),
			want: true,
		},
		{
			a:    zap.Ints("k", []int{1, 2}),
			b:    zap.Ints("k", []int{1, 3}),
			want: false,
		},
		{
			a:    zap.Object("k", users(10)),
			b:    zap.Object("k", users(10)),
			want: true,
		},
		{
			a:    zap.Object("k", users(10)),
			b:    zap.Object("k", users(20)),
			want: false,
		},
		{
			a:    zap.Any("k", map[string]string{"a": "b"}),
			b:    zap.Any("k", map[string]string{"a": "b"}),
			want: true,
		},
		{
			a:    zap.Any("k", map[string]string{"a": "b"}),
			b:    zap.Any("k", map[string]string{"a": "d"}),
			want: false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.a.Equals(tt.b), "a.Equals(b) a: %#v b: %#v", tt.a, tt.b)
		assert.Equal(t, tt.want, tt.b.Equals(tt.a), "b.Equals(a) a: %#v b: %#v", tt.a, tt.b)
	}
}
