
/*
https://leetcode.com/problems/finding-pairs-with-a-certain-sum/submissions/1689011496
55ms -> beats 100%
18.28Mb -> beats 43.75%
*/

type FindSumPairs struct {
    Nums1_values map[int]int
    Nums1_uniq_sorted []int
    Nums2 []int
    Nums2_values map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    nums1_values := make(map[int]int)
    nums2_values := make(map[int]int)
    for _, num := range(nums1) {
        if nums1_val, ok := nums1_values[num]; ok {
            nums1_values[num] = nums1_val + 1
        } else {
            nums1_values[num] = 1
        }
    }
    iter_keys := make([]int, len(nums1_values))
    i := 0
    for k := range nums1_values {
        iter_keys[i] = k
        i++
    }
    slices.Sort(iter_keys)

    for _, num := range(nums2) {
        if nums2_val, ok := nums2_values[num]; ok {
            nums2_values[num] = nums2_val + 1
        } else {
            nums2_values[num] = 1
        }
    }
    return FindSumPairs{
        Nums1_values: nums1_values,
        Nums1_uniq_sorted: iter_keys,
        Nums2: nums2,
        Nums2_values: nums2_values, 
    }
}

func (this *FindSumPairs) Add(index int, val int)  {
    from_val := this.Nums2[index]
    to_val := from_val + val
    this.Nums2[index] = to_val
    this.Nums2_values[from_val] = this.Nums2_values[from_val] - 1
    if to_nums2_val, ok := this.Nums2_values[to_val]; ok {
        this.Nums2_values[to_val] = to_nums2_val + 1
    } else {
        this.Nums2_values[to_val] = 1
    }
}


func (this *FindSumPairs) Count(tot int) int {
    // this might not be worth the effort
    var iter_keys []int
    if this.Nums1_uniq_sorted[len(this.Nums1_uniq_sorted)-1] > tot {
        needle_idx := BinarySearch(this.Nums1_uniq_sorted, tot)
        if needle_idx != -1 {
            iter_keys = this.Nums1_uniq_sorted[:needle_idx]
        }
    } else {
        iter_keys = this.Nums1_uniq_sorted
    }

    totals := 0
    for _, nums_val := range(iter_keys) {
        if nums_val > tot {
            break
        }
        comp_value := tot-nums_val
        if cnt, ok := this.Nums2_values[comp_value]; ok {
            totals = totals + (cnt*this.Nums1_values[nums_val])
        }
    }
    return totals
}

func BinarySearch(sorted_slice []int, needle int) int {
    bottom, idx := 0, 0
    top := len(sorted_slice)-1

    //needle not within range of slice
    if (needle < sorted_slice[0]) || (needle > sorted_slice[top]) {
        return -1
    }
    
    for {
        // not found return tipping point
        if top <= bottom || (top-bottom) <= 1 {
            return idx+1
        }
        idx = int(math.Ceil(float64(top + bottom)/2.0))
        if sorted_slice[idx] == needle {
            return idx
        } else if sorted_slice[idx] > needle {
            top = idx
        } else if sorted_slice[idx] < needle {
            bottom = idx
        }
    }
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */