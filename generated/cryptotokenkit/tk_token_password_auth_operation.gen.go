// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKTokenPasswordAuthOperation */


/* debug [class_header]: Header for TKTokenPasswordAuthOperation */
// The class instance for the [TKTokenPasswordAuthOperation] class.
var (
	TKTokenPasswordAuthOperationClass     _TKTokenPasswordAuthOperationClass
	TKTokenPasswordAuthOperationClassOnce sync.Once
)

func getTKTokenPasswordAuthOperationClass() _TKTokenPasswordAuthOperationClass {
	TKTokenPasswordAuthOperationClassOnce.Do(func() {
		TKTokenPasswordAuthOperationClass = _TKTokenPasswordAuthOperationClass{objc.GetClass("TKTokenPasswordAuthOperation")}
	})
	return TKTokenPasswordAuthOperationClass
}

type _TKTokenPasswordAuthOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenPasswordAuthOperation */
// An interface definition for the [TKTokenPasswordAuthOperation] class.
type ITKTokenPasswordAuthOperation interface {
	ITKTokenAuthOperation
	
/* debug [class_interface_properties]: Properties for TKTokenPasswordAuthOperation */
	// properties:
	Password() objc.IObject /* cross-framework: NSString */
	SetPassword(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenPasswordAuthOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenPasswordAuthOperation */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenPasswordAuthOperationClass) Alloc() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenPasswordAuthOperationClass) New() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenPasswordAuthOperation) Init() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenPasswordAuthOperation) Autorelease() TKTokenPasswordAuthOperation {
	rv := objc.Send[TKTokenPasswordAuthOperation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenPasswordAuthOperation creates a new TKTokenPasswordAuthOperation instance.
func NewTKTokenPasswordAuthOperation() TKTokenPasswordAuthOperation {
	return getTKTokenPasswordAuthOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenPasswordAuthOperation */
// A password-based authentication operation.


// A password-based authentication operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenPasswordAuthOperation
type TKTokenPasswordAuthOperation struct {
	TKTokenAuthOperation
}

// TKTokenPasswordAuthOperationFrom constructs a [TKTokenPasswordAuthOperation] from an unsafe.Pointer.
//
// A password-based authentication operation.
func TKTokenPasswordAuthOperationFrom(ptr unsafe.Pointer) TKTokenPasswordAuthOperation {
	return TKTokenPasswordAuthOperation{
		TKTokenAuthOperation: TKTokenAuthOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenPasswordAuthOperation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenPasswordAuthOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenPasswordAuthOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenPasswordAuthOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenPasswordAuthOperation */

// The password to be filled in when the is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenPasswordAuthOperation/password
func (t_ TKTokenPasswordAuthOperation) Password() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("password"))
	return rv
}/* debug [instance_properties/getter]: password */


// The password to be filled in when the is called.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenPasswordAuthOperation/password
func (t_ TKTokenPasswordAuthOperation) SetPassword(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPassword:"), value)
}/* debug [instance_properties/setter]: password */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenPasswordAuthOperation */



