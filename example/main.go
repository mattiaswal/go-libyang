package main

/*
#include <libyang/libyang.h>
*/
import "C"

import (
	"fmt"
	"log"

	"github.com/mattiaswal/go-libyang/libyang"
)


func load_module(ctx *libyang.Context, m string) {
	module, err := ctx.LoadModule(m, "")
	if err != nil {
		log.Fatalf("Failed to load module: %v", err)
	}

	fmt.Printf("Loaded module: %s (rev: %s)\n", module.Name(), module.Revision())
}
func main() {
	ctx, err := libyang.NewContext(".")
	if err != nil {
		log.Fatalf("Failed to create context: %v", err)
	}
	defer ctx.Free()

	load_module(ctx, "iana-if-type")
	load_module(ctx, "ietf-interfaces")

	jsonData := `
	{
		"ietf-interfaces:interfaces": {
			"interface": [
				{
					"name": "eth0",
                                        "type": "iana-if-type:ethernetCsmacd",
					"enabled": true,
                                        "oper-status": "up",

                                 "statistics": {
                                     "discontinuity-time": "2025-04-17T20:00:00Z",
                               }
                            }
			]
		},
                "ietf-yang-library:yang-library": {
                   "content-id": "some-unique-id"
                },
                "ietf-yang-library:modules-state": {
                   "module-set-id": "some-module-set-id"
                }
	}`

	root, err := ctx.ParseData(jsonData, libyang.DataFormatJSON)
	if err != nil {
		log.Fatalf("Failed to parse data: %v", err)
	}
	defer root.Free()

	if err := root.Validate(); err != C.LY_SUCCESS {
		log.Fatalf("Validation failed: %v", err)
	}
	fmt.Println("Validation successful.")

	result, err := root.Print(libyang.DataFormatJSON)
	if err != nil {
		log.Fatalf("Failed to print data: %v", err)
	}
	fmt.Println("Printed JSON data:")
	fmt.Println(result)
}
