package main

import (
	"log"
	"time"

	"github.com/SeedJobs/devops-go-tests/lock/problem"
	ui "github.com/gizak/termui"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatal("Failed to initialise Ui")
	}
	defer ui.Close()
	d := ui.NewPar("UNLOCKED")
	d.Y = 2
	d.Height = 3
	d.Width = 50
	d.TextFgColor = ui.ColorWhite
	d.BorderLabel = "Lock Display"
	d.BorderFg = ui.ColorCyan

	info := ui.NewPar("Info")
	info.Height = 8
	info.Width = 50
	info.Y = 6
	info.BorderLabel = "Lock Usage"
	info.Text = `[0..9] - Digits to be entered into the lock
s - * button
   Clears the display
h - # button
   Will enter the key to the lock
`
	eLog := problem.NewEventLog()
	event := ui.NewList()
	event.BorderLabel = "Lock event"
	event.Items = eLog.GetBuff()
	event.Y = 14
	event.Width = 50
	event.Height = 10

	lock := problem.NewLock(eLog)
	lock.Initialise()

	ui.Render(d, info, event)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		eLog.Add("Quit")
		event.Items = eLog.GetBuff()
		ui.Render(d, info, event)
		ui.StopLoop()
	})
	ui.Handle("/timer/1s", func(ui.Event) {
		lock.Tick()
		d.Text = lock.GetDisplay()
		event.Items = eLog.GetBuff()
		ui.Render(d, info, event)
	})
	ui.Handle("/sys/kbd", func(key ui.Event) {
		kT := key.Data.(ui.EvtKbd)
		k := rune([]byte(kT.KeyStr)[0])
		if k == 'q' {
			eLog.Add("Quit")
			event.Items = eLog.GetBuff()
			ui.Render(d, info, event)
			ui.StopLoop()
			return
		}
		defer ui.Render(d, info, event)
		if err := lock.EnteredKey(k); err != nil {
			eLog.Add(err.Error())
			event.Items = eLog.GetBuff()
			return
		}
		eLog.Add("Pressed key " + string(k))
		d.Text = lock.GetDisplay()
		event.Items = eLog.GetBuff()
	})
	ui.Loop()
	t := time.After(1 * time.Second)
	<-t
}
