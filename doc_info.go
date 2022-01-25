package printer

//go:generate go run github.com/gorpher/winspool-go/cmd/mkwinsyscall_struct -struct doc

type DOC_INFO_1 struct {
	DocName    *uint16
	OutputFile *uint16
	Datatype   *uint16
}
