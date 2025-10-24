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
// ExampleObject_AccessibilityElementCount demonstrates using AccessibilityElementCount on a Object instance.
//
// Note: This example is not executed because AccessibilityElementCount crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_AccessibilityElementCount() {
	obj := objectivec.NewObject()
	_ = obj.AccessibilityElementCount()
	}

// ExampleObject_ActionProperty demonstrates using ActionProperty on a Object instance.
// Sent to the delegate to request the property the action applies to.
//
// Note: This example is not executed because ActionProperty crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ActionProperty() {
	obj := objectivec.NewObject()
	_ = obj.ActionProperty()
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
//
// Note: This example is not executed because FinalizeForWebScript crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_FinalizeForWebScript() {
	obj := objectivec.NewObject()
	obj.FinalizeForWebScript()
	}

// ExampleObject_ImageRepresentation demonstrates using ImageRepresentation on a Object instance.
// Returns the image to display.
//
// Note: This example is not executed because ImageRepresentation crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ImageRepresentation() {
	obj := objectivec.NewObject()
	_ = obj.ImageRepresentation()
	}

// ExampleObject_ImageRepresentationType demonstrates using ImageRepresentationType on a Object instance.
// Returns the representation type of the image to display.
//
// Note: This example is not executed because ImageRepresentationType crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ImageRepresentationType() {
	obj := objectivec.NewObject()
	_ = obj.ImageRepresentationType()
	}

// ExampleObject_ImageSubtitle demonstrates using ImageSubtitle on a Object instance.
// Returns the display subtitle of the image.
//
// Note: This example is not executed because ImageSubtitle crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ImageSubtitle() {
	obj := objectivec.NewObject()
	_ = obj.ImageSubtitle()
	}

// ExampleObject_ImageTitle demonstrates using ImageTitle on a Object instance.
// Returns the display title of the image.
//
// Note: This example is not executed because ImageTitle crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ImageTitle() {
	obj := objectivec.NewObject()
	_ = obj.ImageTitle()
	}

// ExampleObject_ImageUID demonstrates using ImageUID on a Object instance.
// Returns a unique string that identifies the data source item.
//
// Note: This example is not executed because ImageUID crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ImageUID() {
	obj := objectivec.NewObject()
	_ = obj.ImageUID()
	}

// ExampleObject_ImageVersion demonstrates using ImageVersion on a Object instance.
// Returns the version of the item.
//
// Note: This example is not executed because ImageVersion crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_ImageVersion() {
	obj := objectivec.NewObject()
	_ = obj.ImageVersion()
	}

// ExampleObject_WebPlugInDestroy demonstrates using WebPlugInDestroy on a Object instance.
// Prepares the plug-in for deallocation.
//
// Note: This example is not executed because WebPlugInDestroy crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_WebPlugInDestroy() {
	obj := objectivec.NewObject()
	obj.WebPlugInDestroy()
	}

// ExampleObject_WebPlugInInitialize demonstrates using WebPlugInInitialize on a Object instance.
// Initializes the plug-in.
//
// Note: This example is not executed because WebPlugInInitialize crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_WebPlugInInitialize() {
	obj := objectivec.NewObject()
	obj.WebPlugInInitialize()
	}

// ExampleObject_WebPlugInMainResourceDidFinishLoading demonstrates using WebPlugInMainResourceDidFinishLoading on a Object instance.
// Invoked when the connection successfully finishes loading data.
//
// Note: This example is not executed because WebPlugInMainResourceDidFinishLoading crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_WebPlugInMainResourceDidFinishLoading() {
	obj := objectivec.NewObject()
	obj.WebPlugInMainResourceDidFinishLoading()
	}

// ExampleObject_WebPlugInStart demonstrates using WebPlugInStart on a Object instance.
// Tells the plug-in to start normal operation.
//
// Note: This example is not executed because WebPlugInStart crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_WebPlugInStart() {
	obj := objectivec.NewObject()
	obj.WebPlugInStart()
	}

// ExampleObject_WebPlugInStop demonstrates using WebPlugInStop on a Object instance.
// Tells the plug-in to stop normal operation.
//
// Note: This example is not executed because WebPlugInStop crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObject_WebPlugInStop() {
	obj := objectivec.NewObject()
	obj.WebPlugInStop()
	}

