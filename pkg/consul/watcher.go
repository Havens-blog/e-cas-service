package consul

import (
	"bytes"
	"context"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/pkg/email"
)

func watcher(cli *api.Client, vp *viper.Viper, path string, conf any) {
	time.Sleep(time.Second * 10)
	var g errgroup.Group
	g.Go(func() error {
		w, err := watch.Parse(map[string]interface{}{"type": "keyprefix", "prefix": path})
		if err != nil {
			return err
		}

		w.Handler = func(u uint64, i interface{}) {
			kv := i.(api.KVPairs)
			for _, v := range kv {
				if v.Key == path {
					err = vp.ReadConfig(bytes.NewBuffer(v.Value))
					if err != nil {
						return
					}
					if err = vp.Unmarshal(conf); err != nil {
						return
					}
				}
			}
		}
		err = w.RunWithClientAndHclog(cli, nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		email.SendWarn(context.Background(), consts.EmailTitleViperRemoteWatch, err.Error())
	}
}
