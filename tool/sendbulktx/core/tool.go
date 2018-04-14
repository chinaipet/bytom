package core

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/bytom/util"
)

func init() {
	buildTransactionCmd.PersistentFlags().IntVar(&thdTxNum, "thdtxnum", 10, " The number of transactions per goroutine")
	buildTransactionCmd.PersistentFlags().IntVar(&thdNum, "thdnum", 5, "goroutine num")
	buildTransactionCmd.PersistentFlags().IntVar(&assetNum, "assetnum", 10, "Number of transactions asset")
	buildTransactionCmd.PersistentFlags().StringVar(&sendAcct, "sendAcct", "0CHHJNM3G0A02", "who send btm")
	buildTransactionCmd.PersistentFlags().StringVar(&sendasset, "sendasset", "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "send asset")

}

var (
	acctNum int
	//	btmNum    int
	thdTxNum  int
	thdNum    int
	assetNum  int
	sendAcct  string
	sendasset string
)

var buildTransactionCmd = &cobra.Command{
	Use:   "build-transaction <accountID|alias> <assetID|alias> <amount>",
	Short: "Build one transaction template,default use account id and asset id",
	Args:  cobra.RangeArgs(0, 6),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("type")
		if buildType == "spend" {
			cmd.MarkFlagRequired("receiver")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		controlPrograms := make([]string, acctNum)
		txidChan := make(chan string)
		switch buildType {
		case "issue":

		case "spend":

		case "address":

		default:
			fmt.Println("Invalid transaction template type")
			os.Exit(util.ErrLocalExe)
		}
		txBtm := fmt.Sprintf("%d", assetNum)
		fmt.Println("*****************send tx start*****************")
		// send btm to account
		for i := 0; i < thdNum; i++ {
			go Sendbulktx(thdTxNum, txBtm, sendAcct, sendasset, controlPrograms, txidChan)
		}
		var txid string
		fail := 0
		sucess := 0

		file, error := os.OpenFile("./txid.txt", os.O_RDWR|os.O_CREATE, 0766)
		if error != nil {
			fmt.Println(error)
		}
		for {
			select {
			case txid = <-txidChan:
				if strings.EqualFold(txid, "") {
					fail++
				} else {
					sucess++
					file.WriteString(txid)
					file.WriteString("\n")
				}
			default:
				if (sucess + fail) >= (thdTxNum * thdNum) {
					file.Close()
					os.Exit(0)
				}
				time.Sleep(time.Second * 2)
			}
		}
	},
}

// Execute send tx
func Execute() {

	if _, err := buildTransactionCmd.ExecuteC(); err != nil {
		os.Exit(util.ErrLocalExe)
	}
}
