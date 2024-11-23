package main

import (
	"encoding/json"
	"errors"
	"image/color"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Snippet struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var snippets []Snippet
var filteredSnippets []Snippet
var isEditing bool = false
var editingIndex int = -1

const snippetsFile = "snippets.json"

func loadSnippets() {
	file, err := os.Open(snippetsFile)
	if err != nil {
		if os.IsNotExist(err) {
			snippets = []Snippet{}
			filteredSnippets = snippets
			return
		}
		log.Fatalf("Dosya açılırken hata: %v", err)
	}
	defer file.Close()

	data := json.NewDecoder(file)
	if err := data.Decode(&snippets); err != nil {
		log.Fatalf("JSON ayrıştırılırken hata: %v", err)
	}
	filteredSnippets = snippets
}

func saveSnippets() {
	file, err := os.Create(snippetsFile)
	if err != nil {
		log.Fatalf("Dosya oluşturulurken hata: %v", err)
	}
	defer file.Close()

	data := json.NewEncoder(file)
	if err := data.Encode(&snippets); err != nil {
		log.Fatalf("JSON yazılırken hata: %v", err)
	}
}

type customTheme struct {
	isLight bool
}

func (c customTheme) BackgroundColor() color.Color {
	if c.isLight {
		return color.White
	}
	return color.RGBA{40, 40, 40, 255}
}

func (c customTheme) ButtonColor() color.Color {
	return color.RGBA{60, 60, 60, 255}
}

func (c customTheme) TextColor() color.Color {
	return color.White
}

func (c customTheme) HoverColor() color.Color {
	if c.isLight {
		return color.RGBA{80, 80, 80, 255}
	}
	return color.RGBA{220, 220, 220, 255}
}

func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return c.BackgroundColor()
	case theme.ColorNameButton:
		return c.ButtonColor()
	case theme.ColorNameForeground:
		return c.TextColor()
	case theme.ColorNameHover:
		return c.HoverColor()
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (c customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (c customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (c customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func updateSnippetList(content *fyne.Container, searchQuery string) {
	content.Objects = nil
	filteredSnippets = nil

	for _, snippet := range snippets {
		if strings.Contains(strings.ToLower(snippet.Title), strings.ToLower(searchQuery)) {
			filteredSnippets = append(filteredSnippets, snippet)
		}
	}

	for i, snippet := range filteredSnippets {
		index := i
		item := widget.NewButton(snippet.Title, func() {
			isEditing = true
			editingIndex = index
			titleEntry.SetText(snippet.Title)
			contentEntry.SetText(snippet.Content)
		})
		item.Importance = widget.LowImportance
		content.Add(item)
	}

	content.Refresh()
}

var titleEntry *widget.Entry
var contentEntry *widget.Entry

func main() {
	loadSnippets()

	myApp := app.New()
	myWindow := myApp.NewWindow("FastCode")

	currentTheme := &customTheme{isLight: false}
	myApp.Settings().SetTheme(currentTheme)

	themeDropdown := widget.NewSelect([]string{"Karanlık", "Beyaz"}, func(value string) {
		if value == "Beyaz" {
			currentTheme.isLight = true
		} else {
			currentTheme.isLight = false
		}
		myApp.Settings().SetTheme(currentTheme)
	})

	titleEntry = widget.NewEntry()
	titleEntry.SetPlaceHolder("Başlık")

	contentEntry = widget.NewMultiLineEntry()
	contentEntry.SetPlaceHolder("İçerik")

	listContent := container.NewVBox()
	updateSnippetList(listContent, "")

	listBackground := canvas.NewRectangle(color.RGBA{40, 40, 40, 255})
	listWithBackground := container.NewMax(listBackground, listContent)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Ara...")
	searchEntry.OnChanged = func(query string) {
		updateSnippetList(listContent, query)
	}

	saveButton := widget.NewButton("Kaydet", func() {
		title := strings.TrimSpace(titleEntry.Text)
		content := strings.TrimSpace(contentEntry.Text)

		if title == "" || content == "" {
			dialog.ShowError(errors.New("Başlık ve içerik boş olamaz!"), myWindow)
			return
		}

		if isEditing {
			snippets[editingIndex] = Snippet{Title: title, Content: content}
			isEditing = false
			editingIndex = -1
		} else {
			snippets = append(snippets, Snippet{Title: title, Content: content})
		}

		saveSnippets()
		titleEntry.SetText("")
		contentEntry.SetText("")
		updateSnippetList(listContent, "")
	})

	deleteButton := widget.NewButton("Sil", func() {
		if editingIndex == -1 {
			dialog.ShowError(errors.New("Silmek için bir kayıt seçin!"), myWindow)
			return
		}

		snippets = append(snippets[:editingIndex], snippets[editingIndex+1:]...)
		saveSnippets()
		isEditing = false
		editingIndex = -1
		titleEntry.SetText("")
		contentEntry.SetText("")
		updateSnippetList(listContent, "")
	})

	cancelButton := widget.NewButton("Vazgeç", func() {
		isEditing = false
		editingIndex = -1
		titleEntry.SetText("")
		contentEntry.SetText("")
	})

	content := container.NewVBox(
		themeDropdown,
		searchEntry,
		listWithBackground,
		titleEntry,
		contentEntry,
		container.NewHBox(saveButton, deleteButton, cancelButton),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
