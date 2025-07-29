package main

import (
	"crypto/rand"
	"github.com/andlabs/ui"
	"log"
	"math/big"
	"strconv"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"

func generatePassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[num.Int64()]
	}
	return string(password)
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Генератор паролей", 300, 150, false)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		lengthEntry := ui.NewEntry()
		resultLabel := ui.NewLabel("Пароль появится здесь")

		button := ui.NewButton("Сгенерировать")
		button.OnClicked(func(*ui.Button) {
			length, err := strconv.Atoi(lengthEntry.Text())
			if err != nil || length <= 0 {
				resultLabel.SetText("Ошибка: введите число > 0")
				return
			}
			password := generatePassword(length)
			resultLabel.SetText(password)
		})

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Длина пароля:"), false)
		box.Append(lengthEntry, false)
		box.Append(button, false)
		box.Append(resultLabel, false)

		window.SetChild(box)
		window.Show()
	})
	if err != nil {
		log.Fatal(err)
	}
}
