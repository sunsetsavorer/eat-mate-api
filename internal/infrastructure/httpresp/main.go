package httpresp

import (
	"fmt"
	"net/http"

	"github.com/sunsetsavorer/eat-mate-api/internal/exceptions"
)

func GetError(err error) (int, any) {

	if _, ok := err.(*exceptions.ManyRequestsError); ok {
		return http.StatusTooManyRequests, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.ValidationError); ok {
		return http.StatusUnprocessableEntity, ErrorResp[map[string]string]{
			Errors: err.(*exceptions.ValidationError).Errors(),
		}
	}

	if _, ok := err.(*exceptions.NotFoundError); ok {
		return http.StatusNotFound, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.BadRequestError); ok {
		return http.StatusBadRequest, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	return http.StatusBadRequest, fmt.Errorf("unknow response")
}
