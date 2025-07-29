package main

import (
	"crypto/rand"
	"github.com/andlabs/ui"
	"log"
	"math/big"
	"strconv"
)

func passgen(length int, useUpper, useDigits, useSpecial bool) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()-_=+[]{}<>?"

	charset := lower
	if useUpper {
		charset += upper
	}
	if useDigits {
		charset += digits
	}
	if useSpecial {
		charset += special
	}

	if len(charset) == 0 {
		return "Ошибка: выберите хотя бы 1 тип символов"
	}

	passwrd := make([]byte, length)
	for i := range passwrd {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		passwrd[i] = charset[num.Int64()]
	}
	return string(passwrd)
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Генератор паролей", 400, 250, false)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		lengthEntry := ui.NewEntry()
		resultLabel := ui.NewLabel("Пароль появится здесь")

		checkUpper := ui.NewCheckbox("Включать заглавные буквы (A-Z)")
		checkDigits := ui.NewCheckbox("Включать цифры (0-9)")
		checkSpecial := ui.NewCheckbox("Включать спецсимволы (!@#...)")

		button := ui.NewButton("Сгенерировать")
		button.OnClicked(func(*ui.Button) {
			length, err := strconv.Atoi(lengthEntry.Text())
			if err != nil || length <= 0 {
				resultLabel.SetText("Ошибка: введите число > 0")
				return
			}

			passwrd := passgen(length, checkUpper.Checked(), checkDigits.Checked(), checkSpecial.Checked())
			resultLabel.SetText(passwrd)
		})

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Длина пароля:"), false)
		box.Append(lengthEntry, false)
		box.Append(checkUpper, false)
		box.Append(checkDigits, false)
		box.Append(checkSpecial, false)
		box.Append(button, false)
		box.Append(resultLabel, false)

		window.SetChild(box)
		window.Show()
	})
	if err != nil {
		log.Fatal(err)
	}
}
