type SparseVector struct {
    repr map[int]int
}

func Constructor(nums []int) SparseVector {
    vector := SparseVector{
        repr: make(map[int]int),
    }
    for i, x := range nums {
        if x != 0 {
            vector.repr[i] = x
        }
    }
    return vector
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    product := 0
    for k, v1 := range this.repr {
        product += v1 * vec.repr[k]
    }
    return product
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */