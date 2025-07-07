/*
initial:
 nums1 <- sort nums1
 nums2_sorted <- sort nums2 into copy
add:
 find value in nums2
 apply add/store from value and to value
 binary search in copy update find from value and update to to value
count:
 all values are >=1
 nums1_slice <- slice nums1 from less than tot
 nums2_slice <- slice nums2 from less than tot
 iterate nums2_slice:
   binary search nums1_slice find less than tot-nums2
   tot += len of items less than
 */

 /*
This version:
172 ms beats 6.25% remove printlns -> 166ms beats 6.25%
20.64 Mb Beats 6.25% removing printlns -> 18.5Mb beats 37.5%
 */

type FindSumPairs struct {
    Nums1_values map[int]int
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
    for _, num := range(nums2) {
        if nums2_val, ok := nums2_values[num]; ok {
            nums2_values[num] = nums2_val + 1
        } else {
            nums2_values[num] = 1
        }
    }
    // fmt.Printf("nums1_values: %v nums2_values: %v\n", nums1_values, nums2_values)
    // fmt.Printf("nums1_values len: %d num1 len: %d nums2_values len: %d nums2 len: %d\n",len(nums1_values), len(nums1), len(nums2_values), len(nums2))
    return FindSumPairs{
        Nums1_values: nums1_values,
        Nums2: nums2,
        Nums2_values: nums2_values, 
    }
}

func (this *FindSumPairs) Add(index int, val int)  {
    // fmt.Printf("index: %d  val: %d\n", index, val)
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
    // fmt.Printf("tot: %d\n", tot)
    // get all keys from nums2
    // memory vs cpu tradeoff --> keep a list of these in struct
    var (
        iteration_dict, compare_dict map[int]int
    )
    if len(this.Nums2_values) >= len(this.Nums1_values) {
        // fmt.Println("Nums1 -> iteration")
        iteration_dict = this.Nums1_values
        compare_dict = this.Nums2_values
    } else {
        // fmt.Println("Nums2 -> iteration")
        iteration_dict = this.Nums2_values
        compare_dict = this.Nums1_values
    }
    
    iter_keys := make([]int, len(iteration_dict))
    i := 0
    for k := range iteration_dict {
        iter_keys[i] = k
        i++
    }
    slices.Sort(iter_keys)
    // fmt.Println(iter_keys)
    // this might not be worth the effort
    if iter_keys[len(iter_keys)-1] > tot {
        needle_idx := BinarySearch(iter_keys, tot)
        if needle_idx != -1 {
            iter_keys = iter_keys[:needle_idx]
        }
    }
    totals := 0
    for _, nums_val := range(iter_keys) {
        if nums_val > tot {
            break
        }
        comp_value := tot-nums_val
        if cnt, ok := compare_dict[comp_value]; ok {
            // fmt.Printf("iteration_val count: %d nums_val: %d\n", iteration_dict[nums_val], nums_val)
            // fmt.Printf("compare val: %d count: %d \n", comp_value, cnt)
            totals = totals + (cnt*iteration_dict[nums_val])
        }
    }
    // fmt.Printf("totals: %d\n", totals)
    return totals
}

func BinarySearch(sorted_slice []int, needle int) int {
    // fmt.Println("Binary Searching...")
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
        // fmt.Printf("at idx: %d idx: %d top: %d bottom: %d\n", sorted_slice[idx], idx, top, bottom)
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