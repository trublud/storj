// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"storj.io/storj/internal/fpath"
	libuplink "storj.io/storj/lib/uplink"
	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/macaroon"
	"storj.io/storj/pkg/process"
	"storj.io/storj/uplink"
)

var shareCfg struct {
	DisallowReads     bool     `default:"false" help:"if true, disallow reads"`
	DisallowWrites    bool     `default:"false" help:"if true, disallow writes"`
	DisallowLists     bool     `default:"false" help:"if true, disallow lists"`
	DisallowDeletes   bool     `default:"false" help:"if true, disallow deletes"`
	Readonly          bool     `default:"false" help:"implies disallow_writes and disallow_deletes"`
	Writeonly         bool     `default:"false" help:"implies disallow_reads and disallow_lists"`
	NotBefore         string   `help:"disallow access before this time"`
	NotAfter          string   `help:"disallow access after this time"`
	AllowedPathPrefix []string `help:"whitelist of bucket path prefixes to require"`

	// Share requires information about the current scope
	uplink.ScopeConfig
}

func init() {
	// We skip the use of addCmd here because we only want the configuration options listed
	// above, and addCmd adds a whole lot more than we want.

	shareCmd := &cobra.Command{
		Use:   "share",
		Short: "Creates a possibly restricted api key",
		RunE:  shareMain,
	}
	RootCmd.AddCommand(shareCmd)

	defaultConfDir := fpath.ApplicationDir("storj", "uplink")
	confDirParam := cfgstruct.FindConfigDirParam()
	if confDirParam != "" {
		defaultConfDir = confDirParam
	}

	process.Bind(shareCmd, &shareCfg, defaults, cfgstruct.ConfDir(defaultConfDir))
}

const shareISO8601 = "2006-01-02T15:04:05-0700"

func parseHumanDate(date string, now time.Time) (*time.Time, error) {
	switch {
	case date == "":
		return nil, nil
	case date == "now":
		return &now, nil
	case date[0] == '+':
		d, err := time.ParseDuration(date[1:])
		t := now.Add(d)
		return &t, errs.Wrap(err)
	case date[0] == '-':
		d, err := time.ParseDuration(date[1:])
		t := now.Add(-d)
		return &t, errs.Wrap(err)
	default:
		t, err := time.Parse(shareISO8601, date)
		return &t, errs.Wrap(err)
	}
}

// shareMain is the function executed when shareCmd is called
func shareMain(cmd *cobra.Command, args []string) (err error) {
	now := time.Now()

	notBefore, err := parseHumanDate(shareCfg.NotBefore, now)
	if err != nil {
		return err
	}
	notAfter, err := parseHumanDate(shareCfg.NotAfter, now)
	if err != nil {
		return err
	}

	var restrictions []libuplink.EncryptionRestriction
	for _, path := range shareCfg.AllowedPathPrefix {
		p, err := fpath.New(path)
		if err != nil {
			return err
		}
		if p.IsLocal() {
			return errs.New("required path must be remote: %q", path)
		}

		restrictions = append(restrictions, libuplink.EncryptionRestriction{
			Bucket:     p.Bucket(),
			PathPrefix: p.Path(),
		})
	}

	scope, err := shareCfg.GetScope()
	if err != nil {
		return err
	}
	key, access := scope.APIKey, scope.EncryptionAccess

	if len(restrictions) > 0 {
		key, access, err = access.Restrict(key, restrictions...)
		if err != nil {
			return err
		}
	}

	caveat, err := macaroon.NewCaveat()
	if err != nil {
		return err
	}

	caveat.DisallowDeletes = shareCfg.DisallowDeletes || shareCfg.Readonly
	caveat.DisallowLists = shareCfg.DisallowLists || shareCfg.Writeonly
	caveat.DisallowReads = shareCfg.DisallowReads || shareCfg.Writeonly
	caveat.DisallowWrites = shareCfg.DisallowWrites || shareCfg.Readonly
	caveat.NotBefore = notBefore
	caveat.NotAfter = notAfter

	{
		// Times don't marshal very well with MarshalTextString, and the nonce doesn't
		// matter to humans, so handle those explicitly and then dispatch to the generic
		// routine to avoid having to print all the things individually.
		caveatCopy := proto.Clone(&caveat).(*macaroon.Caveat)
		caveatCopy.Nonce = nil
		if caveatCopy.NotBefore != nil {
			fmt.Println("not before:", caveatCopy.NotBefore.Truncate(0).Format(shareISO8601))
			caveatCopy.NotBefore = nil
		}
		if caveatCopy.NotAfter != nil {
			fmt.Println("not after:", caveatCopy.NotAfter.Truncate(0).Format(shareISO8601))
			caveatCopy.NotAfter = nil
		}
		fmt.Print(proto.MarshalTextString(caveatCopy))
	}

	key, err = key.Restrict(caveat)
	if err != nil {
		return err
	}

	accessData, err := access.Serialize()
	if err != nil {
		return err
	}

	newScope := &libuplink.Scope{
		SatelliteAddr:    scope.SatelliteAddr,
		APIKey:           key,
		EncryptionAccess: access,
	}

	scopeData, err := newScope.Serialize()
	if err != nil {
		return err
	}

	fmt.Println("api key:", key.Serialize())
	fmt.Println("enc ctx:", accessData)
	fmt.Println("scope  :", scopeData)
	return nil
}
