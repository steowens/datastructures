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

	"golang.org/x/exp/constraints"
)

type Bagable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// While golang makes it easy to create Bags on the fly using maps, the code
// that typically gets written when doing this looks like gobledy gook.
// Much cleaner code comes from a properly written interface, thus this file.
// Bags are essential.
type Bag[T Bagable] struct {
	m map[T]int
}

func NewBag[T Bagable]() *Bag[T] {
	return &Bag[T]{
		m: make(map[T]int, 0),
	}
}

func BagOf[T Bagable](vals ...T) *Bag[T] {
	bag := NewBag[T]()
	for _, v := range vals {
		bag.Add(v, 1)
	}
	return bag
}

// Adds n copies of val to Bag.  Returns the total
// number of copies of n in the bag after the addition
// is complete.
func (s *Bag[T]) Add(val T, n int) int {
	count, exists := s.m[val]
	if !exists {
		s.m[val] = n
		return n
	} else {
		s.m[val] = count + n
		return s.m[val]
	}
}

func (s *Bag[T]) AddAll(vals ...T) {
	for _, v := range vals {
		s.Add(v, 1)
	}
}

// Removes n copies of val from Bag.  Returns the number of
// copies of n in the bag remaining after the removal.
func (s *Bag[T]) Remove(val T, n int) int {
	count, exists := s.m[val]
	if exists {
		if n < count {
			s.m[val] = count - n
			return s.m[val]
		} else {
			delete(s.m, val)
		}
	}
	return 0
}

// Returns true if this Bag contains val
func (s *Bag[T]) Contains(val T) int {
	count, exists := s.m[val]
	if exists {
		return count
	}
	return 0
}

// Return an array of all items in this Bag
func (s *Bag[T]) Items() []T {
	a := make([]T, 0)
	for k, v := range s.m {
		for i := 0; i < v; i++ {
			if v != 0 {
				a = append(a, k)
			}
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}

// Returns the Bag of things found in both this Bag and Bag o
// If only one bag (this one or the other one) contains an item
// the result will not contain that item.
func (s *Bag[T]) Intersection(o *Bag[T]) *Bag[T] {
	i := NewBag[T]()
	for k, v := range s.m {
		n := o.Contains(k)
		m := min(v, n)
		if m > 0 {
			i.Add(k, m)
		}
	}
	return i
}

// Combines this Bag and o to form a union of the two Bags.
// The resulting bag will contain the sum of all values from both
// this bag and the other bag.
func (s *Bag[T]) Union(o *Bag[T]) *Bag[T] {
	i := NewBag[T]()
	for k, v := range s.m {
		i.Add(k, v)
	}
	for k, v := range o.m {
		i.Add(k, v)
	}
	return i
}

// Returns a bag which adheres to the following rules.
// If this bag contains elements not in the other bag they are in the result bag.
//
// If the other bag contains elements that are in this bag, the result bag will
// contain the greater of 0 or the number of common elements in this bag minus
// the number of common element in the other bag.
//
//	[ cat, cat, dog, pig ] diff [ cat, dog, elephant ] ==> [ cat, pig ]
//	[ cat, dog, elephant ] diff [cat, cat, dog, pig]  ==> [ elephant ]
func (s *Bag[T]) Difference(o *Bag[T]) *Bag[T] {
	i := NewBag[T]()
	for k, v := range s.m {
		n := o.Contains(k)
		if v > n {
			i.Add(k, v-n)
		}
	}
	return i
}

// Returns the union of the differences between this and o and o and this
//
//	[ cat, cat, dog, pig ] comp [ cat, dog, elephant ] ==> [ cat, pig, elephant ]
//	 [ cat, dog, elephant ] comp [ cat, cat, dog, pig ] ==> [ cat, pig, elephant ]
func (s *Bag[T]) Compliment(o *Bag[T]) *Bag[T] {
	d1 := s.Difference(o)
	d2 := o.Difference(s)
	return d1.Union(d2)
}

// Returns true if there is an exact match between all items in
// this bag and the other bag 'o' as well as an exact match between
// all items in 'o' and this bag.
func (s *Bag[T]) Equals(o *Bag[T]) bool {
	for k, v := range s.m {
		n := o.Contains(k)
		if n != v {
			return false
		}
	}
	for k, v := range o.m {
		n := s.Contains(k)
		if n != v {
			return false
		}
	}
	return true
}

func (s *Bag[T]) String() string {
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

func (s *Bag[T]) Count() int {
	return len(s.m)
}
