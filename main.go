package main

import (
	"crypto/tls"
	"errors"
	"github.com/LianJianTech/lj-go-common/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"lj-go-practice/pkg"
	"lj-go-practice/router"
	"lj-go-practice/router/middle"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "server config file path")
)

func main() {
	pflag.Parse()

	pkg.Init(*cfg)

	runMode := viper.GetString("run_mode")
	gin.SetMode(runMode)
	log.Infof("runMode:%v", runMode)

	g := gin.New()

	router.Load(
		g,
		middle.Logging(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the router has no response, or it might took too long to start up", err)
		}
		log.Info("the router has been deployed successfully")
	}()

	log.Infof("start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(viper.GetString("url") + viper.GetString("addr") + "/gcc/api/check/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("connect to the router error")
}
