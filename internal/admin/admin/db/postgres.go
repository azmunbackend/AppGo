package user

import (
	"context"
	"fmt"
	adminlogin "test/internal/admin/admin"
	"test/pkg/client/postgresql"
	"test/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client,  logger *logging.Logger) adminlogin.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) Login(ctx context.Context, username string) (adminlogin.Login, error){
	
	var result adminlogin.Login
	q:= `select password, username from admin where username =$1;`
	err := r.client.QueryRow(ctx, q, username).Scan(&result.Password, &result.UserName)

	if err != nil  {
		fmt.Println("LOGIN  postgres err:   ", err)
		return result, err
	}
	return result, nil
}