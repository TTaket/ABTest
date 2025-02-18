package commands

import (
	"errors"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	GMainInst MainInstance
	GSignal   chan os.Signal
)

type MainInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

// func init() {
// 	// 设置工作目录
// 	_, filename, _, _ := runtime.Caller(0)
// 	BaseDir := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
// 	if err := os.Chdir(BaseDir); err != nil {
// 		panic(err)
// 	}
// }

func Run(inst MainInstance) {
	if inst == nil {
		panic(errors.New("inst is nil, exit"))
		return
	}

	//rand.Seed(time.Now().UTC().UnixNano())
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := inst.Initialize()
	if err != nil {
		panic(err)
		return
	}
	GMainInst = inst

	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
