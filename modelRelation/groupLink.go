package modelRelation

import (
	"github.com/cayleygraph/cayley/quad"
	"relation-graph/graphRelation/createTriple/conf"
	"github.com/cayleygraph/cayley"
	"relation-graph/graphRelation/createTriple/modelBase"
)

//创建一个群邀请链接
type CreateGroupShareLink struct {
	modelBase.User
	modelBase.Group
	CreateTime int
}

//点击一个群邀请链接
type ClickGroupShareLink struct {
	modelBase.User
	modelBase.Group
	ClickTime int
}

//转为三元组
func (this CreateGroupShareLink) Quad() []quad.Quad {
	//用户创建了这个加群邀请链接
	createQuad := quad.Make(this.User.Id, modelBase.CreateGroupLink.String(), this.Group.Id, nil)
	//在什么时候创建的
	timeQuad := quad.Make(modelBase.UserId_GroupId(this.User.Id, this.Group.Id), modelBase.CreateTime.String(), this.CreateTime, nil)
	return []quad.Quad{createQuad, timeQuad}
}

//添加到cayley
func (this CreateGroupShareLink) AddCreateGroupShareLinkToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(this.Quad())
}

//添加到一个或者多个关系到cayley数据库
func AddCreateGroupShareLinkToCayley(cgsls ...CreateGroupShareLink) error {
	var quadSet []quad.Quad
	for _, cgsl := range cgsls {
		qs := cgsl.Quad()
		for _, q := range qs {
			quadSet = append(quadSet, q)
		}
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}

//转为四元组
func (this ClickGroupShareLink) Quad() []quad.Quad {
	//用户点击这个加群链接
	clickQuad := quad.Make(this.User.Id, modelBase.ClickGroupLink.String(), this.Group.Id, nil)
	//在什么时候点击这个邀请链接的
	timeQuad := quad.Make(modelBase.UserId_GroupId(this.User.Id, this.Group.Id), modelBase.ClickTime.String(), this.ClickTime, nil)
	return []quad.Quad{clickQuad, timeQuad}
}

//添加到cayley
func (this ClickGroupShareLink) AddClickGroupShareLinkToCayley() error {
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(this.Quad())
}

//将一个或者多个关系添加到cayley
func AddClickGroupShareLinkToCayley(cgsls ...ClickGroupShareLink) error {
	var quadSet []quad.Quad
	for _, cgsl := range cgsls {
		qs := cgsl.Quad()
		for _, q := range qs {
			quadSet = append(quadSet, q)
		}
	}
	dbUrl := conf.GetDbUrl()
	store, err := cayley.NewGraph("mongo", dbUrl, nil)
	if err != nil {
		panic(err)
	}
	return store.AddQuadSet(quadSet)
}