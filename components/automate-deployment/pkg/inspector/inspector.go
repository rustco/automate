package inspector

type Inspector interface {
	AddInspection(inspection Inspection)
	ShowInfo() error
	ShowInspectionList()
	Inspect() error
	RollBackChangesOnError() error
	RunExitAction() error
	RunUpgradeInspector(osDestDir string, skipStorageCheck bool) (isError bool)
}

type Inspection interface {
	ShowInfo(index *int) error
}

type SystemInspection interface {
	Inspection
	Inspect() error
	Skip()
	GetShortInfo() []string
	GetInstallationType() InstallationType
	ExitHandler() error
}

type RollbackInspection interface {
	SystemInspection
	RollBackHandler() error
}

type InstallationType int

const (
	EMBEDDED InstallationType = iota
	EXTERNAL InstallationType = iota
	BOTH     InstallationType = iota
)
