package http

import (
	bm "go-common/library/net/http/blademaster"
)

func listBusiness(c *bm.Context) {
	v := new(struct {
		State int32 `form:"state"`
	})
	err := c.Bind(v)
	if err != nil {
		return
	}
	res, err := rpSvc.ListBusiness(c, v.State)
	c.JSON(res, err)
}

func getBusiness(c *bm.Context) {
	v := new(struct {
		Type int32 `form:"type"`
	})
	err := c.Bind(v)
	if err != nil {
		return
	}
	res, err := rpSvc.GetBusiness(c, v.Type)
	c.JSON(res, err)
}

func addBusiness(c *bm.Context) {
	v := new(struct {
		Type   int32  `form:"type" validate:"required"`
		Name   string `form:"name" validate:"required"`
		Appkey string `form:"app_key" validate:"required"`
		Remark string `form:"remark" validate:"required"`
		Alias  string `form:"alias" validate:"required"`
	})
	var err error
	err = c.Bind(v)
	if err != nil {
		return
	}
	_, err = rpSvc.AddBusiness(c, v.Type, v.Name, v.Appkey, v.Remark, v.Alias)
	c.JSON(nil, err)
}

func upBusiness(c *bm.Context) {
	v := new(struct {
		Type   int32  `form:"type" validate:"required"`
		Name   string `form:"name" validate:"required"`
		Appkey string `form:"app_key" validate:"required"`
		Remark string `form:"remark" validate:"required"`
		Alias  string `form:"alias" validate:"required"`
	})
	var err error
	err = c.Bind(v)
	if err != nil {
		return
	}
	_, err = rpSvc.UpBusiness(c, v.Name, v.Appkey, v.Remark, v.Alias, v.Type)
	c.JSON(nil, err)
}

func upBusiState(c *bm.Context) {
	v := new(struct {
		State int32 `form:"state"`
		Type  int32 `form:"type" validate:"required"`
	})
	var err error
	err = c.Bind(v)
	if err != nil {
		return
	}
	_, err = rpSvc.UpBusinessState(c, v.State, v.Type)
	c.JSON(nil, err)
}
