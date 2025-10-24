// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXTransaction */


/* debug [class_header]: Header for CXTransaction */
// The class instance for the [CXTransaction] class.
var (
	CXTransactionClass     _CXTransactionClass
	CXTransactionClassOnce sync.Once
)

func getCXTransactionClass() _CXTransactionClass {
	CXTransactionClassOnce.Do(func() {
		CXTransactionClass = _CXTransactionClass{objc.GetClass("CXTransaction")}
	})
	return CXTransactionClass
}

type _CXTransactionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXTransaction */
// An interface definition for the [CXTransaction] class.
type ICXTransaction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXTransaction */
	// properties:
	IsComplete() bool
	SetIsComplete(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXTransaction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXTransaction */
// Alloc allocates a new instance without initialization.
func (cc _CXTransactionClass) Alloc() CXTransaction {
	rv := objc.Send[CXTransaction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXTransactionClass) New() CXTransaction {
	rv := objc.Send[CXTransaction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXTransaction) Init() CXTransaction {
	rv := objc.Send[CXTransaction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXTransaction) Autorelease() CXTransaction {
	rv := objc.Send[CXTransaction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXTransaction creates a new CXTransaction instance.
func NewCXTransaction() CXTransaction {
	return getCXTransactionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXTransaction */
// An object that contains zero or more action objects for a call controller to perform.


// An object that contains zero or more action objects for a call controller to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction
type CXTransaction struct {
	objectivec.Object
}

// CXTransactionFrom constructs a [CXTransaction] from an unsafe.Pointer.
//
// An object that contains zero or more action objects for a call controller to perform.
func CXTransactionFrom(ptr unsafe.Pointer) CXTransaction {
	return CXTransaction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXTransaction */

// Initializes a new transaction with the specified action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/init(action:)
func NewCXTransactionWithAction(action ICXAction) CXTransaction {
	instance := getCXTransactionClass().Alloc()
	rv := objc.Send[CXTransaction](instance.ID, objc.Sel("initWithAction:"), action)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXTransactionWithAction */


// Initializes a new transaction with the specified actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/init(actions:)
func NewCXTransactionWithActions(actions []CXAction) CXTransaction {
	instance := getCXTransactionClass().Alloc()
	rv := objc.Send[CXTransaction](instance.ID, objc.Sel("initWithActions:"), actions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXTransactionWithActions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXTransaction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXTransaction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXTransaction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXTransaction */

// A Boolean value that indicates whether the transaction has been completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxtransaction/iscomplete
func (c_ CXTransaction) IsComplete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isComplete"))
	return rv
}/* debug [instance_properties/getter]: isComplete */


// A Boolean value that indicates whether the transaction has been completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxtransaction/iscomplete
func (c_ CXTransaction) SetIsComplete(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComplete:"), value)
}/* debug [instance_properties/setter]: isComplete */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXTransaction */


