package module

import (
	"context"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/scanner"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateScanner(ctx context.Context, registration models.ScannerRegistrationReq) string {
	h.mustInitialized(ctx)

	params := operation.NewCreateScannerParams().WithRegistration(&registration)

	res, err := h.api.Scanner.CreateScanner(ctx, params)
	Checkf(ctx, err, "failed to create scanner %s", *registration.Name)

	return NameFromLocation(ctx, res.Location)
}

func (h *Harbor) SetScannerAsDefault(ctx context.Context, registrationID string) {
	h.mustInitialized(ctx)

	params := operation.NewSetScannerAsDefaultParams().
		WithRegistrationID(registrationID).
		WithPayload(&models.IsDefault{IsDefault: true})

	_, err := h.api.Scanner.SetScannerAsDefault(ctx, params)

	Checkf(ctx, err, "failed to set scanner %s as default", registrationID)
}
