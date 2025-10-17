// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Unarchiver] class.
var unarchiverClass = _UnarchiverClass{objc.GetClass("NSUnarchiver")}

type _UnarchiverClass struct {
	class objc.Class
}

// An interface definition for the [Unarchiver] class.
type IUnarchiver interface {
	ICoder
}

// A decoder that restores data from an archive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver

type Unarchiver struct {
	Coder
}

// UnarchiverFrom constructs a [Unarchiver] from an unsafe.Pointer.
//
// A decoder that restores data from an archive.
func UnarchiverFrom(ptr unsafe.Pointer) Unarchiver {
	return Unarchiver{
		Coder: CoderFrom(ptr),
	}
}



