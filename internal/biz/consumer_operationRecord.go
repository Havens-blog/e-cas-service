package biz

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/pkg/kafka"
	"github.com/Havens-blog/e-cas-service/pkg/metric"
)

type OperationRecordReportES struct {
	kafka.WorkerHandler
}

func (h *OperationRecordReportES) Do(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	var data map[string]interface{}
	_ = json.Unmarshal(msg.Value, &data)
	err = repoUsecase.repo.ESInsertDoc(ctx, consts.ESIndexOperationRecord, data)
	if err != nil {
		metric.Count.With(fmt.Sprintf("consumer_%s_to_es_error", msg.Topic)).Inc()
		return err
	}
	return
}

func init() {
	kafka.Register("OperationRecordReportES", &OperationRecordReportES{})
}
