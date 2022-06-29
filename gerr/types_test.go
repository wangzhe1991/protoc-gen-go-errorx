package gerr

import "testing"

func TestTypes(t *testing.T) {
	var (
		input = []error{
			BadRequest(100001, "domain", "reason_400"),
			Unauthorized(100001, "domain", "reason_401"),
			Forbidden(100001, "domain", "reason_403"),
			NotFound(100001, "domain", "reason_404"),
			Conflict(100001, "domain", "reason_409"),
			InternalServer(100001, "domain", "reason_500"),
			ServiceUnavailable(100001, "domain", "reason_503"),
		}
		output = []func(error) bool{
			IsBadRequest,
			IsUnauthorized,
			IsForbidden,
			IsNotFound,
			IsConflict,
			IsInternalServer,
			IsServiceUnavailable,
		}
	)

	for i, in := range input {
		if !output[i](in) {
			t.Errorf("not expect: %v", in)
		}
	}
}
