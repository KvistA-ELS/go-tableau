package tableau

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// GetViewRequest encapsulates the request for getting a single View.
type GetViewRequest struct {
	ID string
}

type viewsResponse struct {
	View *View `json:"view"`
}

// View represents a Tableau view
type View struct {
	ID                  string            `json:"id"`
	Name                string            `json:"name"`
	CertificationNote   string            `json:"CertificationNote"`
	ContentUrl          string            `json:"contentUrl"`
	EncryptExtracts     string            `json:"encryptExtracts"`
	Description         string            `json:"description"`
	WebpageUrl          string            `json:"webpageUrl"`
	IsCertified         bool              `json:"isCertified"`
	UseRemoteQueryAgent bool              `json:"useRemoteQueryAgent"`
	Type                string            `json:"type"`
	Tags                map[string]string `json:"tags"`
	Owner               struct {
		ID string `json:"id"`
	}
	Project struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type viewsService struct {
	client *Client
}

func (dss *viewsService) Get(ctx context.Context, getReq *GetViewRequest) (*View, error) {
	path := fmt.Sprintf("sites/%s/views/%s", dss.client.SiteID, getReq.ID)
	req, err := dss.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for get view")
	}

	ds := &viewsResponse{}
	err = dss.client.do(ctx, req, &ds)
	if err != nil {
		return nil, err
	}

	return ds.View, nil
}
