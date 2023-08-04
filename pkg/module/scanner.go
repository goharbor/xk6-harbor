package module

import (
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/scanner"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateScanner(registration models.ScannerRegistrationReq) string {
	h.mustInitialized()

	params := operation.NewCreateScannerParams().WithRegistration(&registration)

	res, err := h.api.Scanner.CreateScanner(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to create scanner %s", *registration.Name)

	return NameFromLocation(res.Location)
}

func (h *Harbor) SetScannerAsDefault(registrationID string) {
	h.mustInitialized()

	params := operation.NewSetScannerAsDefaultParams().
		WithRegistrationID(registrationID).
		WithPayload(&models.IsDefault{IsDefault: true})

	_, err := h.api.Scanner.SetScannerAsDefault(h.vu.Context(), params)

	Checkf(h.vu.Runtime(), err, "failed to set scanner %s as default", registrationID)
}
