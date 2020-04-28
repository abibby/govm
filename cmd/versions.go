package cmd

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
)

type GoVersion struct {
	major         int
	minor         int
	patch         int
	branch        string
	branchVersion int
}

func availableVersions() ([]*GoVersion, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	dl := path.Join(gopath, "src/golang.org/dl")
	files, err := ioutil.ReadDir(dl)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`^go(\d+)\.(\d+)(\.(\d+))?(([a-z]+)(\d+))?$`)

	versions := []*GoVersion{}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		name := file.Name()
		matches := re.FindStringSubmatch(name)
		if matches == nil {
			continue
		}
		versions = append(versions, &GoVersion{
			major:         mustAtoi(matches[1]),
			minor:         mustAtoi(matches[2]),
			patch:         mustAtoi(matches[4]),
			branch:        matches[6],
			branchVersion: mustAtoi(matches[7]),
		})
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].Greater(versions[j])
	})

	return versions, nil
}

func (v *GoVersion) String() string {
	str := fmt.Sprintf("go%d.%d", v.major, v.minor)
	if v.patch != 0 {
		str += fmt.Sprintf(".%d", v.patch)
	}
	if v.branch != "" {
		str += v.branch
	}
	if v.branchVersion != 0 {
		str += fmt.Sprintf("%d", v.branchVersion)
	}
	return str
}

func (v *GoVersion) branchLevel() int {
	switch v.branch {
	case "beta":
		return 0
	case "rc":
		return 1
	case "":
		return 2
	default:
		return 4
	}
}

func (v *GoVersion) Greater(v2 *GoVersion) bool {
	if v.major > v2.major {
		return true
	} else if v.major < v2.major {
		return false
	}
	if v.minor > v2.minor {
		return true
	} else if v.minor < v2.minor {
		return false
	}
	if v.patch > v2.patch {
		return true
	} else if v.patch > v2.patch {
		return false
	}
	if v.branchLevel() > v2.branchLevel() {
		return true
	} else if v.branchLevel() < v2.branchLevel() {
		return false
	}
	if v.branchVersion > v2.branchVersion {
		return true
	} else if v.branchVersion < v2.branchVersion {
		return false
	}
	return false
}

func mustAtoi(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func VersionExists(v string) (bool, error) {
	versions, err := availableVersions()
	if err != nil {
		return false, err
	}
	for _, version := range versions {
		if version.String() == v {
			return true, nil
		}
	}
	return false, nil
}
