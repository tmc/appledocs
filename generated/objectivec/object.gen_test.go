// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec_test

import (
	"github.com/tmc/appledocs/generated/objectivec"
)

// Suppress unused import errors
var _ = objectivec.NewObject

// ExampleNewObject demonstrates how to create a Object instance.
// Implemented by subclasses to initialize a new object (the receiver) immediately after memory for it has been allocated.
func ExampleNewObject() {
	_ = objectivec.NewObject()
	// Output:
}
// ExampleObject_AccessibilityActivate demonstrates using AccessibilityActivate on a Object instance.
// Tells the element to activate itself and report the success or failure of the operation.
func ExampleObject_AccessibilityActivate() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityActivate()
	// Output:
}

// ExampleObject_AccessibilityAssistiveTechnologyFocusedIdentifiers demonstrates using AccessibilityAssistiveTechnologyFocusedIdentifiers on a Object instance.
// Returns a set of identifier keys indicating which assistive app has focus on the accessibility element.
func ExampleObject_AccessibilityAssistiveTechnologyFocusedIdentifiers() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityAssistiveTechnologyFocusedIdentifiers()
	// Output:
}

// ExampleObject_AccessibilityDecrement demonstrates using AccessibilityDecrement on a Object instance.
// Tells the accessibility element to decrement the value of its content.
func ExampleObject_AccessibilityDecrement() {
	obj := objectivec.NewObject()
	obj.AccessibilityDecrement()
	// Output:
}

// ExampleObject_AccessibilityElementCount demonstrates using AccessibilityElementCount on a Object instance.
func ExampleObject_AccessibilityElementCount() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityElementCount()
	// Output:
}

// ExampleObject_AccessibilityElementDidBecomeFocused demonstrates using AccessibilityElementDidBecomeFocused on a Object instance.
// Sent after an assistive technology has set its virtual focus on the accessibility element.
func ExampleObject_AccessibilityElementDidBecomeFocused() {
	obj := objectivec.NewObject()
	obj.AccessibilityElementDidBecomeFocused()
	// Output:
}

// ExampleObject_AccessibilityElementDidLoseFocus demonstrates using AccessibilityElementDidLoseFocus on a Object instance.
// Sent after an assistive technology has removed its virtual focus from an accessibility element.
func ExampleObject_AccessibilityElementDidLoseFocus() {
	obj := objectivec.NewObject()
	obj.AccessibilityElementDidLoseFocus()
	// Output:
}

// ExampleObject_AccessibilityElementIsFocused demonstrates using AccessibilityElementIsFocused on a Object instance.
// Returns a Boolean value indicating whether an assistive technology is focused on the accessibility element.
func ExampleObject_AccessibilityElementIsFocused() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityElementIsFocused()
	// Output:
}

// ExampleObject_AccessibilityIncrement demonstrates using AccessibilityIncrement on a Object instance.
// Tells the accessibility element to increment the value of its content.
func ExampleObject_AccessibilityIncrement() {
	obj := objectivec.NewObject()
	obj.AccessibilityIncrement()
	// Output:
}

// ExampleObject_AccessibilityLineEndPositionFromCurrentSelection demonstrates using AccessibilityLineEndPositionFromCurrentSelection on a Object instance.
func ExampleObject_AccessibilityLineEndPositionFromCurrentSelection() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityLineEndPositionFromCurrentSelection()
	// Output:
}

// ExampleObject_AccessibilityLineStartPositionFromCurrentSelection demonstrates using AccessibilityLineStartPositionFromCurrentSelection on a Object instance.
func ExampleObject_AccessibilityLineStartPositionFromCurrentSelection() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityLineStartPositionFromCurrentSelection()
	// Output:
}

// ExampleObject_AccessibilityPerformEscape demonstrates using AccessibilityPerformEscape on a Object instance.
// Dismisses a modal view and returns the success or failure of the action.
func ExampleObject_AccessibilityPerformEscape() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityPerformEscape()
	// Output:
}

// ExampleObject_AccessibilityPerformMagicTap demonstrates using AccessibilityPerformMagicTap on a Object instance.
// Performs a salient action.
func ExampleObject_AccessibilityPerformMagicTap() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityPerformMagicTap()
	// Output:
}

// ExampleObject_ActionProperty demonstrates using ActionProperty on a Object instance.
// Sent to the delegate to request the property the action applies to.
func ExampleObject_ActionProperty() {
	obj := objectivec.NewObject()
	_ = obj.ActionProperty()
	// Output:
}

// ExampleObject_BrowserAccessibilitySelectedTextRange demonstrates using BrowserAccessibilitySelectedTextRange on a Object instance.
// Returns the range of selected text in the element.
func ExampleObject_BrowserAccessibilitySelectedTextRange() {
	obj := objectivec.NewObject()
	_ = obj.BrowserAccessibilitySelectedTextRange()
	// Output:
}

// ExampleObject_Dealloc demonstrates using Dealloc on a Object instance.
// Deallocates the memory occupied by the receiver.
func ExampleObject_Dealloc() {
	obj := objectivec.NewObject()
	obj.Dealloc()
	// Output:
}

// ExampleObject_Finalize demonstrates using Finalize on a Object instance.
// The garbage collector invokes this method on the receiver before disposing of the memory it uses.
func ExampleObject_Finalize() {
	obj := objectivec.NewObject()
	obj.Finalize()
	// Output:
}

// ExampleObject_FinalizeForWebScript demonstrates using FinalizeForWebScript on a Object instance.
// Performs cleanup when the scripting environment is reset.
func ExampleObject_FinalizeForWebScript() {
	obj := objectivec.NewObject()
	obj.FinalizeForWebScript()
	// Output:
}

// ExampleObject_ImageRepresentation demonstrates using ImageRepresentation on a Object instance.
// Returns the image to display.
func ExampleObject_ImageRepresentation() {
	obj := objectivec.NewObject()
	_ = obj.ImageRepresentation()
	// Output:
}

// ExampleObject_ImageRepresentationType demonstrates using ImageRepresentationType on a Object instance.
// Returns the representation type of the image to display.
func ExampleObject_ImageRepresentationType() {
	obj := objectivec.NewObject()
	_ = obj.ImageRepresentationType()
	// Output:
}

// ExampleObject_ImageSubtitle demonstrates using ImageSubtitle on a Object instance.
// Returns the display subtitle of the image.
func ExampleObject_ImageSubtitle() {
	obj := objectivec.NewObject()
	_ = obj.ImageSubtitle()
	// Output:
}

// ExampleObject_ImageTitle demonstrates using ImageTitle on a Object instance.
// Returns the display title of the image.
func ExampleObject_ImageTitle() {
	obj := objectivec.NewObject()
	_ = obj.ImageTitle()
	// Output:
}

// ExampleObject_ImageUID demonstrates using ImageUID on a Object instance.
// Returns a unique string that identifies the data source item.
func ExampleObject_ImageUID() {
	obj := objectivec.NewObject()
	_ = obj.ImageUID()
	// Output:
}

// ExampleObject_ImageVersion demonstrates using ImageVersion on a Object instance.
// Returns the version of the item.
func ExampleObject_ImageVersion() {
	obj := objectivec.NewObject()
	_ = obj.ImageVersion()
	// Output:
}

// ExampleObject_WebPlugInDestroy demonstrates using WebPlugInDestroy on a Object instance.
// Prepares the plug-in for deallocation.
func ExampleObject_WebPlugInDestroy() {
	obj := objectivec.NewObject()
	obj.WebPlugInDestroy()
	// Output:
}

// ExampleObject_WebPlugInInitialize demonstrates using WebPlugInInitialize on a Object instance.
// Initializes the plug-in.
func ExampleObject_WebPlugInInitialize() {
	obj := objectivec.NewObject()
	obj.WebPlugInInitialize()
	// Output:
}

// ExampleObject_WebPlugInMainResourceDidFinishLoading demonstrates using WebPlugInMainResourceDidFinishLoading on a Object instance.
// Invoked when the connection successfully finishes loading data.
func ExampleObject_WebPlugInMainResourceDidFinishLoading() {
	obj := objectivec.NewObject()
	obj.WebPlugInMainResourceDidFinishLoading()
	// Output:
}

// ExampleObject_WebPlugInStart demonstrates using WebPlugInStart on a Object instance.
// Tells the plug-in to start normal operation.
func ExampleObject_WebPlugInStart() {
	obj := objectivec.NewObject()
	obj.WebPlugInStart()
	// Output:
}

// ExampleObject_WebPlugInStop demonstrates using WebPlugInStop on a Object instance.
// Tells the plug-in to stop normal operation.
func ExampleObject_WebPlugInStop() {
	obj := objectivec.NewObject()
	obj.WebPlugInStop()
	// Output:
}

