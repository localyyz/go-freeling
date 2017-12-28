package engine

type Context struct {
	Configuration
	*Engine
}

func NewContext(configFile string) *Context {
	return &Context{
		Engine: NewEngine(),
	}
}
