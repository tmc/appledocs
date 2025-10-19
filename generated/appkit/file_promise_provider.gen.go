// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilePromiseProvider] class.
var (
	filePromiseProviderClass     _FilePromiseProviderClass
	filePromiseProviderClassOnce sync.Once
)

func getFilePromiseProviderClass() _FilePromiseProviderClass {
	filePromiseProviderClassOnce.Do(func() {
		filePromiseProviderClass = _FilePromiseProviderClass{objc.GetClass("NSFilePromiseProvider")}
	})
	return filePromiseProviderClass
}

type _FilePromiseProviderClass struct {
	class objc.Class
}

// An interface definition for the [FilePromiseProvider] class.
type IFilePromiseProvider interface {
	objectivec.IObject
}

// An object that provides a promise for the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider
type FilePromiseProvider struct {
	objectivec.Object
}

// FilePromiseProviderFrom constructs a [FilePromiseProvider] from an unsafe.Pointer.
//
// An object that provides a promise for the pasteboard.
func FilePromiseProviderFrom(ptr unsafe.Pointer) FilePromiseProvider {
	return FilePromiseProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilePromiseProviderClass) Alloc() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilePromiseProviderClass) New() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilePromiseProvider) Init() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilePromiseProvider) Autorelease() FilePromiseProvider {
	rv := objc.Send[FilePromiseProvider](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilePromiseProvider creates a new FilePromiseProvider instance.
func NewFilePromiseProvider() FilePromiseProvider {
	return getFilePromiseProviderClass().New()
}




