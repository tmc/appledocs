//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXCallController


// Requests that the actions in the specified transaction be asynchronously performed by the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/request(_:completion:)
func (c_ CXCallController) RequestTransactionCompletion(transaction ICXTransaction, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestTransaction:completion:"), transaction, completion)
}

// Requests that the transaction that contains the specified actions be asynchronously performed by the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/requestTransaction(with:completion:)-4o1m4
func (c_ CXCallController) RequestTransactionWithActionsCompletion(actions []CXAction, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestTransactionWithActions:completion:"), actions, completion)
}

// Requests that the transaction that contains the specified action be asynchronously performed by the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/requestTransaction(with:completion:)-ffme
func (c_ CXCallController) RequestTransactionWithActionCompletion(action ICXAction, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestTransactionWithAction:completion:"), action, completion)
}

// iOS-only properties

// Returns an observer for active calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/callObserver
func (c_ CXCallController) CallObserver() ICXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("callObserver"))
	return rv
}




