package db

import "gorm.io/gorm"

type ParentReply struct {
	gorm.Model
	Content   string
	Like      int64
	ReplyTime int64
	Rpid      int64
}
type ChildReply struct {
	gorm.Model
	Content   string
	Like      int64
	ReplyTime int64
	Rpid      int64
	ParentID  int64
}

func CreateParentReply(reply *ParentReply) error {
	if err := DB.Create(reply).Error; err != nil {
		return err
	}
	return nil
}

func CreateChildReply(reply *ChildReply) error {
	if err := DB.Create(reply).Error; err != nil {
		return err
	}
	return nil
}
