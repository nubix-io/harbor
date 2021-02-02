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

package nubixapp

import (
	"context"
	"encoding/json"

	ps "github.com/goharbor/harbor/src/controller/artifact/processor"
	"github.com/goharbor/harbor/src/controller/artifact/processor/base"
	"github.com/goharbor/harbor/src/lib/log"
	"github.com/goharbor/harbor/src/pkg/artifact"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// const definitions
const (
	ArtifactTypeNubixApp = "NUBIX_APP"
	mediaType            = "application/vnd.nubix.app.config.v1+json"
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
	log.Debugf("nubix_app.AbstractAddtion - addition string: %s\n", addition)

	m, _, err := p.RegCli.PullManifest(artifact.RepositoryName, artifact.Digest)
	if err != nil {
		return nil, err
	}
	_, payload, err := m.Payload()
	if err != nil {
		return nil, err
	}
	manifest := &v1.Manifest{}
	if err := json.Unmarshal(payload, manifest); err != nil {
		return nil, err
	}
	log.Debugf("nubix_app.AbstractAddtion - manifest %v", manifest)

	return nil, nil
}

func (p *processor) GetArtifactType(ctx context.Context, artifact *artifact.Artifact) string {
	return ArtifactTypeNubixApp
}

func (p *processor) ListAdditionTypes(ctx context.Context, artifact *artifact.Artifact) []string {
	return nil
}
