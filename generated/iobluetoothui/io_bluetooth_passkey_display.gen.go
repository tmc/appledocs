// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/iobluetooth"
)

/* debug [class.gen.go]: Generating class IOBluetoothPasskeyDisplay */


/* debug [class_header]: Header for IOBluetoothPasskeyDisplay */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothPasskeyDisplay */
// An interface definition for the [BluetoothPasskeyDisplay] class.
type IBluetoothPasskeyDisplay interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for BluetoothPasskeyDisplay */
	// properties:
	BackgroundImageConstraint() appkit.LayoutConstraint
	SetBackgroundImageConstraint(value appkit.LayoutConstraint)
	CenteredView() appkit.View
	SetCenteredView(value appkit.View)
	IsIncomingRequest() bool
	SetIsIncomingRequest(value bool)
	Passkey() objc.IObject /* cross-framework: NSString */
	SetPasskey(value objc.IObject /* cross-framework: NSString */)
	ReturnHighlightImage() appkit.Image
	SetReturnHighlightImage(value appkit.Image)
	ReturnImage() appkit.Image
	SetReturnImage(value appkit.Image)
	UsePasskeyNotificaitons() bool
	SetUsePasskeyNotificaitons(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothPasskeyDisplay */
	// methods:
	AdvancePasskeyIndicator()
	ResetPasskeyIndicator()
	RetreatPasskeyIndicator()
	SetPasskeyForDeviceUsingSSP(inString objc.IObject /* cross-framework: NSString */, device iobluetooth.BluetoothDevice, isSSP bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothPasskeyDisplay */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothPasskeyDisplayClass) Alloc() BluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothPasskeyDisplay */


// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothPasskeyDisplay *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothPasskeyDisplay */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/sharedDisplayView()
func (bc _BluetoothPasskeyDisplayClass) SharedDisplayView() IBluetoothPasskeyDisplay {
	rv := objc.Send[BluetoothPasskeyDisplay](objc.ID(bc.class), objc.Sel("sharedDisplayView"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedDisplayView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothPasskeyDisplay */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothPasskeyDisplay */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/advancePasskeyIndicator()
func (b_ BluetoothPasskeyDisplay) AdvancePasskeyIndicator() {
	objc.Send[objc.ID](b_.ID, objc.Sel("advancePasskeyIndicator"))
}/* debug [instance_methods/method]: AdvancePasskeyIndicator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/resetPasskeyIndicator()
func (b_ BluetoothPasskeyDisplay) ResetPasskeyIndicator() {
	objc.Send[objc.ID](b_.ID, objc.Sel("resetPasskeyIndicator"))
}/* debug [instance_methods/method]: ResetPasskeyIndicator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/retreatPasskeyIndicator()
func (b_ BluetoothPasskeyDisplay) RetreatPasskeyIndicator() {
	objc.Send[objc.ID](b_.ID, objc.Sel("retreatPasskeyIndicator"))
}/* debug [instance_methods/method]: RetreatPasskeyIndicator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/setPasskey(_:for:usingSSP:)
func (b_ BluetoothPasskeyDisplay) SetPasskeyForDeviceUsingSSP(inString objc.IObject /* cross-framework: NSString */, device iobluetooth.BluetoothDevice, isSSP bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPasskey:forDevice:usingSSP:"), inString, device, isSSP)
}/* debug [instance_methods/method]: SetPasskeyForDeviceUsingSSP */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothPasskeyDisplay */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/backgroundImageConstraint
func (b_ BluetoothPasskeyDisplay) BackgroundImageConstraint() appkit.LayoutConstraint {
	rv := objc.Send[appkit.LayoutConstraint](b_.ID, objc.Sel("backgroundImageConstraint"))
	return rv
}/* debug [instance_properties/getter]: backgroundImageConstraint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/backgroundImageConstraint
func (b_ BluetoothPasskeyDisplay) SetBackgroundImageConstraint(value appkit.LayoutConstraint) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBackgroundImageConstraint:"), value)
}/* debug [instance_properties/setter]: backgroundImageConstraint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/centeredView
func (b_ BluetoothPasskeyDisplay) CenteredView() appkit.View {
	rv := objc.Send[appkit.View](b_.ID, objc.Sel("centeredView"))
	return rv
}/* debug [instance_properties/getter]: centeredView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/centeredView
func (b_ BluetoothPasskeyDisplay) SetCenteredView(value appkit.View) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCenteredView:"), value)
}/* debug [instance_properties/setter]: centeredView */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/isIncomingRequest-swift.property
func (b_ BluetoothPasskeyDisplay) IsIncomingRequest() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isIncomingRequest"))
	return rv
}/* debug [instance_properties/getter]: isIncomingRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/isIncomingRequest-swift.property
func (b_ BluetoothPasskeyDisplay) SetIsIncomingRequest(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsIncomingRequest:"), value)
}/* debug [instance_properties/setter]: isIncomingRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/passkey-swift.property
func (b_ BluetoothPasskeyDisplay) Passkey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("passkey"))
	return rv
}/* debug [instance_properties/getter]: passkey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/passkey-swift.property
func (b_ BluetoothPasskeyDisplay) SetPasskey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPasskey:"), value)
}/* debug [instance_properties/setter]: passkey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/returnHighlightImage
func (b_ BluetoothPasskeyDisplay) ReturnHighlightImage() appkit.Image {
	rv := objc.Send[appkit.Image](b_.ID, objc.Sel("returnHighlightImage"))
	return rv
}/* debug [instance_properties/getter]: returnHighlightImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/returnHighlightImage
func (b_ BluetoothPasskeyDisplay) SetReturnHighlightImage(value appkit.Image) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReturnHighlightImage:"), value)
}/* debug [instance_properties/setter]: returnHighlightImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/returnImage
func (b_ BluetoothPasskeyDisplay) ReturnImage() appkit.Image {
	rv := objc.Send[appkit.Image](b_.ID, objc.Sel("returnImage"))
	return rv
}/* debug [instance_properties/getter]: returnImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/returnImage
func (b_ BluetoothPasskeyDisplay) SetReturnImage(value appkit.Image) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setReturnImage:"), value)
}/* debug [instance_properties/setter]: returnImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/usePasskeyNotificaitons
func (b_ BluetoothPasskeyDisplay) UsePasskeyNotificaitons() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("usePasskeyNotificaitons"))
	return rv
}/* debug [instance_properties/getter]: usePasskeyNotificaitons */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/usePasskeyNotificaitons
func (b_ BluetoothPasskeyDisplay) SetUsePasskeyNotificaitons(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUsePasskeyNotificaitons:"), value)
}/* debug [instance_properties/setter]: usePasskeyNotificaitons */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothPasskeyDisplay */



