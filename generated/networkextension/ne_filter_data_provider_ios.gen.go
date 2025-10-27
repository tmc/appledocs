//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for NEFilterDataProvider


// Handle a remediation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleRemediation(for:)
func (n_ NEFilterDataProvider) HandleRemediationForFlow(flow INEFilterFlow) INEFilterRemediationVerdict {
	rv := objc.Send[NEFilterRemediationVerdict](n_.ID, objc.Sel("handleRemediationForFlow:"), flow)
	return rv
}

// Handle a rules changed event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterDataProvider/handleRulesChanged()
func (n_ NEFilterDataProvider) HandleRulesChanged() {
	objc.Send[objc.ID](n_.ID, objc.Sel("handleRulesChanged"))
}

// iOS-only properties





