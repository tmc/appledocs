// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [BluetoothPasskeyDisplay] class.
var (
	BluetoothPasskeyDisplayClass     _BluetoothPasskeyDisplayClass
	BluetoothPasskeyDisplayClassOnce sync.Once
)

func getBluetoothPasskeyDisplayClass() _BluetoothPasskeyDisplayClass {
	BluetoothPasskeyDisplayClassOnce.Do(func() {
		BluetoothPasskeyDisplayClass = _BluetoothPasskeyDisplayClass{objc.GetClass("IOBluetoothPasskeyDisplay")}
	})
	return BluetoothPasskeyDisplayClass
}

type _BluetoothPasskeyDisplayClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothPasskeyDisplay] class.
type IBluetoothPasskeyDisplay interface {
	appkit.IView
	RetreatPasskeyIndicator()
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay
type BluetoothPasskeyDisplay struct {
	appkit.View
}

// BluetoothPasskeyDisplayFrom constructs a [BluetoothPasskeyDisplay] from an unsafe.Pointer.
func BluetoothPasskeyDisplayFrom(ptr unsafe.Pointer) BluetoothPasskeyDisplay {
	return BluetoothPasskeyDisplay{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothPasskeyDisplayClass) Alloc() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothPasskeyDisplayClass) New() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothPasskeyDisplay) Init() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothPasskeyDisplay) Autorelease() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothPasskeyDisplay creates a new BluetoothPasskeyDisplay instance.
func NewBluetoothPasskeyDisplay() BluetoothPasskeyDisplay {
	return getBluetoothPasskeyDisplayClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/retreatPasskeyIndicator()
func (b_ BluetoothPasskeyDisplay) RetreatPasskeyIndicator() {
	objc.Send[objc.ID](b_.ID, objc.Sel("retreatPasskeyIndicator"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/backgroundimageconstraint
func (b_ BluetoothPasskeyDisplay) BackgroundImageConstraint() appkit.LayoutConstraint {
	rv := objc.Send[appkit.LayoutConstraint](b_.ID, objc.Sel("backgroundImageConstraint"))
	return rv
}


// SetBackgroundImageConstraint sets the value of the backgroundImageConstraint property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/backgroundimageconstraint
func (b_ BluetoothPasskeyDisplay) SetBackgroundImageConstraint(value appkit.ILayoutConstraint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundImageConstraint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/centeredview
func (b_ BluetoothPasskeyDisplay) CenteredView() appkit.View {
	rv := objc.Send[appkit.View](b_.ID, objc.Sel("centeredView"))
	return rv
}


// SetCenteredView sets the value of the centeredView property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/centeredview
func (b_ BluetoothPasskeyDisplay) SetCenteredView(value appkit.IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCenteredView:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/isincomingrequest-swift.property
func (b_ BluetoothPasskeyDisplay) IsIncomingRequest() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncomingRequest"))
	return rv
}


// SetIsIncomingRequest sets the value of the isIncomingRequest property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/isincomingrequest-swift.property
func (b_ BluetoothPasskeyDisplay) SetIsIncomingRequest(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsIncomingRequest:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/passkey-swift.property
func (b_ BluetoothPasskeyDisplay) Passkey() appkit.string {
	rv := objc.Send[appkit.string](b_.ID, objc.Sel("passkey"))
	return rv
}


// SetPasskey sets the value of the passkey property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/passkey-swift.property
func (b_ BluetoothPasskeyDisplay) SetPasskey(value appkit.string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPasskey:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/returnhighlightimage
func (b_ BluetoothPasskeyDisplay) ReturnHighlightImage() appkit.Image {
	rv := objc.Send[appkit.Image](b_.ID, objc.Sel("returnHighlightImage"))
	return rv
}


// SetReturnHighlightImage sets the value of the returnHighlightImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/returnhighlightimage
func (b_ BluetoothPasskeyDisplay) SetReturnHighlightImage(value appkit.IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReturnHighlightImage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/returnimage
func (b_ BluetoothPasskeyDisplay) ReturnImage() appkit.Image {
	rv := objc.Send[appkit.Image](b_.ID, objc.Sel("returnImage"))
	return rv
}


// SetReturnImage sets the value of the returnImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/returnimage
func (b_ BluetoothPasskeyDisplay) SetReturnImage(value appkit.IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReturnImage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/usepasskeynotificaitons
func (b_ BluetoothPasskeyDisplay) UsePasskeyNotificaitons() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("usePasskeyNotificaitons"))
	return rv
}


// SetUsePasskeyNotificaitons sets the value of the usePasskeyNotificaitons property.
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetoothui/iobluetoothpasskeydisplay/usepasskeynotificaitons
func (b_ BluetoothPasskeyDisplay) SetUsePasskeyNotificaitons(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUsePasskeyNotificaitons:"), value)
}



