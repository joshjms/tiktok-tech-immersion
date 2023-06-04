package main

import (
	"context"
	"testing"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"github.com/stretchr/testify/assert"
)

func TestIMServiceImpl_Send(t *testing.T) {
	s := &IMServiceImpl{}

	chat := "joshjms:jvthunder"
	text := "Jvthunder orz"
	sender := "joshjms"

	got, err := s.Send(context.Background(), &rpc.SendRequest{
		Message: &rpc.Message{
			Chat:   chat,
			Text:   text,
			Sender: sender,
		},
	})

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, int32(0), got.Code)
}
