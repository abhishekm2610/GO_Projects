package main

import "fmt"

type SnapshotArray struct {
	arrays map[int][][2]int
	curr   int
}

func Constructor(length int) SnapshotArray {
	obj := SnapshotArray{}
	obj.arrays = make(map[int][][2]int)

	// obj.arrays[0] = make([][2]int, length)

	obj.curr = 0
	return obj

}

func (this *SnapshotArray) Set(index int, val int) {
	// newArr := []int{}
	// newArr = append(newArr, this.arrays[this.curr]...)
	// newArr[index] = val
	// this.curr++
	if _, p := this.arrays[this.curr]; !p {
		// var arr [2]int
		this.arrays[this.curr] = make([][2]int, 0)
	}

	for i, v := range this.arrays[this.curr] {
		if v[0] == index {
			this.arrays[this.curr][i][1] = val
			return
		}
	}
	this.arrays[this.curr] = append(this.arrays[this.curr], [2]int{index, val})

}

func (this *SnapshotArray) Snap() int {
	// newArr := this.arrays[this.curr]
	// this.arrays[this.curr+1] = make([][2]int, 0)
	// this.arrays[this.curr+1] = append(this.arrays[this.curr+1], this.arrays[this.curr]...)

	this.curr++
	// this.arrays[this.curr] = newArr
	return this.curr - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// fmt.Println(this.arrays, snap_id, index)

	ans := 0
	for i := snap_id; i >= 0; i-- {
		allChanges := this.arrays[i]

		for _, v := range allChanges {
			if v[0] == index {
				ans = v[1]
				return ans
			}
		}
	}

	return ans
}

func main() {
	obj := Constructor(4)
	// var p int
	// obj.Set(0, 49352)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 0))
	// obj.Set(0, 5373)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 0))
	// obj.Set(0, 23341)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 0))
	// obj.Set(0, 12771)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 2))
	obj.Set(1, 5)
	fmt.Println(obj.Snap())
	obj.Set(0, 16)
	fmt.Println(obj.Snap())
	obj.Set(2, 15)
	fmt.Println(obj.Snap())
	obj.Set(2, 5)
	fmt.Println(obj.Get(1, 0))
	fmt.Println(obj.Get(0, 2))
	fmt.Println(obj.Snap())

	// fmt.Println(obj.Get(0, 0))
	// obj.Set(0, 5373)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 0))
	// obj.Set(0, 23341)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 0))
	// obj.Set(0, 12771)
	// fmt.Println(obj.Snap())
	// fmt.Println(obj.Get(0, 2))

	fmt.Println(obj.arrays)
}

// ["SnapshotArray","set","snap","set","snap","set","snap","set","get","get","snap"]
// [[4],[1,5],[],[0,16],[],[2,15],[],[2,5],[1,0],[0,2],[]]

// ,"set",		"snap",		"get",		"set",		"snap",		"get",		"set",		"snap",	"get",	"set",		"snap",		"get"
// [0,49352],	[],			[0,0],		[0,5373],	[],			[0,0],		[0,23341],	[],		[0,0],	[0,12771]	,[],		[0,2]
// null,		0,			49352,		null,		1,			49352,		null,		2,		49352,	null,		3,			49352,			null,4,49352,null,5,49352,null,6,49352,null,7,49352,
