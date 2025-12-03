    package main

    import (
        "fmt"
        "log"
        "os"

        "github.com/gin-gonic/gin"
        "github.com/joho/godotenv"

        "github.com/thomasfsr/fitness-app/internal/infrastructure/db"
        repo "github.com/thomasfsr/fitness-app/internal/infrastructure/repository"
        "github.com/thomasfsr/fitness-app/internal/interface/http"
        "github.com/thomasfsr/fitness-app/internal/usecase"
    )

    func main() {
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file loaded, relying on environment")
        }

        d, err := db.NewGormDBFromEnv()
        if err != nil {
            log.Fatalf("failed to connect to db: %v", err)
        }

        // Auto-migrate ORM models
        if err := d.AutoMigrate(&repo.UserModel{}, &repo.WorkoutModel{}, &repo.MessageModel{}); err != nil {
            log.Fatalf("auto migrate failed: %v", err)
        }

        // repositories
        userRepo := repo.NewUserGormRepository(d)
        workoutRepo := repo.NewWorkoutGormRepository(d)
        messageRepo := repo.NewMessageGormRepository(d)

        // usecases
        userUC := usecase.NewUserUseCase(userRepo)
        workoutUC := usecase.NewWorkoutUseCase(workoutRepo, userRepo)
        messageUC := usecase.NewMessageUseCase(messageRepo, userRepo)

        // router
        r := gin.Default()
        api := r.Group("/api")
        {
            http.NewUserHandler(api, userUC)
            http.NewWorkoutHandler(api, workoutUC)
            http.NewMessageHandler(api, messageUC)
        }

        // WhatsApp webhook (example)
        http.NewWhatsAppHandler(r, messageUC)

        port := os.Getenv("PORT")
        if port == "" {
            port = "8080"
        }
        fmt.Printf("listening on :%s", port)
        if err := r.Run(":" + port); err != nil {
            log.Fatalf("server error: %v", err)
        }
    }
