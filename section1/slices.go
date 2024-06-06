// Section 1: Language Fundamentals

// Slices
package section1

func SumEven(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number%2 == 0 {
			sum += number
		}
	}
	return sum
}
