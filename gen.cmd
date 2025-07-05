pushd ..\elorm-gen
go build
popd
..\elorm-gen\elorm-gen.exe entities.json dbcontext.go

