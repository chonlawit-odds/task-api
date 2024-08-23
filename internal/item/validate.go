package item

import (
	"strings"

	"github.com/chonlawit-odds/task-api/internal/constant"
	"github.com/pkg/errors"
)

type Validate struct {
}

func NewValidate() Validate {
	return Validate{}
}

func (validator Validate) ItemStatusFlow(current, next constant.ItemStatus) error {
	switch current {
	case constant.ItemPendingStatus:
		if next == constant.ItemPendingStatus {
			return errors.New("cannot change pending status to pending status")
		}

	default:
		return errors.Errorf(
			"cannot change %s status to %s status",
			strings.ToLower(string(current)),
			strings.ToLower(string(next)),
		)

	}

	return nil
}

func (validator Validate) UpdateItem(current constant.ItemStatus) error {
	switch current {
	case constant.ItemApprovedStatus, constant.ItemRejectedStatus:
		return errors.Errorf("cannot update item when status is %s", strings.ToLower(string(current)))

	}

	return nil
}

func (validator Validate) DeleteItem(current constant.ItemStatus) error {
	switch current {
	case constant.ItemApprovedStatus, constant.ItemRejectedStatus:
		return errors.Errorf("cannot delete item when status is %s", strings.ToLower(string(current)))

	}

	return nil
}
