// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [unused2] class.
var (
	Unused2Class     _unused2Class
	Unused2ClassOnce sync.Once
)

func getunused2Class() _unused2Class {
	Unused2ClassOnce.Do(func() {
		Unused2Class = _unused2Class{objc.GetClass("unused2")}
	})
	return Unused2Class
}

type _unused2Class struct {
	class objc.Class
}





// An interface definition for the [unused2] class.
type Iunused2 interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (uc _unused2Class) Alloc() unused2 {
	rv := objc.Send[unused2](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _unused2Class) New() unused2 {
	rv := objc.Send[unused2](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ unused2) Init() unused2 {
	rv := objc.Send[unused2](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ unused2) Autorelease() unused2 {
	rv := objc.Send[unused2](u_.ID, objc.Sel("autorelease"))
	return rv
}

// Newunused2 creates a new unused2 instance.
func Newunused2() unused2 {
	return getunused2Class().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unused2
type unused2 struct {
	objectivec.Object
}

// unused2From constructs a [unused2] from an unsafe.Pointer.
func unused2From(ptr unsafe.Pointer) unused2 {
	return unused2{objectivec.Object{objc.ID(ptr)}}
}































