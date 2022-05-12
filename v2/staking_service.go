package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListStakingProductsService https://binance-docs.github.io/apidocs/spot/en/#get-staking-product-list-user_data
type ListStakingProductsService struct {
	c       *Client
	product string
	asset   string
	current int64
	size    int64
}

// Status represent the product ("STAKING" for Locked Staking, "F_DEFI" for flexible DeFi Staking, "L_DEFI" for locked DeFi Staking)
func (s *ListStakingProductsService) Product(product string) *ListStakingProductsService {
	s.product = product
	return s
}

func (s *ListStakingProductsService) Asset(asset string) *ListStakingProductsService {
	s.asset = asset
	return s
}

// Current query page. Default: 1, Min: 1
func (s *ListStakingProductsService) Current(current int64) *ListStakingProductsService {
	s.current = current
	return s
}

// Size Default: 10, Max: 100
func (s *ListStakingProductsService) Size(size int64) *ListStakingProductsService {
	s.size = size
	return s
}

// Do send request
func (s *ListStakingProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*StakingProduct, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/productList",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.product != "" {
		m["product"] = s.product
	}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*StakingProduct
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// StakingProduct define a staking product
type StakingProduct struct {
	ProjectId string `json:"projectId"`
	Detail    struct {
		Asset       string `json:"asset"`
		RewardAsset string `json:"rewardAsset"`
		Duration    int    `json:"duration"`
		Renewable   bool   `json:"renewable"`
		Apy         string `json:"apy"`
	} `json:"detail"`
	Quota struct {
		TotalPersonalQuota string `json:"totalPersonalQuota"`
		Minimum            string `json:"minimum"`
	} `json:"quota"`
}

// PurchaseStakingProductService https://binance-docs.github.io/apidocs/spot/en/#purchase-staking-product-user_data
type PurchaseStakingProductService struct {
	c         *Client
	product   string
	productId string
	amount    float64
}

// Product "STAKING" for Locked Staking, "F_DEFI" for flexible DeFi Staking, "L_DEFI" for locked DeFi Staking
func (s *PurchaseStakingProductService) Product(product string) *PurchaseStakingProductService {
	s.product = product
	return s
}

// Product represent the id of the staking product to purchase
func (s *PurchaseStakingProductService) ProductId(productId string) *PurchaseStakingProductService {
	s.productId = productId
	return s
}

// Amount is the quantity of the product to purchase
func (s *PurchaseStakingProductService) Amount(amount float64) *PurchaseStakingProductService {
	s.amount = amount
	return s
}

// Do send request
func (s *PurchaseStakingProductService) Do(ctx context.Context, opts ...RequestOption) (uint64, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/staking/purchase",
		secType:  secTypeSigned,
	}
	m := params{
		"product":   s.product,
		"productId": s.productId,
		"amount":    s.amount,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return 0, err
	}

	var res *PurchaseStakingProductResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return 0, err
	}

	return res.PurchaseId, nil
}

type PurchaseStakingProductResponse struct {
	PurchaseId uint64 `json:"purchaseId"`
}

// GetStakingPersonalLeftQuota https://binance-docs.github.io/apidocs/spot/en/#get-personal-left-quota-of-staking-product-user_data
type GetStakingPersonalLeftQuota struct {
	c         *Client
	product   string
	productId string
}

// Product represent the product ("STAKING" for Locked Staking, "F_DEFI" for flexible DeFi Staking, "L_DEFI" for locked DeFi Staking)
func (s *GetStakingPersonalLeftQuota) Product(product string) *GetStakingPersonalLeftQuota {
	s.product = product
	return s
}

// ProductId represent the id of the staking product for left quota
func (s *GetStakingPersonalLeftQuota) ProductId(productId string) *GetStakingPersonalLeftQuota {
	s.productId = productId
	return s
}

// Do send request
func (s *GetStakingPersonalLeftQuota) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/personalLeftQuota",
		secType:  secTypeSigned,
	}
	m := params{
		"product":   s.product,
		"productId": s.productId,
	}
	r.setFormParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	var res StakingLeftQuotaResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return "", err
	}

	return res.LeftPersonalQuota, nil
}

type StakingLeftQuotaResponse struct {
	LeftPersonalQuota string `json:"leftPersonalQuota"`
}

// GetStakingProductPosition https://binance-docs.github.io/apidocs/spot/en/#get-staking-product-position-user_data
type GetStakingProductPosition struct {
	c         *Client
	product   string
	productId string
	asset     string
	current   int64
	size      int64
}

// Status represent the product ("STAKING" for Locked Staking, "F_DEFI" for flexible DeFi Staking, "L_DEFI" for locked DeFi Staking)
func (s *GetStakingProductPosition) Product(product string) *GetStakingProductPosition {
	s.product = product
	return s
}

// Product represent the id of the staking product to purchase
func (s *GetStakingProductPosition) ProductId(productId string) *GetStakingProductPosition {
	s.productId = productId
	return s
}

func (s *GetStakingProductPosition) Asset(asset string) *GetStakingProductPosition {
	s.asset = asset
	return s
}

// Current query page. Default: 1, Min: 1
func (s *GetStakingProductPosition) Current(current int64) *GetStakingProductPosition {
	s.current = current
	return s
}

// Size Default: 10, Max: 100
func (s *GetStakingProductPosition) Size(size int64) *GetStakingProductPosition {
	s.size = size
	return s
}

// Do send request
func (s *GetStakingProductPosition) Do(ctx context.Context, opts ...RequestOption) ([]*StakingProductPositionResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/position",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.product != "" {
		m["product"] = s.product
	}
	if s.product != "" {
		m["productId"] = s.productId
	}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*StakingProductPositionResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type StakingProductPositionResponse struct {
	PositionID        uint64 `json:"positionId"`
	ProductID         string `json:"productId"`
	Asset             string `json:"asset"`
	Amount            string `json:"amount"`
	PurchaseTime      int64  `json:"purchaseTime"`
	Duration          int    `json:"duration"`
	AccrualDays       int    `json:"accrualDays"`
	RewardAsset       string `json:"rewardAsset"`
	RewardAmt         string `json:"rewardAmt"`
	NextInterestPay   string `json:"nextInterestPay"`
	PayInterestPeriod int    `json:"payInterestPeriod"`
	RedeemAmountEarly string `json:"redeemAmountEarly"`
	InterestEndDate   int64  `json:"interestEndDate"`
	DeliverDate       int64  `json:"deliverDate"`
	RedeemPeriod      int    `json:"redeemPeriod"`
	CanRedeemEarly    bool   `json:"canRedeemEarly"`
	Renewable         bool   `json:"renewable"`
	Type              string `json:"type"`
	Status            string `json:"status"`
}

// GetStakingHistory https://binance-docs.github.io/apidocs/spot/en/#get-staking-history-user_data
type GetStakingHistory struct {
	c         *Client
	product   string
	txnType   string
	asset     string
	startTime *int64
	endTime   *int64
	current   int64
	size      int64
}

// Product set product ("STAKING" for Locked Staking, "F_DEFI" for flexible DeFi Staking, "L_DEFI" for locked DeFi Staking)
func (s *GetStakingHistory) Product(product string) *GetStakingHistory {
	s.product = product
	return s
}

// Type set txn type ("SUBSCRIPTION", "REDEMPTION", "INTEREST")
func (s *GetStakingHistory) Type(accountType string) *GetStakingHistory {
	s.txnType = accountType
	return s
}

// Type set txn type ("SUBSCRIPTION", "REDEMPTION", "INTEREST")
func (s *GetStakingHistory) Asset(asset string) *GetStakingHistory {
	s.asset = asset
	return s
}

// StartTime set starttime
func (s *GetStakingHistory) StartTime(startTime int64) *GetStakingHistory {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *GetStakingHistory) EndTime(endTime int64) *GetStakingHistory {
	s.endTime = &endTime
	return s
}

// Current query page. Default: 1, Min: 1
func (s *GetStakingHistory) Current(current int64) *GetStakingHistory {
	s.current = current
	return s
}

// Size Default: 10, Max: 100
func (s *GetStakingHistory) Size(size int64) *GetStakingHistory {
	s.size = size
	return s
}

// Do send request
func (s *GetStakingHistory) Do(ctx context.Context, opts ...RequestOption) ([]*StakingHistoryResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/stakingRecord",
		secType:  secTypeSigned,
	}
	m := params{
		"product": s.product,
		"txnType": s.txnType,
	}
	if s.asset != "" {
		m["asset"] = s.asset
	}
	if s.startTime != nil {
		m["startTime"] = s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = s.endTime
	}
	if s.current != 0 {
		m["current"] = s.current
	}
	if s.size != 0 {
		m["size"] = s.size
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*StakingHistoryResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type StakingHistoryResponse struct {
	PositionId  string `json:"positionId"`
	Time        int64  `json:"time"`
	Asset       string `json:"asset"`
	Project     string `json:"project"`
	Amount      string `json:"amount"`
	LockPeriod  string `json:"lockPeriod"`
	DeliverDate string `json:"deliverDate"`
	Type        string `json:"auto"`
	Status      string `json:"status"`
}

// GetStakingHistory https://binance-docs.github.io/apidocs/spot/en/#get-staking-history-user_data
type GetStakingLeftDailyPurchaseQuota struct {
	c         *Client
	productId string
}

// ProductId to resolve quota for product
func (s *GetStakingLeftDailyPurchaseQuota) Product(productId string) *GetStakingLeftDailyPurchaseQuota {
	s.productId = productId
	return s
}

// Do send request
func (s *GetStakingLeftDailyPurchaseQuota) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/daily/userLeftQuota",
		secType:  secTypeSigned,
	}
	m := params{
		"productId": s.productId,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	res := struct {
		LeftQuota string
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return "", err
	}

	return res.LeftQuota, nil
}
