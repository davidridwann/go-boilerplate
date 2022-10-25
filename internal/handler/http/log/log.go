package logHandler

import (
	"github.com/davidridwann/wlb-test.git/pkg/helpers"
	mongoConfig "github.com/davidridwann/wlb-test.git/pkg/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var logCollection *mongo.Collection = mongoConfig.GetCollection(mongoConfig.DB, "log")

// Get Log       godoc
// @Summary      Log activity
// @Description  Log activity
// @Tags         Log
// @Produce      json
// @Success      200  {object} map[string]interface{}
// @Router       /log [get]
func Get(c *gin.Context) {
	cursor, err := logCollection.Find(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var logs []bson.M
	if err = cursor.All(c, &logs); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{Message: "Log error"})
	}

	c.JSON(http.StatusOK, helpers.SuccessResponse{Data: logs, Message: "Successfully get log"})
}
