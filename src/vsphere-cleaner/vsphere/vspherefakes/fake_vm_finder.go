// Code generated by counterfeiter. DO NOT EDIT.
package vspherefakes

import (
	"context"
	"github.com/vmware/govmomi/object"
)

type FakeVmFinder struct {
	Err error
}

func (fake FakeVmFinder) FindByIp(arg1 context.Context, arg2 *object.Datacenter, arg3 string, arg4 bool) (object.Reference, error) {
	return nil, fake.Err
}
