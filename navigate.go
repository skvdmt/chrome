package chrome

import "github.com/skvdmt/chrome/internal/devtools/options/page/navigate"

// Navigate Упрощенная сигнатура навигации без возврата лишней информации.
func (d *Driver) Navigate(u string, options ...navigate.Option) error {
	if _, _, _, _, err := d.Page.Navigate(u, options...); err != nil {
		return err
	}
	return nil
}
