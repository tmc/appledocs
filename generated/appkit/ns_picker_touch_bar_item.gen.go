// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPickerTouchBarItem */


/* debug [class_header]: Header for NSPickerTouchBarItem */
// The class instance for the [PickerTouchBarItem] class.
var (
	PickerTouchBarItemClass     _PickerTouchBarItemClass
	PickerTouchBarItemClassOnce sync.Once
)

func getPickerTouchBarItemClass() _PickerTouchBarItemClass {
	PickerTouchBarItemClassOnce.Do(func() {
		PickerTouchBarItemClass = _PickerTouchBarItemClass{objc.GetClass("NSPickerTouchBarItem")}
	})
	return PickerTouchBarItemClass
}

type _PickerTouchBarItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PickerTouchBarItem */
// An interface definition for the [PickerTouchBarItem] class.
type IPickerTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for PickerTouchBarItem */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	CollapsedRepresentationImage() IImage
	SetCollapsedRepresentationImage(value IImage)
	CollapsedRepresentationLabel() objc.IObject /* cross-framework: NSString */
	SetCollapsedRepresentationLabel(value objc.IObject /* cross-framework: NSString */)
	ControlRepresentation() PickerTouchBarItemControlRepresentation
	SetControlRepresentation(value PickerTouchBarItemControlRepresentation)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Enabled() bool
	SetEnabled(value bool)
	NumberOfOptions() int
	SetNumberOfOptions(value int)
	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectionColor() IColor
	SetSelectionColor(value IColor)
	SelectionMode() PickerTouchBarItemSelectionMode
	SetSelectionMode(value PickerTouchBarItemSelectionMode)
	Target() objc.ID
	SetTarget(value objc.ID)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PickerTouchBarItem */
	// methods:
	ImageAtIndex(index int) IImage
	IsEnabledAtIndex(index int) bool
	LabelAtIndex(index int) foundation.String
	SetEnabledAtIndex(enabled bool, index int)
	SetImageAtIndex(image IImage, index int)
	SetLabelAtIndex(label objc.IObject /* cross-framework: NSString */, index int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PickerTouchBarItem */
// Alloc allocates a new instance without initialization.
func (pc _PickerTouchBarItemClass) Alloc() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PickerTouchBarItemClass) New() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PickerTouchBarItem) Init() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PickerTouchBarItem) Autorelease() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPickerTouchBarItem creates a new PickerTouchBarItem instance.
func NewPickerTouchBarItem() PickerTouchBarItem {
	return getPickerTouchBarItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PickerTouchBarItem */
// A bar item that provides a picker control with multiple options.


// A bar item that provides a picker control with multiple options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem
type PickerTouchBarItem struct {
	TouchBarItem
}

// PickerTouchBarItemFrom constructs a [PickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a picker control with multiple options.
func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PickerTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:images:selectionMode:target:action:)
func NewPickerTouchBarItemWithIdentifierImagesSelectionModeTargetAction(identifier TouchBarItemIdentifier /* typedef */, images []Image, selectionMode PickerTouchBarItemSelectionMode, target objc.IObject, action objc.SEL) PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(getPickerTouchBarItemClass().class), objc.Sel("pickerTouchBarItemWithIdentifier:images:selectionMode:target:action:"), identifier, images, selectionMode, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewPickerTouchBarItemWithIdentifierImagesSelectionModeTargetAction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func NewPickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier TouchBarItemIdentifier /* typedef */, labels []string, selectionMode PickerTouchBarItemSelectionMode, target objc.IObject, action objc.SEL) PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(getPickerTouchBarItemClass().class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewPickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PickerTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:images:selectionMode:target:action:)
func (pc _PickerTouchBarItemClass) PickerTouchBarItemWithIdentifierImagesSelectionModeTargetAction(identifier TouchBarItemIdentifier /* typedef */, images []Image, selectionMode PickerTouchBarItemSelectionMode, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pickerTouchBarItemWithIdentifier:images:selectionMode:target:action:"), identifier, images, selectionMode, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PickerTouchBarItemWithIdentifierImagesSelectionModeTargetAction) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func (pc _PickerTouchBarItemClass) PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier TouchBarItemIdentifier /* typedef */, labels []string, selectionMode PickerTouchBarItemSelectionMode, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PickerTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PickerTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/image(at:)
func (p_ PickerTouchBarItem) ImageAtIndex(index int) IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("imageAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: ImageAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/isEnabled(at:)
func (p_ PickerTouchBarItem) IsEnabledAtIndex(index int) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEnabledAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: IsEnabledAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/label(at:)
func (p_ PickerTouchBarItem) LabelAtIndex(index int) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("labelAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: LabelAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/setEnabled(_:at:)
func (p_ PickerTouchBarItem) SetEnabledAtIndex(enabled bool, index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnabled:atIndex:"), enabled, index)
}/* debug [instance_methods/method]: SetEnabledAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/setImage(_:at:)
func (p_ PickerTouchBarItem) SetImageAtIndex(image IImage, index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImage:atIndex:"), image, index)
}/* debug [instance_methods/method]: SetImageAtIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/setLabel(_:at:)
func (p_ PickerTouchBarItem) SetLabelAtIndex(label objc.IObject /* cross-framework: NSString */, index int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLabel:atIndex:"), label, index)
}/* debug [instance_methods/method]: SetLabelAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PickerTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/action
func (p_ PickerTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](p_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/action
func (p_ PickerTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p_ PickerTouchBarItem) CollapsedRepresentationImage() IImage {
	rv := objc.Send[Image](p_.ID, objc.Sel("collapsedRepresentationImage"))
	return rv
}/* debug [instance_properties/getter]: collapsedRepresentationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p_ PickerTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}/* debug [instance_properties/setter]: collapsedRepresentationImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p_ PickerTouchBarItem) CollapsedRepresentationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("collapsedRepresentationLabel"))
	return rv
}/* debug [instance_properties/getter]: collapsedRepresentationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p_ PickerTouchBarItem) SetCollapsedRepresentationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationLabel:"), value)
}/* debug [instance_properties/setter]: collapsedRepresentationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/controlRepresentation-swift.property
func (p_ PickerTouchBarItem) ControlRepresentation() PickerTouchBarItemControlRepresentation {
	rv := objc.Send[PickerTouchBarItemControlRepresentation](p_.ID, objc.Sel("controlRepresentation"))
	return rv
}/* debug [instance_properties/getter]: controlRepresentation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/controlRepresentation-swift.property
func (p_ PickerTouchBarItem) SetControlRepresentation(value PickerTouchBarItemControlRepresentation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlRepresentation:"), value)
}/* debug [instance_properties/setter]: controlRepresentation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/customizationLabel
func (p_ PickerTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/customizationLabel
func (p_ PickerTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/isEnabled
func (p_ PickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/isEnabled
func (p_ PickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/numberOfOptions
func (p_ PickerTouchBarItem) NumberOfOptions() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfOptions"))
	return rv
}/* debug [instance_properties/getter]: numberOfOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/numberOfOptions
func (p_ PickerTouchBarItem) SetNumberOfOptions(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfOptions:"), value)
}/* debug [instance_properties/setter]: numberOfOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectedIndex
func (p_ PickerTouchBarItem) SelectedIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectedIndex"))
	return rv
}/* debug [instance_properties/getter]: selectedIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectedIndex
func (p_ PickerTouchBarItem) SetSelectedIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedIndex:"), value)
}/* debug [instance_properties/setter]: selectedIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectionColor
func (p_ PickerTouchBarItem) SelectionColor() IColor {
	rv := objc.Send[Color](p_.ID, objc.Sel("selectionColor"))
	return rv
}/* debug [instance_properties/getter]: selectionColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectionColor
func (p_ PickerTouchBarItem) SetSelectionColor(value IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionColor:"), value)
}/* debug [instance_properties/setter]: selectionColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectionMode-swift.property
func (p_ PickerTouchBarItem) SelectionMode() PickerTouchBarItemSelectionMode {
	rv := objc.Send[PickerTouchBarItemSelectionMode](p_.ID, objc.Sel("selectionMode"))
	return rv
}/* debug [instance_properties/getter]: selectionMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/selectionMode-swift.property
func (p_ PickerTouchBarItem) SetSelectionMode(value PickerTouchBarItemSelectionMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionMode:"), value)
}/* debug [instance_properties/setter]: selectionMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/target
func (p_ PickerTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/target
func (p_ PickerTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/isenabled
func (p_ PickerTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/isenabled
func (p_ PickerTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPickerTouchBarItem */


