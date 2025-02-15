package api

import (
	"database/sql"

	"github.com/MahatVasudev/liveStreamingProject/server/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

//
// var (
// 	addr = flag.String("addr", ":", os.Getenv("PORT"))
// 	cert = flag.String("cert", "", "")
// 	key  = flag.String("key", "", "")
// )

type APIServer struct {
	Addr  string
	sqlDB *sql.DB
	redDB *redis.Client
}

func NewApiServer(Addr string, sqlDB *sql.DB, redDB *redis.Client) *APIServer {
	return &APIServer{
		Addr,
		sqlDB,
		redDB,
	}
}

func (newapi *APIServer) Run() error {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	// Stores
	userStore := user.NewStore(newapi.sqlDB, newapi.redDB)

	// Handlers
	userHandler := user.NewHandler(userStore)

	// Register Routes
	userHandler.RegisterRoutes("/user", app)

	return app.Listen(newapi.Addr)
}
