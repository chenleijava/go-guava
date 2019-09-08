package counters

import mapset "github.com/deckarep/golang-set"

//base counter set
type CounterSet struct {
	set  mapset.Set
	name string
}

//new counter set
func NewCounterSet(name string) *CounterSet {
	return &CounterSet{
		set:  mapset.NewSet(),
		name: name,
	}
}

// Adds an element to the set. Returns whether
// the item was added.
func (c *CounterSet) Add(i interface{}) bool {
	return c.set.Add(i)
}

// Returns the number of elements in the set.
func (c *CounterSet) Len() int {
	return c.set.Cardinality()
}

// Returns whether the given items
// are all in the set.
func (c *CounterSet) Contains(i ...interface{}) bool {
	return c.set.Contains(i...)
}

// Returns a new set containing only the elements
// that exist only in both sets.
func (c *CounterSet) Intersect(other *CounterSet) mapset.Set {
	return c.set.Intersect(other.set)
}

// Remove a single element from the set.
func (c *CounterSet) Remove(i interface{}) {
	c.set.Remove(i)
}
