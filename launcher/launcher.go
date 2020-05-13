package launcher

import (
	"fmt"
	"whisky-server/launcher/bottle"
)

type Launcher struct {
	bottles []bottle.Bottle
	done    chan struct{}
}

func NewLauncher() Launcher {
	return Launcher{
		bottles: nil,
		done:   make(chan struct{}),
	}
}

func (l *Launcher) Run() {
	for _, data := range l.bottles {
		life, ok := data.Constructor.(bottle.LifeCycle)
		if !ok {
			fmt.Println("lifecycle error")
			continue
		}
		go life.Start()
	}
	for {
		select {
		case <- l.done:
			return
		}
	}
}

func (l *Launcher) AddBottle(constructor interface{}) {
	bot := bottle.Bottle{Constructor:constructor}
	l.bottles = append(l.bottles, bot)
}