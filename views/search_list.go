package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const SearchListName string = "SearchList"

type SearchList struct {
	View        *gocui.View
	Mangas      []nato.Manga
	NameToIDMap map[string]string
}

func GetSearchList(maxX, maxY int, g *gocui.Gui) (*SearchList, error) {
	sl := SearchList{}
	x0, y0, x1, y1 := sl.GetCoords(maxX, maxY)

	slView, err := g.SetView(SearchListName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	slView.Title = SearchListName
	slView.SelFgColor = gocui.ColorBlack
	slView.SelBgColor = gocui.ColorGreen
	slView.BgColor = gocui.ColorBlack
	slView.FgColor = gocui.ColorWhite
	slView.Highlight = true
	slView.Editable = true
	slView.Wrap = true
	slView.Editor = readOnlyEditor

	sl.View = slView
	sl.Mangas = make([]nato.Manga, 0)
	sl.NameToIDMap = make(map[string]string)

	return &sl, err
}

func (sl *SearchList) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return 1, 1, maxX/2 - 1, maxY - SearchBarHeight - 2
}

func (sl *SearchList) FormatMangas() string {
	s := fmt.Sprintf("			press ENTER on the manga title(%s) to start reading\n\n", Selector)

	for _, mg := range sl.Mangas {
		s += fmt.Sprintf("	%s %s\n			Author: %s\n\n", Selector, mg.Name, mg.Author.Name)
		sl.NameToIDMap[mg.Name] = mg.ID
	}

	return s
}

func (sl *SearchList) GetBuf() string {
	return sl.View.Buffer()
}
