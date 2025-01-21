package consul

import (
	"github.com/Havens-blog/e-cas-service/internal/consts"
	"io"
	"testing"

	"github.com/hashicorp/consul/api"
)

// 模拟 Consul 客户端
type mockConsulClient struct {
	kv *mockKV
}

// 模拟 Consul KV 存储
type mockKV struct {
	getErr error
	kvPair *api.KVPair
}

func (m *mockKV) Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error) {
	return m.kvPair, nil, m.getErr
}

// 模拟 Viper
type mockViper struct {
	readErr error
}

func (m *mockViper) ReadConfig(in io.Reader) error {
	return m.readErr
}

func TestNewConsulClient_Success(t *testing.T) {
	cli, err := NewConsulClient("113.44.143.194:8500", "8501f1da-ad0f-401c-9694-6ac1f8ff1fc4")
	if err != nil {
		t.Errorf("NewConsulClient failed: %v", err)
	}
	if cli == nil {
		t.Error("NewConsulClient returned a nil client")
	}
}

func TestGetConsulKV_Success(t *testing.T) {
	cli, err := NewConsulClient("113.44.143.194:8500", "8501f1da-ad0f-401c-9694-6ac1f8ff1fc4")
	if err != nil {
		t.Errorf("NewConsulClient failed: %v", err)
	}
	dict := make(map[string]interface{}, 0)
	_, err = GetConsulKV(cli, "dev/dict.json", &dict)
	if err != nil {
		t.Errorf("Expected config value 'value', got '%v'", consts.Conf.Consul.Kv.DictPath)
	}
	if err == nil {
		t.Error("GetConsulKV returned a nil viper instance")
	}
}
