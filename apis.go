package wxcloudpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	"github.com/wzshiming/requests"
	"github.com/wzshiming/wxcloudpay/authen"
	"github.com/wzshiming/wxcloudpay/proto"
)

var client = requests.NewClient().SetLogLevel(requests.LogMessageAll)

const (
	BaseURL = `https://pay.qcloud.com/cpay/`
)

var (
	ErrResponseValidation = errors.New("cloudpay: Response validation failed")
)

type CloudPay struct {
	cli       *requests.Request
	authenKey []byte
}

func NewCloudPay(authenKey, signatureKey []byte) *CloudPay {
	return NewCloudPayWithURL(BaseURL, authenKey, signatureKey)
}

func NewCloudPayWithURL(baseURL string, authenKey, signatureKey []byte) *CloudPay {
	return &CloudPay{
		cli:       client.NewRequest().SetURLByStr(baseURL),
		authenKey: authenKey,
	}
}

func (c *CloudPay) requests(url string, req, resp interface{}, auth bool) error {
	ctx, err := json.Marshal(req)
	if err != nil {
		return err
	}
	requestsContent := *(*string)(unsafe.Pointer(&ctx))

	authenInfo := &proto.AuthenInfo{
		A: &proto.Authen{
			AuthenCode: authen.HmacSha256(ctx, c.authenKey),
			AuthenType: proto.AuthenTypeHmacSha256,
		},
	}

	request := &proto.Sdk2CloudPayRequest{
		RequestContent: requestsContent,
		AuthenInfo:     authenInfo,
	}

	body, err := c.cli.Clone().
		SetJSON(request).
		Post(url)
	if err != nil {
		return err
	}
	if !auth {
		return json.Unmarshal(body.Body(), resp)
	}
	response := &proto.CloudPay2SdkResponse{}

	err = json.Unmarshal(body.Body(), &response)
	if err != nil {
		return err
	}

	if response.ResponseContent == "" {
		return ErrResponseValidation
	}

	responseContent := *(*[]byte)(unsafe.Pointer(&response.ResponseContent))

	// 返回验证
	if authenInfo := response.AuthenInfo; authenInfo != nil {
		if a := authenInfo.A; a != nil {
			if a.AuthenType == proto.AuthenTypeHmacSha256 {
				if authen.HmacSha256(responseContent, c.authenKey) != a.AuthenCode {
					return ErrResponseValidation
				}
			}
		}
	}

	err = json.Unmarshal(responseContent, resp)
	if err != nil {
		return err
	}
	if resp, ok := resp.(proto.IsResponseStatus); ok {
		responseStatus := resp.ResponseStatus()
		if responseStatus.Status != 0 {
			return fmt.Errorf("Error %d: %s", responseStatus.Status, responseStatus.Description)
		}
	}
	return nil
}

func (c *CloudPay) Ping() error {
	resp := &proto.PingResponse{}
	err := c.requests(PingURL, nil, &resp, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *CloudPay) CloseOrder(req *proto.CloseOrderRequest) (*proto.CloseOrderResponseDetail, error) {
	resp := &proto.CloseOrderResponse{}
	err := c.requests(CloseOrderURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.CloseOrder, nil
}

func (c *CloudPay) MicroPay(req *proto.MicroPayRequest) (*proto.MicroPayResponseDetail, error) {
	resp := &proto.MicroPayResponse{}
	err := c.requests(MicroPayURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.MicroPay, nil
}

func (c *CloudPay) OrderDetail(req *proto.OrderDetailRequest) (*proto.OrderDetailResponseDetail, error) {
	resp := &proto.OrderDetailResponse{}
	err := c.requests(OrderDetailURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.OrderDetailQuery, nil
}

func (c *CloudPay) QueryOrder(req *proto.QueryOrderRequest) (*proto.QueryOrderResponseDetail, error) {
	resp := &proto.QueryOrderResponse{}
	err := c.requests(QueryOrderURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryOrder, nil
}

func (c *CloudPay) QueryRefund(req *proto.QueryRefundRequest) (*proto.QueryRefundResponseDetail, error) {
	resp := &proto.QueryRefundResponse{}
	err := c.requests(QueryRefundURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryRefundOrder, nil
}

func (c *CloudPay) QueryShopInfo(req *proto.QueryShopInfoRequest) (*proto.QueryShopInfoResponseDetail, error) {
	resp := &proto.QueryShopInfoResponse{}
	err := c.requests(QueryShopInfoURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryShopInfo, nil
}

func (c *CloudPay) Refund(req *proto.RefundRequest) (*proto.RefundResponseDetail, error) {
	resp := &proto.RefundResponse{}
	err := c.requests(RefundURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.Refund, nil
}

func (c *CloudPay) ScanCodePay(req *proto.ScanCodePayRequest) (*proto.ScanCodePayResponseDetail, error) {
	resp := &proto.ScanCodePayResponse{}
	err := c.requests(ScanCodePayURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.ScanCodePay, nil
}

func (c *CloudPay) SetShopInfo(req *proto.SetShopInfoRequest) (*proto.SetShopInfoResponseDetail, error) {
	resp := &proto.SetShopInfoResponse{}
	err := c.requests(SetShopInfoURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.SetShopInfo, nil
}
