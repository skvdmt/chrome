package chrome

import (
	"github.com/skvdmt/chrome/internal/model"
	"github.com/skvdmt/chrome/internal/options/select_target"
)

// SelectTarget Переключение вкладки.
func (d *Driver) SelectTarget(options ...select_target.Option) error {
	for _, o := range options {
		if tid, f := o(d.targets); f {
			sid := d.sessions[*tid]
			d.updateCurrent(&sid)
			return nil
		}
	}
	return model.ERR_TARGET_NOT_FOUND
}
