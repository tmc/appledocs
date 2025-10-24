// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IKSaveOptions */


/* debug [class_header]: Header for IKSaveOptions */
// The class instance for the [IKSaveOptions] class.
var (
	IKSaveOptionsClass     _IKSaveOptionsClass
	IKSaveOptionsClassOnce sync.Once
)

func getIKSaveOptionsClass() _IKSaveOptionsClass {
	IKSaveOptionsClassOnce.Do(func() {
		IKSaveOptionsClass = _IKSaveOptionsClass{objc.GetClass("IKSaveOptions")}
	})
	return IKSaveOptionsClass
}

type _IKSaveOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKSaveOptions */
// An interface definition for the [IKSaveOptions] class.
type IIKSaveOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for IKSaveOptions */
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	ImageProperties() objc.IObject /* cross-framework: NSDictionary */
	ImageUTType() objc.IObject /* cross-framework: NSString */
	RememberLastSetting() bool
	SetRememberLastSetting(value bool)
	UserSelection() objc.IObject /* cross-framework: NSDictionary */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKSaveOptions */
	// methods:
	AddSaveOptionsToView(view appkit.View)
	AddSaveOptionsAccessoryViewToSavePanel(savePanel appkit.SavePanel)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKSaveOptions */
// Alloc allocates a new instance without initialization.
func (ic _IKSaveOptionsClass) Alloc() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKSaveOptionsClass) New() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKSaveOptions) Init() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKSaveOptions) Autorelease() IKSaveOptions {
	rv := objc.Send[IKSaveOptions](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKSaveOptions creates a new IKSaveOptions instance.
func NewIKSaveOptions() IKSaveOptions {
	return getIKSaveOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKSaveOptions */
// The class initializes, adds, and manages user interface options for saving image data.


// The class initializes, adds, and manages user interface options for saving image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions
type IKSaveOptions struct {
	objectivec.Object
}

// IKSaveOptionsFrom constructs a [IKSaveOptions] from an unsafe.Pointer.
//
// The class initializes, adds, and manages user interface options for saving image data.
func IKSaveOptionsFrom(ptr unsafe.Pointer) IKSaveOptions {
	return IKSaveOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKSaveOptions */

// Initializes a save options accessory pane for the provided image properties and uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/init(imageProperties:imageUTType:)
func NewIKSaveOptionsWithImagePropertiesImageUTType(imageProperties objc.IObject /* cross-framework: NSDictionary */, imageUTType objc.IObject /* cross-framework: NSString */) IKSaveOptions {
	instance := getIKSaveOptionsClass().Alloc()
	rv := objc.Send[IKSaveOptions](instance.ID, objc.Sel("initWithImageProperties:imageUTType:"), imageProperties, imageUTType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIKSaveOptionsWithImagePropertiesImageUTType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKSaveOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKSaveOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKSaveOptions */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/add(to:)
func (i_ IKSaveOptions) AddSaveOptionsToView(view appkit.View) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addSaveOptionsToView:"), view)
}/* debug [instance_methods/method]: AddSaveOptionsToView */


// Adds accessory view to a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/addAccessoryView(to:)
func (i_ IKSaveOptions) AddSaveOptionsAccessoryViewToSavePanel(savePanel appkit.SavePanel) {
	objc.Send[objc.ID](i_.ID, objc.Sel("addSaveOptionsAccessoryViewToSavePanel:"), savePanel)
}/* debug [instance_methods/method]: AddSaveOptionsAccessoryViewToSavePanel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKSaveOptions */

// Specifies the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/delegate
func (i_ IKSaveOptions) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Specifies the delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/delegate
func (i_ IKSaveOptions) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Returns a dictionary of updated image properties that reflects the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/imageProperties
func (i_ IKSaveOptions) ImageProperties() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](i_.ID, objc.Sel("imageProperties"))
	return rv
}/* debug [instance_properties/getter]: imageProperties */


// Returns the uniform type identifier that reflects the user’s selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/imageUTType
func (i_ IKSaveOptions) ImageUTType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](i_.ID, objc.Sel("imageUTType"))
	return rv
}/* debug [instance_properties/getter]: imageUTType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/rememberLastSetting
func (i_ IKSaveOptions) RememberLastSetting() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("rememberLastSetting"))
	return rv
}/* debug [instance_properties/getter]: rememberLastSetting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/rememberLastSetting
func (i_ IKSaveOptions) SetRememberLastSetting(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRememberLastSetting:"), value)
}/* debug [instance_properties/setter]: rememberLastSetting */


// Returns a dictionary that contains the save options selected by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSaveOptions/userSelection
func (i_ IKSaveOptions) UserSelection() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](i_.ID, objc.Sel("userSelection"))
	return rv
}/* debug [instance_properties/getter]: userSelection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKSaveOptions */


