package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "strings"

    "github.com/FactomProject/factom"
)

const (
    DB_USER     = "factomize"
    DB_PASSWORD = "Everything!"
    DB_NAME     = "factomizeeverything"
)

var usage UsageDate

type Stats struct {
    Records int64
    Entries int64
    Anchors int64
}

type UsageDate struct {
    Date  string       `json:"date"`
    Chain []ChainCount `json:"chain"`
}

type ChainCount struct {
    ChainID string `json:"chainid"`
    Count   int64  `json:"count"`
}

// MAIN ________________________________________________________________________________

func main() {

    //make this callable to reset counters
    fmt.Println(os.Args[0])
    if len(os.Args) > 1 {
        fmt.Println(os.Args[1])

    }
    // make sure needed directories exist.http://localhost:8080/TrustedOracle/?request={Startdate=1,EndDate=2,DataItem=APPL,GreaterThan=100,EqualTo=5,LessThan=6,Answered=7,Outcome=8,OutcomeDate=9,QuestionHash=10,AnswerHash=11}
    // if added to conf file for root directory (or couchibase) add here

    oList := GetPegnetEntries(207465)
    fmt.Println(oList)

    //  err := loadBitcoinAverageData()
    //  if err != nil {
    //      fmt.Println(err)
    //  }

    http.HandleFunc("/BlockStats", BlockStats)     // Returns 201
    http.HandleFunc("/BlockWinners", BlockWinners) // Returns 201
    http.HandleFunc("/Winners", Winners)           // Returns 201
    //http.HandleFunc("/dump", dumpHandler)
    //http.HandleFunc("/dumpdocs", dumpdocHandler)

    err := http.ListenAndServe(":8899", nil)
    if err != nil {
        fmt.Println(err)
    }

}

func BlockStats(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    fmt.Println("Q:", request.URL)
    blockheight := request.URL.Query()["height"]
    i, err := strconv.ParseInt(blockheight[0], 10, 64)
    oList := GetPegnetEntries(i)
    listBytes, err := json.Marshal(oList)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(response, "%s", string(listBytes))
}

func BlockWinners(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    fmt.Println(request.URL.Query())
    blockheight := request.URL.Query()["height"]
    i, err := strconv.ParseInt(blockheight[0], 10, 64)
    wList := GetWinners(i)
    listBytes, err := json.Marshal(wList)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(response, "%s", string(listBytes))
}
func Winners(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "text/html")
    fmt.Println(request.URL.Query())
    blockheight := request.URL.Query()["height"]
    filter := request.URL.Query()["filter"]
    i, err := strconv.ParseInt(blockheight[0], 10, 64)
    wList := GetWinners(i + 1)
    pList := GetPegnetEntries(i)

    page := "<html>"
    bh, _ := strconv.ParseInt(blockheight[0], 10, 64)
    page = page + "<a href=\"" + "http://localhost:8899/Winners?height=" + fmt.Sprintf("%d", bh-1) + "&filter=" + filter[0] + \">Prev</a></html>"
    page = page + "<a href=\"" + "http://localhost:8899/Winners?height=" + fmt.Sprintf("%d", bh+1) + "&filter=" + filter[0] + "\">Next</a></html>"

    page = page + "<table><tr><td>MinerID</td><td>height</td><td>ADA</td><td>BNB</td><td>BRL</td><td>CAD</td><td>CHF</td><td>CNY</td><td>DASH</td><td>DCR</td><td>ETH</td><td>EUR</td><td>FCT</td><td>GBP</td><td>HKD</td><td>INR</td><td>JPY</td><td>KRW</td><td>LTC</td><td>MXN</td><td>PHP</td><td>PNT</td><td>RVN</td><td>SGD</td><td>USD</td><td>XAG</td><td>XAU</td><td>XBC</td><td>XBT</td><td>XLM</td><td>XMR</td><td>XPD</td><td>XPT</td><td>ZEC</td></tr>"

    for _, p := range pList {
        wn := 0
        for _, w := range wList {
            

            if strings.Contains(p.EnryHash, w) {
                wn = wn + 1
 
            }

            fmt.Println(p.MinerID, filter[0])

        }

	       if strings.Contains(p.MinerID, filter[0]) {
		    wn = wn + 2
		}
        fmt.Println("wn", wn)

        if wn == 0 { // dont print
        } else if wn == 1 { // winner
            page = page + "<tr bgcolor=#dddd00>"
        } else if wn == 2 { // filter match
            page = page + "<tr bgcolor=#0dddd0>"
        } else if wn == 3 { //trfilter match and winner
            page = page + "<tr bgcolor=#00aadd>"
        }

        if wn > 0 {
            page = page + "<td>" + p.MinerID + "</td>"
            page = page + "<td>" + fmt.Sprintf("%d", p.BlockHeight) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.ADA) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.BNB) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.BRL) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.CAD) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.CHF) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.CNY) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.DASH) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.DCR) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.ETH) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.EUR) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.FCT) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.GBP) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.HKD) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.INR) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.JPY) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.KRW) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.LTC) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.MXN) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.PHP) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.PNT) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.RVN) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.SGD) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.USD) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XAG) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XAU) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XBC) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XBT) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XLM) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XMR) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XPD) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.XPT) + "</td>"
            page = page + "<td>" + fmt.Sprintf("%f", p.Pegs.ZEC) + "</td></tr>"
        }
    }
    page = page + "</table><br>"
    page = page + "</html>"
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(response, "%s", string(page))
}

func GetPegnetEntries(height int64) []OPR {
    var oprList []OPR
    factom.SetFactomdServer("localhost:8088")

    headstart, err := factom.GetDBlockHead()
    if err != nil {
        fmt.Println("err:", err)
    }
    fmt.Println("headstart:", headstart)
    head, _ := factom.GetDBlock(headstart)

    fmt.Println()
    for (int64)(head.Header.SequenceNumber) > height-1 {
        fmt.Println(head.Header.SequenceNumber)
        if head.Header.SequenceNumber == height {
            fmt.Println(head.Header.SequenceNumber)
            //fmt.Println(s)
            for _, eblockHash := range head.EntryBlockList {
                if eblockHash.ChainID[:] == "a642a8674f46696cc47fdb6b65f9c87b2a19c5ea8123b3d2f0c13b6f33a9d5ef" {
                    //  fmt.Println("eblockHash=",eblockHash.KeyMR)
                    eblock, _ := factom.GetEBlock(eblockHash.KeyMR)
                    //fmt.Println("eblock=",eblock)
                    for _, eb := range eblock.EntryList {
                        entry, _ := factom.GetEntry(eb.EntryHash)
                        var opr OPR
                        _ = json.Unmarshal(entry.Content, &opr)
                        opr.EnryHash = eb.EntryHash
                        oprList = addOPR(oprList, opr)
                    }

                }
            }
        }
        head, _ = factom.GetDBlock(head.Header.PrevBlockKeyMR)

    }
    return oprList
}

func GetWinners(height int64) []string {
    factom.SetFactomdServer("localhost:8088")

    headstart, err := factom.GetDBlockHead()
    if err != nil {
        fmt.Println("err:", err)
    }
    fmt.Println("headstart:", headstart)
    head, _ := factom.GetDBlock(headstart)

    fmt.Println()
    for (int64)(head.Header.SequenceNumber) > height-1 {
        fmt.Println(head.Header.SequenceNumber)
        if head.Header.SequenceNumber == height {
            fmt.Println(head.Header.SequenceNumber)
            //fmt.Println(s)
            for _, eblockHash := range head.EntryBlockList {
                if eblockHash.ChainID[:] == "a642a8674f46696cc47fdb6b65f9c87b2a19c5ea8123b3d2f0c13b6f33a9d5ef" {
                    //  fmt.Println("eblockHash=",eblockHash.KeyMR)
                    eblock, _ := factom.GetEBlock(eblockHash.KeyMR)
                    //fmt.Println("eblock=",eblock)
                    for _, eb := range eblock.EntryList {
                        entry, _ := factom.GetEntry(eb.EntryHash)
                        var opr OPR
                        _ = json.Unmarshal(entry.Content, &opr)
                        if opr.BlockHeight == height {
                            return opr.Winners
                        }
                    }

                }
            }
        }
        head, _ = factom.GetDBlock(head.Header.PrevBlockKeyMR)

    }
    return nil
}

func addOPR(oprList []OPR, opr OPR) []OPR {
    ln := len(oprList) + 1
    oprOut := make([]OPR, ln)

    for i, o := range oprList {
        oprOut[i] = o
    }
    oprOut[len(oprList)] = opr
    return oprOut
}

type OPR struct {
    EnryHash    string   `json:"entryhash"`
    BlockHeight int64    `json:"dbht"`
    MinerID     string   `json:"minerid"`
    Coinbase    string   `json:"coinbase"`
    Pegs        Assets   `json:"assets"`
    Winners     []string `json:"winners"`
}

type Assets struct {
    ADA  float64
    BNB  float64
    BRL  float64
    CAD  float64
    CHF  float64
    CNY  float64
    DASH float64
    DCR  float64
    ETH  float64
    EUR  float64
    FCT  float64
    GBP  float64
    HKD  float64
    INR  float64
    JPY  float64
    KRW  float64
    LTC  float64
    MXN  float64
    PHP  float64
    PNT  float64
    RVN  float64
    SGD  float64
    USD  float64
    XAG  float64
    XAU  float64
    XBC  float64
    XBT  float64
    XLM  float64
    XMR  float64
    XPD  float64
    XPT  float64
    ZEC  float64
}


