package main

import (
	"context"
	"fmt"

	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
)

func main() {
	// クライアントの作成
	cli, err := rawkv.NewClient(context.TODO(), []string{"pd0:2379", "pd1:2379", "pd2:2379"}, config.Default())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	fmt.Printf("cluster ID: %d\n", cli.ClusterID())

	keys := [][]byte {
		[]byte("tbf02"),
		[]byte("tbf03"),
		[]byte("tbf04"),
		[]byte("tbf05"),
		[]byte("tbf06"),
		[]byte("tbf07"),
	}
	values := [][]byte {
		[]byte("2017-04-09"),
		[]byte("2017-10-22"),
		[]byte("2018-04-22"),
		[]byte("2018-10-08"),
		[]byte("2019-04-14"),
		[]byte("2019-09-22"),
	}

	// 6つの値を保存
	err = cli.BatchPut(context.TODO(), keys, values)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully put %d values\n", len(keys))

	// 複数の値を取得
	reqKeys := [][]byte{[]byte("tbf02"), []byte("tbf06")}
	res, err := cli.BatchGet(context.TODO(), reqKeys)
	if err != nil {
		panic(err)
	}
	fmt.Printf("batch get request found values\n")
	for i := range res {
		fmt.Printf("  key: %s value: %s\n", reqKeys[i], res[i])
	}

	// 範囲を指定して値を取得
	resKeys, resValues, err := cli.Scan(context.TODO(), []byte("tbf03"), []byte("tbf06"), 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("range request found values\n")
	for i := range resKeys {
		fmt.Printf("  key: %s value: %s\n", resKeys[i], resValues[i])
	}

	// 値の削除
	err = cli.BatchDelete(context.TODO(), keys)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %d values deleted\n", len(keys))
}
