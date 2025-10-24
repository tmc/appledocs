// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenAuthOperation */


/* debug [class_header]: Header for TKTokenAuthOperation */
// The class instance for the [TKTokenAuthOperation] class.
var (
	TKTokenAuthOperationClass     _TKTokenAuthOperationClass
	TKTokenAuthOperationClassOnce sync.Once
)

func getTKTokenAuthOperationClass() _TKTokenAuthOperationClass {
	TKTokenAuthOperationClassOnce.Do(func() {
		TKTokenAuthOperationClass = _TKTokenAuthOperationClass{objc.GetClass("TKTokenAuthOperation")}
	})
	return TKTokenAuthOperationClass
}

type _TKTokenAuthOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenAuthOperation */
// An interface definition for the [TKTokenAuthOperation] class.
type ITKTokenAuthOperation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenAuthOperation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenAuthOperation */
	// methods:
	FinishWithError(error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenAuthOperation */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenAuthOperationClass) Alloc() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenAuthOperationClass) New() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenAuthOperation) Init() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenAuthOperation) Autorelease() TKTokenAuthOperation {
	rv := objc.Send[TKTokenAuthOperation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenAuthOperation creates a new TKTokenAuthOperation instance.
func NewTKTokenAuthOperation() TKTokenAuthOperation {
	return getTKTokenAuthOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenAuthOperation */
// An authentication operation for a cryptographic token.
//
// The CryptoTokenKit framework provides the following concrete subclasses: , for password-based authentication, and for Smart Card PIN-based authentication.


// An authentication operation for a cryptographic token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenAuthOperation
type TKTokenAuthOperation struct {
	objectivec.Object
}

// TKTokenAuthOperationFrom constructs a [TKTokenAuthOperation] from an unsafe.Pointer.
//
// An authentication operation for a cryptographic token.
func TKTokenAuthOperationFrom(ptr unsafe.Pointer) TKTokenAuthOperation {
	return TKTokenAuthOperation{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenAuthOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenAuthOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenAuthOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenAuthOperation */

// Finishes the authentication operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenAuthOperation/finish()
func (t_ TKTokenAuthOperation) FinishWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("finishWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: FinishWithError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenAuthOperation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenAuthOperation */



