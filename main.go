package main

import (
	"crypto/rand"
	"github.com/andlabs/ui"
	"log"
	"math/big"
	"os/exec"
	"runtime"
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

// copt to clipboard
func ctp(text string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "echo "+text+"| clip")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	default:
		return
	}

	err := cmd.Run()
	if err != nil {
		return
	}
}

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Генератор паролей", 420, 300, false)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		lengthEntry := ui.NewEntry()
		lengthEntry.SetText("12")
		lengthBox := ui.NewHorizontalBox()
		lengthBox.Append(ui.NewLabel("Длина пароля:"), false)
		lengthBox.Append(lengthEntry, true)

		checkUpper := ui.NewCheckbox("Заглавные (A-Z)")
		checkUpper.SetChecked(true)
		checkDigits := ui.NewCheckbox("Цифры (0-9)")
		checkDigits.SetChecked(true)
		checkSpecial := ui.NewCheckbox("Спецсимволы (!@#...)")
		checkSpecial.SetChecked(true)

		checkboxGroup := ui.NewVerticalBox()
		checkboxGroup.Append(checkUpper, false)
		checkboxGroup.Append(checkDigits, false)
		checkboxGroup.Append(checkSpecial, false)

		optionsGroup := ui.NewGroup("Настройки символов")
		optionsGroup.SetChild(checkboxGroup)

		passwordLabel := ui.NewEntry()
		passwordLabel.SetReadOnly(true)

		resultGroup := ui.NewGroup("Сгенерированный пароль")
		resultBox := ui.NewVerticalBox()
		resultBox.Append(passwordLabel, false)
		resultGroup.SetChild(resultBox)

		buttonGenerate := ui.NewButton("Сгенерировать")
		buttonCopy := ui.NewButton("Скопировать")

		buttons := ui.NewHorizontalBox()
		buttons.Append(buttonGenerate, false)
		buttons.Append(buttonCopy, false)

		buttonGenerate.OnClicked(func(*ui.Button) {
			length, err := strconv.Atoi(lengthEntry.Text())
			if err != nil || length <= 0 {
				passwordLabel.SetText("Ошибка: введите число > 0")
				return
			}
			password := passgen(length, checkUpper.Checked(), checkDigits.Checked(), checkSpecial.Checked())
			passwordLabel.SetText(password)
		})

		buttonCopy.OnClicked(func(*ui.Button) {
			text := passwordLabel.Text()
			if text != "" {
				ctp(text)
			}
		})

		//main Window
		mainBox := ui.NewVerticalBox()
		mainBox.SetPadded(true)
		mainBox.Append(lengthBox, false)
		mainBox.Append(optionsGroup, false)
		mainBox.Append(buttons, false)
		mainBox.Append(resultGroup, false)

		window.SetChild(mainBox)
		window.Show()
	})
	if err != nil {
		log.Fatal(err)
	}
}
