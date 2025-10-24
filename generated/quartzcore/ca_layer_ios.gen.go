//go:build darwin && ios

// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Layer


// Adds the specified constraint to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/addConstraint(_:)
func (l_ Layer) AddConstraint(c IConstraint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addConstraint:"), c)
}

// iOS-only properties

// The constraints used to position current layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/constraints
func (l_ Layer) Constraints() []IConstraint {
	rv := objc.Send[[]Constraint](l_.ID, objc.Sel("constraints"))
	return rv
}
func (l_ Layer) SetConstraints(value []IConstraint) {
	l_.ID.Send(objc.RegisterName("setConstraints:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsDynamicContentScaling
func (l_ Layer) WantsDynamicContentScaling() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("wantsDynamicContentScaling"))
	return rv
}
func (l_ Layer) SetWantsDynamicContentScaling(value bool) {
	l_.ID.Send(objc.RegisterName("setWantsDynamicContentScaling:"), value)
}




