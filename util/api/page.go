package api

const DefaultPage = 1
const DefaultLimit = 10

type Page struct {
	Page  int `form:"page,default=1" validate:"omitempty,min=1" json:"page"`            // page, default=1
	Limit int `form:"limit,default=10" validate:"omitempty,min=1,max=100" json:"limit"` // page limit, default=10，min=1, max=100
}

func (p *Page) GetPage() int {
	if p.Page == 0 {
		p.Page = DefaultPage
	}
	return p.Page
}

func (p *Page) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = DefaultLimit
	}
	return p.Limit
}

func (p *Page) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
