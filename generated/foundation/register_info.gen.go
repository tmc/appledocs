// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [registerInfo] class.
var (
	RegisterInfoClass     _registerInfoClass
	RegisterInfoClassOnce sync.Once
)

func getregisterInfoClass() _registerInfoClass {
	RegisterInfoClassOnce.Do(func() {
		RegisterInfoClass = _registerInfoClass{objc.GetClass("registerInfo")}
	})
	return RegisterInfoClass
}

type _registerInfoClass struct {
	class objc.Class
}





// An interface definition for the [registerInfo] class.
type IregisterInfo interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _registerInfoClass) Alloc() registerInfo {
	rv := objc.Send[registerInfo](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _registerInfoClass) New() registerInfo {
	rv := objc.Send[registerInfo](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ registerInfo) Init() registerInfo {
	rv := objc.Send[registerInfo](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ registerInfo) Autorelease() registerInfo {
	rv := objc.Send[registerInfo](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewregisterInfo creates a new registerInfo instance.
func NewregisterInfo() registerInfo {
	return getregisterInfoClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/registerInfo
type registerInfo struct {
	objectivec.Object
}

// registerInfoFrom constructs a [registerInfo] from an unsafe.Pointer.
func registerInfoFrom(ptr unsafe.Pointer) registerInfo {
	return registerInfo{objectivec.Object{objc.ID(ptr)}}
}































