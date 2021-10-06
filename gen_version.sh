#! /bin/bash
VERSION="$(git describe --tags $(git rev-list --tags --max-count=1))"
HASH=`git rev-list --tags --max-count=1`
cat >./VERSION <<EOF
$VERSION
EOF
cat >./typex/version.go <<EOF
//
// Warning:
// This file is generated by go compiler, don't change it!!!
// Generated Time: "$(echo $(date "+%Y-%m-%d %H:%M:%S"))"
//
package typex

type Version struct {
	Version   string
	ReleaseTime string
}

var DefaultVersion = Version{
	Version:   \`${VERSION}-${HASH:0:15}\`,
	ReleaseTime: "$(echo $(date "+%Y-%m-%d %H:%M:%S"))",
}

EOF
echo "Generate Version Susseccfully"
