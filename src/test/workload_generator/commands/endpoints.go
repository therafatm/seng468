package commands

import (
    "log"
    "fmt"
)

func FormatCommandEndpoint(cmd Command) string {
    switch cmd.Name {
        case "BALANCE": 
            return fmt.Sprintf("/api/availableBalance/%s", cmd.Username)

        case "SHARES": 
            return fmt.Sprintf("/api/availableShares/%s/%s", cmd.Username, cmd.Symbol)

        case "ADD":
           return fmt.Sprintf("/api/add/%s/%d", cmd.Username, cmd.Amount)

        case "QUOTE":
			return fmt.Sprintf("/api/getQuote/%s/%s", cmd.Username, cmd.Symbol)

        case "BUY":
			return fmt.Sprintf("/api/buy/%s/%s/%d", cmd.Username, cmd.Symbol, cmd.Amount)

        case "SELL":
			return fmt.Sprintf("/api/sell/%s/%s/%d", cmd.Username, cmd.Symbol, cmd.Amount)

        case "COMMIT_BUY":
			return fmt.Sprintf("/api/commitBuy/%s", cmd.Username)    

        case "COMMIT_SELL":
			return fmt.Sprintf("/api/commitSell/%s", cmd.Username)

        case "CANCEL_BUY":
            return fmt.Sprintf("/api/cancelBuy/%s", cmd.Username)    

        case "CANCEL_SELL":
            return fmt.Sprintf("/api/cancelSell/%s", cmd.Username)

        case "DISPLAY_SUMMARY":
            return ""
            
        case "SET_BUY_TRIGGER":
            return fmt.Sprintf("/api/setBuyTrigger/%s/%s/%d", cmd.Username, cmd.Symbol, cmd.Amount)
            
        case "SET_SELL_TRIGGER":
            return fmt.Sprintf("/api/setSellTrigger/%s/%s/%d", cmd.Username, cmd.Symbol, cmd.Amount)

        case "CANCEL_SET_BUY":
            return fmt.Sprintf("/api/cancelSetBuy/%s/%s", cmd.Username, cmd.Symbol)

        case "CANCEL_SET_SELL":
            return fmt.Sprintf("/api/cancelSetSell/%s/%s", cmd.Username, cmd.Symbol)

        case "SET_BUY_AMOUNT":
            return fmt.Sprintf("/api/setBuyAmount/%s/%s/%d", cmd.Username, cmd.Symbol, cmd.Amount)

        case "SET_SELL_AMOUNT":
            return fmt.Sprintf("/api/setSellAmount/%s/%s/%d", cmd.Username, cmd.Symbol, cmd.Amount)

        case "DUMPLOG":
            return ""

        default:
            log.Fatalf("Invalid command: %s", cmd.Name)
    }
    
    return ""
}