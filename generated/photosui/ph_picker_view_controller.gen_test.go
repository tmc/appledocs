// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui_test

import (
	"github.com/tmc/appledocs/generated/photosui"
)

// Suppress unused import errors
var _ = photosui.NewPHPickerViewController

// ExamplePHPickerViewController_ScrollToInitialPosition demonstrates using ScrollToInitialPosition on a PHPickerViewController instance.
// Resets the visible photo thumbnails by scrolling the view to the picker’s initial position.
func ExamplePHPickerViewController_ScrollToInitialPosition() {
	obj := photosui.NewPHPickerViewController()
	obj.ScrollToInitialPosition()
	// Output:
	}

// ExamplePHPickerViewController_ZoomIn demonstrates using ZoomIn on a PHPickerViewController instance.
// Changes the picker’s content scale by making the photo thumbnails larger in the view.
func ExamplePHPickerViewController_ZoomIn() {
	obj := photosui.NewPHPickerViewController()
	obj.ZoomIn()
	// Output:
	}

// ExamplePHPickerViewController_ZoomOut demonstrates using ZoomOut on a PHPickerViewController instance.
// Changes the picker’s content scale by making the photo thumbnails smaller in the view.
func ExamplePHPickerViewController_ZoomOut() {
	obj := photosui.NewPHPickerViewController()
	obj.ZoomOut()
	// Output:
	}

