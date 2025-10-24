//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Overlay


// Presents an overlay in a window scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/present(in:)
func (o_ Overlay) PresentInScene(scene WindowScene /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("presentInScene:"), scene)
}

// iOS-only properties

// An overlay’s attributes; for example, its position on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/configuration-swift.property
func (o_ Overlay) Configuration() ISKOverlayConfiguration {
	rv := objc.Send[OverlayConfiguration](o_.ID, objc.Sel("configuration"))
	return rv
}

// The overlay’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/delegate
func (o_ Overlay) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("delegate"))
	return rv
}
func (o_ Overlay) SetDelegate(value unsafe.Pointer) {
	o_.ID.Send(objc.RegisterName("setDelegate:"), value)
}




