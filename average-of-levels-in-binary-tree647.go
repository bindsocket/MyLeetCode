package leetcode

import "slices"

/**
https://leetcode.com/problems/average-of-levels-in-binary-tree/

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/

/*
push root
lvl 0 -> root
push root.left
push root.right
lvl 1 -> avg(root.left + root.right)
push root.left.left, root.left.right, root.right.left, root.right.right

push root -> queue
while true:
    lvl = 0
    next_queue = []
    for queue_item in queue:
        next_queue = push queue_item.left, queue_item.right
        sum += queue_item.left + queue_item.right
    queue = next_queue
    lvl n -> avg(sum)
    lvl++
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	var lvlAvgs []float64
	for {
		nextQueue := slices.Clip([]*TreeNode{})
		runSum := 0
		valCount := 0
		for _, queueItem := range queue {
			if queueItem.Left != nil {
				nextQueue = append(nextQueue, queueItem.Left)
			}
			if queueItem.Right != nil {
				nextQueue = append(nextQueue, queueItem.Right)
			}
			runSum += queueItem.Val
			valCount++
		}
		queue = nextQueue
		lvlAvgs = append(lvlAvgs, float64(runSum)/float64(valCount))
		if len(nextQueue) == 0 {
			return lvlAvgs
		}
	}
}
