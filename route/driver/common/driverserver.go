package common

type DriverServer struct {
	Impl Driver
}

func (d *DriverServer) ReadPoints(req Request, rsp *Response) (err error) {
	*rsp, err = d.Impl.ReadPoints(req)
	return
}
