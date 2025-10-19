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

// An interface definition for the [Archiver] class.
type IArchiver interface {
	ICoder
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
// Alloc allocates a new instance without initialization.
func (ac _ArchiverClass) Alloc() Archiver {
	rv := objc.Send[Archiver](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _ArchiverClass) New() Archiver {
	rv := objc.Send[Archiver](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Archiver) Init() Archiver {
	rv := objc.Send[Archiver](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Archiver) Autorelease() Archiver {
	rv := objc.Send[Archiver](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArchiver creates a new Archiver instance.
func NewArchiver() Archiver {
	return archiverClass.New()
}




