// Segment tree with build, update and get function, O(n*4) space
package main

var a []int
var st []int

const infinity = 1 << 9

func build(id, l, r int) {
	if l == r {
		st[id] = a[l]
		return
	}

	mid := (l + r) / 2
	build(id*2, l, mid)
	build(id*2+1, mid+1, r)

	if st[id*2] > st[id*2+1] {
		st[id] = st[id*2]
	} else {
		st[id] = st[id*2+1]
	}
}

func update(id, l, r, i, v int) {
	if i < l || i > r {
		return
	}

	if l == r {
		st[id] = v
		return
	}

	mid := (l + r) / 2
	update(id*2, l, mid, i, v)
	update(id*2+1, mid+1, r, i, v)

	if st[id*2] > st[id*2+1] {
		st[id] = st[id*2]
	} else {
		st[id] = st[id*2+1]
	}
}

func get(id, l, r, u, v int) int {
	if v < l || r < u {
		return -infinity
	}

	if u <= l && v <= r {
		return st[id]
	}

	mid := (l + r) / 2

	max1 := get(id*2, l, mid, u, v)
	max2 := get(id*2+1, mid+1, r, u, v)
	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}

func main() {
}
