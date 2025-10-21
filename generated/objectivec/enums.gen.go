// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

// Enum types and constants
// objc_AssociationPolicy - Type to specify the behavior of an association.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_AssociationPolicy
type objc_AssociationPolicy uint

const (
	// OBJC_ASSOCIATION_ASSIGN - Specifies an unsafe unretained reference to the associated object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_AssociationPolicy/OBJC_ASSOCIATION_ASSIGN
	OBJC_ASSOCIATION_ASSIGN objc_AssociationPolicy = 0
	// OBJC_ASSOCIATION_COPY_NONATOMIC - Specifies that the associated object is copied, and that the association is not made atomically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_AssociationPolicy/OBJC_ASSOCIATION_COPY_NONATOMIC
	OBJC_ASSOCIATION_COPY_NONATOMIC objc_AssociationPolicy = 0
	// OBJC_ASSOCIATION_RETAIN - Specifies a strong reference to the associated object, and that the association is made atomically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_AssociationPolicy/OBJC_ASSOCIATION_RETAIN
	OBJC_ASSOCIATION_RETAIN objc_AssociationPolicy = 0
	// OBJC_ASSOCIATION_RETAIN_NONATOMIC - Specifies a strong reference to the associated object, and that the association is not made atomically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_AssociationPolicy/OBJC_ASSOCIATION_RETAIN_NONATOMIC
	OBJC_ASSOCIATION_RETAIN_NONATOMIC objc_AssociationPolicy = 0
)


