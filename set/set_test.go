package guava

import (
	"fmt"
	"github.com/deckarep/golang-set.git"
	"testing"
)

func TestGoSet(t *testing.T) {

	requiredClasses := mapset.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	scienceSlice := []interface{}{"Biology", "Chemistry", "Math", "Welding"}
	scienceClasses := mapset.NewSetFromSlice(scienceSlice)
	var dd = make([]string, 1)
	dd[0] = "1"
	tt := mapset.NewSet()
	for _, v := range dd {
		tt.Add(v)
	}
	tt.Contains(1)

	electiveClasses := mapset.NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := mapset.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")


	//并集
	allClasses := requiredClasses.Union(scienceClasses)
	fmt.Println(allClasses)
	//Set{Cooking, English, Math, Chemistry, Welding, Biology, Music, Automotive, Go Programming, Python Programming}

	//Is cooking considered a science class?
	fmt.Println(scienceClasses.Contains("Cooking")) //false

	//差集
	fmt.Println(allClasses.Difference(scienceClasses)) //Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	//交集
	fmt.Println(scienceClasses.Intersect(requiredClasses)) //Set{Biology，Math}

	//How many bonus classes do you offer?
	fmt.Println(bonusClasses.Cardinality()) //2

	//Do you have the following classes? Welding, Automotive and English?
	fmt.Println(allClasses.IsSuperset(mapset.NewSetFromSlice([]interface{}{"Welding", "Automotive", "English"}))) //true
}
