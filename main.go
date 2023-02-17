package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Ticker_Upload_To_Hub/app"
	"github.com/Ticker_Upload_To_Hub/internal/config"
	"github.com/Ticker_Upload_To_Hub/internal/utility"
)

func main() {
	opt := utility.NewOption()
	opt.Parse()
	// 初始化项目配置
	cfg, err := config.New(opt.ConfigFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("global config: %+v\n", cfg)

	f, err := os.OpenFile("../log/log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()

	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	t := app.NewMyTick(app.Process, cfg)
	t.Start()
}
