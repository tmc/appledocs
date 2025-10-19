// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileVersion] class.
var (
	fileVersionClass     _FileVersionClass
	fileVersionClassOnce sync.Once
)

func getFileVersionClass() _FileVersionClass {
	fileVersionClassOnce.Do(func() {
		fileVersionClass = _FileVersionClass{objc.GetClass("NSFileVersion")}
	})
	return fileVersionClass
}

type _FileVersionClass struct {
	class objc.Class
}

// An interface definition for the [FileVersion] class.
type IFileVersion interface {
	objectivec.IObject
}

// A snapshot of a file at a specific point in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileVersion
type FileVersion struct {
	objectivec.Object
}

// FileVersionFrom constructs a [FileVersion] from an unsafe.Pointer.
//
// A snapshot of a file at a specific point in time.
func FileVersionFrom(ptr unsafe.Pointer) FileVersion {
	return FileVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileVersionClass) Alloc() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileVersionClass) New() FileVersion {
	rv := objc.Send[FileVersion](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileVersion) Init() FileVersion {
	rv := objc.Send[FileVersion](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileVersion) Autorelease() FileVersion {
	rv := objc.Send[FileVersion](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileVersion creates a new FileVersion instance.
func NewFileVersion() FileVersion {
	return getFileVersionClass().New()
}




