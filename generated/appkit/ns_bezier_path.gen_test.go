// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewBezierPath

// ExampleNewBezierPathWithCGPath demonstrates how to create a BezierPath instance using NewBezierPathWithCGPath.
func ExampleNewBezierPathWithCGPath() {
	_ = appkit.NewBezierPathWithCGPath(
		appkit.PathRef /* not a class type */{}, // cgPath PathRef /* not a class type */
	)
	// Output:
}
// ExampleBezierPath_AddClip demonstrates using AddClip on a BezierPath instance.
// Intersects the area enclosed by the path with the clipping path of the current graphics context and makes the resulting shape the current clipping path.
func ExampleBezierPath_AddClip() {
	obj := appkit.NewBezierPath()
	obj.AddClip()
	// Output:
	}

// ExampleBezierPath_ClosePath demonstrates using ClosePath on a BezierPath instance.
// Closes the most recently added subpath.
func ExampleBezierPath_ClosePath() {
	obj := appkit.NewBezierPath()
	obj.ClosePath()
	// Output:
	}

// ExampleBezierPath_Fill demonstrates using Fill on a BezierPath instance.
// Paints the region enclosed by the path.
func ExampleBezierPath_Fill() {
	obj := appkit.NewBezierPath()
	obj.Fill()
	// Output:
	}

// ExampleBezierPath_RemoveAllPoints demonstrates using RemoveAllPoints on a BezierPath instance.
// Removes all path elements from the path, effectively clearing the path.
func ExampleBezierPath_RemoveAllPoints() {
	obj := appkit.NewBezierPath()
	obj.RemoveAllPoints()
	// Output:
	}

// ExampleBezierPath_SetClip demonstrates using SetClip on a BezierPath instance.
// Replaces the clipping path of the current graphics context with the area inside the path.
func ExampleBezierPath_SetClip() {
	obj := appkit.NewBezierPath()
	obj.SetClip()
	// Output:
	}

// ExampleBezierPath_Stroke demonstrates using Stroke on a BezierPath instance.
// Draws a line along the path using the current stroke color and drawing attributes.
func ExampleBezierPath_Stroke() {
	obj := appkit.NewBezierPath()
	obj.Stroke()
	// Output:
	}

