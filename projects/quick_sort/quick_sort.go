package quick_sort

func partition(arr *[]int, low, high int) int {
	pivot := (*arr)[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]
	return i + 1
}

func quickSort(arr *[]int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func concurrentQuickSortV1(arr *[]int, low, high int, chanSend chan int) {
	if low < high {
		pivot := partition(arr, low, high)
		chanReceive := make(chan int)
		go concurrentQuickSortV2(arr, low, pivot-1, chanReceive)
		go concurrentQuickSortV2(arr, pivot+1, high, chanReceive)
		<-chanReceive
		<-chanReceive
		chanSend <- 0
		return
	} else {
		chanSend <- 0
		return
	}
}

func concurrentQuickSortV2(arr *[]int, low, high int, chanSend chan int) {
	if (high - low) < 1000 {
		quickSort(arr, low, high)
		chanSend <- 0
		return
	}
	if low < high {
		pivot := partition(arr, low, high)
		chanReceive := make(chan int)
		go concurrentQuickSortV2(arr, low, pivot-1, chanReceive)
		go concurrentQuickSortV2(arr, pivot+1, high, chanReceive)
		<-chanReceive
		<-chanReceive
		chanSend <- 0
		return
	} else {
		chanSend <- 0
		return
	}
}
