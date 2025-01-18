package errors

import (
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

// HandleBadRequest handles a bad request error and returns a status with the field violations.
func HandleBadRequest(err error) *errdetails.BadRequest {
	st := status.Convert(err)
	var allErrors []string
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		// If the error is a BadRequest, grab the field violations.
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				allErrors = append(allErrors, violation.GetDescription())
			}
		}
	}

	fieldErr := &errdetails.BadRequest_FieldViolation{
		Field:       "payment",
		Description: strings.Join(allErrors, "\n"),
	}

	// Create a BadRequest with the field violations.
	badReq := &errdetails.BadRequest{}
	badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
	return badReq
}
