package main

import (
	"flag"

	//	"github.com/bytom/blockchain/pseudohsm"
	//	"github.com/bytom/blockchain/txbuilder"
	"github.com/bytom/tool/sendbulktx/core"
)

var (
	acctNum int
	//	btmNum    int
	thdTxNum  int
	thdNum    int
	assetNum  int
	sendAcct  string
	sendasset string
)

func init() {
	//	flag.IntVar(&acctNum, "acctnum", 10, "Number of created accounts")
	//	flag.IntVar(&btmNum, "btmNum", 10000, "Number of btm to send trading accounts")
	flag.IntVar(&thdTxNum, "thdtxnum", 10, " The number of transactions per goroutine")
	flag.IntVar(&thdNum, "thdnum", 5, "goroutine num")
	flag.IntVar(&assetNum, "assetnum", 10, "Number of transactions asset")
	flag.StringVar(&sendAcct, "sendAcct", "0CHHJNM3G0A02", "who send btm")
	flag.StringVar(&sendasset, "sendasset", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "send asset")

}

func main() {
	//flag.Parse()

	core.Execute()
	// read controlPrograms from file
	/*
		// create key
		param := []string{"alice", "123"}
		fmt.Println("*****************create key start*****************")
		var xpub pseudohsm.XPub
		resp, b := core.SendReq(core.CreateKey, param)
		if !b {
			resp, b := core.SendReq(core.ListKeys, param)
			if !b {
				os.Exit(1)
			}
			dataList, _ := resp.([]interface{})
			for _, item := range dataList {
				core.RestoreStruct(item, &xpub)
				if strings.EqualFold(xpub.Alias, param[0]) {
					break
				}

			}
		} else {
			core.RestoreStruct(resp, &xpub)
		}
		fmt.Println("*****************create key end*****************")

		fmt.Println("*****************create account start*****************")
		for i := 0; i < acctNum; i++ {
			// create account
			name := fmt.Sprintf("alice%d", i)
			param = []string{name, xpub.XPub.String()}
			_, b = core.SendReq(core.CreateAccount, param)
			// create receiver
			param = []string{name}
			resp, b = core.SendReq(core.CreateReceiver, param)
			if !b {
				os.Exit(1)
			}
			var recv txbuilder.Receiver
			core.RestoreStruct(resp, &recv)
			recvText, _ := recv.ControlProgram.MarshalText()
			controlPrograms[i] = string(recvText)
		}
		fmt.Println("*****************create account end*****************")
	*/

}
