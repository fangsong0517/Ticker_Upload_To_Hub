package app

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/Ticker_Upload_To_Hub/internal/config"
)

// Fn 定义函数类型
type Fn func(*config.Config) error

// MyTicker 定时器中的成员
type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
	Conf   *config.Config
}

// NewMyTick 新建一个定时器
func NewMyTick(f Fn, cfg *config.Config) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(cfg.Tool.Interval),
		Runner: f,
		Conf:   cfg,
	}
}

// Start 启动定时器需要执行的任务
func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner(t.Conf)
		}
	}
}

// Process 工作主要逻辑
func Process(cfg *config.Config) error {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("====================================" + timeNow + "====================================")
	cmd := exec.Command("bash", cfg.Shell.Path+cfg.Shell.Name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": \n" + string(output))
		return err
	}
	log.Println("Succ Upload!")
	log.Println("\n" + string(output))
	return nil
}
