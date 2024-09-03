//
// Copyright (c) 2015-2022 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//

package madmin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// ProfilerType represents the profiler type
// passed to the profiler subsystem.
type ProfilerType string

// Different supported profiler types.
const (
	ProfilerCPU        ProfilerType = "cpu"        // represents CPU profiler type
	ProfilerCPUIO      ProfilerType = "cpuio"      // represents CPU with IO (fgprof) profiler type
	ProfilerMEM        ProfilerType = "mem"        // represents MEM profiler type
	ProfilerBlock      ProfilerType = "block"      // represents Block profiler type
	ProfilerMutex      ProfilerType = "mutex"      // represents Mutex profiler type
	ProfilerTrace      ProfilerType = "trace"      // represents Trace profiler type
	ProfilerThreads    ProfilerType = "threads"    // represents ThreadCreate profiler type
	ProfilerGoroutines ProfilerType = "goroutines" // represents Goroutine dumps.
)

// StartProfilingResult holds the result of starting
// profiler result in a given node.
type StartProfilingResult struct {
	NodeName string `json:"nodeName"`
	Success  bool   `json:"success"`
	Error    string `json:"error"`
}

// StartProfiling makes an admin call to remotely start profiling on a standalone
// server or the whole cluster in case of a distributed setup.
// Deprecated: use Profile API instead
func (adm *AdminClient) StartProfiling(ctx context.Context, profiler ProfilerType) ([]StartProfilingResult, error) {
	v := url.Values{}
	v.Set("profilerType", string(profiler))
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/profiling/start",
			queryValues: v,
		},
	)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	jsonResult, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var startResults []StartProfilingResult
	err = json.Unmarshal(jsonResult, &startResults)
	if err != nil {
		return nil, err
	}

	return startResults, nil
}

// DownloadProfilingData makes an admin call to download profiling data of a standalone
// server or of the whole cluster in case of a distributed setup.
// Deprecated: use Profile API instead
func (adm *AdminClient) DownloadProfilingData(ctx context.Context) (io.ReadCloser, error) {
	path := fmt.Sprintf(adminAPIPrefix + "/profiling/download")
	resp, err := adm.executeMethod(ctx,
		http.MethodGet, requestData{
			relPath: path,
		},
	)
	if err != nil {
		closeResponse(resp)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	if resp.Body == nil {
		return nil, errors.New("body is nil")
	}

	return resp.Body, nil
}

// Profile makes an admin call to remotely start profiling on a standalone
// server or the whole cluster in  case of a distributed setup for a specified duration.
func (adm *AdminClient) Profile(ctx context.Context, profiler ProfilerType, duration time.Duration) (io.ReadCloser, error) {
	v := url.Values{}
	v.Set("profilerType", string(profiler))
	v.Set("duration", duration.String())
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/profile",
			queryValues: v,
		},
	)
	if err != nil {
		closeResponse(resp)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	if resp.Body == nil {
		return nil, errors.New("body is nil")
	}
	return resp.Body, nil
}

/* trinet */
func (adm *AdminClient) StopProfile(ctx context.Context) error {
	v := url.Values{}
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/profile/stop",
			queryValues: v,
		},
	)
	if err != nil {
		closeResponse(resp)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}
	return nil
}

type StartProfileResponse struct {
	Result string `json:"result"`
}

func (adm *AdminClient) StartProfile(ctx context.Context, profiler ProfilerType) (StartProfileResponse, error) {
	v := url.Values{}
	v.Set("profilerType", string(profiler))
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/profile/start",
			queryValues: v,
		},
	)
	if err != nil {
		closeResponse(resp)
		return StartProfileResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return StartProfileResponse{}, httpRespToErrorResponse(resp)
	}

	var startResults StartProfileResponse
	startResults.Result = "StartProfile success"

	return startResults, nil
}

type ListProfileResponse struct {
	ProfileFiles []ProfileFile `json:"ProfileFile"`
}

type ProfileFile struct {
	FileName string `json:"fileName"`
}

func (adm *AdminClient) ListProfile(ctx context.Context) (ListProfileResponse, error) {
	v := url.Values{}
	resp, err := adm.executeMethod(ctx,
		http.MethodGet, requestData{
			relPath:     adminAPIPrefix + "/profile/list",
			queryValues: v,
		},
	)
	if err != nil {
		closeResponse(resp)
		return ListProfileResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ListProfileResponse{}, httpRespToErrorResponse(resp)
	}

	if resp.Body == nil {
		return ListProfileResponse{}, errors.New("body is nil")
	}

	jsonResult, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListProfileResponse{}, err
	}

	var ListResults ListProfileResponse
	err = json.Unmarshal(jsonResult, &ListResults)
	if err != nil {
		return ListProfileResponse{}, err
	}

	return ListResults, nil
}

func (adm *AdminClient) GetProfile(ctx context.Context, profileName string) (io.ReadCloser, error) {
	v := url.Values{}
	v.Set("profileName", profileName)
	resp, err := adm.executeMethod(ctx,
		http.MethodGet, requestData{
			relPath:     adminAPIPrefix + "/profile/get",
			queryValues: v,
		},
	)
	if err != nil {
		closeResponse(resp)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	if resp.Body == nil {
		return nil, errors.New("body is nil")
	}
	return resp.Body, nil
}

/* trinet */
