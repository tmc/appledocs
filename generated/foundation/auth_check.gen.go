// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [authCheck] class.
var (
	AuthCheckClass     _authCheckClass
	AuthCheckClassOnce sync.Once
)

func getauthCheckClass() _authCheckClass {
	AuthCheckClassOnce.Do(func() {
		AuthCheckClass = _authCheckClass{objc.GetClass("authCheck")}
	})
	return AuthCheckClass
}

type _authCheckClass struct {
	class objc.Class
}





// An interface definition for the [authCheck] class.
type IauthCheck interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _authCheckClass) Alloc() authCheck {
	rv := objc.Send[authCheck](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _authCheckClass) New() authCheck {
	rv := objc.Send[authCheck](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ authCheck) Init() authCheck {
	rv := objc.Send[authCheck](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ authCheck) Autorelease() authCheck {
	rv := objc.Send[authCheck](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewauthCheck creates a new authCheck instance.
func NewauthCheck() authCheck {
	return getauthCheckClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/authCheck
type authCheck struct {
	objectivec.Object
}

// authCheckFrom constructs a [authCheck] from an unsafe.Pointer.
func authCheckFrom(ptr unsafe.Pointer) authCheck {
	return authCheck{objectivec.Object{objc.ID(ptr)}}
}































