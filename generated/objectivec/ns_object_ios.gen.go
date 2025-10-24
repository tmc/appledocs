//go:build darwin && ios

// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for Object


// Tells the element to activate itself and report the success or failure of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivate()
func (o_ Object) AccessibilityActivate() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityActivate"))
	return rv
}

// Returns a set of identifier keys indicating which assistive app has focus on the accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAssistiveTechnologyFocusedIdentifiers()
func (o_ Object) AccessibilityAssistiveTechnologyFocusedIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAssistiveTechnologyFocusedIdentifiers"))
	return rv
}

// Tells the accessibility element to decrement the value of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDecrement()
func (o_ Object) AccessibilityDecrement() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityDecrement"))
}

// Sent after an assistive technology has set its virtual focus on the accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementDidBecomeFocused()
func (o_ Object) AccessibilityElementDidBecomeFocused() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementDidBecomeFocused"))
}

// Sent after an assistive technology has removed its virtual focus from an accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementDidLoseFocus()
func (o_ Object) AccessibilityElementDidLoseFocus() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementDidLoseFocus"))
}

// Returns a Boolean value indicating whether an assistive technology is focused on the accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementIsFocused()
func (o_ Object) AccessibilityElementIsFocused() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityElementIsFocused"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHitTest(_:event:)
func (o_ Object) AccessibilityHitTestWithEvent(point IObject, event IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityHitTest:withEvent:"), point, event)
	return Object{ID: rv}
}

// Tells the accessibility element to increment the value of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIncrement()
func (o_ Object) AccessibilityIncrement() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityIncrement"))
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineEndPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineEndPositionFromCurrentSelection() int {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityLineEndPositionFromCurrentSelection"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineRange(forPosition:)
func (o_ Object) AccessibilityLineRangeForPosition(position int) IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityLineRangeForPosition:"), position)
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineStartPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineStartPositionFromCurrentSelection() int {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityLineStartPositionFromCurrentSelection"))
	return rv
}

// Dismisses a modal view and returns the success or failure of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPerformEscape()
func (o_ Object) AccessibilityPerformEscape() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityPerformEscape"))
	return rv
}

// Performs a salient action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPerformMagicTap()
func (o_ Object) AccessibilityPerformMagicTap() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityPerformMagicTap"))
	return rv
}

// Scrolls screen content in an application-specific way and returns the success or failure of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityScroll(_:)
func (o_ Object) AccessibilityScroll(direction IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityScroll:"), direction)
	return rv
}

// Zooms in on the content at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityZoomIn(at:)
func (o_ Object) AccessibilityZoomInAtPoint(point IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityZoomInAtPoint:"), point)
	return rv
}

// Zooms out from the content at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityZoomOut(at:)
func (o_ Object) AccessibilityZoomOutAtPoint(point IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityZoomOutAtPoint:"), point)
	return rv
}

// Returns the value for this element within the given range, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityAttributedValue(in:)
func (o_ Object) BrowserAccessibilityAttributedValueInRange(range_ IObject) IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("browserAccessibilityAttributedValueInRange:"), range_)
	return rv
}

// Deletes text from the element at the current cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityDeleteTextAtCursor(numberOfCharacters:)
func (o_ Object) BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityDeleteTextAtCursor:"), numberOfCharacters)
}

// Inserts text into the element at the current cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityInsertTextAtCursor(text:)
func (o_ Object) BrowserAccessibilityInsertTextAtCursor(text IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityInsertTextAtCursor:"), text)
}

// Returns the range of selected text in the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySelectedTextRange()
func (o_ Object) BrowserAccessibilitySelectedTextRange() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("browserAccessibilitySelectedTextRange"))
	return rv
}

// Updates the element’s selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySetSelectedTextRange(_:)
func (o_ Object) BrowserAccessibilitySetSelectedTextRange(range_ IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilitySetSelectedTextRange:"), range_)
}

// Returns this element’s value in the given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityValue(in:)
func (o_ Object) BrowserAccessibilityValueInRange(range_ IObject) IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("browserAccessibilityValueInRange:"), range_)
	return rv
}

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivateBlock
func (o_ Object) AccessibilityActivateBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityActivateBlock"))
	return rv
}
func (o_ Object) SetAccessibilityActivateBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityActivateBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivationPoint
func (o_ Object) AccessibilityActivationPoint() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
}
func (o_ Object) SetAccessibilityActivationPoint(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityActivationPoint:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivationPointBlock
func (o_ Object) AccessibilityActivationPointBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityActivationPointBlock"))
	return rv
}
func (o_ Object) SetAccessibilityActivationPointBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityActivationPointBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedHint
func (o_ Object) AccessibilityAttributedHint() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityAttributedHint"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedHint(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedHint:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedHintBlock
func (o_ Object) AccessibilityAttributedHintBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedHintBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedHintBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedHintBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedLabel
func (o_ Object) AccessibilityAttributedLabel() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityAttributedLabel"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedLabel(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedLabel:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedLabelBlock
func (o_ Object) AccessibilityAttributedLabelBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedLabelBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedLabelBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedLabelBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedUserInputLabels
func (o_ Object) AccessibilityAttributedUserInputLabels() []IObject {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	result := make([]IObject, len(rv))
	for i, id := range rv {
		result[i] = Object{ID: id}
	}
	return result
}
func (o_ Object) SetAccessibilityAttributedUserInputLabels(value []IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedUserInputLabels:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedUserInputLabelsBlock
func (o_ Object) AccessibilityAttributedUserInputLabelsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedUserInputLabelsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedUserInputLabelsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedUserInputLabelsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedValue
func (o_ Object) AccessibilityAttributedValue() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityAttributedValue"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedValue(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedValue:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedValueBlock
func (o_ Object) AccessibilityAttributedValueBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedValueBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedValueBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedValueBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityContainerType
func (o_ Object) AccessibilityContainerType() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityContainerType"))
	return rv
}
func (o_ Object) SetAccessibilityContainerType(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityContainerType:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityContainerTypeBlock
func (o_ Object) AccessibilityContainerTypeBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityContainerTypeBlock"))
	return rv
}
func (o_ Object) SetAccessibilityContainerTypeBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityContainerTypeBlock:"), value)
}

// An array of custom actions to display along with the built-in actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomActions
func (o_ Object) AccessibilityCustomActions() []IObject {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("accessibilityCustomActions"))
	result := make([]IObject, len(rv))
	for i, id := range rv {
		result[i] = Object{ID: id}
	}
	return result
}
func (o_ Object) SetAccessibilityCustomActions(value []IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomActions:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomActionsBlock
func (o_ Object) AccessibilityCustomActionsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityCustomActionsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityCustomActionsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomActionsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomRotors
func (o_ Object) AccessibilityCustomRotors() []IObject {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("accessibilityCustomRotors"))
	result := make([]IObject, len(rv))
	for i, id := range rv {
		result[i] = Object{ID: id}
	}
	return result
}
func (o_ Object) SetAccessibilityCustomRotors(value []IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomRotors:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomRotorsBlock
func (o_ Object) AccessibilityCustomRotorsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityCustomRotorsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityCustomRotorsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomRotorsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDecrementBlock
func (o_ Object) AccessibilityDecrementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityDecrementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityDecrementBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDecrementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDirectTouchOptions
func (o_ Object) AccessibilityDirectTouchOptions() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityDirectTouchOptions"))
	return rv
}
func (o_ Object) SetAccessibilityDirectTouchOptions(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDirectTouchOptions:"), value)
}

// An array of location descriptor objects that you use to define what drags are possible from this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDragSourceDescriptors
func (o_ Object) AccessibilityDragSourceDescriptors() []IObject {
	rv := objc.Send[[]IObject](o_.ID, objc.Sel("accessibilityDragSourceDescriptors"))
	return rv
}
func (o_ Object) SetAccessibilityDragSourceDescriptors(value []IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDragSourceDescriptors:"), value)
}

// An array of location descriptor objects that you use to define where drops are possible on this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDropPointDescriptors
func (o_ Object) AccessibilityDropPointDescriptors() []IObject {
	rv := objc.Send[[]IObject](o_.ID, objc.Sel("accessibilityDropPointDescriptors"))
	return rv
}
func (o_ Object) SetAccessibilityDropPointDescriptors(value []IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDropPointDescriptors:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementsBlock
func (o_ Object) AccessibilityElementsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityElementsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityElementsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityElementsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementsHidden
func (o_ Object) AccessibilityElementsHidden() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityElementsHidden"))
	return rv
}
func (o_ Object) SetAccessibilityElementsHidden(value bool) {
	o_.ID.Send(objc.RegisterName("setAccessibilityElementsHidden:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementsHiddenBlock
func (o_ Object) AccessibilityElementsHiddenBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityElementsHiddenBlock"))
	return rv
}
func (o_ Object) SetAccessibilityElementsHiddenBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityElementsHiddenBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityExpandedStatus
func (o_ Object) AccessibilityExpandedStatus() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityExpandedStatus"))
	return rv
}
func (o_ Object) SetAccessibilityExpandedStatus(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityExpandedStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityExpandedStatusBlock
func (o_ Object) AccessibilityExpandedStatusBlock() func() unsafe.Pointer {
	rv := objc.Send[func() unsafe.Pointer](o_.ID, objc.Sel("accessibilityExpandedStatusBlock"))
	return rv
}
func (o_ Object) SetAccessibilityExpandedStatusBlock(value func() unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityExpandedStatusBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityFrameBlock
func (o_ Object) AccessibilityFrameBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityFrameBlock"))
	return rv
}
func (o_ Object) SetAccessibilityFrameBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityFrameBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHeaderElementsBlock
func (o_ Object) AccessibilityHeaderElementsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityHeaderElementsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityHeaderElementsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityHeaderElementsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHintBlock
func (o_ Object) AccessibilityHintBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityHintBlock"))
	return rv
}
func (o_ Object) SetAccessibilityHintBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityHintBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIdentifierBlock
func (o_ Object) AccessibilityIdentifierBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityIdentifierBlock"))
	return rv
}
func (o_ Object) SetAccessibilityIdentifierBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityIdentifierBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIncrementBlock
func (o_ Object) AccessibilityIncrementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityIncrementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityIncrementBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityIncrementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLabelBlock
func (o_ Object) AccessibilityLabelBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityLabelBlock"))
	return rv
}
func (o_ Object) SetAccessibilityLabelBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityLabelBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLanguageBlock
func (o_ Object) AccessibilityLanguageBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityLanguageBlock"))
	return rv
}
func (o_ Object) SetAccessibilityLanguageBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityLanguageBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityMagicTapBlock
func (o_ Object) AccessibilityMagicTapBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityMagicTapBlock"))
	return rv
}
func (o_ Object) SetAccessibilityMagicTapBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityMagicTapBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNavigationStyle
func (o_ Object) AccessibilityNavigationStyle() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityNavigationStyle"))
	return rv
}
func (o_ Object) SetAccessibilityNavigationStyle(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityNavigationStyle:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNavigationStyleBlock
func (o_ Object) AccessibilityNavigationStyleBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityNavigationStyleBlock"))
	return rv
}
func (o_ Object) SetAccessibilityNavigationStyleBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityNavigationStyleBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNextTextNavigationElement
func (o_ Object) AccessibilityNextTextNavigationElement() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityNextTextNavigationElement"))
	return Object{ID: rv}
}
func (o_ Object) SetAccessibilityNextTextNavigationElement(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityNextTextNavigationElement:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNextTextNavigationElementBlock
func (o_ Object) AccessibilityNextTextNavigationElementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityNextTextNavigationElementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityNextTextNavigationElementBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityNextTextNavigationElementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPath
func (o_ Object) AccessibilityPath() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityPath"))
	return rv
}
func (o_ Object) SetAccessibilityPath(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPath:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPathBlock
func (o_ Object) AccessibilityPathBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPathBlock"))
	return rv
}
func (o_ Object) SetAccessibilityPathBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPathBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPerformEscapeBlock
func (o_ Object) AccessibilityPerformEscapeBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPerformEscapeBlock"))
	return rv
}
func (o_ Object) SetAccessibilityPerformEscapeBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPerformEscapeBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPreviousTextNavigationElement
func (o_ Object) AccessibilityPreviousTextNavigationElement() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityPreviousTextNavigationElement"))
	return Object{ID: rv}
}
func (o_ Object) SetAccessibilityPreviousTextNavigationElement(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPreviousTextNavigationElement:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPreviousTextNavigationElementBlock
func (o_ Object) AccessibilityPreviousTextNavigationElementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPreviousTextNavigationElementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityPreviousTextNavigationElementBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPreviousTextNavigationElementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityRespondsToUserInteraction
func (o_ Object) AccessibilityRespondsToUserInteraction() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityRespondsToUserInteraction"))
	return rv
}
func (o_ Object) SetAccessibilityRespondsToUserInteraction(value bool) {
	o_.ID.Send(objc.RegisterName("setAccessibilityRespondsToUserInteraction:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityRespondsToUserInteractionBlock
func (o_ Object) AccessibilityRespondsToUserInteractionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityRespondsToUserInteractionBlock"))
	return rv
}
func (o_ Object) SetAccessibilityRespondsToUserInteractionBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityRespondsToUserInteractionBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityShouldGroupAccessibilityChildrenBlock
func (o_ Object) AccessibilityShouldGroupAccessibilityChildrenBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityShouldGroupAccessibilityChildrenBlock"))
	return rv
}
func (o_ Object) SetAccessibilityShouldGroupAccessibilityChildrenBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityShouldGroupAccessibilityChildrenBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextInputResponder
func (o_ Object) AccessibilityTextInputResponder() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityTextInputResponder"))
	return Object{ID: rv}
}
func (o_ Object) SetAccessibilityTextInputResponder(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextInputResponder:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextInputResponderBlock
func (o_ Object) AccessibilityTextInputResponderBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTextInputResponderBlock"))
	return rv
}
func (o_ Object) SetAccessibilityTextInputResponderBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextInputResponderBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextualContext
func (o_ Object) AccessibilityTextualContext() IObject {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityTextualContext"))
	return rv
}
func (o_ Object) SetAccessibilityTextualContext(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextualContext:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextualContextBlock
func (o_ Object) AccessibilityTextualContextBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTextualContextBlock"))
	return rv
}
func (o_ Object) SetAccessibilityTextualContextBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextualContextBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTraitsBlock
func (o_ Object) AccessibilityTraitsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTraitsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityTraitsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTraitsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityUserInputLabels
func (o_ Object) AccessibilityUserInputLabels() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("accessibilityUserInputLabels"))
	return rv
}
func (o_ Object) SetAccessibilityUserInputLabels(value []string) {
	o_.ID.Send(objc.RegisterName("setAccessibilityUserInputLabels:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityUserInputLabelsBlock
func (o_ Object) AccessibilityUserInputLabelsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityUserInputLabelsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityUserInputLabelsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityUserInputLabelsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityValueBlock
func (o_ Object) AccessibilityValueBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityValueBlock"))
	return rv
}
func (o_ Object) SetAccessibilityValueBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityValueBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityViewIsModal
func (o_ Object) AccessibilityViewIsModal() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityViewIsModal"))
	return rv
}
func (o_ Object) SetAccessibilityViewIsModal(value bool) {
	o_.ID.Send(objc.RegisterName("setAccessibilityViewIsModal:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityViewIsModalBlock
func (o_ Object) AccessibilityViewIsModalBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityViewIsModalBlock"))
	return rv
}
func (o_ Object) SetAccessibilityViewIsModalBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityViewIsModalBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/automationElementsBlock
func (o_ Object) AutomationElementsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("automationElementsBlock"))
	return rv
}
func (o_ Object) SetAutomationElementsBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAutomationElementsBlock:"), value)
}

// The kind of container that contains this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityContainerType
func (o_ Object) BrowserAccessibilityContainerType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityContainerType"))
	return rv
}
func (o_ Object) SetBrowserAccessibilityContainerType(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setBrowserAccessibilityContainerType:"), value)
}

// A Boolean value that indicates whether the element has native focus in the browser Document Object Model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityHasDOMFocus
func (o_ Object) BrowserAccessibilityHasDOMFocus() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("browserAccessibilityHasDOMFocus"))
	return rv
}
func (o_ Object) SetBrowserAccessibilityHasDOMFocus(value bool) {
	o_.ID.Send(objc.RegisterName("setBrowserAccessibilityHasDOMFocus:"), value)
}

// A Boolean value that’s the element’s value for aria-required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityIsRequired
func (o_ Object) BrowserAccessibilityIsRequired() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("browserAccessibilityIsRequired"))
	return rv
}
func (o_ Object) SetBrowserAccessibilityIsRequired(value bool) {
	o_.ID.Send(objc.RegisterName("setBrowserAccessibilityIsRequired:"), value)
}

// The element’s value for aria-pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityPressedState
func (o_ Object) BrowserAccessibilityPressedState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityPressedState"))
	return rv
}
func (o_ Object) SetBrowserAccessibilityPressedState(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setBrowserAccessibilityPressedState:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isAccessibilityElementBlock
func (o_ Object) IsAccessibilityElementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("isAccessibilityElementBlock"))
	return rv
}
func (o_ Object) SetIsAccessibilityElementBlock(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setIsAccessibilityElementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/shouldGroupAccessibilityChildren
func (o_ Object) ShouldGroupAccessibilityChildren() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldGroupAccessibilityChildren"))
	return rv
}
func (o_ Object) SetShouldGroupAccessibilityChildren(value bool) {
	o_.ID.Send(objc.RegisterName("setShouldGroupAccessibilityChildren:"), value)
}




