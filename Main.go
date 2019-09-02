package main

import (
	"crypto/tls"
	"github.com/coreos/etcd/pkg/transport"
	"go.uber.org/zap"
	"strings"
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

	var tlsConfig *tls.Config

	tlsInfo := transport.TLSInfo{
		CertFile:       etcd_certfile ,
		KeyFile:        etcd_keyfile ,
		TrustedCAFile:  etcd_cafile ,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		fmt.Println("get client config error...")
		panic(err)
	}

	etcd_servers_array := strings.Split(etcd_servers, ",")
	if len(etcd_servers_array) == 0 {
		fmt.Println("server size is 0 ...")
		panic(err)
	}

	config := clientv3.Config{
		Endpoints:   etcd_servers_array,
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
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
	RootCmd.Flags().StringVarP(&etcd_keyfile, "etcd-keyfile", "", "", "etcd_keyfile")
}

var RootCmd = &cobra.Command{
	Use:   "etcdSnapshot",
	Short: "etcdSnapshot",
	Long:  `etcdSnapshot`,
	Run: func(cmd *cobra.Command, args []string) {
		if etcd_servers == "" || etcd_cafile == "" || etcd_certfile == "" || etcd_keyfile == "" {
			fmt.Println("param is not correct")
			panic("param is not correct")
		}
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


