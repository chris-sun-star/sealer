// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filesystem

import (
	"fmt"

	"github.com/sealerio/sealer/pkg/imageengine"

	"github.com/sealerio/sealer/pkg/filesystem/cloudfilesystem"
	"github.com/sealerio/sealer/pkg/filesystem/clusterimage"
	"github.com/sealerio/sealer/pkg/runtime"
)

// NewClusterImageMounter mount and unmount ClusterImage.
func NewClusterImageMounter(engine imageengine.Interface) (clusterimage.Interface, error) {
	return clusterimage.NewClusterImageMounter(engine)
}

// NewFilesystem :according to the Metadata file content to determine what kind of Filesystem will be load.
func NewFilesystem(rootfs string) (cloudfilesystem.Interface, error) {
	md, err := runtime.LoadMetadata(rootfs)
	if err != nil {
		return nil, fmt.Errorf("failed to load Metadata file to determine the filesystem type: %v", err)
	}

	if md == nil || !md.NydusFlag {
		return cloudfilesystem.NewOverlayFileSystem()
	}

	return cloudfilesystem.NewNydusFileSystem()
}
