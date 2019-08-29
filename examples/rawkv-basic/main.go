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

	key := []byte("tbf07")
	val := []byte("2019-09-22")

	// 値の保存
	err = cli.Put(context.TODO(), key, val)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully put %s:%s to tikv\n", key, val)

	// 値の取得
	val, err = cli.Get(context.TODO(), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("found val: %s for key: %s\n", val, key)

	// 値の削除
	err = cli.Delete(context.TODO(), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %s deleted\n", key)

	// 値の保存 (既に削除されているのでnilが返る)
	val, err = cli.Get(context.TODO(), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("found val: %s for key: %s\n", val, key)
}
