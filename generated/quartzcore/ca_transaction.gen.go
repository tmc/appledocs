// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Transaction] class.
var (
	TransactionClass     _TransactionClass
	TransactionClassOnce sync.Once
)

func getTransactionClass() _TransactionClass {
	TransactionClassOnce.Do(func() {
		TransactionClass = _TransactionClass{objc.GetClass("CATransaction")}
	})
	return TransactionClass
}

type _TransactionClass struct {
	class objc.Class
}

// An interface definition for the [Transaction] class.
type ITransaction interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A mechanism for grouping multiple layer-tree operations into atomic updates to the render tree.
//
// is the Core Animation mechanism for batching multiple layer-tree operations into atomic updates to the render tree. Every modification to a layer tree must be part of a transaction. Nested transactions are supported. Core Animation supports two types of transactions: transactions and transactions. Implicit transactions are created automatically when the layer tree is modified by a thread without an active transaction and are committed automatically when the thread’s runloop next iterates. Explicit transactions occur when the the application sends the class a message before modifying the layer tree, and a message afterwards. allows you to override default animation properties that are set for animatable properties. You can customize duration, timing function, whether changes to properties trigger animations, and provide a handler that informs you when all animations from the transaction group are completed. During a transaction you can temporarily acquire a recursive spin lock for managing property atomicity. supports nested transactions. The following code shows how you can fade out a layer (named ) over a 2 second duration while scaling it to three times its original size. The scale animation is within a nested transaction with its own duration of 1 second. After the outer transaction completes, a completion block removes from its parent layer.


// A mechanism for grouping multiple layer-tree operations into atomic updates to the render tree.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getTransactionClass().New()
}



// Sets the arbitrary keyed-data for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setValue(_:forKey:)
func (tc _TransactionClass) SetValueForKey(anObject objectivec.IObject, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setValue:forKey:"), anObject, key)
}



