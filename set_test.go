/*
github.com/steowens/datastructures (c) by Stephen Owens>

github.com/steowens/datastructures is licensed under a
Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.

You should have received a copy of the license (LICENSE) along with this
work. If not, see <http://creativecommons.org/licenses/by-nc-sa/4.0/>.
*/
package collection

import (
	"strings"
	"testing"
)

func TestAddSet(t *testing.T) {
	set := NewSet[string]()
	set.AddAll("cat", "dog")
	if set.Count() != 2 {
		t.Fatalf(`Set returns wrong count value %d, expected: 2`, set.Count())
	}

	set.Add("dog")
	if set.Count() != 2 {
		t.Fatalf(`Set returns wrong count value %d, expected: 2`, set.Count())
	}

	set.Add("mouse")
	if set.Count() != 3 {
		t.Fatalf(`Set returns wrong count value %d, expected: 3`, set.Count())
	}

	if !set.Contains("mouse") {
		t.Fatal(`Set.Contains('mouse') returns false`)
	}
	if set.Contains("fork") {
		t.Fatal(`Set.Contains('fork') returns true `)
	}
}

func TestSetRemove(t *testing.T) {
	set := SetOf("cat", "mouse", "dog", "sheep", "lion")
	if !set.Contains("sheep") {
		t.Fatal(`Set.Contains('sheep') returns false`)
	}
	set.Remove("sheep")
	if set.Contains("sheep") {
		t.Fatal(`Set.Contains('sheep') returns true after removal`)
	}
	if set.Count() != 4 {
		t.Fatalf(`Expected set.Count() %d, actual count %d`, 4, set.Count())
	}
	items := set.Items()
	if len(items) != 4 {
		t.Fatalf(`Expected len(set.Items()) %d, actual length %d`, 4, len(items))
	}
	expected := [4]string{"cat", "dog", "lion", "mouse"}
	compareSlices(items, expected[:], t)
}

func TestSetIntersection(t *testing.T) {
	pets := SetOf("cat", "parakeet", "sheepdog", "goldfish")
	farmanimals := SetOf("cow", "chicken", "pig", "sheepdog")

	intersects := pets.Intersection(farmanimals)
	expected := [1]string{"sheepdog"}
	items := intersects.Items()
	compareSlices(items, expected[:], t)
}

func TestSetUnion(t *testing.T) {
	pets := SetOf("cat", "parakeet", "sheepdog", "goldfish")
	farmanimals := SetOf("cow", "chicken", "pig", "sheepdog")

	union := pets.Union(farmanimals)
	expected := [7]string{"cat", "chicken", "cow", "goldfish", "parakeet", "pig", "sheepdog"}
	items := union.Items()
	compareSlices(items, expected[:], t)
}

func TestSetCompliment(t *testing.T) {
	pets := SetOf("cat", "parakeet", "sheepdog", "goldfish")
	farmanimals := SetOf("cow", "chicken", "pig", "sheepdog")

	compliment := pets.Compliment(farmanimals)
	expected := [3]string{"cat", "goldfish", "parakeet"}
	items := compliment.Items()
	compareSlices(items, expected[:], t)

	compliment = farmanimals.Compliment(pets)
	expected = [3]string{"chicken", "cow", "pig"}
	items = compliment.Items()
	compareSlices(items, expected[:], t)
}

func compareSlices(items []string, expected []string, t *testing.T) {
	if len(items) != len(expected) {
		t.Fatalf(`Length of items %d is not expected value of %d`, len(items), len(expected))
	}
	for x := 0; x < len(items); x++ {
		if items[x] != expected[x] {
			t.Fatalf(`Expected items [%s], actual items [%s]`, strings.Join(expected, ", "), strings.Join(items, ", "))
		}
	}
}
