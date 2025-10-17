// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Archiver] class.
var archiverClass = _ArchiverClass{objc.GetClass("NSArchiver")}

type _ArchiverClass struct {
	class objc.Class
}

// A coder that stores an object’s data to an archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver

type Archiver struct {
	Coder
}

// ArchiverFrom constructs a [Archiver] from an unsafe.Pointer.
//
// A coder that stores an object’s data to an archive.
func ArchiverFrom(ptr unsafe.Pointer) Archiver {
	return Archiver{
		Coder: CoderFrom(ptr),
	}
}



