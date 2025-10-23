// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CXTransaction] class.
type ICXTransaction interface {
	objectivec.IObject
	Actions() []CXAction
	Complete() bool
	UUID() foundation.UUID
	IsComplete() bool
	SetIsComplete(value bool)
	AddAction(action ICXAction)
}

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

// Alloc allocates a new instance without initialization.
func (cc _CXTransactionClass) Alloc() CXTransaction {
	rv := objc.Send[CXTransaction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a new transaction with the specified action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/init(action:)
func NewCXTransactionWithAction(action ICXAction) CXTransaction {
	instance := getCXTransactionClass().Alloc()
	rv := objc.Send[CXTransaction](instance.ID, objc.Sel("initWithAction:"), action)
	rv.Autorelease()
	return rv
}


// Initializes a new transaction with the specified actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/init(actions:)
func NewCXTransactionWithActions(actions []CXAction) CXTransaction {
	instance := getCXTransactionClass().Alloc()
	rv := objc.Send[CXTransaction](instance.ID, objc.Sel("initWithActions:"), actions)
	rv.Autorelease()
	return rv
}



// Adds the specified action to the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/addAction(_:)
func (c_ CXTransaction) AddAction(action ICXAction) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addAction:"), action)
}


// The actions added to a transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/actions
func (c_ CXTransaction) Actions() []CXAction {
	rv := objc.Send[[]CXAction](c_.ID, objc.Sel("actions"))
	return rv
}


// A Boolean value that indicates whether the transaction has been completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/isComplete
func (c_ CXTransaction) Complete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("complete"))
	return rv
}


// The unique identifier of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/uuid
func (c_ CXTransaction) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("UUID"))
	return rv
}


// A Boolean value that indicates whether the transaction has been completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxtransaction/iscomplete
func (c_ CXTransaction) IsComplete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isComplete"))
	return rv
}


// A Boolean value that indicates whether the transaction has been completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxtransaction/iscomplete
func (c_ CXTransaction) SetIsComplete(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComplete:"), value)
}


