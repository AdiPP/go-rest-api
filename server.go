package main

import (
	"fmt"
	"net/http"

	"github.com/AdiPP/go-rest-api/controller"
	router "github.com/AdiPP/go-rest-api/http"
	"github.com/AdiPP/go-rest-api/repository"
	"github.com/AdiPP/go-rest-api/service"
)

var (
	httpRouter router.Router = router.NewChiRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService service.PostService = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main () {
	const port string = ":8000"

	httpRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and running")
	})

	httpRouter.Get("/posts", postController.GetPosts)

	httpRouter.Post("/posts", postController.CreatePost)

	httpRouter.Serve(port)
}