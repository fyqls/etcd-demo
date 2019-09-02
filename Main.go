package main

import (
	"go.uber.org/zap"
	"time"

	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/snapshot"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	etcd_servers string
	etcd_cafile string
	etcd_certfile string
	etcd_keyfile string
)

func create_etcd_snapshot() {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:12379"},
		DialTimeout: 5 * time.Second,
	}

	sp := snapshot.NewV3(zap.NewExample())
	dpPath := filepath.Join(os.TempDir(), fmt.Sprintf("snapshot%d.db", time.Now().Nanosecond()))
	if err := sp.Save(context.Background(), config, dpPath); err != nil {
		fmt.Println("snapshot error...")
		panic(err)
	}
	fmt.Println("backup sucess!")

}

func init() {
	RootCmd.Flags().StringVarP(&etcd_servers, "etcd_servers", "", "", "etcd servers address")
	RootCmd.Flags().StringVarP(&etcd_cafile, "etcd-cafile", "", "", "etcd-cafile")
	RootCmd.Flags().StringVarP(&etcd_certfile, "etcd-certfile", "", "", "etcd-certfile")
	RootCmd.Flags().StringVarP(&etcd_certfile, "etcd_keyfile", "", "", "etcd_keyfile")
}

var RootCmd = &cobra.Command{
	Use:   "etcdSnapshot",
	Short: "etcdSnapshot",
	Long:  `etcdSnapshot`,
	Run: func(cmd *cobra.Command, args []string) {
		create_etcd_snapshot()
	},
}

func main() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//create_etcd_snapshot()
	//fmt.Println("backup sucess!")
}


//- --etcd_servers=https://cicd01:2379,https://cicd02:2379,https://cicd03:2379
//- --etcd-cafile=/etc/kubernetes/ssl/etcd-ca.pem
//- --etcd-certfile=/etc/kubernetes/ssl/etcd.pem
//- --etcd-keyfile=/etc/kubernetes/ssl/etcd-key.pem


