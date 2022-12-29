package main

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := sessionConfig()
	return sessions.Sessions(keyPairs, store)
}

// 配置session参数
func sessionConfig() sessions.Store {
	maxAge := 3600
	secret := "max"
	// var store sessions.Store
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		MaxAge: maxAge,
		Path:   "/",
	})
	return store
}

// 图片生成器
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	} else if len(length) == 2 {
		l = length[1]
	} else if len(length) == 3 {
		l = length[2]
	}
	captchaID := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaID)
	_ = session.Save()
	_ = server(c.Writer, c.Request, captchaID, ".png", "zh", false, w, h)
}

func server(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "nocache")
	w.Header().Set("Expires", "0")

	var cont bytes.Buffer

	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&cont, id, width, height)
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(cont.Bytes()))
	return nil
}

func captchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./*.html")
	r.Use(Session("max"))
	r.GET("/captcha", func(ctx *gin.Context) {
		Captcha(ctx, 4)
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/captcha/verify/:value", func(ctx *gin.Context) {
		value := ctx.Param("value")
		if captchaVerify(ctx, value) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "suc",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "fai",
			})
		}
	})
	r.Run(":9090")
}
