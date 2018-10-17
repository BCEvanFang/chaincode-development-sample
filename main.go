/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"

	sa "simplecontract/simpleasset"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(sa.SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
