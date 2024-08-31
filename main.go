package main

import (
	"math/rand/v2"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	application fyne.App
	window      fyne.Window

	numbers    = false
	numbersArr = []rune("1234567890")
	special    = false
	specialArr = []rune("!@#$%&")
	capital    = false
	capitalArr = []rune("QWERTYUIOPASDFGHJKLZXCVBNM")
	passwd     = ""
)

func main() {
	application = app.New()
	window = application.NewWindow("password generator")

	numbersChek := widget.NewCheck("numbers", func(b bool) {
		numbers = b
	})

	specialCheck := widget.NewCheck("special chars", func(b bool) {
		special = b
	})

	capitalCheck := widget.NewCheck("capital letters", func(b bool) {
		capital = b
	})

	checks := container.NewHBox(numbersChek, specialCheck, capitalCheck)
	checksCent := container.NewCenter(checks)

	passText := widget.NewLabel("4")
	passSlid := widget.NewSlider(4, 40)
	passSlid.OnChanged = func(f float64) {
		passText.Text = strconv.Itoa(int(f))
		passText.Refresh()
	}
	passCont := container.New(&InputLayout{}, passSlid, passText)

	output := widget.NewEntry()

	genPassBtn := widget.NewButton("Generate", func() {
		chars := []rune("qwertyuiopasdfghjklzxcvbnm")
		passwd = ""

		if numbers {
			chars = append(chars, numbersArr...)
		}
		if special {
			chars = append(chars, specialArr...)
		}
		if capital {
			chars = append(chars, capitalArr...)
		}

		for range int(passSlid.Value) {
			passwd += string(chars[rand.IntN(len(chars))])
		}
		output.Text = passwd
		output.Refresh()
	})

	mainBox := container.NewVBox(checksCent, passCont, genPassBtn, output)

	window.SetContent(mainBox)
	window.ShowAndRun()
}
