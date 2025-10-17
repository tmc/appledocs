// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileVersion] class.
var fileVersionClass = _FileVersionClass{objc.GetClass("NSFileVersion")}

type _FileVersionClass struct {
	class objc.Class
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



