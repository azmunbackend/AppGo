package adminlogin

import (
	"context"
	"net/http"
	"test/internal/appresult"
	"test/internal/config"
	"test/internal/handlers"
	"test/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	URL = "/login"
)

type handler struct {
	repository Repository
	logger     *logging.Logger
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h handler) Register(router *gin.RouterGroup) {
	router.POST(URL, h.Login)
}

// CreateAccount creates a new account
//
// @Summary      Create account
// @Description  Create a new account
// @Tags         accounts
// @Accept       json
// @Produce      json    
// @Param 		request body Login true "Create news"
// @Success      200
// @Router       /api/v1/admin/login [post]
func (h handler) Login(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		h.logger.Error("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	data, err := h.repository.Login(context.TODO(), login.UserName)
	if err != nil {
		h.logger.Error("Error fetching user data: ", err)
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	if data.UserName == "" {
		h.logger.Error("Username not found")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or password not correct"})
		return
	}

	if login.Password != "" && data.Password != "" {
		err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(login.Password))
		if err != nil {
			h.logger.Error("Password not correct: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password not correct"})
			return
		}

		atClaims := jwt.MapClaims{}
		atClaims["exp"] = time.Now().Add(time.Minute * 60 * 12).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

		cfg := config.GetConfig()
		tokenDTO, err := at.SignedString([]byte(cfg.JwtKey))
		if err != nil {
			h.logger.Error("Token cannot be generated: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token cannot be generated"})
			return
		}

		successResult := appresult.Success
		successResult.Data = ResultTokenDTO{
			Token: tokenDTO,
		}

		c.JSON(http.StatusOK, successResult)
	} else {
		c.JSON(http.StatusUnauthorized, appresult.ErrNotData)
	}
}
