package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"main/config"
	ent "main/internal/auth/domain/entities"
	"main/internal/auth/domain/ports"
	"main/pkg/logger"
	cm "main/pkg/utils/common"
	"main/pkg/utils/jwt"
	"main/pkg/utils/validator"
)

// handlerHttp Auth handlers
type handlerHttp struct {
	ctx     context.Context
	cfg     *config.Config
	service ports.IService
	logger  logger.Logger
}

var (
	Responser  fiber.Map
	StatusCode int
)

// NewHttpHandler Auth Domain HTTP handlers constructor
func NewHttpHandler(ctx context.Context, cfg *config.Config, service ports.IService, logger logger.Logger) ports.IHandlers {
	return &handlerHttp{ctx: ctx, cfg: cfg, service: service, logger: logger}
}

// Create godoc
// @Summary Register
// @Description Registration process
// @Tags Auth
// @Param Body body entities.RegisterReq true "`Body for user registration`"
// @Accept json
// @Produce json
// @Success 200 {object} entities.HandlerResponse{}
// @Router /auth/register [post]
func (h handlerHttp) Register(c *fiber.Ctx) error {
	dat := ent.RegisterReq{}

	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = cm.HTTPResponser(nil, StatusCode, true, errParser.Error())
	}

	errValidate := validator.ValidateStruct(c.Context(), &dat)
	if errValidate != nil {
		StatusCode = fiber.StatusBadRequest
		Responser = cm.HTTPResponser(nil, StatusCode, true, errValidate.Error())
	}

	if errParser == nil && errValidate == nil {
		record := h.service.Register(h.ctx, dat)

		if record == -3 {
			StatusCode = fiber.StatusConflict
			Responser = cm.HTTPResponser(nil, fiber.StatusConflict, true, "Bu kullanıcı adı zaten kayıtlı")
		} else if record == -2 {
			StatusCode = fiber.StatusConflict
			Responser = cm.HTTPResponser(nil, fiber.StatusConflict, true, "Girdiğiniz kod eşleşmedi")
		} else if record == -1 {
			StatusCode = fiber.StatusInternalServerError
			Responser = cm.HTTPResponser(nil, StatusCode, true, "Sistemde Hata Oluştu")
		} else if record == 0 {
			StatusCode = fiber.StatusNotFound
			Responser = cm.HTTPResponser(nil, StatusCode, true, "Girdiğiniz doğrulama kodunun süresi dolmuş")
		} else {
			StatusCode = fiber.StatusOK
			Responser = cm.HTTPResponser(nil, StatusCode, false, "İşlem Başarılı")
		}
	}

	return c.Status(StatusCode).JSON(Responser)
}

// Create godoc
// @Summary Login
// @Description Login process
// @Tags Auth
// @Param Body body entities.LoginReq true "`Body for user registration`"
// @Accept json
// @Produce json
// @Success 200 {object} entities.HandlerResponse{}
// @Router /auth/login [post]
func (h handlerHttp) Login(c *fiber.Ctx) error {
	dat := ent.LoginReq{}

	errParser := c.BodyParser(&dat)
	if errParser != nil {
		StatusCode = fiber.StatusBadGateway
		Responser = cm.HTTPResponser(nil, StatusCode, true, errParser.Error())
	}

	errValidate := validator.ValidateStruct(c.Context(), &dat)
	if errValidate != nil {
		StatusCode = fiber.StatusBadRequest
		Responser = cm.HTTPResponser(nil, StatusCode, true, errValidate.Error())
	}

	if errParser == nil && errValidate == nil {
		record, claims := h.service.Login(h.ctx, dat)
		if record == -1 {
			StatusCode = fiber.StatusInternalServerError
			Responser = cm.HTTPResponser(nil, StatusCode, true, "Sistemde Hata Oluştu")
		} else if record == 0 {
			StatusCode = fiber.StatusOK
			Responser = cm.HTTPResponser(nil, StatusCode, true, "Kayıt Bulunamadı")
		} else {
			claim := jwt.TokenClaim{
				AuthId:      claims.AuthId,
				Lang:        claims.Lang,
				ParentId:    claims.ParentId,
				ManagerId:   claims.ManagerId,
				AccountType: claims.AccountType,
				UserType:    claims.UserType,
				CompanyName: claims.CompanyName,
				UserTitle:   claims.UserTitle,
				UserName:    claims.UserName,
				IsDemo:      claims.IsDemo,
				UniqueId:    claims.UniqueId,
				Status:      claims.Status,
				Exp:         0,
				Iat:         0,
			}
			Token, Error := jwt.GenerateToken(h.cfg, claim)
			if Error != nil {
				StatusCode = fiber.StatusConflict
				Responser = cm.HTTPResponser(nil, StatusCode, true, Token)
			} else {
				StatusCode = fiber.StatusOK
				Responser = cm.HTTPResponser(nil, StatusCode, false, Token)
			}
		}
	}

	return c.Status(StatusCode).JSON(Responser)

}
