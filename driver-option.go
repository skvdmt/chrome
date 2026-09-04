package chrome

import "github.com/skvdmt/chrome/internal/model"

// driverOption Функциональная опция драйвера.
type driverOption func(d *Driver) error

// WithPath Указание пути к google-chrome.
func WithPath(filename string) driverOption {
	return func(d *Driver) error {
		d.path = filename
		return nil
	}
}

// Указание аргументов из файла.
func WithArgsFromFile(filename string) driverOption {
	return func(d *Driver) error {
		a, err := model.NewArgsFromFile(filename)
		if err != nil {
			return err
		}
		d.args = a
		return nil
	}
}

// Указание аргумента и при необходимости его значения.
func SetArg(name string, value ...string) driverOption {
	return func(d *Driver) error {
		if len(value) > 0 {
			d.args.Set(name, value[0])
			return nil
		}
		d.args.Set(name)
		return nil
	}
}

// WithDebug Включение отладки.
func WithDebug() driverOption {
	return func(d *Driver) error {
		d.debug.Enable()
		return nil
	}
}

// WithRemoveUserDataDirAfterClose Удаление пользовательской
// дириктории chrome после закрытия.
func WithRemoveUserDataDirAfterClose() driverOption {
	return func(d *Driver) error {
		d.removeUserDataDirAfterClose = true
		return nil
	}
}
