package sort

// O(n)logn
func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	QuickSort(nums[:left])
	QuickSort(nums[left+1:])
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j <= len(nums)-1-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

type Student struct {
	Age int
}

type Students []*Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func HeapSort(nums []int) {
	// implement heap sort
	// 1. build heap
	// 2. pop root, swap with last element, heapify
	// 3. repeat step 2
	Heaptify(nums)
	for last := len(nums) - 1; last > 0; last-- {
		nums[0], nums[last] = nums[last], nums[0]
		down(nums, 0, last-1)
	}
}

func Heaptify(nums []int) {
	n := len(nums) - 1
	for i := (n - 1) / 2; i >= 0; i-- {
		down(nums, i, n)
	}
}

func down(nums []int, i, n int) {
	for {
		left := 2*i + 1
		if left > n || left < 0 {
			break
		}
		min := left
		right := left + 1
		if right <= n && right > 0 && nums[right] < nums[min] {
			min = right
		}
		if nums[i] <= nums[min] {
			break
		}
		nums[i], nums[min] = nums[min], nums[i]
		i = min
	}
}
