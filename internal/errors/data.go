package errors

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
)

func NewDuplicatedKeyError(format string, values ...interface{}) error {
	return errors.BadRequest("DUPLICATED_VALUE", fmt.Sprintf(format, values...))
}

func NewOperationForbiddenError(format string, values ...interface{}) error {
	return errors.Forbidden("FORBIDDEN_OPERATION", fmt.Sprintf(format, values...))
}

func NewDefaultResourceForbiddenError(format string, values ...interface{}) error {
	return errors.Forbidden("DEFAULT_RESOURCE", fmt.Sprintf(format, values...))
}

func NewResourceNotFoundError(format string, values ...interface{}) error {
	return errors.NotFound("RESOURCE_NOT_FOUND", fmt.Sprintf(format, values...))
}

func NewInvalidArgumentError(format string, values ...interface{}) error {
	return errors.BadRequest("INVALID_ARGUMENT", fmt.Sprintf(format, values...))
}
