package alimns

import (
	"github.com/bigtree8/flow-core/config"
	"testing"
)

func TestNewMnsClient(t *testing.T) {
	conf := config.MnsConfig{
		Url:             "",
		AccessKeyId:     "",
		AccessKeySecret: "",
	}
	c, err := NewMnsClient(conf)
	if err != nil {
		t.Error(err)
		return
	} else if c != nil {
		t.Error("client is not nil")
		return
	}
}

func TestNewMnsClient2(t *testing.T) {
	conf := config.MnsConfig{
		Url:             "http://www.baidu.com",
		AccessKeyId:     "1",
		AccessKeySecret: "2",
	}

	_, err := NewMnsClient(conf)
	if err == nil {
		t.Error("invalid config must return err")
	}
}

func TestGetMnsBasicQueue(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Error("not panic")
		}
	}()
	GetMnsBasicQueue(nil, "test")
}
