package main

import (
	"strings"

	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func switchView(g *gocui.Gui, v *gocui.View) error {
	for i, name := range viewNames {
		if name == v.Name() {
			g.SetCurrentView(viewNames[(i+1)%len(viewNames)])
			break
		}
	}
	return nil
}

func enterCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()

	valid, cmd, args := validateCommand(s)
	if valid {
		screen.sb.SaveCommand(s)

		if err := runCommand(cmd, args); err != nil {
			return err
		}
	}

	x, y := v.Origin()
	if err := v.SetCursor(x, y); err != nil {
		return err
	}

	v.Clear()

	return nil
}

func getPrevCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()
	s = screen.sb.GetPrevCommand(s)

	v.Clear()
	v.Write([]byte(s))

	return nil
}

func getNextCommand(g *gocui.Gui, v *gocui.View) error {
	s := v.Buffer()
	s = screen.sb.GetNextCommand(s)

	v.Clear()
	v.Write([]byte(s))

	return nil
}

func pickManga(g *gocui.Gui, v *gocui.View) error {
	_, y := v.Cursor()

	s, _ := v.Line(y)
	s = strings.TrimSpace(s)

	err := getMangaScreen(s)

	return err
}
