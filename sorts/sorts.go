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
	h := buildMaxHeap(a) // heapを構成
	for i := len(a) - 1; i >= 1; i-- {
		h.swap(0, i)    // Maxヒープの根は最大値なので、最後尾と交換しつつ、順にヒープの大きさを縮めると昇順ソートが可能
		h.HeapSize--    // 後ろに現ヒープ内での最大値を持ってきたのでヒープサイズを一つ狭める
		h.maxHeapify(0) // ヒープの根をから順にMaxヒープを再構成
	}
	return
}

type heap struct {
	a        []int // heapの実態
	HeapSize int   // ヒープの大きさ（ヒープ最後尾のインデックス値）
}

// 指定されたノードの親ノードを返す
func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

// 指定された親ノードの左ノードを返す
func (h *heap) left(i int) int {
	return 2*i + 1
}

// 指定された親ノードの右ノードを返す
func (h *heap) right(i int) int {
	return 2*i + 2
}

// 指定されたノードを根とする二分木をmax-heapとなるように構成する
func (h *heap) maxHeapify(i int) {
	var largest int // 左右ノードもしくは親ノードのうち最も大きなノードのインデックス値
	l := h.left(i)
	r := h.right(i)
	if l <= h.HeapSize && h.a[l] > h.a[i] {
		largest = l // 親と左側では左側のほうが要素が大きい
	} else {
		largest = i
	}
	if r <= h.HeapSize && h.a[r] > h.a[largest] {
		largest = r // 親もしくは左側よりも右側のほうが大きい
	}
	if largest != i { // 親ノードよりも左右ノードのほうが大きいときは再帰
		h.swap(i, largest)
		h.maxHeapify(largest)
	}
}

// ヒープ内の要素を交換する
func (h *heap) swap(i, j int) {
	tmp := h.a[i]
	h.a[i] = h.a[j]
	h.a[j] = tmp
}

func buildMaxHeap(a []int) heap {
	h := heap{}
	h.a = a
	h.HeapSize = len(a) - 1                      // ヒープの最後尾を指定（配列の最後尾）
	for i := (h.HeapSize - 1) / 2; i >= 0; i-- { // ヒープ最後尾の親ノードから順にmax-heapifyを掛けていってMaxヒープを構成
		h.maxHeapify(i)
	}
	return h
}

// QuickSort クイックソート
func QuickSort(a []int) {
	quicksort(a, 0, len(a)-1) // 最初は配列の先頭と最後尾を範囲に指定
	return
}

func quicksort(a []int, p, r int) {
	if p < r { // 開始地点よりも終了地点の方が大きいとき（範囲内の大きさが１以上のとき）
		q := partition(a, p, r) // 最後尾の値をピボットにして仕切りとなるインデックス値を求める
		quicksort(a, p, q-1)    // 仕切りより前の範囲でクイックソート
		quicksort(a, q+1, r)    // 仕切りより後の範囲でクイックソート
	}
}

// 最後尾を基準として仕切りの場所を定める
func partition(a []int, p, r int) int {
	var tmp int              // 値入れ替え用バッファ
	x := a[r]                // 最後尾をピボットにする
	i := p - 1               // 仕切りより前の範囲の最後尾を示すインデックス値
	for j := p; r > j; j++ { // jによってピボットとしている最後尾rの前までピボットとの大小を比較する
		if a[j] <= x { // ピボット以下であるとき仕切りとなるインデックス値を一つ増やしてそこに値を差し込む
			i++        // 仕切りを一つ分動かす
			tmp = a[i] // jにあった値を仕切り内部の最後尾と入れ替え（元々iにあった値はピボットよりも大きいことを確認済み）
			a[i] = a[j]
			a[j] = tmp
		}
	}
	i++        // 全て判定を終えたら仕切り内部の最後尾を指したiを一つ動かし仕切りの場所を決定する
	tmp = a[i] // 仕切りの場所にピボットを配置
	a[i] = a[r]
	a[r] = tmp
	return i // 仕切りの場所を返す
}
