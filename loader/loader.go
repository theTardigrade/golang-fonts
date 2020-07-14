// Package loader allows fonts to be dynamically loaded.
package loader

import (
	"github.com/theTardigrade/goFonts/data/opensansbold"
	"github.com/theTardigrade/goFonts/data/opensansbolditalic"
	"github.com/theTardigrade/goFonts/data/opensansextrabold"
	"github.com/theTardigrade/goFonts/data/opensansextrabolditalic"
	"github.com/theTardigrade/goFonts/data/opensansitalic"
	"github.com/theTardigrade/goFonts/data/opensanslight"
	"github.com/theTardigrade/goFonts/data/opensanslightitalic"
	"github.com/theTardigrade/goFonts/data/opensansregular"
	"github.com/theTardigrade/goFonts/data/opensanssemibold"
	"github.com/theTardigrade/goFonts/data/opensanssemibolditalic"
	"github.com/theTardigrade/goFonts/data/playfairdisplayblack"
	"github.com/theTardigrade/goFonts/data/playfairdisplayblackitalic"
	"github.com/theTardigrade/goFonts/data/playfairdisplaybold"
	"github.com/theTardigrade/goFonts/data/playfairdisplaybolditalic"
	"github.com/theTardigrade/goFonts/data/playfairdisplayitalic"
	"github.com/theTardigrade/goFonts/data/playfairdisplayregular"
	"github.com/theTardigrade/goFonts/data/ralewayblack"
	"github.com/theTardigrade/goFonts/data/ralewayblackitalic"
	"github.com/theTardigrade/goFonts/data/ralewaybold"
	"github.com/theTardigrade/goFonts/data/ralewaybolditalic"
	"github.com/theTardigrade/goFonts/data/ralewayextrabold"
	"github.com/theTardigrade/goFonts/data/ralewayextrabolditalic"
	"github.com/theTardigrade/goFonts/data/ralewayextralight"
	"github.com/theTardigrade/goFonts/data/ralewayextralightitalic"
	"github.com/theTardigrade/goFonts/data/ralewayitalic"
	"github.com/theTardigrade/goFonts/data/ralewaylight"
	"github.com/theTardigrade/goFonts/data/ralewaylightitalic"
	"github.com/theTardigrade/goFonts/data/ralewaymedium"
	"github.com/theTardigrade/goFonts/data/ralewaymediumitalic"
	"github.com/theTardigrade/goFonts/data/ralewayregular"
	"github.com/theTardigrade/goFonts/data/ralewaysemibold"
	"github.com/theTardigrade/goFonts/data/ralewaysemibolditalic"
	"github.com/theTardigrade/goFonts/data/ralewaythin"
	"github.com/theTardigrade/goFonts/data/ralewaythinitalic"
	"github.com/theTardigrade/goFonts/data/robotoblack"
	"github.com/theTardigrade/goFonts/data/robotoblackitalic"
	"github.com/theTardigrade/goFonts/data/robotobold"
	"github.com/theTardigrade/goFonts/data/robotobolditalic"
	"github.com/theTardigrade/goFonts/data/robotocondensedbold"
	"github.com/theTardigrade/goFonts/data/robotocondensedbolditalic"
	"github.com/theTardigrade/goFonts/data/robotocondenseditalic"
	"github.com/theTardigrade/goFonts/data/robotocondensedlight"
	"github.com/theTardigrade/goFonts/data/robotocondensedlightitalic"
	"github.com/theTardigrade/goFonts/data/robotocondensedregular"
	"github.com/theTardigrade/goFonts/data/robotoitalic"
	"github.com/theTardigrade/goFonts/data/robotolight"
	"github.com/theTardigrade/goFonts/data/robotolightitalic"
	"github.com/theTardigrade/goFonts/data/robotomedium"
	"github.com/theTardigrade/goFonts/data/robotomediumitalic"
	"github.com/theTardigrade/goFonts/data/robotoregular"
	"github.com/theTardigrade/goFonts/data/robotothin"
	"github.com/theTardigrade/goFonts/data/robotothinitalic"
)

// TTF retrieves the data for a given font by name at runtime.
func TTF(pkgName string) (data []byte, found bool) {
	found = true

	if len(pkgName) >= 2 {
		switch pkgName[0] {
		case 'p':
			switch pkgName[1:] {
			case "layfairdisplayblack":
				data = playfairdisplayblack.TTF
			case "layfairdisplayblackitalic":
				data = playfairdisplayblackitalic.TTF
			case "layfairdisplaybold":
				data = playfairdisplaybold.TTF
			case "layfairdisplaybolditalic":
				data = playfairdisplaybolditalic.TTF
			case "layfairdisplayitalic":
				data = playfairdisplayitalic.TTF
			case "layfairdisplayregular":
				data = playfairdisplayregular.TTF
			default:
				found = false
			}
		case 'r':
			switch pkgName[1:] {
			case "alewayblack":
				data = ralewayblack.TTF
			case "alewayblackitalic":
				data = ralewayblackitalic.TTF
			case "alewaybold":
				data = ralewaybold.TTF
			case "alewaybolditalic":
				data = ralewaybolditalic.TTF
			case "alewayextrabold":
				data = ralewayextrabold.TTF
			case "alewayextrabolditalic":
				data = ralewayextrabolditalic.TTF
			case "alewayextralight":
				data = ralewayextralight.TTF
			case "alewayextralightitalic":
				data = ralewayextralightitalic.TTF
			case "alewayitalic":
				data = ralewayitalic.TTF
			case "alewaylight":
				data = ralewaylight.TTF
			case "alewaylightitalic":
				data = ralewaylightitalic.TTF
			case "alewaymedium":
				data = ralewaymedium.TTF
			case "alewaymediumitalic":
				data = ralewaymediumitalic.TTF
			case "alewayregular":
				data = ralewayregular.TTF
			case "alewaysemibold":
				data = ralewaysemibold.TTF
			case "alewaysemibolditalic":
				data = ralewaysemibolditalic.TTF
			case "alewaythin":
				data = ralewaythin.TTF
			case "alewaythinitalic":
				data = ralewaythinitalic.TTF
			case "obotoblack":
				data = robotoblack.TTF
			case "obotoblackitalic":
				data = robotoblackitalic.TTF
			case "obotobold":
				data = robotobold.TTF
			case "obotobolditalic":
				data = robotobolditalic.TTF
			case "obotocondensedbold":
				data = robotocondensedbold.TTF
			case "obotocondensedbolditalic":
				data = robotocondensedbolditalic.TTF
			case "obotocondenseditalic":
				data = robotocondenseditalic.TTF
			case "obotocondensedlight":
				data = robotocondensedlight.TTF
			case "obotocondensedlightitalic":
				data = robotocondensedlightitalic.TTF
			case "obotocondensedregular":
				data = robotocondensedregular.TTF
			case "obotoitalic":
				data = robotoitalic.TTF
			case "obotolight":
				data = robotolight.TTF
			case "obotolightitalic":
				data = robotolightitalic.TTF
			case "obotomedium":
				data = robotomedium.TTF
			case "obotomediumitalic":
				data = robotomediumitalic.TTF
			case "obotoregular":
				data = robotoregular.TTF
			case "obotothin":
				data = robotothin.TTF
			case "obotothinitalic":
				data = robotothinitalic.TTF
			default:
				found = false
			}
		case 'o':
			switch pkgName[1:] {
			case "pensansbold":
				data = opensansbold.TTF
			case "pensansbolditalic":
				data = opensansbolditalic.TTF
			case "pensansextrabold":
				data = opensansextrabold.TTF
			case "pensansextrabolditalic":
				data = opensansextrabolditalic.TTF
			case "pensansitalic":
				data = opensansitalic.TTF
			case "pensanslight":
				data = opensanslight.TTF
			case "pensanslightitalic":
				data = opensanslightitalic.TTF
			case "pensansregular":
				data = opensansregular.TTF
			case "pensanssemibold":
				data = opensanssemibold.TTF
			case "pensanssemibolditalic":
				data = opensanssemibolditalic.TTF
			default:
				found = false
			}
		default:
			found = false
		}
	}

	return
}
