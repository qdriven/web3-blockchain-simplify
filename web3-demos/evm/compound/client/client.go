package client

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"go-flashrunner/compound/models"
	"strconv"
)

// Client is used to interact with the compound.finance api
type Client struct {
	url         string
	client      *http.Client
	restyClient *resty.Client
}

// NewClient is used to instantiate a new go-compound client
func NewClient(url string) *Client {
	return &Client{url: url,
		client:      &http.Client{},
		restyClient: resty.New(),
	}
}

// GetTotalCollateralValueInEth is used to retrieve the total collateral value
// in eth that is owned by this account
func (c *Client) GetTotalCollateralValueInEth(address string) (float64, error) {
	resp, err := c.GetAccount(address)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(resp.Accounts[0].TotalCollateralValueInEth.Value, 64)
}

// GetTotalBorrowValueInEth is used to retrieve the total collateral value
// in eth that is owned by this account
func (c *Client) GetTotalBorrowValueInEth(address string) (float64, error) {
	resp, err := c.GetAccount(address)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(resp.Accounts[0].TotalBorrowValueInEth.Value, 64)
}

// GetAccount is used to retrieve information on a single account
func (c *Client) GetAccount(address string) (*models.AccountResponse, error) {
	call := fmt.Sprintf("/account?addresses[]=%s", address)
	apiURL := fmt.Sprintf("%s/%s", c.url, call)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.AccountResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetUnhealthyAccount(maxHealth, minBorrowValueInEth string) ([]models.Account, error) {
	request := &models.AccountRequest{
		MaxHealth:  models.ValueField{Value:maxHealth},
		MinBorrowValueInEth: models.ValueField{Value:minBorrowValueInEth},
		PageSize: 50,
		PageNumber: 1,
	}
	bytesReq,_:=json.Marshal(request)
	fmt.Println(string(bytesReq))
	accountRes := &models.AccountResponse{}
	apiURL := fmt.Sprintf("%s/account", c.url)
	fmt.Println(apiURL)
	resp, err := c.restyClient.R().EnableTrace().SetBody(request).SetResult(accountRes).Post(apiURL)
	c.printHttpTrace(resp,err)

	if err != nil {
		return nil, nil
	}
	log.Info(resp.String())
	unhealthyAccount := make([]models.Account,0)
	for _, account := range accountRes.Accounts {
		//log.Info("account index is ",index)
		if account.Health.Value!=""{
			unhealthyAccount= append(unhealthyAccount, account)
		}
	}
	return unhealthyAccount, nil
}

func (c *Client) printHttpTrace(resp *resty.Response,err error){
	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println(resp.Request.Body)
	fmt.Println(resp.Request.URL)
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

// GetAccounts is used to retrieve information on many accounts
func (c *Client) GetAccounts(pageSize, pageNum string) (*models.AccountResponse, error) {
	// https://api.compound.finance/api/v2/account
	if pageSize == "" {
		pageSize = "10"
	}
	if pageNum == "" {
		pageNum = "0"
	}
	apiURL := fmt.Sprintf("%s/account?page_size=%s&page_num=%s", c.url, pageSize, pageNum)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.AccountResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCToken is used to get information on a particular ctoken
func (c *Client) GetCToken(address string) (*models.CTokenResponse, error) {
	call := fmt.Sprintf("/ctoken?addresses[]=%s", address)
	apiURL := fmt.Sprintf("%s/%s", c.url, call)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.CTokenResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCTokens is used to retrieve information on many ctokens
func (c *Client) GetCTokens() (*models.CTokenResponse, error) {
	apiURL := fmt.Sprintf("%s/ctoken", c.url)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.CTokenResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}
