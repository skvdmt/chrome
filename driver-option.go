package chrome

import "github.com/skvdmt/chrome/internal/model"

// DriverOption Функциональная опция драйвера.
type DriverOption func(d *Driver) error

// WithPath Указание пути к google-chrome.
func WithPath(filename string) DriverOption {
	return func(d *Driver) error {
		d.path = filename
		return nil
	}
}

// WithArgsFromFile Указание аргументов из файла.
func WithArgsFromFile(filename string) DriverOption {
	return func(d *Driver) error {
		a, err := model.NewArgsFromFile(filename)
		if err != nil {
			return err
		}
		d.args = a
		return nil
	}
}

// SetArg Указание аргумента и при необходимости его значения.
func SetArg(name string, value ...string) DriverOption {
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
func WithDebug() DriverOption {
	return func(d *Driver) error {
		d.debug.Enable()
		return nil
	}
}

// WithRemoveUserDataDirAfterClose Удаление пользовательской
// дириктории chrome после закрытия.
func WithRemoveUserDataDirAfterClose() DriverOption {
	return func(d *Driver) error {
		d.removeUserDataDirAfterClose = true
		return nil
	}
}
