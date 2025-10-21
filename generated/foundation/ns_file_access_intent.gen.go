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
	FileAccessIntentClass     _FileAccessIntentClass
	FileAccessIntentClassOnce sync.Once
)

func getFileAccessIntentClass() _FileAccessIntentClass {
	FileAccessIntentClassOnce.Do(func() {
		FileAccessIntentClass = _FileAccessIntentClass{objc.GetClass("NSFileAccessIntent")}
	})
	return FileAccessIntentClass
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
// Use this class when performing asynchronous operations with a file coordinator using the coordinator’s method.
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


// Returns a file access intent object for reading the given URL with the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent/readingIntent(with:options:)
func (fc _FileAccessIntentClass) ReadingIntentWithURLOptions(url IURL, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("readingIntentWithURL:options:"), url, options)
	return rv
}

// The current URL for the item managed by the file access intent instance. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileAccessIntent/url
func (f_ FileAccessIntent) URL() URL {
	rv := objc.Send[URL](f_.ID, objc.Sel("URL"))
	return rv
}



