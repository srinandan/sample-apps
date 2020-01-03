// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/srinandan/sample-apps/tracking/pkg/api/v1"
)

type Tracking struct {
	TrackingId      string `json:"trackingId,omitempty"`
	Status          string `json:"status,omitempty"`
	Created         string `json:"created,omitempty"`
	Updated         string `json:"updated,omitempty"`
	Signed          string `json:"signed,omitempty"`
	Weight          string `json:"weight,omitempty"`
	EstDeliveryDate string `json:"est_delivery_date,omitempty"`
	Carrier         string `json:"carrier,omitempty"`
}

var trackingItems = []Tracking{}

func ReadTrackingFile() error {
	trackingBytes, err := ioutil.ReadFile("tracking.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(trackingBytes, &trackingItems); err != nil {
		return err
	}

	return nil
}

// server is used to implement TrackingServer
type trackingServer struct {
	v1.UnimplementedTrackingServer
}

func NewTrackingService() (v1.TrackingServer, error) {
	err := ReadTrackingFile()
	if err != nil {
		return &trackingServer{}, err
	}

	return &trackingServer{}, err
}

func (s *trackingServer) GetTrackingDetails(ctx context.Context, req *v1.TrackingRequest) (*v1.TrackingResponse, error) {
	for _, trackingItem := range trackingItems {
		if trackingItem.TrackingId == req.TrackingId {
			return &v1.TrackingResponse{
				TrackingId:      trackingItem.TrackingId,
				Status:          trackingItem.Status,
				Created:         trackingItem.Created,
				Updated:         trackingItem.Updated,
				Weight:          trackingItem.Weight,
				EstDeliveryDate: trackingItem.EstDeliveryDate,
				Carrier:         trackingItem.Carrier,
			}, nil
		}
	}
	return &v1.TrackingResponse{}, fmt.Errorf("tracking item not found")
}

func (s *trackingServer) ListTrackingDetails(ctx context.Context, empty *empty.Empty) (*v1.TrackingListResponse, error) {
	trackingListResponse := v1.TrackingListResponse{}
	
	if len(trackingItems) == 0 {
		return &trackingListResponse, fmt.Errorf("tracking items not found")
	}

	for _, trackingItem := range trackingItems {
		trackingResponse := v1.TrackingResponse{}
		trackingResponse.TrackingId = trackingItem.TrackingId
		trackingResponse.Status = trackingItem.Status
		trackingResponse.Created = trackingItem.Created
		trackingResponse.Updated = trackingItem.Updated
		trackingResponse.Weight = trackingItem.Weight
		trackingResponse.EstDeliveryDate = trackingItem.EstDeliveryDate
		trackingResponse.Carrier = trackingItem.Carrier
		trackingListResponse.TrackingResponse = append(trackingListResponse.TrackingResponse, &trackingResponse)
	}
	return &trackingListResponse, nil
}
