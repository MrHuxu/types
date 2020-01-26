package types

import (
	"strconv"
	"strings"
)

// TreeNodeWithParentPointer defines the struct of a node in a numeric binary tree
type TreeNodeWithParentPointer struct {
	Val int

	Parent, Left, Right *TreeNodeWithParentPointer
}

// BuildTreeWithParentPointer builds a binary tree by the given numeric values
func BuildTreeWithParentPointer(vals []interface{}) *TreeNodeWithParentPointer {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNodeWithParentPointer{Val: vals[0].(int)}
	vals = vals[1:len(vals)]

	level := []*TreeNodeWithParentPointer{root}
	for len(vals) != 0 {
		var nextLevel []*TreeNodeWithParentPointer

		for _, node := range level {
			if node != nil && len(vals) != 0 {
				if vals[0] == nil {
					nextLevel = append(nextLevel, nil)
				} else {
					node.Left = &TreeNodeWithParentPointer{Val: vals[0].(int)}
					node.Left.Parent = node
					nextLevel = append(nextLevel, node.Left)
				}
				vals = vals[1:len(vals)]

				if len(vals) == 0 {
					break
				} else {
					if vals[0] == nil {
						nextLevel = append(nextLevel, nil)
					} else {
						node.Right = &TreeNodeWithParentPointer{Val: vals[0].(int)}
						node.Right.Parent = node
						nextLevel = append(nextLevel, node.Right)
					}
					vals = vals[1:len(vals)]
				}
			}
		}

		level = nextLevel
	}

	return root
}

// String formats the print content of a node
// the data will be outputed like a REAL tree in console
//               1
//           /       \
//       2               3
//     /   \           /   \
//   4       5       6       7
//          / \
//         8   9
func (node *TreeNodeWithParentPointer) String() string {
	if node == nil {
		return "[nil]"
	}

	var levels [][]string

	level := []*TreeNodeWithParentPointer{node}
	for {
		var strs []string
		var nextLevel []*TreeNodeWithParentPointer

		for _, node := range level {
			if node == nil {
				strs = append(strs, " ")
				nextLevel = append(nextLevel, nil, nil)
			} else {
				strs = append(strs, strconv.Itoa(node.Val))
				nextLevel = append(nextLevel, node.Left, node.Right)
			}
		}

		levels = append(levels, strs)

		var hasNode bool
		for _, node := range nextLevel {
			if node != nil {
				hasNode = true
			}
		}
		if hasNode {
			level = nextLevel
		} else {
			break
		}
	}

	var output []string
	var downLine string
	var pos []int
	for i, str := range levels[len(levels)-1] {
		pos = append(pos, i*4)

		if i == 0 {
			downLine = downLine + str
		} else {
			downLine = downLine + "   " + str
		}
	}
	output = append(output, downLine)
	lineLen := len(downLine)

	for i := len(levels) - 2; i >= 0; i-- {
		var upPos []int
		var upLine, wireLine []string

		var j int
		for k := 0; k < lineLen; k++ {
			upLine = append(upLine, " ")
			wireLine = append(wireLine, " ")

			if j == len(pos) || k != pos[j] {
				continue
			} else if j%2 == 0 {
				j++
				continue
			}

			parentPos := (pos[j-1] + pos[j]) / 2
			leftPos := pos[j-1]
			rightPos := pos[j]
			upLine[parentPos] = levels[i][(j-1)/2]
			if downLine[leftPos] != ' ' {
				wireLine[(parentPos+leftPos)/2] = "/"
			}
			if downLine[rightPos] != ' ' {
				wireLine[(parentPos+rightPos)/2] = "\\"
			}

			upPos = append(upPos, parentPos)

			j++
		}

		output = append([]string{strings.Join(upLine, ""), strings.Join(wireLine, "")}, output...)
		pos = upPos
		downLine = strings.Join(upLine, "")
	}

	return strings.Join(output, "\n")
}
