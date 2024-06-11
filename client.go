package ninjarmm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sourcegraph/conc/iter"
	"github.com/stellaraf/go-ninjarmm/internal/auth"
	"github.com/stellaraf/go-ninjarmm/internal/check"
	"github.com/stellaraf/go-ninjarmm/internal/types"
	"github.com/stellaraf/go-ninjarmm/internal/util"
	"github.com/stellaraf/go-utils"
)

type Client struct {
	auth       *auth.Auth
	baseURL    string
	httpClient *resty.Client
}

func handleResponse(response *resty.Response, data any) error {
	err := check.ForError(response)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response.Body(), data)
	if err != nil {
		return err
	}
	return nil
}

// Location retrieves a location's details.
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

// OrganizationLocations retrieves all locations belonging to an organization.
func (client *Client) OrganizationLocations(orgID int) ([]Location, error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d/locations", orgID))
	if err != nil {
		return nil, err
	}
	var locations []Location
	err = handleResponse(res, &locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// OrganizationDevices retrieves all devices belonging to an organization.
func (client *Client) OrganizationDevices(orgID int) ([]Device, error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d/devices", orgID))
	if err != nil {
		return nil, err
	}
	var devices []Device
	err = handleResponse(res, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

// Organizations retrieves a summary for all organizations.
func (client *Client) Organizations() ([]OrganizationSummary, error) {
	res, err := client.httpClient.R().Get("/api/v2/organizations")
	if err != nil {
		return nil, err
	}
	var orgs []OrganizationSummary
	err = handleResponse(res, &orgs)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

// Organization retrieves an organization's details.
func (client *Client) Organization(id int) (org Organization, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/organization/%d", id))
	if err != nil {
		return
	}
	err = handleResponse(res, &org)
	return
}

// Device retrieves a device's details.
func (client *Client) Device(id int) (device DeviceDetails, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/device/%d", id))
	if err != nil {
		return
	}
	err = handleResponse(res, &device)
	return
}

// Devices retrieves all devices matching the filter. If no filter is provided, all devices will be
// returned.
func (client *Client) Devices(df *deviceFilter) (Devices, error) {
	req := client.httpClient.R()
	if df != nil {
		req.SetQueryParam("df", df.String())
	}
	res, err := req.Get("/api/v2/devices")
	if err != nil {
		return nil, err
	}
	err = check.ForError(res)
	if err != nil {
		return nil, err
	}
	var devices []Device
	err = json.Unmarshal(res.Body(), &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

// DeviceCustomFields retrieves custom fields for a device.
func (client *Client) DeviceCustomFields(id int) (customFields map[string]any, err error) {
	res, err := client.httpClient.R().Get(fmt.Sprintf("/api/v2/device/%d/custom-fields", id))
	if err != nil {
		return
	}
	err = handleResponse(res, &customFields)
	return
}

// Roles retrieves roles and optionally filters the results based on role name. The role name
// provided will be matched against the exact name provided as well as the name in uppercase and
// underscore separated. For example:
//
//	"Windows Server"
//	// will match both:
//	"Windows Server"
//	"WINDOWS_SERVER"
func (client *Client) Roles(filter ...string) ([]Role, error) {
	req := client.httpClient.R().SetResult([]Role{})
	res, err := req.Get("/api/v2/roles")
	if err != nil {
		return nil, err
	}
	var roles []Role
	err = handleResponse(res, &roles)
	if err != nil {
		return nil, err
	}
	if len(filter) != 0 {
		filtered := make([]Role, 0, len(roles))
		for _, f := range filter {
			for _, role := range roles {
				if util.MatchWithUpper(role.Name, f) {
					filtered = append(filtered, role)
				}
			}
		}
		return filtered, nil
	}
	return roles, nil
}

// Role retrieves a role by ID.
func (client *Client) Role(id int) (*Role, error) {
	roles, err := client.Roles()
	if err != nil {
		return nil, err
	}
	for _, role := range roles {
		if role.ID == id {
			return &role, nil
		}
	}
	return nil, fmt.Errorf("role with ID '%d' not found", id)
}

// SetDeviceRole sets the role of a device.
func (client *Client) SetDeviceRole(deviceID int, roleID int) error {
	b := make(map[string]int, 1)
	b["nodeRoleId"] = roleID
	req := client.httpClient.R().SetBody(b)
	res, err := req.Patch(fmt.Sprintf("/api/v2/device/%d", deviceID))
	if err != nil {
		return err
	}
	if res.IsError() {
		err = check.ForError(res)
		if err != nil {
			return err
		}
	}
	return nil
}

// OSPatches retrieves an OS patch summary for an organization.
func (client *Client) OSPatches(orgID int) (patchReport OSPatchReportQuery, err error) {
	df := NewDeviceFilter().Org(EQ, orgID)
	res, err := client.httpClient.R().SetQueryParam("df", df.Encode()).Get("/api/v2/queries/os-patches")
	if err != nil {
		return
	}
	err = handleResponse(res, &patchReport)
	return
}

// OSPatchReport retrieves a patch report for an organization.
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

// CreateOrganization creates a new organization.
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
	err = handleResponse(res, &org)
	return
}

// ScheduleMaintenance schedules a maintenance window for a device.
func (client *Client) ScheduleMaintenance(deviceID int, start, end time.Time, disabledFeatures []string) error {
	body := &MaintenanceRequest{
		Start:            start,
		End:              end,
		DisabledFeatures: disabledFeatures,
	}
	req := client.httpClient.R().SetError(&types.NinjaRMMPutError{}).SetBody(body)
	res, err := req.Put(fmt.Sprintf("/api/v2/device/%d/maintenance", deviceID))
	if err != nil {
		return err
	}
	if res.IsError() {
		parsed := res.Error().(*types.NinjaRMMPutError)
		err = fmt.Errorf(parsed.GetErrorMessage(deviceID))
		return err
	}
	return nil
}

// CancelMaintenance cancels a scheduled maintenance for a device.
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

// SoftwareInventory retrieves the software inventory for a set of devices.
func (client *Client) SoftwareInventory(filter *deviceFilter) ([]SoftwareInventoryResult, error) {
	qc := NewQueryClient[SoftwareInventoryResult](client, DefaultQueryBatchSize)
	var q map[string]string = nil
	if filter != nil {
		q = map[string]string{"df": filter.String()}
	}
	res, err := qc.Do("/api/v2/queries/software", q)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

// DevicesWithSoftware retrieves the software inventory for a set of devices and returns the device
// IDs for devices that have a specific software.
func (client *Client) DevicesWithSoftware(name string, filter *deviceFilter) ([]int, error) {
	all, err := client.Devices(filter)
	if err != nil {
		return nil, err
	}
	final := make([]int, 0, len(all))
	chunks := utils.ChunkSlice(all, DefaultQueryBatchSize)

	n := 0
	results, err := iter.MapErr(chunks, func(devices *[]Device) ([]int, error) {
		n++
		ids := make([]int, 0, len(*devices))
		for _, d := range *devices {
			d := d
			ids = append(ids, d.ID)
		}
		df := NewDeviceFilter().ID(IN, ids...)
		log.Println(n, "-", df.String())
		inventory, err := client.SoftwareInventory(df)
		if err != nil {
			return nil, err
		}
		log.Println(n, "-", fmt.Sprintf("len(inventory)=%d", len(inventory)))
		r := make([]int, 0, len(*devices))
		for _, pkg := range inventory {
			pkg := pkg
			did := int(pkg.DeviceID)
			if pkg.Name == name {
				log.Println(n, "-", pkg.DeviceID, pkg.Name)
				r = append(r, did)
			}
		}
		return r, nil
	})
	if err != nil {
		return nil, err
	}
	for _, batch := range results {
		batch := batch
		final = append(final, batch...)
	}
	return utils.Set(final), nil
}

// New creates a new NinjaRMMClient.
func New(
	baseURL, clientID, clientSecret string,
	encryption *string,
	tokenCache TokenCache) (*Client, error) {

	auth, err := auth.New(
		baseURL,
		clientID,
		clientSecret,
		encryption,
		tokenCache,
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
