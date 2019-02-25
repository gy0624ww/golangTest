package main

import (
	"os/exec"
	"context"
	"fmt"
	"time"
)

type message struct {
	content []byte
	err error
}

func main() {
	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		contentChannel chan *message
		readMessage *message
	)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	//ctx, _ = context.WithCancel(context.TODO())
	contentChannel = make(chan *message, 1000)
	go func(ctx context.Context) {
		var (
			cmd *exec.Cmd
			content []byte
			err error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 5;echo hello world;")

		content, err = cmd.CombinedOutput()
		// 输出写到channel 到main函数读取
		contentChannel <- &message{
			content:content,
			err:err,
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancelFunc()

	readMessage = <- contentChannel
	fmt.Println(111)
	fmt.Println(readMessage.err,string(readMessage.content))
}
