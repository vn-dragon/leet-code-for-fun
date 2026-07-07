package daily

import "sort"

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	type car struct{ pos, spd int }
	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{position[i], speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	stack := []float64{}
	for _, c := range cars {
		t := float64(target-c.pos) / float64(c.spd)
		if len(stack) == 0 || t > stack[len(stack)-1] {
			stack = append(stack, t)
		}
	}
	return len(stack)
}
