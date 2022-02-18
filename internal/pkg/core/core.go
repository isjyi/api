package core

const _UI = `
              .__ 
_____  ______ |__|
\__  \ \____ \|  |
 / __ \|  |_> >  |
(____  /   __/|__|
     \/|__|       
`

type Option func(*option)

type option struct {
	disableSwagger    bool
	enableCors        bool
	enableRate        bool
	enableOpenBrowser string
}

// WithDisableSwagger 禁用 swagger
func WithDisableSwagger() Option {
	return func(opt *option) {
		opt.disableSwagger = true
	}
}

// WithEnableCors 设置支持跨域
func WithEnableCors() Option {
	return func(opt *option) {
		opt.enableCors = true
	}
}

// WithEnableRate 设置支持限流
func WithEnableRate() Option {
	return func(opt *option) {
		opt.enableRate = true
	}
}
