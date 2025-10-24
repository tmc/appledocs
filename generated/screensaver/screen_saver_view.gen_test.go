// Code generated from Apple documentation for ScreenSaver. DO NOT EDIT.

package screensaver_test

import (
	"github.com/tmc/appledocs/generated/screensaver"
)

// Suppress unused import errors
var _ = screensaver.NewScreenSaverView

// ExampleNewScreenSaverViewWithFrameIsPreview demonstrates how to create a ScreenSaverView instance using NewScreenSaverViewWithFrameIsPreview.
// Creates a newly allocated screen saver view with the specified frame rectangle and preview information.
func ExampleNewScreenSaverViewWithFrameIsPreview() {
	_ = screensaver.NewScreenSaverViewWithFrameIsPreview(
		screensaver.Rect /* not a class type */{}, // frame Rect /* not a class type */
		false, // isPreview bool
	)
	// Output:
}
// ExampleScreenSaverView_AnimateOneFrame demonstrates using AnimateOneFrame on a ScreenSaverView instance.
// Advances the screen saver’s animation by a single frame.
func ExampleScreenSaverView_AnimateOneFrame() {
	obj := screensaver.NewScreenSaverView()
	obj.AnimateOneFrame()
	// Output:
	}

// ExampleScreenSaverView_StartAnimation demonstrates using StartAnimation on a ScreenSaverView instance.
// Activates the periodic timer that animates the screen saver.
func ExampleScreenSaverView_StartAnimation() {
	obj := screensaver.NewScreenSaverView()
	obj.StartAnimation()
	// Output:
	}

// ExampleScreenSaverView_StopAnimation demonstrates using StopAnimation on a ScreenSaverView instance.
// Deactivates the timer that advances the animation.
func ExampleScreenSaverView_StopAnimation() {
	obj := screensaver.NewScreenSaverView()
	obj.StopAnimation()
	// Output:
	}




