package main

func findMaxSlidingWindow(nums []int, windowSize int) []int {
  result := make([]int, 0)

  if len(nums) == 0 {
    return result
  }

  if len(nums) < windowSize {
    windowSize = len(nums)
  }

  window := NewDeque()

  for i := 0; i < windowSize; i++ {
    for !window.Empty() && nums[i] >= nums[window.Back()] {
      window.PopBack()
    }

    window.PushBack(i)
  }

  result = append(result, nums[window.Front()])

  for i := windowSize; i < len(nums); i++ {
    for !window.Empty() && nums[i] >= nums[window.Back()] {
      window.PopBack()
    }

    if !window.Empty() && window.Front() <= (i - windowSize) {
      window.PopFront()
    }

    window.PushBack(i)

    result = append(result, nums[window.Front()])
  } 

  return result
}
