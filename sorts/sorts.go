package sorts

import "math"

// InsertionSort 挿入ソート
func InsertionSort(a []int) {
	var (
		i, j, key int
	)

	for j = 1; j < len(a); j++ {
		key = a[j]
		i = j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	return
}

// MergeSort マージソート
func MergeSort(a []int) {
	mergesort(a, 0, len(a)-1)
	return
}

func mergesort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2     // 分割
		mergesort(a, p, q)   // 前半部ソート
		mergesort(a, q+1, r) // 後半部ソート
		merge(a, p, q, r)    // 統合
	}
}

func merge(a []int, p, q, r int) {
	n1 := q - p + 1                         // p~qまでの範囲分の配列
	n2 := r - q                             // q+1~rまでの配列
	larr := make([]int, n1)                 // 前半部分配列
	rarr := make([]int, n2)                 //後半部分配列を代入するための変数
	copy(larr, a[p:q+1])                    // 元配列の前半部分の入れる
	larr = append(larr, int(math.MaxInt64)) // 門番（とても大きな数字）を入れ実質終端とする
	copy(rarr, a[q+1:r+1])                  // 元配列の後半部分
	rarr = append(rarr, int(math.MaxInt64)) // 門番を代入
	i := 0                                  // larrのインデックス値
	j := 0                                  // rarrのインデックス値
	for k := p; r >= k; k++ {               // p~rの範囲内で部分配列を比較しながら統合
		// 昇順になっている部分配列同士を随時比較しながら統合する処理（元配列に挿入された要素を持つ部分配列はインデックス値を進める）
		if larr[i] < rarr[j] { // 小さな要素を優先的に代入（昇順ソート）
			a[k] = larr[i]
			i++
		} else {
			a[k] = rarr[j]
			j++
		}
	}
}

// HeapSort ヒープソート
func HeapSort(a []int) {
	return
}

// QuickSort クイックソート
func QuickSort(a []int) {
	return
}
