// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package nsh

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/edwarnicke/govpp/binapi/vpe"
)

// RPCService defines RPC service nsh.
type RPCService interface {
	NshAddDelEntry(ctx context.Context, in *NshAddDelEntry) (*NshAddDelEntryReply, error)
	NshAddDelMap(ctx context.Context, in *NshAddDelMap) (*NshAddDelMapReply, error)
	NshEntryDump(ctx context.Context, in *NshEntryDump) (RPCService_NshEntryDumpClient, error)
	NshMapDump(ctx context.Context, in *NshMapDump) (RPCService_NshMapDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) NshAddDelEntry(ctx context.Context, in *NshAddDelEntry) (*NshAddDelEntryReply, error) {
	out := new(NshAddDelEntryReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NshAddDelMap(ctx context.Context, in *NshAddDelMap) (*NshAddDelMapReply, error) {
	out := new(NshAddDelMapReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NshEntryDump(ctx context.Context, in *NshEntryDump) (RPCService_NshEntryDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_NshEntryDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_NshEntryDumpClient interface {
	Recv() (*NshEntryDetails, error)
	api.Stream
}

type serviceClient_NshEntryDumpClient struct {
	api.Stream
}

func (c *serviceClient_NshEntryDumpClient) Recv() (*NshEntryDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *NshEntryDetails:
		return m, nil
	case *vpe.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) NshMapDump(ctx context.Context, in *NshMapDump) (RPCService_NshMapDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_NshMapDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_NshMapDumpClient interface {
	Recv() (*NshMapDetails, error)
	api.Stream
}

type serviceClient_NshMapDumpClient struct {
	api.Stream
}

func (c *serviceClient_NshMapDumpClient) Recv() (*NshMapDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *NshMapDetails:
		return m, nil
	case *vpe.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
