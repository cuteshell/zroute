package route

import "strings"

type RouteHook func(srcPath string, dstPath string, data interface{}) error

type RouteRule struct {
	SrcPath string //identifier, such as topic path or channel name
	DstPath string //identifier, such as topic path or channel name
	Type    int32  // type of the route
}

type RouteTable struct {
	Rules       []RouteRule
	BeforeRoute RouteHook
	AfterRoute  RouteHook
}

func (r *RouteTable) Route(srcPath string, data interface{}) {
	if r.BeforeRoute != nil {
		r.BeforeRoute(srcPath, "", data)
	}

	for _, rule := range r.Rules {
		if rule.SrcPath == srcPath {
			r.SendTo(rule.DstPath, data)
			continue
		}
		if len(srcPath) > len(rule.SrcPath) && 
			strings.HasPrefix(srcPath, rule.SrcPath) && srcPath[len(rule.SrcPath)] == '/' {
			r.SendTo(rule.DstPath, data)
		}
		if r.AfterRoute != nil {
			r.AfterRoute(srcPath, rule.DstPath, data)
		}
	}
}

func (r *RouteTable) SendTo(dstPath string, data interface{}) {

}
