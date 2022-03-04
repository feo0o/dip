package app

var (
	Name      = "Dip"
	majorVer  = "0"
	minorVer  = "0"
	buildVer  = "0"
	gitCommit string
)

func VersionSlim() string {
	return majorVer + "." + minorVer + "." + buildVer
}

func Version() string {
	s := Name + " v" + VersionSlim()
	if gitCommit != "" {
		return s + "-" + gitCommit
	}
	return s
}
