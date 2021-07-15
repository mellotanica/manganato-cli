package views

import (
	"github.com/jroimartin/gocui"
)

const (
	SearchBarHeight int    = 2
	searchBarName   string = "SearchBar"
)

type SearchBar struct {
	View     *gocui.View
	Name     string
	Commands *[]string
}

func GetSearchBar(maxX, maxY int, g *gocui.Gui) (*SearchBar, error) {
	sb := SearchBar{}
	x0, y0, x1, y1 := sb.GetCoords(maxX, maxY)

	sbView, err := g.SetView(searchBarName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	sbView.Title = searchBarName
	sbView.SelFgColor = gocui.ColorGreen
	sbView.BgColor = gocui.ColorBlack
	sbView.FgColor = gocui.ColorWhite
	sbView.Editable = true

	cmds := make([]string, 0)

	sb.View = sbView
	sb.Name = searchBarName
	sb.Commands = &cmds
	return &sb, err
}

func (sb *SearchBar) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return 1, maxY - SearchBarHeight - 1, maxX - 1, maxY - 1
}

func (sb *SearchBar) SaveCommand(cmd string) {
	if cmd == "" {
		return
	}

	for i, v := range *sb.Commands {
		if v == removeNewline(cmd) {
			(*sb.Commands) = append((*sb.Commands)[:i], (*sb.Commands)[i+1:]...)
			break
		}
	}

	*sb.Commands = append(*sb.Commands, removeNewline(cmd))
}

func (sb *SearchBar) GetPrevCommand(cmd string) string {
	cmds := *sb.Commands

	if len(cmds) == 0 {
		return ""
	}

	if cmd == "" {
		return cmds[len(cmds)-1]
	}

	for i, v := range cmds {
		if v == removeNewline(cmd) {
			if i-1 < 0 {
				return ""
			}
			return cmds[i-1]
		}
	}

	return cmds[len(cmds)-1]
}

func (sb *SearchBar) GetNextCommand(cmd string) string {
	cmds := *sb.Commands

	if len(cmds) == 0 || cmd == "" {
		return ""
	}

	for i, v := range cmds {
		if v == removeNewline(cmd) {
			if i+1 == len(cmds) {
				return ""
			}
			return cmds[i+1]
		}
	}

	return ""
}

func removeNewline(s string) string {
	return s[:len(s)-1]
}
