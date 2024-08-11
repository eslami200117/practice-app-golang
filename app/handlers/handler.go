package handler

import "rest.gtld.test/realTimeApp/app/model"

type handler interface {
	GetCurrenct(username string,user *model.Login)
}