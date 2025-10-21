// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FSStatFSResult] class.
var (
	FSStatFSResultClass     _FSStatFSResultClass
	FSStatFSResultClassOnce sync.Once
)

func getFSStatFSResultClass() _FSStatFSResultClass {
	FSStatFSResultClassOnce.Do(func() {
		FSStatFSResultClass = _FSStatFSResultClass{objc.GetClass("FSStatFSResult")}
	})
	return FSStatFSResultClass
}

type _FSStatFSResultClass struct {
	class objc.Class
}

// An interface definition for the [FSStatFSResult] class.
type IFSStatFSResult interface {
	objectivec.IObject
}

// A type used to report a volume’s statistics.
//
// The names of this type’s properties match those in the structure in , which reports these values for an FSKit file system. All numeric properties default to . Override these values, unless a given property has no meaningful value to provide. For the read-only , set this value with the designated initializer.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSStatFSResult
type FSStatFSResult struct {
	objectivec.Object
}

// FSStatFSResultFrom constructs a [FSStatFSResult] from an unsafe.Pointer.
//
// A type used to report a volume’s statistics.
func FSStatFSResultFrom(ptr unsafe.Pointer) FSStatFSResult {
	return FSStatFSResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSStatFSResultClass) Alloc() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSStatFSResultClass) New() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSStatFSResult) Init() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSStatFSResult) Autorelease() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSStatFSResult creates a new FSStatFSResult instance.
func NewFSStatFSResult() FSStatFSResult {
	return getFSStatFSResultClass().New()
}




