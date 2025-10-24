//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXTransaction


// Adds the specified action to the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXTransaction/addAction(_:)
func (c_ CXTransaction) AddAction(action ICXAction) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addAction:"), action)
}

// iOS-only properties

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




