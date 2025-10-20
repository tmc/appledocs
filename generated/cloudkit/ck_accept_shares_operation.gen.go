// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKAcceptSharesOperation] class.
var (
	CKAcceptSharesOperationClass     _CKAcceptSharesOperationClass
	CKAcceptSharesOperationClassOnce sync.Once
)

func getCKAcceptSharesOperationClass() _CKAcceptSharesOperationClass {
	CKAcceptSharesOperationClassOnce.Do(func() {
		CKAcceptSharesOperationClass = _CKAcceptSharesOperationClass{objc.GetClass("CKAcceptSharesOperation")}
	})
	return CKAcceptSharesOperationClass
}

type _CKAcceptSharesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKAcceptSharesOperation] class.
type ICKAcceptSharesOperation interface {
	ICKOperation
}

// An operation that confirms a user’s participation in a share.
//
// Use this operation to accept participation in one or more shares. You create the operation with an array of share metadatas, which CloudKit provides to your app when the user taps or clicks a share’s . The method CloudKit calls varies by platform and app configuration. For more information, see . You can also fetch a share’s metadata using . If there are several metadatas, group them by their and create an operation for each container. Then add the operation to each container’s operation queue to run it. The operation executes its callbacks on a private serial queue. The operation calls once for each metadata you provide. CloudKit returns the metadata and its related share, or an error if it can’t accept the share. CloudKit also batches per-metadata errors. If the operation completes with errors, it returns a error. The error stores individual errors in its dictionary. Use the key to extract them. After CloudKit applies all record changes, the operation calls . When the closure executes, the server may continue processing residual tasks of the operation, such as creating the record zone in the user’s private database. The following example demonstrates how to accept a share that CloudKit provides to your window scene delegate. It shows how to create the operation, configure it, and execute it in the correct container:
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation
type CKAcceptSharesOperation struct {
	CKOperation
}

// CKAcceptSharesOperationFrom constructs a [CKAcceptSharesOperation] from an unsafe.Pointer.
//
// An operation that confirms a user’s participation in a share.
func CKAcceptSharesOperationFrom(ptr unsafe.Pointer) CKAcceptSharesOperation {
	return CKAcceptSharesOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKAcceptSharesOperationClass) Alloc() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKAcceptSharesOperationClass) New() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKAcceptSharesOperation) Init() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKAcceptSharesOperation) Autorelease() CKAcceptSharesOperation {
	rv := objc.Send[CKAcceptSharesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAcceptSharesOperation creates a new CKAcceptSharesOperation instance.
func NewCKAcceptSharesOperation() CKAcceptSharesOperation {
	return getCKAcceptSharesOperationClass().New()
}




