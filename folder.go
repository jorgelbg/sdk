package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

import "time"

// PermissionLevel is a type for representing the permission levels for the permission field
type PermissionLevel int

// Permissions levels as described in the doc:
// https://grafana.com/docs/grafana/latest/http_api/folder_permissions/#folder-permissions-api
const (
	View PermissionLevel = iota + 1
	Edit
	_
	Admin
)

// Folder as described in the doc
// https://grafana.com/docs/grafana/latest/http_api/folder/#get-all-folders
type Folder struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	HasAcl    bool   `json:"hasAcl"`
	CanSave   bool   `json:"canSave"`
	CanEdit   bool   `json:"canEdit"`
	CanAdmin  bool   `json:"canAdmin"`
	CreatedBy string `json:"createdBy"`
	Created   string `json:"created"`
	UpdatedBy string `json:"updatedBy"`
	Updated   string `json:"updated"`
	Version   int    `json:"version"`
	Overwrite bool   `json:"overwrite"`
}

// FolderPermission as described in the doc
// https://grafana.com/docs/grafana/latest/http_api/folder_permissions/
type FolderPermission struct {
	ID             int             `json:"ID"`
	FolderID       int             `json:"folderId"`
	Created        time.Time       `json:"created"`
	Updated        time.Time       `json:"updated"`
	UserID         int             `json:"userId"`
	UserLogin      string          `json:"userLogin"`
	UserEmail      string          `json:"userEmail"`
	UserAvatarURL  string          `json:"userAvatarUrl"`
	TeamID         int             `json:"teamId"`
	TeamEmail      string          `json:"teamEmail"`
	TeamAvatarURL  string          `json:"teamAvatarUrl"`
	Team           string          `json:"team"`
	Role           string          `json:"role"`
	Permission     PermissionLevel `json:"permission"`
	PermissionName string          `json:"permissionName"`
	UID            string          `json:"uid"`
	Title          string          `json:"title"`
	Slug           string          `json:"slug"`
	IsFolder       bool            `json:"isFolder"`
	URL            string          `json:"url"`
	Inherited      bool            `json:"inherited"`
}

// FolderPermissionItem as described in the doc
// https://grafana.com/docs/grafana/latest/http_api/folder_permissions/
//
// Each item has a pair of [Role|TeamID|UserID] and Permission.
// unnecessary fields are omitted.
type FolderPermissionItem struct {
	Role       string          `json:"role,omitempty"`
	TeamID     int             `json:"teamId,omitempty"`
	UserID     int             `json:"userId,omitempty"`
	Permission PermissionLevel `json:"permission,omitempty"`
}

// DefaultRolesFolderPermissions represents the default role permissions for a given folder
// Viewer has View permissions
// Editor has Edit permissions
var DefaultRolesFolderPermissions = []FolderPermissionItem{
	{Role: UserViewer, Permission: View},
	{Role: UserEditor, Permission: Edit},
}

type FolderPermissionList struct {
	Items []FolderPermissionItem `json:"items"`
}
