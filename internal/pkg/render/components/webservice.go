package components

import (
	"github.com/kiaedev/kiae/api/app"
	"github.com/kiaedev/kiae/api/project"
	"github.com/kiaedev/kiae/internal/pkg/render/traits"
	"github.com/kiaedev/kiae/internal/pkg/render/utils"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
	v1 "k8s.io/api/core/v1"
)

type KWebservice struct {
	Name             string                  `json:"name"`
	Labels           map[string]string       `json:"labels,omitempty"`
	Annotations      map[string]string       `json:"annotations,omitempty"`
	Envs             map[string]string       `json:"envs,omitempty"`
	Image            string                  `json:"image"`
	ImagePullPolicy  string                  `json:"imagePullPolicy,omitempty"`
	ImagePullSecrets []string                `json:"imagePullSecrets,omitempty"`
	Ports            []*project.Port         `json:"ports"`
	Replicas         uint32                  `json:"replicas"`
	Resources        v1.ResourceRequirements `json:"resources"`
	LivenessProbe    *project.HealthProbe    `json:"livenessProbe,omitempty"`
	ReadinessProbe   *project.HealthProbe    `json:"readinessProbe,omitempty"`

	Traits []common.ApplicationTrait
}

func NewKWebservice(kApp *app.Application, proj *project.Project) *KWebservice {
	traits := make([]common.ApplicationTrait, 0)
	if len(kApp.Configs) > 0 {
		traits = append(traits, common.ApplicationTrait{
			Type: "k-config", Properties: util.Object2RawExtension(map[string]interface{}{"configs": kApp.Configs}),
		})
	}

	return &KWebservice{
		Name: kApp.Name,
		// Labels:      map[string]string{"kiae.dev/test": "test"},
		Annotations: kApp.Annotations,
		Image:       kApp.Image,
		// ImagePullSecrets: "",
		Replicas:  kApp.Replicas,
		Ports:     kApp.Ports,
		Resources: utils.BuildResources(kApp.Size, 0.5),
		Envs:      map[string]string{},
		Traits:    traits,
	}
}

func (c *KWebservice) GetName() string {
	return c.Name
}

func (c *KWebservice) GetType() string {
	return "k-webservice"
}

func (c *KWebservice) SetupTrait(trait traits.Trait) {
	c.Traits = append(c.Traits, common.ApplicationTrait{Type: trait.GetType(), Properties: util.Object2RawExtension(trait)})
}
