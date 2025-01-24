package rest

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	openapiErrors "github.com/go-openapi/errors"
	"net/http"
	"strings"

	"github.com/dimuska139/urlshortener/internal/service/server/rest/gen/models"
	"github.com/dimuska139/urlshortener/pkg/logging"
)

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func flattenComposite(errs *openapiErrors.CompositeError) *openapiErrors.CompositeError {
	var res []error

	for _, er := range errs.Errors {
		var compositeError *openapiErrors.CompositeError

		if errors.As(er, &compositeError) && len(compositeError.Errors) > 0 {
			flat := flattenComposite(compositeError)

			if len(flat.Errors) > 0 {
				res = append(res, flat.Errors...)
			}
		} else {
			res = append(res, er)
		}
	}

	return openapiErrors.CompositeValidationError(res...)
}

func getValidationErrors(errs *openapiErrors.CompositeError) map[string]string {
	errsMap := make(map[string]string)

	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *openapiErrors.ParseError:
			errsMap["common"] = "Invalid JSON format."

		case *openapiErrors.Validation:
			errKey := strings.ReplaceAll(e.Name, "body.", "")
			msg := e.Error()

			if strings.Contains(msg, "must be of type int") {
				msg = "This value should be of type integer."
			}

			if strings.Contains(msg, "in query is required") {
				msg = "This query parameter is missing."
			}

			if strings.Contains(msg, "in body is required") {
				msg = "This parameter is missing."
			}

			errsMap[errKey] = msg

		default:
			if len(errs.Errors) > 0 {
				errsMap["common"] = errs.Errors[0].Error()
			}
		}
	}

	return errsMap
}

func handleErrors(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	internalErrorResponse := &models.InternalError{
		Common: "Something went wrong.",
	}

	switch e := err.(type) {
	case *openapiErrors.CompositeError:
		er := flattenComposite(e)
		if len(er.Errors) > 0 {
			resp := ErrorResponse{
				Errors: getValidationErrors(er),
			}

			encoded, _ := json.Marshal(resp)

			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write(encoded)
		} else {
			handleErrors(rw, r, nil)
		}

	case *openapiErrors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(err.(*openapiErrors.MethodNotAllowedError).Allowed, ","))
		rw.WriteHeader(int(e.Code()))

		if r == nil || r.Method != http.MethodHead {
			resp := models.InternalError{
				Common: "Method not allowed.",
			}

			encoded, _ := json.Marshal(resp)

			_, _ = rw.Write(encoded)
		}

	case *jwt.ValidationError:
		rw.WriteHeader(http.StatusUnauthorized)

		encoded, _ := json.Marshal(models.InternalError{
			Common: "Authorization required.",
		})

		_, _ = rw.Write(encoded)

		return

	case openapiErrors.Error:
		switch e.Code() {
		case http.StatusUnauthorized:
			rw.WriteHeader(int(e.Code()))

			encoded, _ := json.Marshal(models.InternalError{
				Common: "Authorization required.",
			})

			_, _ = rw.Write(encoded)

		case http.StatusForbidden:
			rw.WriteHeader(int(e.Code()))

			encoded, _ := json.Marshal(models.InternalError{
				Common: e.Error(),
			})

			_, _ = rw.Write(encoded)

		case http.StatusNotFound:
			rw.WriteHeader(int(e.Code()))

			encoded, _ := json.Marshal(models.InternalError{
				Common: "Not found.",
			})

			_, _ = rw.Write(encoded)

		default:
			logging.Error(ctx, "internal server error",
				"err", err.Error())
			rw.WriteHeader(http.StatusInternalServerError)

			encoded, _ := json.Marshal(internalErrorResponse)

			_, _ = rw.Write(encoded)
		}

		return

	case nil:
		rw.WriteHeader(http.StatusInternalServerError)

		encoded, _ := json.Marshal(internalErrorResponse)
		_, _ = rw.Write(encoded)

	default:
		rw.WriteHeader(http.StatusInternalServerError)

		if r == nil || r.Method != http.MethodHead {
			encoded, _ := json.Marshal(internalErrorResponse)
			_, _ = rw.Write(encoded)
		}
	}
}
