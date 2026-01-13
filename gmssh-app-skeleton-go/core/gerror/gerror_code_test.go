package gerror

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	reflect.DeepEqual(CodeNil.String(), "-1")
	reflect.DeepEqual(CodeInternalError.String(), "200050:Internal Error")
	reflect.DeepEqual(CodeInternalError.Message(), "Internal Error")
	reflect.DeepEqual(CodeInternalError.Code(), 200050)
}
