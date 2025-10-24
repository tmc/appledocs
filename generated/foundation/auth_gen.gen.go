// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [authGen] class.
var (
	AuthGenClass     _authGenClass
	AuthGenClassOnce sync.Once
)

func getauthGenClass() _authGenClass {
	AuthGenClassOnce.Do(func() {
		AuthGenClass = _authGenClass{objc.GetClass("authGen")}
	})
	return AuthGenClass
}

type _authGenClass struct {
	class objc.Class
}

// An interface definition for the [authGen] class.
type IauthGen interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/authGen
type authGen struct {
	objectivec.Object
}

// authGenFrom constructs a [authGen] from an unsafe.Pointer.
func authGenFrom(ptr unsafe.Pointer) authGen {
	return authGen{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _authGenClass) Alloc() authGen {
	rv := objc.Send[authGen](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _authGenClass) New() authGen {
	rv := objc.Send[authGen](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ authGen) Init() authGen {
	rv := objc.Send[authGen](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ authGen) Autorelease() authGen {
	rv := objc.Send[authGen](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewauthGen creates a new authGen instance.
func NewauthGen() authGen {
	return getauthGenClass().New()
}




