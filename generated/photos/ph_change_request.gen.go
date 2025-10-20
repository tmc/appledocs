// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHChangeRequest] class.
var (
	PHChangeRequestClass     _PHChangeRequestClass
	PHChangeRequestClassOnce sync.Once
)

func getPHChangeRequestClass() _PHChangeRequestClass {
	PHChangeRequestClassOnce.Do(func() {
		PHChangeRequestClass = _PHChangeRequestClass{objc.GetClass("PHChangeRequest")}
	})
	return PHChangeRequestClass
}

type _PHChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHChangeRequest] class.
type IPHChangeRequest interface {
	objectivec.IObject
}

// The abstract base class of the framework’s photo library change requests.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHChangeRequest
type PHChangeRequest struct {
	objectivec.Object
}

// PHChangeRequestFrom constructs a [PHChangeRequest] from an unsafe.Pointer.
//
// The abstract base class of the framework’s photo library change requests.
func PHChangeRequestFrom(ptr unsafe.Pointer) PHChangeRequest {
	return PHChangeRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHChangeRequestClass) Alloc() PHChangeRequest {
	rv := objc.Send[PHChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHChangeRequestClass) New() PHChangeRequest {
	rv := objc.Send[PHChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHChangeRequest) Init() PHChangeRequest {
	rv := objc.Send[PHChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHChangeRequest) Autorelease() PHChangeRequest {
	rv := objc.Send[PHChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHChangeRequest creates a new PHChangeRequest instance.
func NewPHChangeRequest() PHChangeRequest {
	return getPHChangeRequestClass().New()
}




