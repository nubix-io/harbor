// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nubiximage

import (
	"context"

	ps "github.com/goharbor/harbor/src/controller/artifact/processor"
	"github.com/goharbor/harbor/src/controller/artifact/processor/base"
	"github.com/goharbor/harbor/src/lib/log"
	"github.com/goharbor/harbor/src/pkg/artifact"
)

// const definitions
const (
	ArtifactTypeNubixImage = "NUBIX_IMAGE"
	mediaType              = "application/vnd.nubix.image.config.v1+json"
)

func init() {
	pc := &processor{
		manifestProcessor: base.NewManifestProcessor(),
	}
	pc.IndexProcessor = base.NewIndexProcessor()
	if err := ps.Register(pc, mediaType); err != nil {
		log.Errorf("failed to register processor for media type %s: %v", mediaType, err)
		return
	}
}

type processor struct {
	*base.IndexProcessor
	manifestProcessor *base.ManifestProcessor
}

func (p *processor) AbstractMetadata(ctx context.Context, artifact *artifact.Artifact, manifest []byte) error {
	if err := p.manifestProcessor.AbstractMetadata(ctx, artifact, manifest); err != nil {
		return err
	}
	return nil
}

func (p *processor) AbstractAddition(ctx context.Context, artifact *artifact.Artifact, addition string) (*ps.Addition, error) {
	return nil, nil
}

func (p *processor) GetArtifactType(ctx context.Context, artifact *artifact.Artifact) string {
	return ArtifactTypeNubixImage
}

func (p *processor) ListAdditionTypes(ctx context.Context, artifact *artifact.Artifact) []string {
	return nil
}
