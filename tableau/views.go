package tableau

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// GetViewRequest encapsulates the request for getting a single View.
type GetViewRequest struct {
	ID      string
	Filters []ViewFilter
}

type ViewFilter struct {
	Name  string
	Value string
}

// GetViewsRequest encapsulates the request for getting a single View.
type GetViewsRequest struct {
	PageSize   int
	PageNumber int
}

type viewsService struct {
	client *Client
}

func (dss *viewsService) GetViewData(ctx context.Context, getViewReq *GetViewRequest) (interface{}, error) {
	requestFilter := ""
	if len(getViewReq.Filters) > 0 {
		var temp []string
		for _, filter := range getViewReq.Filters {
			temp = append(temp, filter.Name+"="+filter.Value)
		}
		requestFilter = strings.Join(temp, "&")
	}

	path := fmt.Sprintf("sites/%s/views/%s/data?%s", dss.client.SiteID, getViewReq.ID, requestFilter)
	req, err := dss.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get view")
	}

	body, err := dss.client.doRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (dss *viewsService) GetViews(ctx context.Context, getViewsReq *GetViewsRequest) (interface{}, error) {
	path := fmt.Sprintf("sites/%s/views/?pageSize=%d&pageNumber=%d", dss.client.SiteID, getViewsReq.PageSize, getViewsReq.PageNumber)
	req, err := dss.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get view")
	}

	var views interface{}
	err = dss.client.do(ctx, req, &views)
	if err != nil {
		return nil, err
	}

	return views, nil
}
