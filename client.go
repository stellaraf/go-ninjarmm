package ninjarmm

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-utils"
)

type Client struct {
	auth       *authT
	baseURL    string
	httpClient *resty.Client
}

func (client *Client) handleResponse(response *resty.Response, data any) (err error) {
	err = checkForError(response)
	if err != nil {
		return
	}
	err = json.Unmarshal(response.Body(), data)
	return
}

func (client *Client) Location(orgID, locID int) (location *Location, err error) {
	locations, err := client.OrganizationLocations(orgID)
	if err != nil {
		return
	}
	for _, loc := range locations {
		if loc.ID == locID {
			location = &loc
			return
		}
	}
	if location == nil {
		err = fmt.Errorf("location with id '%d' not found in organization '%d'", locID, orgID)
	}
	return
}

func (client *Client) OrganizationLocations(orgID int) (locations []Location, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d/locations", orgID))
	if err != nil {
		return
	}
	err = client.handleResponse(res, &locations)
	return
}

func (client *Client) Organizations() (orgs []OrganizationSummary, err error) {
	res, err := client.httpClient.R().Get("/api/v2/organizations")
	if err != nil {
		return
	}
	err = client.handleResponse(res, &orgs)
	return
}

func (client *Client) Organization(id int) (org Organization, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d", id))
	if err != nil {
		return
	}
	err = client.handleResponse(res, &org)
	return
}

func (client *Client) Device(id int) (device DeviceDetails, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/device/%d", id))
	if err != nil {
		return
	}
	err = client.handleResponse(res, &device)
	return
}

func (client *Client) OSPatches(orgId int) (patchReport OSPatchReportQuery, err error) {
	q := url.Values{}
	q.Add("org", fmt.Sprintf("%d", orgId))
	res, err := client.httpClient.R().SetQueryParam("df", q.Encode()).Get("/api/v2/queries/os-patches")
	if err != nil {
		return
	}
	err = client.handleResponse(res, &patchReport)
	return
}

func (client *Client) OSPatchReport(orgId int) (patchReport []OSPatchReportDetail, err error) {
	reports, err := client.OSPatches(orgId)
	if err != nil {
		return
	}
	devicesToCollect := []int{}
	for _, report := range reports.Results {
		if !utils.SliceContains(devicesToCollect, report.DeviceID) {
			devicesToCollect = append(devicesToCollect, report.DeviceID)
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

		device, hasKey := deviceMap[report.DeviceID]
		if !hasKey {
			err = fmt.Errorf("failed to get details for device '%d'", report.DeviceID)
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

func (client *Client) CreateOrganization(name string) (org Organization, err error) {
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

// New creates a new NinjaRMMClient.
func New(
	baseURL, clientID, clientSecret string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback) (client *Client, err error) {

	auth, err := newAuth(
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
	client = &Client{auth: auth, baseURL: baseURL, httpClient: httpClient}
	return
}
