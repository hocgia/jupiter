package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/douyu/jupiter/pkg/util/xtest/proto/testproto"
	"github.com/douyu/jupiter/pkg/util/xtest/server/yell"
	"github.com/stretchr/testify/assert"
)

// TestBase test direct dial with New()
func TestDirectGrpc(t *testing.T) {
	t.Run("test direct grpc", func(t *testing.T) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		res, err := directClient.SayHello(ctx, &testproto.HelloRequest{
			Name: "hello",
		})
		assert.Nil(t, err)
		assert.Equal(t, res.Message, yell.RespFantasy.Message)
	})
}

func TestConfigBlockTrue(t *testing.T) {
	t.Run("test no address block", func(t *testing.T) {
		flag := false
		defer func() {
			if r := recover(); r != nil {
				flag = true
			}
			assert.True(t, flag)
		}()
		cfg := DefaultConfig()
		cfg.OnDialError = "panic"
		newGRPCClient(cfg)
	})
}

func TestConfigBlockFalse(t *testing.T) {
	t.Run("test no address and no block", func(t *testing.T) {
		cfg := DefaultConfig()
		cfg.OnDialError = "panic"
		cfg.Block = false
		conn := newGRPCClient(cfg)
		assert.Equal(t, conn.GetState().String(), "IDLE")
	})
}
