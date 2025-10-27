//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for NEFilterControlProvider


// Handle a request for new filtering rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/handleNewFlow(_:completionHandler:)
func (n_ NEFilterControlProvider) HandleNewFlowCompletionHandler(flow INEFilterFlow, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("handleNewFlow:completionHandler:"), flow, completionHandler)
}

// Handle a request for remediation from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/handleRemediation(for:completionHandler:)
func (n_ NEFilterControlProvider) HandleRemediationForFlowCompletionHandler(flow INEFilterFlow, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("handleRemediationForFlow:completionHandler:"), flow, completionHandler)
}

// Notify the Filter Data Provider that the filtering rules have changed on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/notifyRulesChanged()
func (n_ NEFilterControlProvider) NotifyRulesChanged() {
	objc.Send[objc.ID](n_.ID, objc.Sel("notifyRulesChanged"))
}

// iOS-only properties

// A dictionary containing sets of strings used to customize the remediation portion of the block page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/remediationMap
func (n_ NEFilterControlProvider) RemediationMap() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("remediationMap"))
	return rv
}
func (n_ NEFilterControlProvider) SetRemediationMap(value foundation.IDictionary) {
	n_.ID.Send(objc.RegisterName("setRemediationMap:"), value)
}

// A dictionary containing strings to be appended to URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterControlProvider/urlAppendStringMap
func (n_ NEFilterControlProvider) URLAppendStringMap() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("URLAppendStringMap"))
	return rv
}
func (n_ NEFilterControlProvider) SetURLAppendStringMap(value foundation.IDictionary) {
	n_.ID.Send(objc.RegisterName("setURLAppendStringMap:"), value)
}





