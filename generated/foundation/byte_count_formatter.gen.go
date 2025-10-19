// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ByteCountFormatter] class.
var (
	byteCountFormatterClass     _ByteCountFormatterClass
	byteCountFormatterClassOnce sync.Once
)

func getByteCountFormatterClass() _ByteCountFormatterClass {
	byteCountFormatterClassOnce.Do(func() {
		byteCountFormatterClass = _ByteCountFormatterClass{objc.GetClass("NSByteCountFormatter")}
	})
	return byteCountFormatterClass
}

type _ByteCountFormatterClass struct {
	class objc.Class
}

// An interface definition for the [ByteCountFormatter] class.
type IByteCountFormatter interface {
	IFormatter
}

// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ByteCountFormatter
type ByteCountFormatter struct {
	Formatter
}

// ByteCountFormatterFrom constructs a [ByteCountFormatter] from an unsafe.Pointer.
//
// A formatter that converts a byte count value into a localized description that is formatted with the appropriate byte modifier (KB, MB, GB and so on).
func ByteCountFormatterFrom(ptr unsafe.Pointer) ByteCountFormatter {
	return ByteCountFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ByteCountFormatterClass) Alloc() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ByteCountFormatterClass) New() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ByteCountFormatter) Init() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ByteCountFormatter) Autorelease() ByteCountFormatter {
	rv := objc.Send[ByteCountFormatter](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewByteCountFormatter creates a new ByteCountFormatter instance.
func NewByteCountFormatter() ByteCountFormatter {
	return getByteCountFormatterClass().New()
}




