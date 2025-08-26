// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"github.com/stainless-sdks/serval-go/option"
)

// EntitlementService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementService] method instead.
type EntitlementService struct {
	Options []option.RequestOption
}

// NewEntitlementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntitlementService(opts ...option.RequestOption) (r EntitlementService) {
	r = EntitlementService{}
	r.Options = opts
	return
}
