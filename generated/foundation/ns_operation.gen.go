// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Operation] class.
var (
	OperationClass     _OperationClass
	OperationClassOnce sync.Once
)

func getOperationClass() _OperationClass {
	OperationClassOnce.Do(func() {
		OperationClass = _OperationClass{objc.GetClass("NSOperation")}
	})
	return OperationClass
}

type _OperationClass struct {
	class objc.Class
}





// An interface definition for the [Operation] class.
type IOperation interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _OperationClass) Alloc() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OperationClass) New() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Operation) Init() Operation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Operation) Autorelease() Operation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOperation creates a new Operation instance.
func NewOperation() Operation {
	return getOperationClass().New()
}





// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type Operation struct {
	objectivec.Object
}

// OperationFrom constructs a [Operation] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{objectivec.Object{objc.ID(ptr)}}
}































