package engine

type Context struct {
	*Engine
}

func NewContext() *Context {
	return &Context{
		Engine: NewEngine(),
	}
}
