package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"

	"github.com/go-resty/resty/v2"
)

type NinjaRMMClient struct {
	auth       *NinjaRMMAuth
	baseURL    string
	httpClient *resty.Client
}

func (client *NinjaRMMClient) handleResponse(response *resty.Response, data any) (err error) {
	err = checkForError(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response.Body(), data)
	return
}

func (client *NinjaRMMClient) Organizations() (orgs []OrganizationSummary, err error) {
	res, err := client.httpClient.R().Get("/api/v2/organizations")
	if err != nil {
		return
	}
	err = client.handleResponse(res, &orgs)
	return
}

func (client *NinjaRMMClient) Organization(id int) (org Organization, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d", id))
	if err != nil {
		return
	}
	err = client.handleResponse(res, &org)
	return
}

func (client *NinjaRMMClient) Device(id int) (device DeviceDetails, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/device/%d", id))
	if err != nil {
		return
	}
	err = client.handleResponse(res, &device)
	return
}

func (client *NinjaRMMClient) OSPatches(orgId int) (patchReport OSPatchReportQuery, err error) {
	q := url.Values{}
	q.Add("org", fmt.Sprintf("%d", orgId))
	res, err := client.httpClient.R().SetQueryParam("df", q.Encode()).Get("/api/v2/queries/os-patches")
	if err != nil {
		return
	}
	err = client.handleResponse(res, &patchReport)
	return
}

func (client *NinjaRMMClient) OSPatchReport(orgId int) (patchReport []OSPatchReportDetail, err error) {
	reports, err := client.OSPatches(orgId)
	if err != nil {
		return
	}
	devicesToCollect := []int{}
	for _, report := range reports.Results {
		_id, err := report.DeviceID.Int64()
		if err != nil {
			return patchReport, err
		}
		id := int(_id)
		if !arrayContains(devicesToCollect, id) {
			devicesToCollect = append(devicesToCollect, id)
		}
	}
	sort.Ints(devicesToCollect)
	deviceMap := make(map[int]DeviceDetails)
	for _, deviceId := range devicesToCollect {
		device, err := client.Device(deviceId)
		if err != nil {
			return patchReport, err
		}
		deviceMap[deviceId] = device
	}
	if len(deviceMap) != len(devicesToCollect) {
		err = fmt.Errorf("failed to collect device details for Organization '%d'", orgId)
		return
	}
	for _, report := range reports.Results {
		_id, err := report.DeviceID.Int64()
		if err != nil {
			return patchReport, err
		}
		id := int(_id)
		device, hasKey := deviceMap[id]
		if !hasKey {
			err = fmt.Errorf("failed to get details for device '%d'", id)
			return patchReport, err
		}
		result := OSPatchReportDetail{
			ID:        report.ID,
			Name:      report.Name,
			Severity:  report.Severity,
			Status:    report.Status,
			Type:      report.Type,
			KBNumber:  report.KBNumber,
			Timestamp: report.Timestamp,
			Device:    device,
		}
		patchReport = append(patchReport, result)
	}
	return
}

func (client *NinjaRMMClient) CreateOrganization(name string) (org Organization, err error) {
	orgs, err := client.Organizations()
	if err != nil {
		return
	}
	matchingOrgName := ""
	matchingOrgID := 0
	for _, org := range orgs {
		if org.Name == name {
			matchingOrgName = org.Name
			matchingOrgID = org.ID
		}
	}
	if matchingOrgName != "" {
		err = fmt.Errorf("object with name '%s' already exists (ID '%d'). A new object will not be created", matchingOrgName, matchingOrgID)
		return
	}
	res, err := client.httpClient.R().SetBody(struct{ Name string }{Name: name}).Post("/v2/organizations")
	if err != nil {
		return
	}
	err = client.handleResponse(res, &org)
	return
}

func CreateNinjaRMMClient(
	baseURL, clientID, clientSecret string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback) (client *NinjaRMMClient, err error) {

	auth, err := createNinjaRMMAuth(
		baseURL,
		clientID,
		clientSecret,
		encryption,
		getAccessTokenCallback,
		setAccessTokenCallback,
		getRefreshTokenCallback,
		setRefreshTokenCallback,
	)
	if err != nil {
		return
	}
	if auth == nil {
		err = fmt.Errorf("failed to initialize authentication")
		return
	}
	httpClient := resty.New()
	httpClient.SetBaseURL(baseURL)
	token, err := auth.GetAccessToken()
	if err != nil {
		return
	}
	httpClient.SetAuthToken(token)
	client = &NinjaRMMClient{auth: auth, baseURL: baseURL, httpClient: httpClient}
	return
}
