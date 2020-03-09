package rbuser

const (
	timeLayout = "2006-01-02 15:04:05"

	// Shell defaults
	expiredShell = "/usr/local/shells/expired"
	noLoginShell = "/usr/local/shells/disusered"
	defaultShell = "/usr/local/shells/shell"
)

// User groups
var (
	associatGroup  = group{"associat", 107}
	clubGroup      = group{"club", 102}
	committeeGroup = group{"committe", 100}
	dcuGroup       = group{"dcu", 31382}
	founderGroup   = group{"founder", 105}
	intersocGroup  = group{"intersoc", 1016}
	memberGroup    = group{"member", 103}
	projectsGroup  = group{"projects", 1014}
	redbrickGroup  = group{"redbrick", 1017}
	societyGroup   = group{"society", 101}
	staffGroup     = group{"staff", 109}

	groups = []group{
		associatGroup,
		clubGroup,
		committeeGroup,
		dcuGroup,
		founderGroup,
		intersocGroup,
		memberGroup,
		projectsGroup,
		redbrickGroup,
		societyGroup,
		staffGroup,
	}
)
