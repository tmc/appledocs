// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ExpressionDescription] class.
var expressionDescriptionClass = _ExpressionDescriptionClass{objc.GetClass("NSExpressionDescription")}

type _ExpressionDescriptionClass struct {
	class objc.Class
}

// An object that describes an expression to include with a fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSExpressionDescription

type ExpressionDescription struct {
	PropertyDescription
}

// ExpressionDescriptionFrom constructs a [ExpressionDescription] from an unsafe.Pointer.
//
// An object that describes an expression to include with a fetch request.
func ExpressionDescriptionFrom(ptr unsafe.Pointer) ExpressionDescription {
	return ExpressionDescription{
		PropertyDescription: PropertyDescriptionFrom(ptr),
	}
}



