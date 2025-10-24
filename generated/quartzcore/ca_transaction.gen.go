// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CATransaction */


/* debug [class_header]: Header for CATransaction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Transaction */
// An interface definition for the [Transaction] class.
type ITransaction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Transaction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Transaction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Transaction */
// Alloc allocates a new instance without initialization.
func (tc _TransactionClass) Alloc() Transaction {
	rv := objc.Send[Transaction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Transaction */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Transaction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Transaction */

// Returns the animation duration used by all animations within this transaction group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/animationDuration()
func (tc _TransactionClass) AnimationDuration() float64 {
	rv := objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("animationDuration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnimationDuration) */


// Returns the timing function used for all animations within this transaction group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/animationTimingFunction()
func (tc _TransactionClass) AnimationTimingFunction() IMediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(tc.class), objc.Sel("animationTimingFunction"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnimationTimingFunction) */


// Begin a new transaction for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/begin()
func (tc _TransactionClass) Begin() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("begin"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Begin) */


// Commit all changes made during the current transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/commit()
func (tc _TransactionClass) Commit() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("commit"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Commit) */


// Returns the completion block object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/completionBlock()
func (tc _TransactionClass) CompletionBlock() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("completionBlock"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompletionBlock) */


// Returns whether actions triggered as a result of property changes made within this transaction group are suppressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/disableActions()
func (tc _TransactionClass) DisableActions() bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("disableActions"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DisableActions) */


// Flushes any extant implicit transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/flush()
func (tc _TransactionClass) Flush() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("flush"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Flush) */


// Attempts to acquire a recursive spin-lock lock, ensuring that returned layer values are valid until unlocked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/lock()
func (tc _TransactionClass) Lock() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("lock"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Lock) */


// Sets the animation duration used by all animations within this transaction group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setAnimationDuration(_:)
func (tc _TransactionClass) SetAnimationDuration(dur float64) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setAnimationDuration:"), dur)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetAnimationDuration) */


// Sets the timing function used for all animations within this transaction group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setAnimationTimingFunction(_:)
func (tc _TransactionClass) SetAnimationTimingFunction(function IMediaTimingFunction) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setAnimationTimingFunction:"), function)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetAnimationTimingFunction) */


// Sets the completion block object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setCompletionBlock(_:)
func (tc _TransactionClass) SetCompletionBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setCompletionBlock:"), block)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetCompletionBlock) */


// Sets whether actions triggered as a result of property changes made within this transaction group are suppressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setDisableActions(_:)
func (tc _TransactionClass) SetDisableActions(flag bool) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setDisableActions:"), flag)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetDisableActions) */


// Sets the arbitrary keyed-data for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setValue(_:forKey:)
func (tc _TransactionClass) SetValueForKey(anObject objc.IObject, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setValue:forKey:"), anObject, key)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetValueForKey) */


// Relinquishes a previously acquired transaction lock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/unlock()
func (tc _TransactionClass) Unlock() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("unlock"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Unlock) */


// Returns the arbitrary keyed-data specified by the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/value(forKey:)
func (tc _TransactionClass) ValueForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("valueForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ValueForKey) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Transaction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Transaction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Transaction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CATransaction */



