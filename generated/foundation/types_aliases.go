// Type aliases for compatibility with objectivec package
package foundation

import "github.com/tmc/appledocs/generated/objectivec"

// These aliases allow foundation types to be used where objectivec expects them
type (
	_ = objectivec.Point
	_ = objectivec.Range
)
