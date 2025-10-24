// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewLayer

// ExampleNewLayer demonstrates how to create a Layer instance.
// Returns an initialized   object.
func ExampleNewLayer() {
	_ = quartzcore.NewLayer()
	// Output:
}
// ExampleNewLayerWithRemoteClientId demonstrates how to create a Layer instance using NewLayerWithRemoteClientId.
// Initializes a layer with a remote client ID.
func ExampleNewLayerWithRemoteClientId() {
	_ = quartzcore.NewLayerWithRemoteClientId(
		quartzcore.uint32 /* not a class type */{}, // client_id uint32 /* not a class type */
	)
	// Output:
}
// ExampleLayer_AffineTransform demonstrates using AffineTransform on a Layer instance.
// Returns an affine version of the layer’s transform.
func ExampleLayer_AffineTransform() {
	obj := quartzcore.NewLayer()
	_ = obj.AffineTransform()
	// Output:
	}

// ExampleLayer_AnimationKeys demonstrates using AnimationKeys on a Layer instance.
// Returns an array of strings that identify the animations currently attached to the layer.
func ExampleLayer_AnimationKeys() {
	obj := quartzcore.NewLayer()
	_ = obj.AnimationKeys()
	// Output:
	}

// ExampleLayer_ContentsAreFlipped demonstrates using ContentsAreFlipped on a Layer instance.
// Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
func ExampleLayer_ContentsAreFlipped() {
	obj := quartzcore.NewLayer()
	_ = obj.ContentsAreFlipped()
	// Output:
	}

// ExampleLayer_Display demonstrates using Display on a Layer instance.
// Reloads the content of this layer.
func ExampleLayer_Display() {
	obj := quartzcore.NewLayer()
	obj.Display()
	// Output:
	}

// ExampleLayer_DisplayIfNeeded demonstrates using DisplayIfNeeded on a Layer instance.
// Initiates the update process for a layer if it is currently marked as needing an update.
func ExampleLayer_DisplayIfNeeded() {
	obj := quartzcore.NewLayer()
	obj.DisplayIfNeeded()
	// Output:
	}

// ExampleLayer_LayoutIfNeeded demonstrates using LayoutIfNeeded on a Layer instance.
// Recalculate the receiver’s layout, if required.
func ExampleLayer_LayoutIfNeeded() {
	obj := quartzcore.NewLayer()
	obj.LayoutIfNeeded()
	// Output:
	}

// ExampleLayer_LayoutSublayers demonstrates using LayoutSublayers on a Layer instance.
// Tells the layer to update its layout.
func ExampleLayer_LayoutSublayers() {
	obj := quartzcore.NewLayer()
	obj.LayoutSublayers()
	// Output:
	}

// ExampleLayer_ModelLayer demonstrates using ModelLayer on a Layer instance.
// Returns the model layer object associated with the receiver, if any.
func ExampleLayer_ModelLayer() {
	obj := quartzcore.NewLayer()
	_ = obj.ModelLayer()
	// Output:
	}

// ExampleLayer_NeedsDisplay demonstrates using NeedsDisplay on a Layer instance.
// Returns a Boolean indicating whether the layer has been marked as needing an update.
func ExampleLayer_NeedsDisplay() {
	obj := quartzcore.NewLayer()
	_ = obj.NeedsDisplay()
	// Output:
	}

// ExampleLayer_NeedsLayout demonstrates using NeedsLayout on a Layer instance.
// Returns a Boolean indicating whether the layer has been marked as needing a layout update.
func ExampleLayer_NeedsLayout() {
	obj := quartzcore.NewLayer()
	_ = obj.NeedsLayout()
	// Output:
	}

// ExampleLayer_PreferredFrameSize demonstrates using PreferredFrameSize on a Layer instance.
// Returns the preferred size of the layer in the coordinate space of its superlayer.
func ExampleLayer_PreferredFrameSize() {
	obj := quartzcore.NewLayer()
	_ = obj.PreferredFrameSize()
	// Output:
	}

// ExampleLayer_PresentationLayer demonstrates using PresentationLayer on a Layer instance.
// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
func ExampleLayer_PresentationLayer() {
	obj := quartzcore.NewLayer()
	_ = obj.PresentationLayer()
	// Output:
	}

// ExampleLayer_RemoveAllAnimations demonstrates using RemoveAllAnimations on a Layer instance.
// Remove all animations attached to the layer.
func ExampleLayer_RemoveAllAnimations() {
	obj := quartzcore.NewLayer()
	obj.RemoveAllAnimations()
	// Output:
	}

// ExampleLayer_RemoveFromSuperlayer demonstrates using RemoveFromSuperlayer on a Layer instance.
// Detaches the layer from its parent layer.
func ExampleLayer_RemoveFromSuperlayer() {
	obj := quartzcore.NewLayer()
	obj.RemoveFromSuperlayer()
	// Output:
	}

// ExampleLayer_SetNeedsDisplay demonstrates using SetNeedsDisplay on a Layer instance.
// Marks the layer’s contents as needing to be updated.
func ExampleLayer_SetNeedsDisplay() {
	obj := quartzcore.NewLayer()
	obj.SetNeedsDisplay()
	// Output:
	}

// ExampleLayer_SetNeedsLayout demonstrates using SetNeedsLayout on a Layer instance.
// Invalidates the layer’s layout and marks it as needing an update.
func ExampleLayer_SetNeedsLayout() {
	obj := quartzcore.NewLayer()
	obj.SetNeedsLayout()
	// Output:
	}

