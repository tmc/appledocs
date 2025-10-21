package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tmc/appledocs/occ2go"
)

// prepareClassMethods filters methods to return only class methods, deduplicated by Go method name.
// When multiple methods would generate the same Go method name (e.g., foo and foo:),
// the methods are disambiguated by appending parameter labels and the Name field is updated.
func prepareClassMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seenSelectors := make(map[string]bool)

	// First pass: collect all class methods
	var classMethods []*occ2go.ParsedMethod
	for _, m := range methods {
		if m.IsClassMethod && !seenSelectors[m.Selector] {
			classMethods = append(classMethods, m)
			seenSelectors[m.Selector] = true
		}
	}

	// Second pass: detect Go method name collisions and disambiguate
	goNameCounts := make(map[string]int)

	// Count how many methods map to each Go name
	for _, m := range classMethods {
		goName := selectorToGoName(m.Selector)
		goNameCounts[goName]++
	}

	// Third pass: build result with disambiguation, updating .Name as needed
	seenGoNames := make(map[string]bool)
	for _, m := range classMethods {
		goName := selectorToGoName(m.Selector)

		// If this Go name has duplicates, disambiguate using parameter labels
		if goNameCounts[goName] > 1 {
			// Create a copy of the method and update its Name field
			methodCopy := *m
			methodCopy.Name = disambiguateMethodName(m)
			goName = methodCopy.Name

			// Skip if we've already seen this exact Go name (shouldn't happen after disambiguation)
			if seenGoNames[goName] {
				continue
			}

			result = append(result, &methodCopy)
			seenGoNames[goName] = true
		} else {
			// No collision, use original method
			if seenGoNames[goName] {
				continue
			}
			result = append(result, m)
			seenGoNames[goName] = true
		}
	}

	return result
}

// prepareInstanceMethods filters methods to return only instance methods,
// excluding ALL init methods (which are converted to constructors), deduplicated by Go method name.
// When multiple methods would generate the same Go method name (e.g., foo and foo:),
// the methods are disambiguated by appending parameter labels and the Name field is updated.
func prepareInstanceMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seenSelectors := make(map[string]bool)

	// First pass: collect all instance methods (excluding init)
	var instanceMethods []*occ2go.ParsedMethod
	for _, m := range methods {
		if !m.IsClassMethod && !seenSelectors[m.Selector] {
			// Skip ALL init methods - they're converted to package-level constructors
			if strings.HasPrefix(m.Selector, "init") {
				continue
			}
			instanceMethods = append(instanceMethods, m)
			seenSelectors[m.Selector] = true
		}
	}

	// Second pass: detect Go method name collisions and disambiguate
	goNameCounts := make(map[string]int)

	// Count how many methods map to each Go name
	for _, m := range instanceMethods {
		goName := selectorToGoName(m.Selector)
		goNameCounts[goName]++
	}

	// Third pass: build result with disambiguation, updating .Name as needed
	seenGoNames := make(map[string]bool)
	for _, m := range instanceMethods {
		goName := selectorToGoName(m.Selector)

		// If this Go name has duplicates, disambiguate using parameter labels
		if goNameCounts[goName] > 1 {
			// Create a copy of the method and update its Name field
			methodCopy := *m
			methodCopy.Name = disambiguateMethodName(m)
			goName = methodCopy.Name

			// Skip if we've already seen this exact Go name (shouldn't happen after disambiguation)
			if seenGoNames[goName] {
				continue
			}

			result = append(result, &methodCopy)
			seenGoNames[goName] = true
		} else {
			// No collision, use original method
			if seenGoNames[goName] {
				continue
			}
			result = append(result, m)
			seenGoNames[goName] = true
		}
	}

	return result
}

// prepareInitMethods filters methods to return only init methods (for constructor generation), deduplicated by constructor name.
// Includes both:
//   - Instance methods with selectors starting with "init" (traditional init methods)
//   - Class methods marked as initializers in documentation (factory methods like buttonWithTitle:target:action:)
//
// When multiple methods would generate the same constructor name (e.g., initWithContentsOfURL: and arrayWithContentsOfURL:),
// the instance method is preferred.
func prepareInitMethods(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	result := make([]*occ2go.ParsedMethod, 0)
	seen := make(map[string]*occ2go.ParsedMethod)

	// First pass: collect all potential init methods with their constructor names
	for _, m := range methods {
		isInit := false

		// Traditional instance init methods
		if !m.IsClassMethod && strings.HasPrefix(m.Selector, "init") {
			isInit = true
		}

		// Class factory methods marked as initializers in docs (e.g., buttonWithTitle:target:action:)
		if m.IsClassMethod && m.IsInitializer {
			isInit = true
		}

		if isInit {
			// We don't have the className here, so we'll use a simple dedup strategy:
			// Prefer instance methods over class methods with similar signatures
			key := m.Selector
			existing, exists := seen[key]

			if !exists {
				seen[key] = m
			} else {
				// If we have both an instance and class method, prefer instance
				// Instance methods take precedence because they're the "real" initializers
				if !m.IsClassMethod && existing.IsClassMethod {
					seen[key] = m
				}
			}
		}
	}

	// Second pass: deduplicate by actual constructor name
	constructorNames := make(map[string]bool)
	for _, m := range methods {
		if seenMethod, exists := seen[m.Selector]; exists && seenMethod == m {
			// Generate constructor name (we need className, but we don't have it here)
			// So we'll do a simpler check: deduplicate by parameter signature
			paramSig := fmt.Sprintf("%d", len(m.Parameters))
			for _, p := range m.Parameters {
				paramSig += ":" + p.Type
			}
			constructorKey := m.Selector + paramSig

			if !constructorNames[constructorKey] {
				result = append(result, m)
				constructorNames[constructorKey] = true
			}
		}
	}
	return result
}

// prepareInitMethodsWithClassName properly deduplicates init methods by their generated constructor names.
// This fixes cases where both instance and class factory methods would generate the same constructor name
// (e.g., initWithContentsOfURL: and arrayWithContentsOfURL: both map to NewArrayWithContentsOfURL).
// Instance methods are preferred over class factory methods when deduplicating.
func prepareInitMethodsWithClassName(className string, methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	// First get all potential init methods
	initMethods := prepareInitMethods(methods)

	// Deduplicate by constructor name
	seen := make(map[string]*occ2go.ParsedMethod)
	for _, m := range initMethods {
		constructorName := initMethodToConstructorName(className, m.Selector)

		existing, exists := seen[constructorName]
		if !exists {
			seen[constructorName] = m
		} else {
			// If we have both an instance and class method mapping to the same name,
			// prefer the instance method
			if !m.IsClassMethod && existing.IsClassMethod {
				seen[constructorName] = m
			}
		}
	}

	// Convert map back to slice, sorted by constructor name for stable output
	constructorNames := make([]string, 0, len(seen))
	for name := range seen {
		constructorNames = append(constructorNames, name)
	}
	sort.Strings(constructorNames)

	result := make([]*occ2go.ParsedMethod, 0, len(seen))
	for _, name := range constructorNames {
		result = append(result, seen[name])
	}

	return result
}

// sortMethodsByName sorts methods by name for consistent output
func sortMethodsByName(methods []*occ2go.ParsedMethod) []*occ2go.ParsedMethod {
	sorted := make([]*occ2go.ParsedMethod, len(methods))
	copy(sorted, methods)

	// Simple bubble sort by Name
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i].Name > sorted[j].Name {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	return sorted
}

// isConstructor checks if a method is a constructor (returns instance of class).
// Delegates to occ2go.IsFactoryMethodForClass.
func isConstructor(method *occ2go.ParsedMethod, className string) bool {
	return occ2go.IsFactoryMethodForClass(method.Selector, className, method.IsClassMethod)
}

// classDependsOnCoreGraphics returns true if any method in the class uses CoreGraphics types
// that actually require the coregraphics import (i.e., not mapped to unsafe.Pointer).
func classDependsOnCoreGraphics(methods []*occ2go.ParsedMethod, framework string) bool {
	if framework == "CoreGraphics" {
		return false
	}

	for _, m := range methods {
		// Check return type - use the same mapping logic as getRequiredImports
		if m.ReturnType != "" && m.ReturnType != "void" {
			goType := mapObjCTypeToGo(m.ReturnType, framework)
			// Check if this actually needs coregraphics import (not unsafe.Pointer)
			if strings.HasPrefix(goType, "coregraphics.") {
				return true
			}
		}

		// Check parameters - use the same mapping logic as getRequiredImports
		for _, p := range m.Parameters {
			goType := mapObjCTypeToGo(p.Type, framework)
			// Check if this actually needs coregraphics import (not unsafe.Pointer)
			if strings.HasPrefix(goType, "coregraphics.") {
				return true
			}
		}
	}

	return false
}

// isInheritedFromNSObject checks if a selector is likely inherited from NSObject.
// These methods should not be redeclared in child interfaces.
// Returns true for common NSObject methods that appear in most/all classes.
func isInheritedFromNSObject(selector string) bool {
	return occ2go.IsInheritedFromNSObject(selector)
}

// filterPropertyMethods removes methods that are generated from properties (getters/setters)
// to prevent duplicate generation. Property methods are generated separately in the properties section.
func filterPropertyMethods(class *occ2go.ParsedClass) []*occ2go.ParsedMethod {
	if class == nil {
		return nil
	}

	// Build set of property selectors (getter and setter)
	// Also build a set of property setter METHOD names (without colon) that would collide with generated setters
	propertySelectors := make(map[string]bool)
	propertySetterMethods := make(map[string]bool) // Maps "setFoo" -> true if property "foo" exists
	for _, prop := range class.Properties {
		// Getter selector is just the property name
		propertySelectors[prop.Name] = true

		// Setter selector is "set<CapitalizedName>:"
		capitalizedName := strings.ToUpper(prop.Name[:1]) + prop.Name[1:]
		setterSelector := "set" + capitalizedName + ":"
		propertySelectors[setterSelector] = true

		// Also track the setter method name without colon for collision detection
		// E.g., property "accessibilityFrameInParentSpace" -> method "setAccessibilityFrameInParentSpace"
		setterMethodName := "set" + capitalizedName
		propertySetterMethods[setterMethodName] = true
	}

	// Filter methods, excluding those that match property selectors
	var filtered []*occ2go.ParsedMethod

	// Build a map of method selectors to detect setter-like collisions
	methodSelectors := make(map[string]*occ2go.ParsedMethod)
	for _, m := range class.Methods {
		methodSelectors[m.Selector] = m
	}

	for _, m := range class.Methods {
		if !m.IsClassMethod && !propertySelectors[m.Selector] {
			// Skip parameterless set* methods that would collide with property setters
			// Example: method setAccessibilityFrameInParentSpace() collides with property accessibilityFrameInParentSpace's setter
			if strings.HasPrefix(m.Selector, "set") && len(m.Parameters) == 0 && !strings.HasSuffix(m.Selector, ":") {
				if propertySetterMethods[m.Selector] {
					// Skip this method because it collides with a generated property setter
					continue
				}

				// Also check if there's a setter version with a colon (method overload case)
				setterVersion := m.Selector + ":"
				if setterMethod, exists := methodSelectors[setterVersion]; exists && len(setterMethod.Parameters) > 0 {
					// Skip this parameterless version as it collides with the parameterized method
					continue
				}
			}
			filtered = append(filtered, m)
		} else if m.IsClassMethod {
			// Always include class methods
			filtered = append(filtered, m)
		}
	}

	// Now call prepareInstanceMethods on the filtered list
	return prepareInstanceMethods(filtered)
}
