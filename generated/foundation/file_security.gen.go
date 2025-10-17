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



