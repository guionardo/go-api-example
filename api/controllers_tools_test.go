package api

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getQueryVars(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Params = []gin.Param{{Key: "registro", Value: "1"}}
	ctx.Request = httptest.NewRequest("GET", "/feiras/?regiao5=2&nome_feira=3", nil)

	t.Run("default", func(t *testing.T) {
		gotValues := getQueryVars(ctx, "registro", "regiao5", "nome_feira", "bairro")
		if len(gotValues) != 4 {
			t.Errorf("getQueryVars() = %v, want %v", gotValues, 4)
		}
		var registro, regiao5, nomeFeira, bairro string
		unpackVars(gotValues, &registro, &regiao5, &nomeFeira, &bairro)

		if (registro != "1") || (regiao5 != "2") || (nomeFeira != "3") || (bairro != "") {
			t.Errorf("getQueryVars() = %v, want %v", gotValues, "[1 2 3]")
		}

	})

}
