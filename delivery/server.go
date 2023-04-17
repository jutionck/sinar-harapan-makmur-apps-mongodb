package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/go-sinar-makmur-mongodb/config"
	"github.com/jutionck/go-sinar-makmur-mongodb/delivery/controller"
	"github.com/jutionck/go-sinar-makmur-mongodb/manager"
	"log"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) initController() {
	controller.NewBrandController(s.engine, s.ucManager.BrandUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to serve config: %v", err)
	}

	infra, err := manager.NewInfraManager(cfg)
	fmt.Println("err:", err)
	if err != nil {
		log.Fatalf("failed to connecto infra: %v", err)
	}
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	r := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	return &Server{
		ucManager: uc,
		engine:    r,
		host:      host,
	}
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
