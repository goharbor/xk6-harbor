package module

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/loadimpact/k6/js/common"
	"github.com/opencontainers/go-digest"
)

func getEnv(key string, defaults ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	panic(fmt.Errorf("%s envirument is required", key))
}

func writeBlob(rootPath string, data []byte) (digest.Digest, error) {
	dgt := digest.FromBytes(data)

	dir := path.Join(rootPath, "blobs", dgt.Algorithm().String())

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	filename := path.Join(dir, dgt.Hex())
	if err := ioutil.WriteFile(filename, data, 0664); err != nil {
		return "", err
	}

	return dgt, nil
}

func getErrors(i interface{}) *models.Errors {
	if v, ok := i.(interface {
		GetPayload() *models.Errors
	}); ok {
		return v.GetPayload()
	}

	return nil
}

func getErrorMessage(err error) string {
	if errs := getErrors(err); errs != nil && len(errs.Errors) > 0 {
		return errs.Errors[0].Message
	}

	return err.Error()
}

func Check(ctx context.Context, err error) {
	if err == nil {
		return
	}

	common.Throw(common.GetRuntime(ctx), errors.New(getErrorMessage(err)))
}

func Checkf(ctx context.Context, err error, format string, a ...interface{}) {
	if err == nil {
		return
	}

	common.Throw(
		common.GetRuntime(ctx),
		fmt.Errorf("%s, error: %s", fmt.Sprintf(format, a...), getErrorMessage(err)),
	)
}

func Throwf(ctx context.Context, format string, a ...interface{}) {
	common.Throw(common.GetRuntime(ctx), fmt.Errorf(format, a...))
}

func IDFromLocation(ctx context.Context, loc string) int64 {
	parts := strings.Split(loc, "/")

	id, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	Check(ctx, err)

	return id
}

func NameFromLocation(ctx context.Context, loc string) string {
	parts := strings.Split(loc, "/")

	return parts[len(parts)-1]
}
