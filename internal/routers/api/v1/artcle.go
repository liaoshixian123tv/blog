package v1

import "github.com/gin-gonic/gin"

type Artcle struct{}

func NewArtcle() Artcle {
	return Artcle{}
}
func (t Artcle) Get(c *gin.Context)    {}
func (t Artcle) List(c *gin.Context)   {}
func (t Artcle) Create(c *gin.Context) {}
func (t Artcle) Update(c *gin.Context) {}
func (t Artcle) Delete(c *gin.Context) {}
