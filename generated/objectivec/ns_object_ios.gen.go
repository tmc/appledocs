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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElement(at:)
func (o_ Object) AccessibilityElementAtIndex(index int) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementAtIndex:"), index)
	return Object{ID: rv}
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementCount()
func (o_ Object) AccessibilityElementCount() int {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityElementCount"))
	return rv
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/index(ofAccessibilityElement:)
func (o_ Object) IndexOfAccessibilityElement(element IObject) int {
	rv := objc.Send[int](o_.ID, objc.Sel("indexOfAccessibilityElement:"), element)
	return rv
}

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivateBlock
func (o_ Object) AccessibilityActivateBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityActivateBlock"))
	return rv
}
func (o_ Object) SetAccessibilityActivateBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityActivateBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivationPointBlock
func (o_ Object) AccessibilityActivationPointBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityActivationPointBlock"))
	return rv
}
func (o_ Object) SetAccessibilityActivationPointBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityActivationPointBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedHintBlock
func (o_ Object) AccessibilityAttributedHintBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityAttributedHintBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedHintBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedHintBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedLabelBlock
func (o_ Object) AccessibilityAttributedLabelBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityAttributedLabelBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedLabelBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedLabelBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedUserInputLabels
func (o_ Object) AccessibilityAttributedUserInputLabels() []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedUserInputLabels(value []objc.ID) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedUserInputLabels:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedUserInputLabelsBlock
func (o_ Object) AccessibilityAttributedUserInputLabelsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityAttributedUserInputLabelsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedUserInputLabelsBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedUserInputLabelsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAttributedValueBlock
func (o_ Object) AccessibilityAttributedValueBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityAttributedValueBlock"))
	return rv
}
func (o_ Object) SetAccessibilityAttributedValueBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityAttributedValueBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityContainerType
func (o_ Object) AccessibilityContainerType() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityContainerType"))
	return rv
}
func (o_ Object) SetAccessibilityContainerType(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityContainerType:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityContainerTypeBlock
func (o_ Object) AccessibilityContainerTypeBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityContainerTypeBlock"))
	return rv
}
func (o_ Object) SetAccessibilityContainerTypeBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityContainerTypeBlock:"), value)
}

// An array of custom actions to display along with the built-in actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomActions
func (o_ Object) AccessibilityCustomActions() []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("accessibilityCustomActions"))
	return rv
}
func (o_ Object) SetAccessibilityCustomActions(value []objc.ID) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomActions:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomActionsBlock
func (o_ Object) AccessibilityCustomActionsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityCustomActionsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityCustomActionsBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomActionsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomRotors
func (o_ Object) AccessibilityCustomRotors() []objc.ID {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("accessibilityCustomRotors"))
	return rv
}
func (o_ Object) SetAccessibilityCustomRotors(value []objc.ID) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomRotors:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityCustomRotorsBlock
func (o_ Object) AccessibilityCustomRotorsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityCustomRotorsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityCustomRotorsBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityCustomRotorsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDecrementBlock
func (o_ Object) AccessibilityDecrementBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityDecrementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityDecrementBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDecrementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDirectTouchOptions
func (o_ Object) AccessibilityDirectTouchOptions() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityDirectTouchOptions"))
	return rv
}
func (o_ Object) SetAccessibilityDirectTouchOptions(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDirectTouchOptions:"), value)
}

// An array of location descriptor objects that you use to define what drags are possible from this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDragSourceDescriptors
func (o_ Object) AccessibilityDragSourceDescriptors() []objc.ID {
	rv := objc.Send[[]IObject](o_.ID, objc.Sel("accessibilityDragSourceDescriptors"))
	return rv
}
func (o_ Object) SetAccessibilityDragSourceDescriptors(value []objc.ID) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDragSourceDescriptors:"), value)
}

// An array of location descriptor objects that you use to define where drops are possible on this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDropPointDescriptors
func (o_ Object) AccessibilityDropPointDescriptors() []objc.ID {
	rv := objc.Send[[]IObject](o_.ID, objc.Sel("accessibilityDropPointDescriptors"))
	return rv
}
func (o_ Object) SetAccessibilityDropPointDescriptors(value []objc.ID) {
	o_.ID.Send(objc.RegisterName("setAccessibilityDropPointDescriptors:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementsBlock
func (o_ Object) AccessibilityElementsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityElementsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityElementsBlock(value IObject) {
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
func (o_ Object) AccessibilityElementsHiddenBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityElementsHiddenBlock"))
	return rv
}
func (o_ Object) SetAccessibilityElementsHiddenBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityElementsHiddenBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityExpandedStatus
func (o_ Object) AccessibilityExpandedStatus() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityExpandedStatus"))
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
func (o_ Object) AccessibilityFrameBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityFrameBlock"))
	return rv
}
func (o_ Object) SetAccessibilityFrameBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityFrameBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHeaderElementsBlock
func (o_ Object) AccessibilityHeaderElementsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityHeaderElementsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityHeaderElementsBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityHeaderElementsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHintBlock
func (o_ Object) AccessibilityHintBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityHintBlock"))
	return rv
}
func (o_ Object) SetAccessibilityHintBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityHintBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIdentifierBlock
func (o_ Object) AccessibilityIdentifierBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityIdentifierBlock"))
	return rv
}
func (o_ Object) SetAccessibilityIdentifierBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityIdentifierBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIncrementBlock
func (o_ Object) AccessibilityIncrementBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityIncrementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityIncrementBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityIncrementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLabelBlock
func (o_ Object) AccessibilityLabelBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityLabelBlock"))
	return rv
}
func (o_ Object) SetAccessibilityLabelBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityLabelBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLanguageBlock
func (o_ Object) AccessibilityLanguageBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityLanguageBlock"))
	return rv
}
func (o_ Object) SetAccessibilityLanguageBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityLanguageBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityMagicTapBlock
func (o_ Object) AccessibilityMagicTapBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityMagicTapBlock"))
	return rv
}
func (o_ Object) SetAccessibilityMagicTapBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityMagicTapBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNavigationStyle
func (o_ Object) AccessibilityNavigationStyle() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityNavigationStyle"))
	return rv
}
func (o_ Object) SetAccessibilityNavigationStyle(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityNavigationStyle:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNavigationStyleBlock
func (o_ Object) AccessibilityNavigationStyleBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityNavigationStyleBlock"))
	return rv
}
func (o_ Object) SetAccessibilityNavigationStyleBlock(value IObject) {
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
func (o_ Object) AccessibilityNextTextNavigationElementBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityNextTextNavigationElementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityNextTextNavigationElementBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityNextTextNavigationElementBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPath
func (o_ Object) AccessibilityPath() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityPath"))
	return rv
}
func (o_ Object) SetAccessibilityPath(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPath:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPathBlock
func (o_ Object) AccessibilityPathBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityPathBlock"))
	return rv
}
func (o_ Object) SetAccessibilityPathBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityPathBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPerformEscapeBlock
func (o_ Object) AccessibilityPerformEscapeBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityPerformEscapeBlock"))
	return rv
}
func (o_ Object) SetAccessibilityPerformEscapeBlock(value IObject) {
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
func (o_ Object) AccessibilityPreviousTextNavigationElementBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityPreviousTextNavigationElementBlock"))
	return rv
}
func (o_ Object) SetAccessibilityPreviousTextNavigationElementBlock(value IObject) {
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
func (o_ Object) AccessibilityRespondsToUserInteractionBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityRespondsToUserInteractionBlock"))
	return rv
}
func (o_ Object) SetAccessibilityRespondsToUserInteractionBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityRespondsToUserInteractionBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityShouldGroupAccessibilityChildrenBlock
func (o_ Object) AccessibilityShouldGroupAccessibilityChildrenBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityShouldGroupAccessibilityChildrenBlock"))
	return rv
}
func (o_ Object) SetAccessibilityShouldGroupAccessibilityChildrenBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityShouldGroupAccessibilityChildrenBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextInputResponder
func (o_ Object) AccessibilityTextInputResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTextInputResponder"))
	return rv
}
func (o_ Object) SetAccessibilityTextInputResponder(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextInputResponder:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextInputResponderBlock
func (o_ Object) AccessibilityTextInputResponderBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityTextInputResponderBlock"))
	return rv
}
func (o_ Object) SetAccessibilityTextInputResponderBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextInputResponderBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextualContext
func (o_ Object) AccessibilityTextualContext() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityTextualContext"))
	return rv
}
func (o_ Object) SetAccessibilityTextualContext(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextualContext:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTextualContextBlock
func (o_ Object) AccessibilityTextualContextBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityTextualContextBlock"))
	return rv
}
func (o_ Object) SetAccessibilityTextualContextBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTextualContextBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTraits
func (o_ Object) AccessibilityTraits() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityTraits"))
	return rv
}
func (o_ Object) SetAccessibilityTraits(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityTraits:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityTraitsBlock
func (o_ Object) AccessibilityTraitsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityTraitsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityTraitsBlock(value IObject) {
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
func (o_ Object) AccessibilityUserInputLabelsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityUserInputLabelsBlock"))
	return rv
}
func (o_ Object) SetAccessibilityUserInputLabelsBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityUserInputLabelsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityValueBlock
func (o_ Object) AccessibilityValueBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityValueBlock"))
	return rv
}
func (o_ Object) SetAccessibilityValueBlock(value IObject) {
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
func (o_ Object) AccessibilityViewIsModalBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityViewIsModalBlock"))
	return rv
}
func (o_ Object) SetAccessibilityViewIsModalBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAccessibilityViewIsModalBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/automationElementsBlock
func (o_ Object) AutomationElementsBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("automationElementsBlock"))
	return rv
}
func (o_ Object) SetAutomationElementsBlock(value IObject) {
	o_.ID.Send(objc.RegisterName("setAutomationElementsBlock:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isAccessibilityElement
func (o_ Object) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAccessibilityElement"))
	return rv
}
func (o_ Object) SetIsAccessibilityElement(value bool) {
	o_.ID.Send(objc.RegisterName("setIsAccessibilityElement:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isAccessibilityElementBlock
func (o_ Object) IsAccessibilityElementBlock() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("isAccessibilityElementBlock"))
	return rv
}
func (o_ Object) SetIsAccessibilityElementBlock(value IObject) {
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




