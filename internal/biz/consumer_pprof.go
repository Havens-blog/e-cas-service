package biz

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/Havens-blog/e-cas-service/pkg/kafka"
)

type PprofES struct {
	kafka.WorkerHandler
}

func (h *PprofES) Do(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	fmt.Println("pprof写入es...")
	/* fmt.Printf("Message Value:%s,Message topic:%q partition:%d offset:%d\n", string(msg.Value), msg.Topic, msg.Partition, msg.Offset) */
	return
}

func init() {
	kafka.Register("PprofES", &PprofES{})
}
