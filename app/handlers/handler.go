package handler

import "rest.gtld.test/realTimeApp/app/model"

type handler interface {
	GetCurrenctNode(username string,user *model.Login)
}