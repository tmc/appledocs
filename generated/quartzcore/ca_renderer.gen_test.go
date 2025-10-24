// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewRenderer

// ExampleRenderer_EndFrame demonstrates using EndFrame on a Renderer instance.
// Release any data associated with the current frame.
func ExampleRenderer_EndFrame() {
	obj := quartzcore.NewRenderer()
	obj.EndFrame()
	// Output:
	}

// ExampleRenderer_NextFrameTime demonstrates using NextFrameTime on a Renderer instance.
// Returns the time at which the next update should happen.
func ExampleRenderer_NextFrameTime() {
	obj := quartzcore.NewRenderer()
	_ = obj.NextFrameTime()
	// Output:
	}

// ExampleRenderer_Render demonstrates using Render on a Renderer instance.
// Render the update region of the current frame to the target context.
func ExampleRenderer_Render() {
	obj := quartzcore.NewRenderer()
	obj.Render()
	// Output:
	}

// ExampleRenderer_UpdateBounds demonstrates using UpdateBounds on a Renderer instance.
// Returns the bounds of the update region that contains all pixels that will be rendered by the current frame.
func ExampleRenderer_UpdateBounds() {
	obj := quartzcore.NewRenderer()
	_ = obj.UpdateBounds()
	// Output:
	}

