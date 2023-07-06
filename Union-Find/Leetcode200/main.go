package main

type UnionSet struct {
	root  []int
	count int
}

func Constructor(rows, cols int) UnionSet {
	count := rows * cols
	root := make([]int, count)
	for i := 0; i < count; i++ {
		root[i] = i
	}
	return UnionSet{
		root:  root,
		count: count,
	}
}

//  寻找到该值的祖先
func (this *UnionSet) find(x int) int {
	if x == this.root[x] {
		return x
	} else {
		this.root[x] = this.find(this.root[x])
		return this.root[x]
	}
}

func (this *UnionSet) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX != rootY {
		this.root[rootX] = rootY
		this.count -= 1
	}
}

func (this *UnionSet) getCount() int {
	return this.count
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	water := 0
	rows, cols := len(grid), len(grid[0])
	uf := Constructor(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				water += 1
			} else {
				directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
				for _, d := range directions {
					x := i + d[0]
					y := j + d[1]
					if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == '1' {
						uf.union(x*cols+y, i*cols+j)
					}
				}
			}
		}
	}
	return uf.getCount() - water
}
