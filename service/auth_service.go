package service

import (
	"myauth/ent"
	"myauth/repository"
	"myauth/util"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	authRepository *repository.AuthRepository
}

func (s *AuthService) RegisterUser(ctx *gin.Context) {
	// TODO: insert daba to database
	// client := ctx.MustGet("client").(*ent.Client)
	// client := ent.NewClient()
	// bg := context.Background()

	// tx, err := client.Tx(bg)
	// if err != nil {
	// 	ctx.JSON(500, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	user := ent.User{}
	user.Username = "test"
	user.Password = util.CreateHash("test")
	user.Email = "test@test.com"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := s.authRepository.RegisterUser(&user).Save(ctx)
	if err != nil {
		// tx.Rollback()
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, "success")
}
