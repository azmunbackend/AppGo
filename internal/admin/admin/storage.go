package adminlogin

import "context"

type Repository interface{
	Login(ctx context.Context, username string) (Login, error)
	
}