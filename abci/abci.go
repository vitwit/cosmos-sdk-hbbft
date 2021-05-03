package abci

import (
	"context"
	"errors"

	"google.golang.org/grpc"
)

type Application struct {
	// TODO: Define params necessary for ABCI implementation
}

func New() *grpc.Server {
	srv := grpc.NewServer()
	abciSvr := Application{}
	RegisterABCIApplicationServer(srv, abciSvr)
	return srv
}

func (a Application) Echo(ctx context.Context, req *RequestEcho) (*ResponseEcho, error) {
	return nil, errors.New("not implemented")
}

func (a Application) Flush(ctx context.Context, req *RequestFlush) (*ResponseFlush, error) {
	return nil, errors.New("not implemented")
}

func (a Application) Info(ctx context.Context, req *RequestInfo) (*ResponseInfo, error) {
	return nil, errors.New("not implemented")
}

func (a Application) DeliverTx(ctx context.Context, req *RequestDeliverTx) (*ResponseDeliverTx, error) {
	return nil, errors.New("not implemented")
}

func (a Application) CheckTx(ctx context.Context, req *RequestCheckTx) (*ResponseCheckTx, error) {
	return nil, errors.New("not implemented")
}

func (a Application) Query(ctx context.Context, req *RequestQuery) (*ResponseQuery, error) {
	return nil, errors.New("not implemented")
}

func (a Application) Commit(ctx context.Context, req *RequestCommit) (*ResponseCommit, error) {
	return nil, errors.New("not implemented")
}

func (a Application) InitChain(ctx context.Context, req *RequestInitChain) (*ResponseInitChain, error) {
	return nil, errors.New("not implemented")
}

func (a Application) BeginBlock(ctx context.Context, req *RequestBeginBlock) (*ResponseBeginBlock, error) {
	return nil, errors.New("not implemented")
}

func (a Application) EndBlock(ctx context.Context, req *RequestEndBlock) (*ResponseEndBlock, error) {
	return nil, errors.New("not implemented")
}

func (a Application) ListSnapshots(ctx context.Context, req *RequestListSnapshots) (*ResponseListSnapshots, error) {
	return nil, errors.New("not implemented")
}

func (a Application) OfferSnapshot(ctx context.Context, req *RequestOfferSnapshot) (*ResponseOfferSnapshot, error) {
	return nil, errors.New("not implemented")
}

func (a Application) LoadSnapshotChunk(ctx context.Context, req *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error) {
	return nil, errors.New("not implemented")
}

func (a Application) ApplySnapshotChunk(ctx context.Context, req *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) {
	return nil, errors.New("not implemented")
}
