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

// Sets the arbitrary keyed-data for the specified key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setValue(_:forKey:)
func (tc _TransactionClass) SetValueForKey(anObject objc.ID, key string) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("setValue:forKey:"), anObject, key)
}


