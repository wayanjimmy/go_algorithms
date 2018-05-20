package main

func swap(a []int, i, j int) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}

func BubbleSort(a []int) {
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(a)-1; i++ {
			if a[i+1] < a[i] {
				swap(a, i+1, i)
				swapped = true
			}
		}
	}

}
