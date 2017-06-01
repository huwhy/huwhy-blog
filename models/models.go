package models

import "time"

type User struct {
	Id       int
	Username string
	Password string
}

type Article struct {
	Id        int
	FirstCid  int
	SecondCid int
	ThirdCid  int
	Title     string
	Tags      string
	Summary   string
	Content   string
	Status    string
	ImgUrl    string
	Author    string
	Comments  int
	Views     int
	Updated   time.Time
	Created   time.Time
}

type Seo struct {
	Id          int
	SeoKey      string
	TargetId    int
	Title       string
	Description string
	Keywords    string
	Updated     time.Time
	Created     time.Time
}

type Category struct {
	Id      int
	Pid     int
	Name    string
	Parent  bool
	level   int
	Updated time.Time
	Created time.Time
	//Children []Category
}

type Link struct {
	Id   int
	Name string
	Url  string
}
