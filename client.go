package ninjarmm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-utils"
)

type Client struct {
	auth       *authT
	baseURL    string
	httpClient *resty.Client
}

func (client *Client) handleResponse(response *resty.Response, data any) error {
	err := checkForError(response)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response.Body(), data)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) Location(orgID, locID int) (*Location, error) {
	locations, err := client.OrganizationLocations(orgID)
	if err != nil {
		return nil, err
	}
	for _, loc := range locations {
		if loc.ID == locID {
			return &loc, nil
		}
	}
	err = fmt.Errorf("location with id '%d' not found in organization '%d'", locID, orgID)
	return nil, err
}

func (client *Client) OrganizationLocations(orgID int) ([]Location, error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d/locations", orgID))
	if err != nil {
		return nil, err
	}
	var locations []Location
	err = client.handleResponse(res, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (client *Client) OrganizationDevices(orgID int) ([]Device, error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d/devices", orgID))
	if err != nil {
		return nil, err
	}
	var devices []Device
	err = client.handleResponse(res, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (client *Client) Organizations() ([]OrganizationSummary, error) {
	res, err := client.httpClient.R().Get("/api/v2/organizations")
	if err != nil {
		return nil, err
	}
	var orgs []OrganizationSummary
	err = client.handleResponse(res, &orgs)
	if err != nil {
		return nil, err
	}
	return orgs, nil
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

func (client *Client) DeviceCustomFields(id int) (customFields map[string]any, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/device/%d/custom-fields", id))
	if err != nil {
		return
	}
	err = client.handleResponse(res, &customFields)
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

func (client *Client) OSPatchReport(orgId int) ([]OSPatchReportDetail, error) {
	reports, err := client.OSPatches(orgId)
	if err != nil {
		return nil, err
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
			return nil, err
		}
		deviceMap[deviceId] = device
	}
	if len(deviceMap) != len(devicesToCollect) {
		err = fmt.Errorf("failed to collect device details for Organization '%d'", orgId)
		return nil, err
	}
	patchReport := make([]OSPatchReportDetail, 0, len(reports.Results))
	for _, report := range reports.Results {
		device, hasKey := deviceMap[report.DeviceID]
		if !hasKey {
			err = fmt.Errorf("failed to get details for device '%d'", report.DeviceID)
			return nil, err
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
	return patchReport, nil
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

func (client *Client) ScheduleMaintenance(deviceID int, start, end time.Time, disabledFeatures []string) error {
	body := &MaintenanceRequest{
		Start:            start,
		End:              end,
		DisabledFeatures: disabledFeatures,
	}
	req := client.httpClient.R().SetError(&NinjaRMMPutError{}).SetBody(body)
	res, err := req.Put(fmt.Sprintf("/api/v2/device/%d/maintenance", deviceID))
	if err != nil {
		return err
	}
	if res.IsError() {
		parsed := res.Error().(*NinjaRMMPutError)
		err = fmt.Errorf(parsed.GetErrorMessage(deviceID))
		return err
	}
	return nil
}

func (client *Client) CancelMaintenance(deviceID int) error {
	req := client.httpClient.R()
	res, err := req.Delete(fmt.Sprintf("/api/v2/device/%d/maintenance", deviceID))
	if err != nil {
		return err
	}
	if res.StatusCode() > 299 {
		b := string(res.Body())
		err = fmt.Errorf("failed to delete maintenance for device '%d' due to error '%s'", deviceID, b)
		return err
	}
	return nil
}

// New creates a new NinjaRMMClient.
func New(
	baseURL, clientID, clientSecret string,
	encryption *string,
	getAccessTokenCallback CachedTokenCallback,
	setAccessTokenCallback SetTokenCallback,
	getRefreshTokenCallback CachedTokenCallback,
	setRefreshTokenCallback SetTokenCallback) (*Client, error) {

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
		return nil, err
	}
	if auth == nil {
		err = fmt.Errorf("failed to initialize authentication")
		return nil, err
	}
	httpClient := resty.New()
	httpClient.SetBaseURL(baseURL)
	token, err := auth.GetAccessToken()
	if err != nil {
		return nil, err
	}
	httpClient.SetAuthToken(token)
	httpClient.AddRetryCondition(func(res *resty.Response, err error) bool {
		return res.StatusCode() == http.StatusUnauthorized
	})
	httpClient.AddRetryHook(func(res *resty.Response, err error) {
		if res.StatusCode() == http.StatusUnauthorized {
			token, err := auth.GetAccessToken()
			if err == nil {
				httpClient.SetAuthToken(token)
			}
		}
	})
	client := &Client{auth: auth, baseURL: baseURL, httpClient: httpClient}
	return client, err
}
