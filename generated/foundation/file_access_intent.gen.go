// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileAccessIntent] class.
var (
	fileAccessIntentClass     _FileAccessIntentClass
	fileAccessIntentClassOnce sync.Once
)

func getFileAccessIntentClass() _FileAccessIntentClass {
	fileAccessIntentClassOnce.Do(func() {
		fileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}
	})
	return fileAccessIntentClass
}

type _FileAccessIntentClass struct {
	class objc.Class
}

// An interface definition for the [FileAccessIntent] class.
type IFileAccessIntent interface {
	objectivec.IObject
}

// The details of a coordinated-read or coordinated-write operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent
type FileAccessIntent struct {
	objectivec.Object
}

// FileAccessIntentFrom constructs a [FileAccessIntent] from an unsafe.Pointer.
//
// The details of a coordinated-read or coordinated-write operation.
func FileAccessIntentFrom(ptr unsafe.Pointer) FileAccessIntent {
	return FileAccessIntent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileAccessIntentClass) Alloc() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileAccessIntentClass) New() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileAccessIntent) Init() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileAccessIntent) Autorelease() FileAccessIntent {
	rv := objc.Send[FileAccessIntent](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileAccessIntent creates a new FileAccessIntent instance.
func NewFileAccessIntent() FileAccessIntent {
	return getFileAccessIntentClass().New()
}




