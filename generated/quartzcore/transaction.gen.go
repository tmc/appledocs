// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Transaction] class.
var transactionClass = _TransactionClass{objc.GetClass("CATransaction")}

type _TransactionClass struct {
	class objc.Class
}

// An interface definition for the [Transaction] class.
type ITransaction interface {
	objectivec.IObject
}

// A mechanism for grouping multiple layer-tree operations into atomic updates to the render tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction

type Transaction struct {
	objectivec.Object
}

// TransactionFrom constructs a [Transaction] from an unsafe.Pointer.
//
// A mechanism for grouping multiple layer-tree operations into atomic updates to the render tree.
func TransactionFrom(ptr unsafe.Pointer) Transaction {
	return Transaction{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _TransactionClass) Alloc() Transaction {
	rv := objc.Send[Transaction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _TransactionClass) New() Transaction {
	rv := objc.Send[Transaction](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Transaction) Init() Transaction {
	rv := objc.Send[Transaction](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Transaction) Autorelease() Transaction {
	rv := objc.Send[Transaction](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTransaction creates a new Transaction instance.
func NewTransaction() Transaction {
	return transactionClass.New()
}


// Sets the arbitrary keyed-data for the specified key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setValue(_:forKey:)
func (tc _TransactionClass) SetValueForKey(anObject objc.ID, key string) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setValue:forKey:"), anObject, key)
}


