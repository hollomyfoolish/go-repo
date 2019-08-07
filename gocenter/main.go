package gcheck

import (
	"fmt"

	"github.com/MixinNetwork/go-number"
	"github.com/fxtlabs/date"
	"github.com/heketi/utils"
	"github.com/naoina/go-stringutil"
	p_string "github.com/pefish/go-string"
	"rsc.io/quote"
)

func Check() {
	fmt.Println(stringutil.ToUpperCamelCase("username"))
	fmt.Println(number.FromString("123"))
	fmt.Println(p_string.String.DesensitizeMobile(`18317042247`))
	fmt.Println(date.Today())
	fmt.Println(utils.GenUUID())
	fmt.Println(quote.Hello())
}
