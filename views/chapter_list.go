package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
	nato "github.com/kaikaew13/manganato-api"
)

const ChapterListName = "ChapterList"

type ChapterList struct {
	View        *gocui.View
	Chapters    []nato.Chapter
	NameToIDMap map[string]string
}

func GetChapterList(maxX, maxY int, g *gocui.Gui) (*ChapterList, error) {
	cl := ChapterList{}
	x0, y0, x1, y1 := cl.GetCoords(maxX, maxY)

	clView, err := g.SetView(ChapterListName, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	clView.Title = ChapterListName
	clView.SelFgColor = gocui.ColorGreen
	clView.BgColor = gocui.ColorBlack
	clView.FgColor = gocui.ColorWhite
	clView.Editable = true
	clView.Editor = selectingEditor

	cl.View = clView
	cl.Chapters = make([]nato.Chapter, 0)
	cl.NameToIDMap = make(map[string]string)

	return &cl, err
}

func (cl *ChapterList) GetCoords(maxX, maxY int) (x0, y0, x1, y1 int) {
	return maxX / 2, (maxY - SearchBarHeight - 1) / 2, maxX - 1, maxY - SearchBarHeight - 2
}

func (cl *ChapterList) FormatChapters() string {
	s := "\n		CHAPTER NAME					VIEWS			UPLOADED\n\n"

	for _, chapter := range cl.Chapters {
		s += fmt.Sprintf("		%s				%s			%s\n\n", chapter.ChapterName, chapter.Views, chapter.Uploaded)
	}

	return s
}
