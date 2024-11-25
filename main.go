package main

import (
	"errors"
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"snipper/services"
)

var snippetService *services.SnippetService
var isEditing bool = false
var editingIndex int = -1
var currentTheme *customTheme
var deleteButton, cancelButton *widget.Button

func main() {
	snippetService = services.NewSnippetService("snippets.json")
	err := snippetService.Load()
	if err != nil {
		log.Fatalf("Error loading snippets: %v", err)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("FastPaste")

	currentTheme = &customTheme{isLight: false}
	myApp.Settings().SetTheme(currentTheme)

	searchEntry := createStyledInput("Search...")
	titleEntry := createStyledInput("Subject")
	contentEntry := createStyledInput("Snippet")

	listContent := container.NewVBox()

	updateSnippetList := func(content *fyne.Container, query string) {
		content.Objects = nil
		snippets := snippetService.GetAll()

		for i, snippet := range snippets {
			if strings.Contains(strings.ToLower(snippet.Title), strings.ToLower(query)) {
				index := i
				item := widget.NewButton(snippet.Title, func() {
					isEditing = true
					editingIndex = index
					titleEntry.Entry.SetText(snippet.Title)
					contentEntry.Entry.SetText(snippet.Content)
					updateButtonVisibility()
				})
				content.Add(item)
			}
		}

		content.Refresh()
	}

	themeDropdown := widget.NewSelect([]string{"White Theme", "Dark Theme"}, func(selected string) {
		if selected == "White Theme" {
			currentTheme.isLight = true
		} else {
			currentTheme.isLight = false
		}
		myApp.Settings().SetTheme(currentTheme)
		applyInputStyles(searchEntry, titleEntry, contentEntry)
	})
	themeDropdown.SetSelected("Dark Theme")

	updateSnippetList(listContent, "")

	searchEntry.Entry.OnChanged = func(query string) {
		updateSnippetList(listContent, query)
	}

	saveButton := widget.NewButton("Save", func() {
		title := strings.TrimSpace(titleEntry.Entry.Text)
		content := strings.TrimSpace(contentEntry.Entry.Text)

		if title == "" || content == "" {
			dialog.ShowError(errors.New("Title and content cannot be empty!"), myWindow)
			return
		}

		if isEditing {
			_ = snippetService.Update(editingIndex, title, content)
			isEditing = false
		} else {
			_ = snippetService.Add(title, content)
		}

		editingIndex = -1
		titleEntry.Entry.SetText("")
		contentEntry.Entry.SetText("")
		updateSnippetList(listContent, "")
		updateButtonVisibility()
	})

	deleteButton = widget.NewButton("Delete", func() {
		if editingIndex != -1 {
			_ = snippetService.Delete(editingIndex)
			editingIndex = -1
			isEditing = false
			titleEntry.Entry.SetText("")
			contentEntry.Entry.SetText("")
			updateSnippetList(listContent, "")
			updateButtonVisibility()
		}
	})

	cancelButton = widget.NewButton("Cancel", func() {
		editingIndex = -1
		isEditing = false
		titleEntry.Entry.SetText("")
		contentEntry.Entry.SetText("")
		updateButtonVisibility()
	})

	buttons := container.NewHBox(
		layout.NewSpacer(),
		saveButton,
		deleteButton,
		cancelButton,
	)

	updateButtonVisibility()

	content := container.NewVBox(
		themeDropdown,
		fixedSpacer(10),
		searchEntry.Container,
		fixedSpacer(5),
		listContent,
		fixedSpacer(10),
		titleEntry.Container,
		fixedSpacer(5),
		contentEntry.Container,
		fixedSpacer(20),
		buttons,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func updateButtonVisibility() {
	if isEditing {
		deleteButton.Show()
		cancelButton.Show()
	} else {
		deleteButton.Hide()
		cancelButton.Hide()
	}
}

type styledInput struct {
	Container *fyne.Container
	Entry     *widget.Entry
}

func createStyledInput(placeholder string) *styledInput {
	entry := widget.NewEntry()
	entry.SetPlaceHolder(placeholder)

	bg := canvas.NewRectangle(color.White)
	container := container.NewMax(bg, entry)

	return &styledInput{
		Container: container,
		Entry:     entry,
	}
}

func fixedSpacer(height float32) fyne.CanvasObject {
	spacer := canvas.NewRectangle(color.Transparent)
	spacer.SetMinSize(fyne.NewSize(0, height))
	return spacer
}

func applyInputStyles(inputs ...*styledInput) {
	for _, input := range inputs {
		if currentTheme.isLight {
			input.Container.Objects[0] = canvas.NewRectangle(color.RGBA{240, 240, 240, 255})
		} else {
			input.Container.Objects[0] = canvas.NewRectangle(color.RGBA{60, 60, 60, 255})
		}
		input.Container.Refresh()
	}
}

type customTheme struct {
	isLight bool
}

func (c customTheme) BackgroundColor() color.Color {
	if c.isLight {
		return color.RGBA{245, 245, 245, 255}
	}
	return color.RGBA{40, 40, 40, 255}
}

func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return c.BackgroundColor()
	case theme.ColorNameButton:
		if c.isLight {
			return color.RGBA{200, 200, 200, 255}
		}
		return color.RGBA{60, 60, 60, 255}
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
