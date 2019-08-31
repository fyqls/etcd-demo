package main

import (
	"go.uber.org/zap"
	"time"

	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/snapshot"
	"os"
	"path/filepath"
)


func create_etcd_snapshot() {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:12379"},
		DialTimeout: 5 * time.Second,
	}

	//cli, err := clientv3.New(config)
	//if err!=nil{
	//}

	//logger := zap.NewExample()
	sp := snapshot.NewV3(zap.NewExample())
	dpPath := filepath.Join(os.TempDir(), fmt.Sprintf("snapshot%d.db", time.Now().Nanosecond()))
	if err := sp.Save(context.Background(), config, dpPath); err != nil {
		fmt.Println("snapshot error...")
		panic(err)
	}
	fmt.Println("backup sucess!")

}


func main() {
	create_etcd_snapshot()
	fmt.Println("backup sucess!")
}
