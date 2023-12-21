package main

import "fmt"

func main() {
	number := []int{5, 6, 2, 3, 4, 1}
	insertSort(number)
	fmt.Println(number)
}

/*
									{5, 6, 2, 3, 4, 1}
i = 1    j=0   value = arr[1] = 6
i = 2    j=1  value = arr[2] = 2    {2, 5, 6, 3, 4, 1}
i = 3    j=2   value = arr[3] = 3     {2, 3, 5, 6, 4, 1}



// 插入排序，a表示数组，n表示数组大小
public void insertionSort(int[] a, int n) {
  if (n <= 1) return;

  for (int i = 1; i < n; ++i) {
    int value = a[i];
    int j = i - 1;
    // 查找插入的位置
    for (; j >= 0; --j) {
      if (a[j] > value) {
        a[j+1] = a[j];  // 数据移动
      } else {
        break;
      }
    }
    a[j+1] = value; // 插入数据
  }
}


*/

func insertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
				fmt.Println(arr[j+1],arr[j])
			} else {
				break
			}
		}
		arr[j+1] = value
		fmt.Println(arr)
	}

}

//// {5, 9, 2, 3, 4, 1}  value = arr[2] = 2;    j = 1 ; arr[1]= 6 ;
//i = 1   value = 9  j = 0
//i = 2   value = 2  j = 1    {5, 9, 9, 3, 4, 1}    {5, 5, 9, 3, 4, 1}    {2, 5, 9, 3, 4, 1}
//i = 3   value = 3  j = 2



