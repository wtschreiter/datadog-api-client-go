// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// IncidentAttachmentLinkAttributesAttachmentObject The link attachment.
type IncidentAttachmentLinkAttributesAttachmentObject struct {
	// The URL of this link attachment.
	DocumentUrl string `json:"documentUrl"`
	// The title of this link attachment.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewIncidentAttachmentLinkAttributesAttachmentObject instantiates a new IncidentAttachmentLinkAttributesAttachmentObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentAttachmentLinkAttributesAttachmentObject(documentUrl string, title string) *IncidentAttachmentLinkAttributesAttachmentObject {
	this := IncidentAttachmentLinkAttributesAttachmentObject{}
	this.DocumentUrl = documentUrl
	this.Title = title
	return &this
}

// NewIncidentAttachmentLinkAttributesAttachmentObjectWithDefaults instantiates a new IncidentAttachmentLinkAttributesAttachmentObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentAttachmentLinkAttributesAttachmentObjectWithDefaults() *IncidentAttachmentLinkAttributesAttachmentObject {
	this := IncidentAttachmentLinkAttributesAttachmentObject{}
	return &this
}

// GetDocumentUrl returns the DocumentUrl field value.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) GetDocumentUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) GetDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentUrl, true
}

// SetDocumentUrl sets field value.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) SetDocumentUrl(v string) {
	o.DocumentUrl = v
}

// GetTitle returns the Title field value.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentAttachmentLinkAttributesAttachmentObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["documentUrl"] = o.DocumentUrl
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentAttachmentLinkAttributesAttachmentObject) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		DocumentUrl *string `json:"documentUrl"`
		Title       *string `json:"title"`
	}{}
	all := struct {
		DocumentUrl string `json:"documentUrl"`
		Title       string `json:"title"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.DocumentUrl == nil {
		return fmt.Errorf("required field documentUrl missing")
	}
	if required.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.DocumentUrl = all.DocumentUrl
	o.Title = all.Title
	return nil
}
