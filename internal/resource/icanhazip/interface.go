package icanhazip

import "context"

type Resource interface {
	GetIP(context.Context, GetIPInput) (*GetIPReply, error)
}
