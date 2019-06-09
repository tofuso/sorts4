package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/tofuso/sorts4/sorts"
)

func main() {
	var (
		fn            = flag.String("file", "", "ソート対象を記述したファイルを指定する。")
		flaginsertion = flag.Bool("i", false, "insertion sortを実行します。")
		flagmerge     = flag.Bool("m", false, "merge sortを実行します。")
		flagheap      = flag.Bool("h", false, "heap sortを実行します。")
		flagquick     = flag.Bool("q", false, "quick sortを実行します。")
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
	fmt.Println("元の配列: " + showArray(keys))
	fmt.Println("要素数: "+strconv.Itoa(len(keys)))
	insertionKeys := make([]int, len(keys))
	mergeKeys := make([]int, len(keys))
	heapKeys := make([]int, len(keys))
	quickKeys := make([]int, len(keys))
	copy(insertionKeys, keys)
	copy(mergeKeys, keys)
	copy(heapKeys, keys)
	copy(quickKeys, keys)

	if *flaginsertion {
		sorts.InsertionSort(&insertionKeys)
	}
	if *flagmerge {
		sorts.MergeSort(&mergeKeys)
	}
	if *flagheap {
		sorts.HeapSort(&heapKeys)
	}
	if *flagquick {
		sorts.QuickSort(&quickKeys)
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
