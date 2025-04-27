package bubbleup

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	_ tea.Model = BubbleUp[string]{}
	_ tea.Model = BubbleUp[int]{}
)

type BubbleUp[T any] struct {
	title      string
	items      []BubbleUpItem[T]
	curentItem int
	grabbed    bool
	theme      *huh.Theme
	showHelp   bool
	submitted  bool
}

type BubbleUpItem[T any] struct {
	key   T
	value T
}

func (bui BubbleUpItem[T]) KeyAsString() string {
	return valueAsString(bui.key)
}

func (bui BubbleUpItem[T]) ValueAsString() string {
	return valueAsString(bui.value)
}

func New[T any]() BubbleUp[T] {
	return BubbleUp[T]{}
}

func NewItem[T any](key, value T) BubbleUpItem[T] {
	return BubbleUpItem[T]{
		key:   key,
		value: value,
	}
}

func (b BubbleUp[T]) Down() BubbleUp[T] {
	return b.move(1)

}

func (b BubbleUp[T]) Up() BubbleUp[T] {
	return b.move(-1)
}

func (b BubbleUp[T]) move(dx int) BubbleUp[T] {
	next := b.curentItem + dx
	if next >= len(b.items) {
		next = 0
	}
	if next < 0 {
		next = len(b.items) - 1
	}
	if b.grabbed {
		prev := b.curentItem
		b.curentItem = next
		return b.swap(prev, next)
	}
	b.curentItem = next
	return b
}

func (b BubbleUp[T]) swap(i, j int) BubbleUp[T] {
	var tmp BubbleUpItem[T]
	tmp = b.items[i]
	b.items[i] = b.items[j]
	b.items[j] = tmp
	return b
}

func (b BubbleUp[T]) WithItems(opt ...BubbleUpItem[T]) BubbleUp[T] {
	b.items = opt
	return b
}

func (b BubbleUp[T]) WithTitle(title string) BubbleUp[T] {
	b.title = title
	return b
}

func (b BubbleUp[T]) WithTheme(theme *huh.Theme) BubbleUp[T] {
	b.theme = theme
	return b
}

func (b BubbleUp[T]) WithHelp(ok bool) BubbleUp[T] {
	b.showHelp = ok
	return b
}

func (b BubbleUp[T]) Items() []BubbleUpItem[T] {
	return b.items
}

func (b BubbleUp[T]) Values() (out []T) {
	out = make([]T, 0, len(b.items))
	for _, item := range b.items {
		out = append(out, item.value)
	}
	return out
}

func (b BubbleUp[T]) IsSubmitted() bool {
	return b.submitted
}

// Init implements tea.Model.
func (b BubbleUp[T]) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (b BubbleUp[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case " ", "space", "tab":
			b.grabbed = !b.grabbed
			return b, nil
		case "up":
			return b.Up(), nil
		case "down":
			return b.Down(), nil
		case "enter", "return":
			b.submitted = true
			return b, nil
		}
	}

	return b, nil
}

// View implements tea.Model.
func (b BubbleUp[T]) View() string {

	if b.submitted {
		return ""
	}

	var sb strings.Builder

	sb.WriteString(b.theme.Group.Title.Render(b.title))
	sb.WriteRune('\n')
	sb.WriteRune('\n')

	var style lipgloss.Style
	for i, item := range b.items {
		switch i {
		case b.curentItem:
			if b.grabbed {
				style = b.theme.Focused.SelectedPrefix
			} else {
				style = b.theme.Focused.UnselectedPrefix.Foreground(b.theme.Focused.SelectedOption.GetForeground())
			}
		default:
			// style = b.theme.Focused.UnselectedOption
			style = b.theme.Focused.UnselectedPrefix
		}
		sb.WriteString(
			style.Render(
				item.KeyAsString()))
		sb.WriteRune('\n')
	}

	sb.WriteRune('\n')
	sb.WriteString(b.theme.Help.FullDesc.Render("↑ up • ↓ down • [space] (de)select • [enter] submit"))
	sb.WriteRune('\n')

	return sb.String()
}

func valueAsString(v any) string {
	switch vv := v.(type) {
	case string:
		return vv
	case fmt.Stringer:
		return vv.String()
	case error:
		return vv.Error()
	default:
		return fmt.Sprintf("%v", vv)
	}
}
