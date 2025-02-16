package api

import (
	"database/sql"

	"github.com/MahatVasudev/liveStreamingProject/server/services/room"
	"github.com/MahatVasudev/liveStreamingProject/server/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	cert  string
	key   string
}

func NewApiServer(
	Addr string,
	sqlDB *sql.DB,
	redDB *redis.Client,
	cert string,
	key string,
) *APIServer {
	return &APIServer{
		Addr,
		sqlDB,
		redDB,
		cert,
		key,
	}
}

func (newapi *APIServer) Run() error {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	// Stores
	userStore := user.NewStore(newapi.sqlDB, newapi.redDB)

	// Handlers
	userHandler := user.NewHandler(userStore)
	roomHandler := room.NewHandler()

	// Register Routes
	userHandler.RegisterRoutes("/user", app)
	roomHandler.RegisterRoutes("/room", app)

	// go dispatchKeyFrames()

	if newapi.cert != "" {
		return app.ListenTLS(newapi.Addr, newapi.cert, newapi.key)
	}

	return app.Listen(newapi.Addr)
}

// func dispatchKeyFrames() {
//   for range time.NewTicker(time.Second * 3).C {
//     for _, room := range w.Rooms {
//       room.Peers.DispatchKeyFrame()
//     }
//   }
// }
