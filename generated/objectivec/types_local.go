// Local type definitions to avoid circular import with foundation package
package objectivec

// Point represents an NSPoint/CGPoint structure
type Point struct {
	X float64
	Y float64
}

// Range represents an NSRange structure
type Range struct {
	Location uint
	Length   uint
}
