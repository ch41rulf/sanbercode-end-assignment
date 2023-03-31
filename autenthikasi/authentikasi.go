package autenthikasi

import "github.com/gin-gonic/gin"

func AuthMiddleware(users map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, pwd, ok := c.Request.BasicAuth()
		if !ok {
			c.Writer.Write([]byte("Username atau Password tidak boleh kosong"))
			c.Abort()
			return
		}

		if val, found := users[uname]; found && val == pwd {
			c.Next()
			return
		}

		c.Writer.Write([]byte("Username atau Password tidak sesuai"))
		c.Abort()
		return
	}
}
