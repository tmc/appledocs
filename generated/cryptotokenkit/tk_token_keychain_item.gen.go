// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenKeychainItem */


/* debug [class_header]: Header for TKTokenKeychainItem */
// The class instance for the [TKTokenKeychainItem] class.
var (
	TKTokenKeychainItemClass     _TKTokenKeychainItemClass
	TKTokenKeychainItemClassOnce sync.Once
)

func getTKTokenKeychainItemClass() _TKTokenKeychainItemClass {
	TKTokenKeychainItemClassOnce.Do(func() {
		TKTokenKeychainItemClass = _TKTokenKeychainItemClass{objc.GetClass("TKTokenKeychainItem")}
	})
	return TKTokenKeychainItemClass
}

type _TKTokenKeychainItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenKeychainItem */
// An interface definition for the [TKTokenKeychainItem] class.
type ITKTokenKeychainItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenKeychainItem */
	// properties:
	Constraints() foundation.IDictionary
	SetConstraints(value foundation.IDictionary)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	ObjectID() TKTokenObjectID /* typedef */
	KeychainContents() ITKTokenKeychainContents
	SetKeychainContents(value ITKTokenKeychainContents)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenKeychainItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenKeychainItem */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeychainItemClass) Alloc() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenKeychainItemClass) New() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeychainItem) Init() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeychainItem) Autorelease() TKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainItem creates a new TKTokenKeychainItem instance.
func NewTKTokenKeychainItem() TKTokenKeychainItem {
	return getTKTokenKeychainItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenKeychainItem */
// An abstract base class for managing a token’s contents as keychain items.
//
// Don’t use this base class directly. Instead, use one of its subclasses, such as for managing certificates or for managing cryptographic keys.


// An abstract base class for managing a token’s contents as keychain items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem
type TKTokenKeychainItem struct {
	objectivec.Object
}

// TKTokenKeychainItemFrom constructs a [TKTokenKeychainItem] from an unsafe.Pointer.
//
// An abstract base class for managing a token’s contents as keychain items.
func TKTokenKeychainItemFrom(ptr unsafe.Pointer) TKTokenKeychainItem {
	return TKTokenKeychainItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenKeychainItem */

// Initializes a token keychain item with the specified object ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/init(objectID:)
func NewTKTokenKeychainItemWithObjectID(objectID TKTokenObjectID /* typedef */) TKTokenKeychainItem {
	instance := getTKTokenKeychainItemClass().Alloc()
	rv := objc.Send[TKTokenKeychainItem](instance.ID, objc.Sel("initWithObjectID:"), objectID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKTokenKeychainItemWithObjectID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenKeychainItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenKeychainItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenKeychainItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenKeychainItem */

// Access constraints for the keychain item, keyed by values wrapped in objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/constraints
func (t_ TKTokenKeychainItem) Constraints() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("constraints"))
	return rv
}/* debug [instance_properties/getter]: constraints */


// Access constraints for the keychain item, keyed by values wrapped in objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/constraints
func (t_ TKTokenKeychainItem) SetConstraints(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConstraints:"), value)
}/* debug [instance_properties/setter]: constraints */


// The user-visible label for the keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/label
func (t_ TKTokenKeychainItem) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The user-visible label for the keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/label
func (t_ TKTokenKeychainItem) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Returns the object ID used for keychain item identification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainItem/objectID
func (t_ TKTokenKeychainItem) ObjectID() TKTokenObjectID /* typedef */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("objectID"))
	return rv
}/* debug [instance_properties/getter]: objectID */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainItem) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}/* debug [instance_properties/getter]: keychainContents */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainItem) SetKeychainContents(value ITKTokenKeychainContents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainContents:"), value)
}/* debug [instance_properties/setter]: keychainContents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenKeychainItem */


