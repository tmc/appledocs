//go:build darwin && ios

// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for INInteraction


// Returns the value of the specified parameter of this interaction object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/parameterValue(for:)
func (i_ INInteraction) ParameterValueForParameter(parameter INParameter) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("parameterValueForParameter:"), parameter)
	return rv
}

// iOS-only properties





