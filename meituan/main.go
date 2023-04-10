package main

import (
	"container/heap"
	"fmt"
)

type cup struct {
	id, vol, water, cost int
}

type cups []cup

func (cs cups) Len() int {
	return len(cs)
}

func (cs cups) Less(i, j int) bool {
	return cs[i].cost < cs[j].cost
}

func (cs cups) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs *cups) Push(x interface{}) {
	*cs = append(*cs, x.(cup))
}

func (cs *cups) Pop() interface{} {
	old := *cs
	n := len(old)
	x := old[n-1]
	*cs = old[:n-1]
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	cs := make(cups, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cs[i].vol)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&cs[i].water)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&cs[i].cost)
	}

	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var q int
		fmt.Scan(&q)

		pq := make(cups, 0)
		used := make(map[int]bool)

		heap.Push(&pq, cs[q-1])
		used[q-1] = true

		cost := 0
		for len(pq) > 0 {
			cur := heap.Pop(&pq).(cup)
			if cur.water == cur.vol {
				continue
			}
			if cur.id > 0 {
				left := &cs[cur.id-1]
				if !used[cur.id-1] {
					heap.Push(&pq, *left)
					used[cur.id-1] = true
				}
				diff := min(left.vol-left.water, cur.water)
				left.water += diff
				cur.water -= diff
				cost += diff * cur.cost
			}
			if cur.id < n-1 {
				right := &cs[cur.id+1]
				if !used[cur.id+1] {
					heap.Push(&pq, *right)
					used[cur.id+1] = true
				}
				diff := min(right.vol-right.water, cur.water)
				right.water += diff
				cur.water -= diff
				cost += diff * cur.cost
			}
		}
		fmt.Printf("%d ", cost)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
