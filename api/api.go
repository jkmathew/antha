//go:generate protoc -I/usr/local/include -I${GOPATH}/src -I. --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/jkmathew/antha/api/v1/coord.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/inventory.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/measurement.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/message.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/state.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/task.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/blob.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/polynomial.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/workflow.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/empty.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/element.proto ${GOPATH}/src/github.com/jkmathew/antha/api/v1/device.proto

package api
