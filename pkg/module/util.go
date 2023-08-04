package module

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/content/local"
	"github.com/dop251/goja"
	"github.com/opencontainers/go-digest"
	"go.k6.io/k6/js/common"
)

func newLocalStore(rt *goja.Runtime, name string) (string, content.Store) {
	rootPath := filepath.Join(DefaultRootPath, name)

	store, err := local.NewStore(rootPath)
	Check(rt, err)

	return rootPath, store
}

func isDigest(reference string) bool {
	i := strings.Index(reference, ":")
	return i > 0 && i+1 != len(reference)
}

func getDistrubtionRef(projectName, repositoryName, reference string) string {
	if isDigest(reference) {
		return fmt.Sprintf("%s/%s@%s", projectName, repositoryName, reference)
	}

	return fmt.Sprintf("%s/%s:%s", projectName, repositoryName, reference)
}

func writeBlob(rootPath string, data []byte) (digest.Digest, error) {
	dgt := digest.FromBytes(data)

	dir := path.Join(rootPath, "blobs", dgt.Algorithm().String())

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	filename := path.Join(dir, dgt.Hex())
	if err := os.WriteFile(filename, data, 0664); err != nil {
		return "", err
	}

	return dgt, nil
}

// func getErrors(i interface{}) *models.Errors {
// 	if v, ok := i.(interface {
// 		GetPayload() *models.Errors
// 	}); ok {
// 		return v.GetPayload()
// 	}

// 	return nil
// }

// func getErrorMessage(err error) string {
// 	if errs := getErrors(err); errs != nil && len(errs.Errors) > 0 {
// 		return errs.Errors[0].Message
// 	}

// 	return err.Error()
// }

func Check(rt *goja.Runtime, err error) {
	if err == nil {
		return
	}

	common.Throw(rt, err)
}

func Checkf(rt *goja.Runtime, err error, format string, a ...interface{}) {
	if err == nil {
		return
	}

	common.Throw(
		rt,
		fmt.Errorf("%s, error: %s", fmt.Sprintf(format, a...), err),
	)
}

func Throwf(rt *goja.Runtime, format string, a ...interface{}) {
	common.Throw(rt, fmt.Errorf(format, a...))
}

func ExportTo(rt *goja.Runtime, target interface{}, args ...goja.Value) {
	if len(args) > 0 {
		if err := rt.ExportTo(args[0], target); err != nil {
			common.Throw(rt, err)
		}
	}
}

func IDFromLocation(rt *goja.Runtime, loc string) int64 {
	parts := strings.Split(loc, "/")

	id, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	Check(rt, err)

	return id
}

func NameFromLocation(loc string) string {
	parts := strings.Split(loc, "/")

	return parts[len(parts)-1]
}
