// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "7e0d319a09eca6c34035a0900757b60f"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"793124cfbcf9ae17892b017c666be8f1": "1f8b08000000000000ff74ceb16eac301085e1daf314471457700bd323a54bda25054abb726016ac8d31190f8922c4bb4746284d94d6fff1a7595c7f772363db606717d83e9f0ffb4ee4c312455192299493fa792c884c317a9dd657dbc7502715d67e92fae8b7af5af87df5c2055544b775eed171d26bbba88f73babeb8373f38e552f1ff146d57612323681e707eb617fe2cb5223271d194c3bf13d876322cc73627fbe35564c43e89442959a42202ea1a5dfbd83608eecef08a8f3cb668329c91df0620f612ff60dc3020c5c00851182e2596e3264c2c6c69a7ef000000ffff04fb4b164b010000",
		"920fdc8a8285bc442cfb791bcafe45b6": "1f8b08000000000000ff9c51c14a033110bde72bc69c9252b2e051e945416f55a8e8395d67d3b0bbc93299d595b2ff2e49ab524105afefbd79f3e6cd60ebd63a84fd1e4cb03d9afb2330cf42f87e88c4a00480749e77e3d6d4b1af5cdc8e4d63bb58390ce14dfe480fdd98760ec32f0a5bb754bd9cffe1f19d1e5a572151a4248516a219430d6b7c5571e0048bbb817d0c49835a9480e6160392e5484b28531af602c0c1c50a0e823cab8500f04d5664225b9947dbf967cba8f465c1cf56107c57a6010879a400ee689acc93e7dd866ddd2a24d202603eb574e62a4eaa9c5c36ca5cfa364e6bdbe7bee512a431d5e7276e7c97e18ab11f3acb98a4fe6f0c809aa71ca1b499975fc7c038b1d207ce6c9095cc37cb65393de3ce3c900da989d423a98f679e80354f3a4bbf4204df895988f7000000ffffa276d8b858020000",
		"a61b972dce26544398f591b4e2f715f3": "1f8b08000000000000ff548fc16ae3301086cfd653081d16693172f6bab0a7253d95b484400fa514c51d2ba296e48c4634c1f8dd8b1453da8b0edf7c9aff9fc9f4efc6029f67ae83f1a01f57b02c8c393f45242e59230812b96005638db08e4ef9a8fbe83b1b8f7918cc183b0b215ccb5b4cf1d34a8440fd09bbba65b87608e7ec1004538c0d39f4fc00895e77f02189ff5ea3f441f19935c8fffee3abafaba1186b6ccb01eba8a05f0f13b918d2bc28d6a0dec52d62440988c5c51c8ab8562b3bf6390440a9ea4c3f393a495bcdafaf85ef7390aa6248b5446190f24849deec7b081221e9ffd17b13de52cb37ea3bbf7323a496ff29f670bb6385cf9b97226ecfd98c52c0c5f869044d17122d1ff4ce7828d10bfb0c0000fffffb4a36cd9f010000",
		"a95877cfc2957859690620af727512eb": "1f8b08000000000000fff248cdc9c957482bcacf55a8ae56d0cb4bcc4dd50bc92cc949cdac4a55a8ade502040000ffff8a41cdb420000000",
		"cb9f04739ceb356f960fa53f13efd3f3": "1f8b08000000000000ff3ccd310ec2300c85e1dda77863cbd01c052676d3bab4a22491630f5594bb2350c4faa4ff7b99e7173f05b5628afc96e9d687d688eccc826bb63dc58262eab3a1121002785970265714f375c5262ad4882804dcf9d81736816d6c483d671578717e1c42abc71943ca5670e9fcf8cf8611a29af4f7a462ae11713fbefa270000ffff1dfc7cc2ad000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("github.com/gobuffalo/genny/genny/new", "../new/templates")
		b.SetResolver("-name-/-name-.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "920fdc8a8285bc442cfb791bcafe45b6"})
		b.SetResolver("-name-/-name-_test.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "a61b972dce26544398f591b4e2f715f3"})
		b.SetResolver("-name-/options.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "cb9f04739ceb356f960fa53f13efd3f3"})
		b.SetResolver("-name-/options_test.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "793124cfbcf9ae17892b017c666be8f1"})
		b.SetResolver("-name-/templates/example.txt.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "a95877cfc2957859690620af727512eb"})
	}()

	return nil
}()