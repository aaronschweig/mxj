package mxj_test

import (
	"os"
	"testing"

	"github.com/aaronschweig/mxj/v2"
)

func TestSetPathSeparator(t *testing.T) {

	f, err := os.Open("config.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	mxj.SetPathSeparator("@")

	m, err := mxj.NewMapXmlReader(f)
	if err != nil {
		panic(err)
	}

	includesPath := "jenkins.branch.OrganizationFolder@navigators@org.jenkinsci.plugins.github__branch__source.GitHubSCMNavigator@traits@jenkins.scm.impl.trait.WildcardSCMSourceFilterTrait@includes"

	m.UpdateValuesForPath("includes:test test2 test3", includesPath)

	m.XmlIndentWriter(os.Stdout, "", "  ")
}
