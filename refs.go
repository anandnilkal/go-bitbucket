package bitbucket

import (
	"encoding/json"
)

type Ref struct {
	c *Client
}

func (r *Ref) Create(ro *RefsOptions) (*RepositoryBranch, error) {
	data := r.buildRefsBody(ro)
	urlStr := r.c.requestUrl("/repositories/%s/%s/refs/branches", ro.Owner, ro.RepoSlug)
	response, err := r.c.execute("POST", urlStr, data)
	if err != nil {
		return nil, err
	}

	var rb RepositoryBranch
	err = json.Unmarshal(response.([]byte), &rb) 
	if err != nil {
		return nil, err
	}

	return &rb, nil
} 

func (r *Ref) Delete(ro *RefsOptions) (interface{}, error) {
	urlStr := r.c.requestUrl("/repositories/%s/%s/refs/branches/%s", ro.Owner, ro.RepoSlug, ro.BranchName)
	return r.c.execute("DELETE", urlStr, "")
} 

func (r *Ref) Get(ro *RefsOptions) (*RepositoryBranch, error) {
	urlStr := r.c.requestUrl("/repositories/%s/%s/refs/branches/%s", ro.Owner, ro.RepoSlug, ro.BranchName)
	response, err := r.c.execute("GET", urlStr, "")
	if err != nil {
		return nil, err
	}	
	var rb RepositoryBranch
	err = json.Unmarshal(response.([]byte), &rb) 
	if err != nil {
		return nil, err
	}

	return &rb, nil
} 
