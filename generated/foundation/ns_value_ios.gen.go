//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Value


// iOS-only properties

// Returns the CoreGraphics affine transform representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/cgAffineTransformValue
func (v_ Value) CGAffineTransformValue() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](v_.ID, objc.Sel("CGAffineTransformValue"))
	return rv
}

// Returns the CoreGraphics point structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/cgPointValue
func (v_ Value) CGPointValue() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v_.ID, objc.Sel("CGPointValue"))
	return rv
}

// Returns the CoreGraphics rectangle structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/cgRectValue
func (v_ Value) CGRectValue() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v_.ID, objc.Sel("CGRectValue"))
	return rv
}

// Returns the CoreGraphics size structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/cgSizeValue
func (v_ Value) CGSizeValue() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("CGSizeValue"))
	return rv
}

// Returns the CoreGraphics vector structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/cgVectorValue
func (v_ Value) CGVectorValue() corefoundation.CGVector {
	rv := objc.Send[corefoundation.CGVector](v_.ID, objc.Sel("CGVectorValue"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/directionalEdgeInsetsValue
func (v_ Value) DirectionalEdgeInsetsValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("directionalEdgeInsetsValue"))
	return rv
}

// Returns the UIKit edge insets structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/uiEdgeInsetsValue
func (v_ Value) UIEdgeInsetsValue() UIEdgeInsets /* not a class type */ {
	rv := objc.Send[EdgeInsets](v_.ID, objc.Sel("UIEdgeInsetsValue"))
	return rv
}

// Returns the UIKit offset structure representation of the value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue/uiOffsetValue
func (v_ Value) UIOffsetValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("UIOffsetValue"))
	return rv
}




