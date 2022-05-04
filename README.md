# golang_median_of_two_sorted_array

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

## Examples

**Example 1:**

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

```

**Example 2:**

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

```

**Constraints:**

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- $`-10^6$ <= nums1[i], nums2[i] <= $10^6$`

## 解析

題目給定兩個排序過的整數陣列 num1, num2 要找出這兩個陣列組成新的排序陣列的中間數

首先要知道中間數的定義

假設把 num1, num2 照大小順序合成一個新的陣列 merged

merged 的中間數就是 

((merged[Math.floor((merge.length-1)/2)])+ (merged[Math.floor((merge.length)/2)]))/2

其中一個作法是照大小依序找出 前 (m+n)/2+1 小的數字

然後把 第(m+n-1)/2 , (m+n)/2 小的數相加除以 2 就是結果

這樣的作法就是 O(m+n)

所以如果需要降低複雜度需要找出其他關係式來使用 binary search

假設 num1, num2 如下圖一樣

![two_sorted_array.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9744c443-74ae-49c2-9586-ee756144e44d/two_sorted_array.png)

可以看出對於兩個陣列 large, small 的中間數。

假設最後找到的中間數，在 small 找到的是 small_idx ，在 large 找到的是 large_idx

會有以下關係 small_idx + large_idx = mid - 2

所以只找從 small 去找 small_idx , large_idx 基本上就是固定了

並且可以知道如果 small_idx 是中間數，則 Large[large_idx] ≤ small[small_idx+1] 且 small[small_idx] ≤ Large[large_idx+1]

利用上面這個條件，就可以對 small 的 array 去做 binary search 

當發現 small[small_idx] > large[large_idx+1] 則把 右界左移，否則左界右移

## 程式碼

```go
import "math"
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
   if len(num1) > len(num2) {
      return FindMedian(&num2, &num1)
   } else {
      return FindMedian(&num1, &num2)
   }
}

func FindMedian(small *[]int, large *[]int) float64 {
   sLen := len(*small)
   bLen := len(*large)
   total := sLen + bLen
   half := total/2
   if bLen == 0 {
      return 0
   }
   if sLen == 0 {
      return float64(((*large)[(bLen-1)/2]+ (*large)[bLen/2])/2)
   }
   left := 0
   right := sLen - 1
   small_idx := 0
   large_idx := 0
   small_left := math.MinInt32
   small_right := math.MaxInt32
   large_left := math.MinInt32
   large_right := math.MaxInt32
   // check left <= right+1
   for left <= right + 1 {
      if left + right >= 0 { // left 還沒移到左界之外
        small_idx = (left+ right)/2
      } else {
         small_idx = -1
      }
      large_idx = half - small_idx - 2
      
      if small_idx >= 0 {
        small_left = (*small)[small_idx]
      } else {
        small_left = math.MinInt32 
      }
      if small_idx + 1 < sLen {
        small_right = (*small)[small_idx+1]
      } else {
        small_right = math.MaxInt32
      }
      if large_idx >= 0 {
        large_left = (*large)[large_idx]
      } else {
        large_left = math.MinInt32 
      }
      if large_idx + 1 < bLen {
        large_right = (*large)[large_idx+1] 
      } else {
        large_right = math.MaxInt32
      }
      if small_left <= large_right && small_right >= large_left {
        if total % 2 == 1 {
           return float64(Min(large_right, small_right)) 
        } else {
           return float64(Min(large_right, small_right)+ Max(large_left, small_left)) /2
        }
      }
      if small_left > large_right {
         right = small_idx - 1
      } else {
         left = small_idx + 1
      }
   }
   return 0
}

func Max(a, b int)int {
  if a > b {
     return a
  }
  return b
}

func Min(a, b int) int{
  if a < b {
    return a
  }
  return b
}
```

## 困難點

1. 需要知道中間數的定義
2. 需要知道兩個陣列中間數的關係
3. 需要知道邊界處理方式

## Solve Point

- [x]  Understand what the problem would like to solve
- [x]  Analysis Complexity