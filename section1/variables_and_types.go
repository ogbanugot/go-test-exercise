// Section 1: Language Fundamentals

// Variables and Types
package section1

func Swap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
