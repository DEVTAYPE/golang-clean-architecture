package handlers

import (
	"api-basico-dev/models"
	"api-basico-dev/server"
	"api-basico-dev/services"
	"net/http"
	"strconv"
)

type PostHandler struct {
	service *services.PostService //TODO: usar una interfaz en lugar de una estructura concreta para el servicio de posts
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

// Aquí puedes agregar métodos para manejar las solicitudes relacionadas con los posts, como crear un post, obtener posts, actualizar un post, eliminar un post, etc.
func (h *PostHandler) CreatePostHandler(ctx *server.Context) {

	var req models.CreatePostRequest
	userID := ctx.GetUserUID()

	if err := ctx.BindJSON(&req); err != nil {
		RespondError(ctx, NewAppError("Datos de solicitud inválidos", 400))
		return
	}

	if req.Title == "" || req.Content == "" {
		RespondError(ctx, NewAppError("Todos los campos son obligatorios", 400))
		return
	}

	post := &models.Post{
		UserId:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.service.CreatePost(ctx.Ctx, post); err != nil {
		RespondError(ctx, NewAppError("Error al crear el post", 500))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post creado exitosamente",
		"post": map[string]interface{}{
			"id":         post.ID,
			"title":      post.Title,
			"content":    post.Content,
			"user_id":    post.UserId,
			"created_at": post.CreatedAt,
			"updated_at": post.UpdatedAt,
		},
	})
}

func (h *PostHandler) GetAllPostsHandler(ctx *server.Context) {
	posts, err := h.service.GetAllPosts(ctx.Ctx)

	if err != nil {
		RespondError(ctx, NewAppError("Error al obtener los posts", 500))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Posts obtenidos exitosamente",
		"posts":   posts,
	})
}

func (h *PostHandler) GetPostByIDHandler(ctx *server.Context) {

	postIDStr := ctx.GetParam("id")
	postIDInt, err := strconv.Atoi(postIDStr)
	if err != nil {
		RespondError(ctx, NewAppError("ID de post inválido", 400))
		return
	}

	postID := uint(postIDInt)

	if err != nil {
		RespondError(ctx, NewAppError("ID de post inválido", 400))
		return
	}

	post, err := h.service.GetPostByID(ctx.Ctx, postID)

	if err != nil {
		RespondError(ctx, NewAppError("Error al obtener el post", 500))
		return
	}

	if post == nil {
		RespondError(ctx, NewAppError("Post no encontrado", 404))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post obtenido exitosamente",
		"post":    post,
	})
}

func (h *PostHandler) GetPostsByUserIDHandler(ctx *server.Context) {

	userID := ctx.GetUserUID()

	posts, err := h.service.GetAllPostsByUserID(ctx.Ctx, userID)

	if err != nil {
		RespondError(ctx, NewAppError("Error al obtener los posts del usuario", 500))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Posts del usuario obtenidos exitosamente",
		"posts":   posts,
	})
}

func (h *PostHandler) UpdatePostHandler(ctx *server.Context) {

	postIDStr := ctx.GetParam("id")
	postIDInt, err := strconv.Atoi(postIDStr)
	if err != nil {
		RespondError(ctx, NewAppError("ID de post inválido", 400))
		return
	}

	postID := uint(postIDInt)

	var req models.UpdatePostRequest

	if err := ctx.BindJSON(&req); err != nil {
		RespondError(ctx, NewAppError("Datos de solicitud inválidos", 400))
		return
	}

	if req.Title == "" || req.Content == "" {
		RespondError(ctx, NewAppError("Todos los campos son obligatorios", 400))
		return
	}

	post := &models.Post{
		ID:      postID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := h.service.UpdatePost(ctx.Ctx, post); err != nil {
		RespondError(ctx, NewAppError("Error al actualizar el post", 500))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post actualizado exitosamente",
		"post": map[string]interface{}{
			"id":         post.ID,
			"title":      post.Title,
			"content":    post.Content,
			"user_id":    post.UserId,
			"created_at": post.CreatedAt,
			"updated_at": post.UpdatedAt,
		},
	})
}

func (h *PostHandler) DeletePostHandler(ctx *server.Context) {

	postIDStr := ctx.GetParam("id")
	postIDInt, err := strconv.Atoi(postIDStr)
	if err != nil {
		RespondError(ctx, NewAppError("ID de post inválido", 400))
		return
	}

	postID := uint(postIDInt)

	if err := h.service.DeletePost(ctx.Ctx, postID); err != nil {
		RespondError(ctx, NewAppError("Error al eliminar el post", 500))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post eliminado exitosamente",
	})
}
