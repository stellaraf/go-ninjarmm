package ninjarmm

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sourcegraph/conc/iter"
	"github.com/stellaraf/go-ninjarmm/internal/auth"
	"github.com/stellaraf/go-ninjarmm/internal/util"
	"github.com/stellaraf/go-utils"
)

type Client struct {
	auth       *auth.Auth
	baseURL    string
	httpClient *resty.Client
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
	res, err := client.httpClient.R().
		SetResult([]Location{}).
		SetError(&Error{}).
		Get(fmt.Sprintf("/api/v2/organization/%d/locations", orgID))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return *res.Result().(*[]Location), nil
}

// OrganizationDevices retrieves all devices belonging to an organization.
func (client *Client) OrganizationDevices(orgID int) ([]Device, error) {
	res, err := client.httpClient.R().
		SetResult([]Device{}).
		SetError(&Error{}).
		Get(fmt.Sprintf("/api/v2/organization/%d/devices", orgID))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return *res.Result().(*[]Device), nil
}

// Organizations retrieves a summary for all organizations.
func (client *Client) Organizations() ([]OrganizationSummary, error) {
	res, err := client.httpClient.R().
		SetResult([]OrganizationSummary{}).
		SetError(&Error{}).
		Get("/api/v2/organizations")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return *res.Result().(*[]OrganizationSummary), nil
}

// Organization retrieves an organization's details.
func (client *Client) Organization(id int) (*Organization, error) {
	res, err := client.httpClient.R().SetResult(Organization{}).SetError(&Error{}).Get(fmt.Sprintf("/api/v2/organization/%d", id))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return res.Result().(*Organization), nil
}

// Device retrieves a device's details.
func (client *Client) Device(id int) (*DeviceDetails, error) {
	res, err := client.httpClient.R().SetResult(&DeviceDetails{}).SetError(&Error{}).Get(fmt.Sprintf("/api/v2/device/%d", id))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return res.Result().(*DeviceDetails), nil
}

// Devices retrieves all devices matching the filter. If no filter is provided, all devices will be
// returned.
func (client *Client) Devices(df *deviceFilter) (Devices, error) {
	req := client.httpClient.R().SetResult(Devices{}).SetError(&Error{})
	if df != nil {
		req.SetQueryParam("df", df.String())
	}
	res, err := req.Get("/api/v2/devices")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return *res.Result().(*Devices), nil
}

// SearchDevices finds devices with a System Name matching a regex pattern.
func (client *Client) SearchDevices(name *regexp.Regexp, df *deviceFilter) (Devices, error) {
	all, err := client.Devices(df)
	if err != nil {
		return nil, err
	}
	return all.MatchName(name), nil
}

// DeviceCustomFields retrieves custom fields for a device.
func (client *Client) DeviceCustomFields(id int) (map[string]any, error) {
	res, err := client.httpClient.R().
		SetResult(map[string]any{}).
		SetError(&Error{}).
		Get(fmt.Sprintf("/api/v2/device/%d/custom-fields", id))
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return *res.Result().(*map[string]any), nil
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
	req := client.httpClient.R().
		SetResult([]Role{}).
		SetError(&Error{})
	res, err := req.Get("/api/v2/roles")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	roles := *res.Result().(*[]Role)
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
	req := client.httpClient.R().SetBody(b).SetError(&Error{})
	res, err := req.Patch(fmt.Sprintf("/api/v2/device/%d", deviceID))
	if err != nil {
		return err
	}
	if res.IsError() {
		return res.Error().(*Error)
	}
	return nil
}

// OSPatches retrieves an OS patch summary for devices matching a device filter.
func (client *Client) OSPatches(df *deviceFilter) (*OSPatchReportQuery, error) {
	req := client.httpClient.R().SetResult(OSPatchReportQuery{}).SetError(&Error{})
	if df != nil {
		req.SetQueryParam("df", df.String())
	}
	res, err := req.Get("/api/v2/queries/os-patches")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return res.Result().(*OSPatchReportQuery), nil
}

// OSPatchReport retrieves a patch report for devices matching a device filter.
func (client *Client) OSPatchReport(df *deviceFilter) ([]OSPatchReportDetail, error) {
	reports, err := client.OSPatches(df)
	if err != nil {
		return nil, err
	}
	devicesToCollect := make([]int, 0)
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
		deviceMap[deviceId] = *device
	}
	if len(deviceMap) != len(devicesToCollect) {
		err = fmt.Errorf("failed to collect details for devices matching filter '%s'", df.String())
		return nil, err
	}
	patchReport := make([]OSPatchReportDetail, 0, len(reports.Results))
	for _, report := range reports.Results {
		device, hasKey := deviceMap[report.DeviceID]
		if !hasKey {
			err = fmt.Errorf("failed to get details for device %d", report.DeviceID)
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
func (client *Client) CreateOrganization(name string) (*Organization, error) {
	orgs, err := client.Organizations()
	if err != nil {
		return nil, err
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
		return nil, err
	}
	res, err := client.httpClient.R().SetBody(struct{ Name string }{Name: name}).SetResult(&Organization{}).SetError(&Error{}).Post("/v2/organizations")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, res.Error().(*Error)
	}
	return res.Result().(*Organization), nil
}

// ScheduleMaintenance schedules a maintenance window for a device.
func (client *Client) ScheduleMaintenance(deviceID int, start, end time.Time, disabledFeatures []string) error {
	body := &MaintenanceRequest{
		Start:            start,
		End:              end,
		DisabledFeatures: disabledFeatures,
	}
	req := client.httpClient.R().SetBody(body).SetError(&Error{})
	res, err := req.Put(fmt.Sprintf("/api/v2/device/%d/maintenance", deviceID))
	if err != nil {
		return err
	}
	if res.IsError() {
		return res.Error().(*Error)
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
func (client *Client) DevicesWithSoftware(devices Devices, pattern *regexp.Regexp) ([]int, error) {
	final := make([]int, 0, len(devices))
	chunks := utils.ChunkSlice(devices, DefaultQueryBatchSize)

	results, err := iter.MapErr(chunks, func(filtered *[]Device) ([]int, error) {
		ids := make([]int, 0, len(*filtered))
		for _, d := range *filtered {
			d := d
			ids = append(ids, d.ID)
		}
		df := NewDeviceFilter().ID(IN, ids...)
		inventory, err := client.SoftwareInventory(df)
		if err != nil {
			return nil, err
		}
		r := make([]int, 0, len(*filtered))
		for _, pkg := range inventory {
			pkg := pkg
			did := int(pkg.DeviceID)
			if pattern.MatchString(pkg.Name) {
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
	tokenCache TokenCache,
) (*Client, error) {

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
	token, err := auth.GetAccessToken()
	if err != nil {
		return nil, err
	}

	httpClient := resty.New().
		SetBaseURL(baseURL).
		SetError(&Error{}).
		SetAuthToken(token)

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
