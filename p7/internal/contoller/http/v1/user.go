package v1

import (
	"log"
	"net/http"
	"os"
	"time"

	"p7/internal/entity"
	"p7/internal/middleware"
	"p7/internal/usecase"
	"p7/internal/usecase/repo"
	"p7/pkg/postgres"
	"p7/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type userRoutes struct {
	t usecase.UserInterface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.UserInterface) {
	r := &userRoutes{t: t}

	h := handler.Group("/users")
	{
		public := h.Group("/")
		public.Use(middleware.NewRateLimiter(rate.Every(10*time.Second), 3).Middleware())
		{
			public.POST("/", r.RegisterUser)
			public.POST("/login", r.LoginUser)
		}

		protected := h.Group("/")
		protected.Use(utils.JWTAuthMiddleware())
		protected.Use(middleware.NewRateLimiter(rate.Every(10*time.Second), 5).Middleware())
		{
			protected.PATCH("promote/:id", middleware.RoleMiddleware("admin"), r.PromoteUser)
			protected.GET("/me", r.GetMe)
			protected.GET("/protected/hello", r.ProtectedFunc)
		}
	}
}

func (r *userRoutes) PromoteUser(c *gin.Context) {
	username := c.Param("id")
	log.Println(username)
	status, err := r.t.Promote(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"promoted": status})
}

func (r *userRoutes) GetMe(c *gin.Context) {
	userId := c.Value("userID")

	me, err := r.t.GetMe(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"me": me})
}

func (r *userRoutes) ProtectedFunc(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func StartServer() {
	r := gin.Default()

	r2 := r.Group("/api")

	dbUrl := os.Getenv("DB_URL")
	newUserRoutes(r2, usecase.NewUserUseCase(repo.NewUserRepo(postgres.NewPostgresConn(dbUrl))))

	log.Fatal(http.ListenAndServe(":8080", r))

}

func (r *userRoutes) RegisterUser(c *gin.Context) {
	var createUserDTO entity.CreateUserDTO
	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(createUserDTO.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := entity.User{
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
		Password: hashedPassword,
		Role:     createUserDTO.Role,
	}

	createdUser, sessionId, err := r.t.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "User registered successfully. Please check your email for verification code.",
		"session_id": sessionId,
		"user":       createdUser,
	})
}

func (r *userRoutes) LoginUser(c *gin.Context) {
	var input entity.LoginUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := r.t.LoginUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
