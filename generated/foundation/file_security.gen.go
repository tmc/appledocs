// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileSecurity] class.
var fileSecurityClass = _FileSecurityClass{objc.GetClass("NSFileSecurity")}

type _FileSecurityClass struct {
	class objc.Class
}

// An interface definition for the [FileSecurity] class.
type IFileSecurity interface {
	objectivec.IObject
}

// A stub class that encapsulates security information about a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileSecurity

type FileSecurity struct {
	objectivec.Object
}

// FileSecurityFrom constructs a [FileSecurity] from an unsafe.Pointer.
//
// A stub class that encapsulates security information about a file.
func FileSecurityFrom(ptr unsafe.Pointer) FileSecurity {
	return FileSecurity{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (fc _FileSecurityClass) Alloc() FileSecurity {
	rv := objc.Send[FileSecurity](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (fc _FileSecurityClass) New() FileSecurity {
	rv := objc.Send[FileSecurity](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileSecurity) Init() FileSecurity {
	rv := objc.Send[FileSecurity](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileSecurity) Autorelease() FileSecurity {
	rv := objc.Send[FileSecurity](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileSecurity creates a new FileSecurity instance.
func NewFileSecurity() FileSecurity {
	return fileSecurityClass.New()
}




