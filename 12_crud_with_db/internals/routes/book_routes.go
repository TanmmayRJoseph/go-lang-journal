package bookroutes

import (
	"book_store/internals/handlers"
	"book_store/internals/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterBookRoutes(r *gin.Engine, db *mongo.Database) *gin.Engine {

	// creating repo and handler once on startup
	repo := repository.NewRepo(db)
	h := handlers.NewHandler(repo)

	// Book Group /books
	bookGrp := r.Group("/books")
	{

		bookGrp.POST("/create", h.CreateBook)
		bookGrp.GET("/all", h.GetAllBook)
		bookGrp.GET("/:id", h.GetBookById)
		bookGrp.PUT("/:id", h.UpdateBookById)
		bookGrp.DELETE("/:id", h.DeleteBookById)
	}

	return r
}
