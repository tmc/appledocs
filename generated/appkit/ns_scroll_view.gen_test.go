// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewScrollView

// ExampleScrollView_FlashScrollers demonstrates using FlashScrollers on a ScrollView instance.
// Flash the overlay scroll bars.
func ExampleScrollView_FlashScrollers() {
	obj := appkit.NewScrollView()
	obj.FlashScrollers()
	// Output:
	}

// ExampleScrollView_Tile demonstrates using Tile on a ScrollView instance.
// Lays out the components of the receiver: the content view, the scrollers, and the ruler views.
func ExampleScrollView_Tile() {
	obj := appkit.NewScrollView()
	obj.Tile()
	// Output:
	}

