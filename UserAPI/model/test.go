package model

import "testing"

func TestSetSession(t *testing.T) {
	s := &Session{
		SessionID: "20180916153001",
		UserType:  "admin",
		NickName:  "df",//
	}
	ss := SessionService{}
	err := ss.SetSession(s)
	if err != nil {
		t.Errorf("fail to add one session(%+v): %s", s, err)
		t.FailNow()
	}
}

func TestDelSession(t *testing.T){
	var key = "20180916153001"
	ss := SessionService{}
	err := ss.DelSession(key)
	if err != nil {
		t.Errorf("fail to delete one session(%s):%s",key,err)
		t.FailNow()
	}
}

func TestGetSession(t *testing.T){
	var key = "20180916153001"
	ss := SessionService{}
	value,err := ss.GetSession(key)
	if err != nil {
		t.Errorf("fail to delete one session(%s):%s",key,err)
		t.FailNow()
	}
	t.Log("get session(%s):%s",key,value)
}