package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tofuso/sorts4/sorts"
)

func main() {
	var (
		fn            = flag.String("file", "", "ソート対象を記述したファイルを指定する。")
		flagview      = flag.Bool("view", false, "ソートの実行結果を表示します。")
		flaginsertion = flag.Bool("insertion", false, "insertion sortを実行します。")
		flagmerge     = flag.Bool("merge", false, "merge sortを実行します。")
		flagheap      = flag.Bool("heap", false, "heap sortを実行します。")
		flagquick     = flag.Bool("quick", false, "quick sortを実行します。")
	)
	flag.Parse()

	if *fn == "" {
		fmt.Fprintln(os.Stderr, "ファイル名を指定してください。")
		os.Exit(1)
	}
	file, err := os.Open(*fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ファイル: "+*fn+"を開くことに失敗しました。")
	}
	defer file.Close()

	sc := bufio.NewScanner(file) // ファイルからの読み込み処理
	var keys []int
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, sc.Text()+"を整数に変換できませんでした。")
		}
		keys = append(keys, num)
	}
	if *flagview {
		fmt.Println("元の配列: " + showArray(keys))
	}
	fmt.Println("要素数: " + strconv.Itoa(len(keys)))

	var start, end time.Time // 実行時間測定
	insertionKeys := make([]int, len(keys))
	mergeKeys := make([]int, len(keys))
	heapKeys := make([]int, len(keys))
	quickKeys := make([]int, len(keys))
	copy(insertionKeys, keys)
	copy(mergeKeys, keys)
	copy(heapKeys, keys)
	copy(quickKeys, keys)

	// 各ソートアルゴリズムで実行した結果を表示する
	if *flaginsertion {
		start = time.Now() // 測定開始
		sorts.InsertionSort(insertionKeys)
		end = time.Now() // 測定終了
		fmt.Println("insertion sort:")
		if *flagview {
			fmt.Println(showArray(insertionKeys))
		}
		fmt.Println("実行時間: ", end.Sub(start).Seconds()*1000, "ms")
	}
	if *flagmerge {
		start = time.Now() // 測定開始
		sorts.MergeSort(mergeKeys)
		end = time.Now() // 測定終了
		fmt.Println("merge sort:")
		if *flagview {
			fmt.Println(showArray(mergeKeys))
		}
		fmt.Println("実行時間: ", end.Sub(start).Seconds()*1000, "ms")
	}
	if *flagheap {
		start = time.Now() // 測定開始
		sorts.HeapSort(heapKeys)
		end = time.Now() // 測定終了
		fmt.Println("heap sort:")
		if *flagview {
			fmt.Println(showArray(heapKeys))
		}
		fmt.Println("実行時間: ", end.Sub(start).Seconds()*1000, "ms")
	}
	if *flagquick {
		start = time.Now() // 測定開始
		sorts.QuickSort(quickKeys)
		end = time.Now() // 測定終了
		fmt.Println("quick sort:")
		if *flagview {
			fmt.Println(showArray(quickKeys))
		}
		fmt.Println("実行時間: ", end.Sub(start).Seconds()*1000, "ms")
	}
	return
}

func showArray(arr []int) string {
	var l string
	for i, num := range arr {
		if i != len(arr)-1 {
			l += strconv.Itoa(num) + ", "
		} else {
			l += strconv.Itoa(num)
		}
	}
	return l
}
