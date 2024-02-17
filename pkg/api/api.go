package api

import (
	"fmt"

	"github.com/xmp-er/temp-mod/constants"
)

func CallAPI() {
	fmt.Println("API called successfully")
	fmt.Println("Log Level", constants.LOG_LEVEL)
}
