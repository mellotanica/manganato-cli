package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const searchListName string = "SearchList"

type SearchList struct {
	View        *gocui.View
	Name        string
	Mangas      []nato.Manga
	NameToIDMap map[string]string
	OriginX     int
	OriginY     int
}

func GetSearchList(maxX, maxY int, g *gocui.Gui) (*SearchList, error) {
	sl := SearchList{}
	x0, y0, x1, y1 := sl.GetCoords(maxX, maxY)

	slView, err := g.SetView(searchListName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	slView.Title = searchListName
	slView.SelFgColor = gocui.ColorBlack
	slView.SelBgColor = gocui.ColorGreen
	slView.BgColor = gocui.ColorBlack
	slView.FgColor = gocui.ColorWhite
	slView.Highlight = true
	slView.Editable = true
	slView.Wrap = true
	slView.Editor = readOnlyEditor

	sl.View = slView
	sl.Name = searchListName
	sl.Mangas = make([]nato.Manga, 0)
	sl.NameToIDMap = make(map[string]string)
	sl.OriginX, sl.OriginY = slView.Origin()

	return &sl, err
}

func (sl *SearchList) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return 1, 1, maxX/2 - 1, maxY - SearchBarHeight - 2
}

func (sl *SearchList) FormatMangas() string {
	s := fmt.Sprintf("\t\t\tpress ENTER on the manga title(%s) to start reading\n\n", Selector)

	for _, mg := range sl.Mangas {
		s += fmt.Sprintf("\t%s %s\n\t\tAuthor: %s\n\n", Selector, mg.Name, mg.Author.Name)
		sl.NameToIDMap[mg.Name] = mg.ID
		sl.NameToIDMap[mg.Author.Name] = mg.Author.ID
		sl.NameToIDMap[strings.ToLower(mg.Author.Name)] = mg.Author.ID
		sl.NameToIDMap[strings.ToUpper(mg.Author.Name)] = mg.Author.ID
	}

	return s
}

func (sl *SearchList) GetBuf() string {
	return sl.View.Buffer()
}
