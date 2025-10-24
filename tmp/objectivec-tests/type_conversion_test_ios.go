//go:build ios || tvos || visionos

// Type conversion tests for ObjectiveC framework
// Tests the type mapping scenarios we fixed for appledocs-531

package objectivec_test

import (
	"testing"
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
)

// TestTypeConversionCompiles verifies that all type conversions compile correctly
// This test validates the fix for appledocs-531 where array returns were incorrectly typed
func TestTypeConversionCompiles(t *testing.T) {
	obj := objectivec.NewObject()

	// Test 1: []objc.ID -> []IObject conversion compiles
	// Before fix: would error "cannot use rv (variable of type []objc.ID) as IObject value"
	var labels []objectivec.IObject = obj.AccessibilityAttributedUserInputLabels()
	_ = labels
	t.Log("✓ []objc.ID -> []IObject conversion compiles")

	// Test 2: objc.ID -> IObject conversion compiles
	// Before fix: would error "cannot use rv (variable of type objc.ID) as IObject value"
	var frame objectivec.IObject = obj.AccessibilityFrame()
	_ = frame
	t.Log("✓ objc.ID -> IObject conversion compiles")

	// Test 3: []IObject parameter type compiles
	newLabels := make([]objectivec.IObject, 0)
	obj.SetAccessibilityAttributedUserInputLabels(newLabels)
	t.Log("✓ []IObject parameter type compiles")

	// Test 4: IObject parameter type compiles
	newFrame := objectivec.NewObject()
	obj.SetAccessibilityFrame(newFrame)
	t.Log("✓ IObject parameter type compiles")

	// Test 5: Primitives are not converted
	var version uint = obj.ImageVersion()
	_ = version
	t.Log("✓ Primitive types are not converted")

	// Test 6: Function types are preserved
	var block func() unsafe.Pointer = obj.AccessibilityExpandedStatusBlock()
	_ = block
	t.Log("✓ Function types are preserved")

	t.Log("All type conversions from appledocs-531 fix are working correctly!")
}

// TestInterfaceAssignment verifies that IObject interface can be used
func TestInterfaceAssignment(t *testing.T) {
	obj := objectivec.NewObject()

	// Can assign concrete Object to IObject
	var iobj objectivec.IObject = obj
	t.Log("✓ Object → IObject assignment works")

	// Can call property getters on interface
	_ = iobj.AccessibilityFrame()
	t.Log("✓ Property getters work on IObject interface")

	// Can call property setters on interface
	iobj.SetAccessibilityFrame(obj)
	t.Log("✓ Property setters work on IObject interface")

	// Can assign property getter result to IObject
	var frame objectivec.IObject = iobj.AccessibilityFrame()
	_ = frame
	t.Log("✓ Property getter results assign to IObject interface")
}

// TestArrayTypeConversion verifies array wrapper generation
func TestArrayTypeConversion(t *testing.T) {
	obj := objectivec.NewObject()

	// Get array property
	labels := obj.AccessibilityAttributedUserInputLabels()

	// Verify it's typed as []IObject, not []objc.ID
	var _ []objectivec.IObject = labels

	// Can iterate as IObject interface
	for i, label := range labels {
		var _ objectivec.IObject = label
		t.Logf("✓ Array element %d is IObject interface type", i)
	}

	if len(labels) == 0 {
		t.Log("✓ Array property returns empty array (expected for base NSObject)")
	}
}

// TestObjectIDAccess verifies we can access the underlying objc.ID
func TestObjectIDAccess(t *testing.T) {
	obj := objectivec.NewObject()

	// Should be able to access the ID field (Object embeds objc.ID)
	id := obj.ID
	if id == 0 {
		t.Error("Object ID should not be zero")
	}
	t.Logf("✓ Can access Object.ID field: %d", id)
}

// TestCrossFrameworkFallback verifies types from other frameworks map to IObject
// This validates the framework hierarchy violation fix
func TestCrossFrameworkFallback(t *testing.T) {
	obj := objectivec.NewObject()

	// This method would return foundation.Number in Foundation framework
	// In ObjectiveC framework, it should fall back to []IObject
	var indices []objectivec.IObject = obj.IndicesOfObjectsByEvaluatingObjectSpecifier(obj)
	_ = indices

	t.Log("✓ Cross-framework types properly fall back to IObject in ObjectiveC framework")
}
