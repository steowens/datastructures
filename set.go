/*
github.com/steowens/datastructures (c) by Stephen Owens>

github.com/steowens/datastructures is licensed under a
Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.

You should have received a copy of the license (LICENSE) along with this
work. If not, see <http://creativecommons.org/licenses/by-nc-sa/4.0/>.
*/
package collection

import (
	"fmt"
	"sort"
	"strings"
)

// While golang makes it easy to create sets on the fly using maps, the code
// that typically gets written when doing this looks like gobledy gook.
// Much cleaner code comes from a properly written interface, thus this file.
// Sets are essential.
type Set[T Bagable] struct {
	m map[T]bool
}

func NewSet[T Bagable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]bool, 0),
	}
}

func SetOf[T Bagable](vals ...T) *Set[T] {
	result := NewSet[T]()
	for _, val := range vals {
		result.Add(val)
	}
	return result
}

// Adds val to set if not already present.
// Returns true if string was added
// else returns false if string was alreay present.
func (s *Set[T]) Add(val T) bool {
	_, exists := s.m[val]
	if !exists {
		s.m[val] = true
		return true
	}
	return false
}

func (s *Set[T]) AddAll(vals ...T) int {
	count := 0
	for _, val := range vals {
		if s.Add(val) {
			count++
		}
	}
	return count
}

// Removes val from set if it is present in set.
// Returns true if the string was present and is now
// removed, false otherwise.
func (s *Set[T]) Remove(val T) bool {
	_, exists := s.m[val]
	if exists {
		delete(s.m, val)
		return true
	}
	return false
}

// Returns true if this set contains val
func (s *Set[T]) Contains(val T) bool {
	_, exists := s.m[val]
	return exists
}

// Return an array of all items in this set
func (s *Set[T]) Items() []T {
	a := make([]T, 0)
	for k, v := range s.m {
		if v {
			a = append(a, k)
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}

// Returns the set of things found in both this set and set o
func (s *Set[T]) Intersection(o *Set[T]) *Set[T] {
	i := NewSet[T]()
	for k := range s.m {
		if o.Contains(k) {
			i.Add(k)
		}
	}
	return i
}

// Combines this set and o to form a union of the two sets
func (s *Set[T]) Union(o *Set[T]) *Set[T] {
	i := NewSet[T]()
	for k := range s.m {
		i.Add(k)
	}
	for k := range o.m {
		i.Add(k)
	}
	return i
}

// Returns the compliment of the intersection of this set and set 'o',
// where this set is considered to be the universal set.  Basically
// cuts o out of this set where 'o' intersects with this set.
func (s *Set[T]) Compliment(o *Set[T]) *Set[T] {
	i := NewSet[T]()
	for k := range s.m {
		if !o.Contains(k) {
			i.Add(k)
		}
	}
	return i
}

// Return true if both sets contain the same elements
func (s *Set[T]) Equals(o *Set[T]) bool {
	for k := range s.m {
		if !o.Contains(k) {
			return false
		}
	}
	for k := range o.m {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}

func (s *Set[T]) String() string {
	items := s.Items()
	var sb strings.Builder

	prefix := ""
	sb.WriteString("[")
	for _, val := range items {
		sb.WriteString(prefix)
		str := fmt.Sprintf("%v", val)
		sb.WriteString(str)
		prefix = ", "
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *Set[T]) Count() int {
	return len(s.m)
}
