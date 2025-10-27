// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [unused1] class.
var (
	Unused1Class     _unused1Class
	Unused1ClassOnce sync.Once
)

func getunused1Class() _unused1Class {
	Unused1ClassOnce.Do(func() {
		Unused1Class = _unused1Class{objc.GetClass("unused1")}
	})
	return Unused1Class
}

type _unused1Class struct {
	class objc.Class
}





// An interface definition for the [unused1] class.
type Iunused1 interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (uc _unused1Class) Alloc() unused1 {
	rv := objc.Send[unused1](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _unused1Class) New() unused1 {
	rv := objc.Send[unused1](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ unused1) Init() unused1 {
	rv := objc.Send[unused1](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ unused1) Autorelease() unused1 {
	rv := objc.Send[unused1](u_.ID, objc.Sel("autorelease"))
	return rv
}

// Newunused1 creates a new unused1 instance.
func Newunused1() unused1 {
	return getunused1Class().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/unused1
type unused1 struct {
	objectivec.Object
}

// unused1From constructs a [unused1] from an unsafe.Pointer.
func unused1From(ptr unsafe.Pointer) unused1 {
	return unused1{objectivec.Object{objc.ID(ptr)}}
}































