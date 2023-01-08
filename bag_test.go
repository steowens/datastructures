/*
github.com/steowens/datastructures (c) by Stephen Owens>

github.com/steowens/datastructures is licensed under a
Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.

You should have received a copy of the license (LICENSE) along with this
work. If not, see <http://creativecommons.org/licenses/by-nc-sa/4.0/>.
*/
package collection

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestAddBag(t *testing.T) {
	bag := NewBag[string]()
	n := bag.Add("cat", 1)
	if n != 1 {
		t.Fatalf(`Add operation returns wrong result value %d`, n)
	}
	n = bag.Add("cow", 2)
	if n != 2 {
		t.Fatalf(`Add operation returns wrong result value %d`, n)
	}
	items := bag.Items()

	cats := 0
	cows := 0
	for _, v := range items {
		if v == "cat" {
			cats++
		}
		if v == "cow" {
			cows++
		}
	}
	if cats != 1 {
		t.Fatalf(`Items operation returns wrong number of cats %d`, n)
	}
	if cows != 2 {
		t.Fatalf(`Items operation returns wrong number of cows %d`, n)
	}
	if bag.Contains("cat") != 1 {
		t.Fatalf(`Contains operation returns wrong number of cats %d`, n)
	}
	if bag.Contains("cow") != 2 {
		t.Fatalf(`Contains operation returns wrong number of cows %d`, n)
	}
	if bag.Contains("martian") != 0 {
		t.Fatalf(`Contains operation returns wrong number of martians %d`, n)
	}
}

func TestRemoveBag(t *testing.T) {
	bag := NewBag[string]()
	n := bag.Add("cat", 3)
	if n != 3 {
		t.Fatalf(`Add operation returns wrong result value %d, expected: %d`, n, 3)
	}
	n = bag.Remove("cat", 1)
	if n != 2 {
		t.Fatalf(`Remove 1 operation returns wrong result value %d, expected: %d`, n, 2)
	}
	n = bag.Remove("cat", 2)
	if n != 0 {
		t.Fatalf(`Remove 2 operation returns wrong result value %d, expected: %d`, n, 0)
	}
	n = bag.Remove("cat", 3)
	if n != 0 {
		t.Fatalf(`Remove 3 operation returns wrong result value %d, expected: %d`, n, 0)
	}
}

func TestBagIntersection(t *testing.T) {
	bag1 := NewBag[string]()
	bag1.AddAll("cat", "cat", "cow", "cow", "cow")
	bag2 := NewBag[string]()
	bag2.AddAll("cat", "martian", "cow")

	intersection := bag1.Intersection(bag2)
	n := intersection.Contains("cat")
	if n != 1 {
		t.Fatalf(`Intersection should have %d cat but has %d`, 1, n)
	}
	n = intersection.Contains("cow")
	if n != 1 {
		t.Fatalf(`Intersection should have %d cow but has %d`, 1, n)
	}
	n = intersection.Contains("martian")
	if n != 0 {
		t.Fatalf(`Intersection should have %d martians but has %d`, 1, n)
	}

}

func TestBagUnion(t *testing.T) {
	bag1 := NewBag[string]()
	bag1.AddAll("cat", "cat", "cow", "cow", "cow", "chicken")
	bag2 := NewBag[string]()
	bag2.AddAll("cat", "martian", "cow", "pig", "pig")

	union := bag1.Union(bag2)
	n := union.Contains("cat")
	if n != 3 {
		t.Fatalf(`Union should have %d cats but has %d`, 3, n)
	}
	n = union.Contains("cow")
	if n != 4 {
		t.Fatalf(`Union should have %d cows but has %d`, 4, n)
	}
	n = union.Contains("martian")
	if n != 1 {
		t.Fatalf(`Union should have %d martians but has %d`, 1, n)
	}
	n = union.Contains("pig")
	if n != 2 {
		t.Fatalf(`Union should have %d pigs but has %d`, 2, n)
	}
	n = union.Contains("chicken")
	if n != 1 {
		t.Fatalf(`Union should have %d chickens but has %d`, 1, n)
	}
}

func TestBagDifference(t *testing.T) {
	bag1 := NewBag[string]()
	bag1.AddAll("cat", "cat", "cow", "cow", "cow", "chicken")
	bag2 := NewBag[string]()
	bag2.AddAll("cat", "martian", "cow", "pig", "pig")
	resultType := "Difference"
	result := bag1.Difference(bag2)
	n := result.Contains("cat")
	if n != 1 {
		t.Fatalf(`%s should have %d cats but has %d`, resultType, 1, n)
	}
	n = result.Contains("cow")
	if n != 2 {
		t.Fatalf(`%s should have %d cows but has %d`, resultType, 2, n)
	}
	n = result.Contains("martian")
	if n != 0 {
		t.Fatalf(`%s should have %d martians but has %d`, resultType, 0, n)
	}
	n = result.Contains("pig")
	if n != 0 {
		t.Fatalf(`%s should have %d pigs but has %d`, resultType, 0, n)
	}
	n = result.Contains("chicken")
	if n != 1 {
		t.Fatalf(`%s should have %d chickens but has %d`, resultType, 1, n)
	}

	result = bag2.Difference(bag1)
	n = result.Contains("cat")
	if n != 0 {
		t.Fatalf(`%s should have %d cats but has %d`, resultType, 0, n)
	}
	n = result.Contains("cow")
	if n != 0 {
		t.Fatalf(`%s should have %d cows but has %d`, resultType, 0, n)
	}
	n = result.Contains("martian")
	if n != 1 {
		t.Fatalf(`%s should have %d martians but has %d`, resultType, 1, n)
	}
	n = result.Contains("pig")
	if n != 2 {
		t.Fatalf(`%s should have %d pigs but has %d`, resultType, 2, n)
	}
	n = result.Contains("chicken")
	if n != 0 {
		t.Fatalf(`%s should have %d chickens but has %d`, resultType, 0, n)
	}
}

func TestBagCompliment(t *testing.T) {
	bag1 := BagOf("cat", "cat", "dog", "pig")
	bag2 := BagOf("cat", "dog", "elephant")

	comp1 := bag1.Compliment(bag2)
	comp2 := bag2.Compliment(bag1)

	expected := BagOf("cat", "pig", "elephant")

	if !comp1.Equals(expected) {
		t.Fatalf(`compliment %s does not match the expected set %s`, comp1, expected)
	}
	if !comp2.Equals(expected) {
		t.Fatalf(`compliment %s does not match the expected set %s`, comp2, expected)
	}
	if comp1.Count() != expected.Count() {
		t.Fatalf(`compliment count %d is not the expected value %d`, comp1.Count(), expected.Count())
	}
	_ = comp1.String()
	t.Logf("Compliments %s, %s", comp1, comp2)
}
