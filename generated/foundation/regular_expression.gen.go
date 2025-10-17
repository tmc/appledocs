// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RegularExpression] class.
var regularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}

type _RegularExpressionClass struct {
	class objc.Class
}

// An interface definition for the [RegularExpression] class.
type IRegularExpression interface {
	objectivec.IObject
}

// An immutable representation of a compiled regular expression that you apply to Unicode strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRegularExpression

type RegularExpression struct {
	objectivec.Object
}

// RegularExpressionFrom constructs a [RegularExpression] from an unsafe.Pointer.
//
// An immutable representation of a compiled regular expression that you apply to Unicode strings.
func RegularExpressionFrom(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{objectivec.Object{objc.ID(ptr)}}
}



