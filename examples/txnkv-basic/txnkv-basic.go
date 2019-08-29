package main

import (
	"context"
	"fmt"
	"binary"

	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/txnkv"
)

func main() {
	// クライアントの作成
	cli, err := txnkv.NewClient(context.TODO(), []string{"pd0:2379", "pd1:2379", "pd2:2379"}, config.Default())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	
	key := []byte("tbf07")
	value := []byte("2019-09-22")
	



	// 値の削除
	err = cli.BatchDelete(context.TODO(), keys)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %d values deleted\n", len(keys))
}

func intToBytes(input int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, input)
	return bs
}