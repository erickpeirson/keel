package k8s

import (
	"fmt"
	"reflect"
	"strings"

	apps_v1 "k8s.io/api/apps/v1"
	batch_v1 "k8s.io/api/batch/v1"
	core_v1 "k8s.io/api/core/v1"
)

// GenericResource - generic resource,
// used to work with multiple kinds of k8s resources
type GenericResource struct {
	// original resource
	obj interface{}

	Identifier string
	Namespace  string
	Name       string
}

type genericResource []*GenericResource

func (c genericResource) Len() int {
	return len(c)
}

func (c genericResource) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c genericResource) Less(i, j int) bool {
	return c[i].Identifier < c[j].Identifier
}

// NewGenericResource - create new generic k8s resource
func NewGenericResource(obj interface{}) (*GenericResource, error) {

	switch obj.(type) {
	case *apps_v1.Deployment, *apps_v1.StatefulSet, *apps_v1.DaemonSet:
		// ok
	case *batch_v1.CronJob:
		// ok
	default:
		return nil, fmt.Errorf("unsupported resource type: %v", reflect.TypeOf(obj).Kind())
	}

	gr := &GenericResource{
		obj: obj,
	}

	gr.Identifier = gr.GetIdentifier()
	gr.Namespace = gr.GetNamespace()
	gr.Name = gr.GetName()

	return gr, nil
}

func (r *GenericResource) String() string {
	return fmt.Sprintf("%s/%s/%s images: %s", r.Kind(), r.Namespace, r.Name, strings.Join(r.GetImages(), ", "))
}

// DeepCopy uses an autogenerated deepcopy functions, copying the receiver, creating a new GenericResource
func (r *GenericResource) DeepCopy() *GenericResource {
	gr := new(GenericResource)
	if r.obj == nil {
		return gr
	}
	gr.Identifier = r.Identifier
	gr.Namespace = r.Namespace
	gr.Name = r.Name

	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		gr.obj = obj.DeepCopy()
	case *apps_v1.StatefulSet:
		gr.obj = obj.DeepCopy()
	case *apps_v1.DaemonSet:
		gr.obj = obj.DeepCopy()
	case *batch_v1.CronJob:
		gr.obj = obj.DeepCopy()
	}

	return gr
}

// GetIdentifier returns resource identifier
func (r *GenericResource) GetIdentifier() string {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getDeploymentIdentifier(obj)
	case *apps_v1.StatefulSet:
		return getStatefulSetIdentifier(obj)
	case *apps_v1.DaemonSet:
		return getDaemonsetSetIdentifier(obj)
	case *batch_v1.CronJob:
		return getCronJobIdentifier(obj)
	}
	return ""
}

// GetName returns resource name
func (r *GenericResource) GetName() string {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.GetName()
	case *apps_v1.StatefulSet:
		return obj.GetName()
	case *apps_v1.DaemonSet:
		return obj.GetName()
	case *batch_v1.CronJob:
		return obj.GetName()
	}
	return ""
}

// GetNamespace returns resource namespace
func (r *GenericResource) GetNamespace() string {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.GetNamespace()
	case *apps_v1.StatefulSet:
		return obj.GetNamespace()
	case *apps_v1.DaemonSet:
		return obj.GetNamespace()
	case *batch_v1.CronJob:
		return obj.GetNamespace()
	}
	return ""
}

// Kind returns a type of resource that this structure represents
func (r *GenericResource) Kind() string {
	switch r.obj.(type) {
	case *apps_v1.Deployment:
		return "deployment"
	case *apps_v1.StatefulSet:
		return "statefulset"
	case *apps_v1.DaemonSet:
		return "daemonset"
	case *batch_v1.CronJob:
		return "cronjob"
	}
	return ""
}

// GetResource - get resource
func (r *GenericResource) GetResource() interface{} {
	return r.obj
}

// GetLabels - get resource labels
func (r *GenericResource) GetLabels() (labels map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getOrInitialise(obj.GetLabels())
	case *apps_v1.StatefulSet:
		return getOrInitialise(obj.GetLabels())
	case *apps_v1.DaemonSet:
		return getOrInitialise(obj.GetLabels())
	case *batch_v1.CronJob:
		return getOrInitialise(obj.GetLabels())
	}
	return
}

// SetLabels - set resource labels
func (r *GenericResource) SetLabels(labels map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		obj.SetLabels(labels)
	case *apps_v1.StatefulSet:
		obj.SetLabels(labels)
	case *apps_v1.DaemonSet:
		obj.SetLabels(labels)
	case *batch_v1.CronJob:
		obj.SetLabels(labels)
	}
}

// GetSpecAnnotations - get resource spec template annotations
func (r *GenericResource) GetSpecAnnotations() (annotations map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getOrInitialise(obj.Spec.Template.GetAnnotations())
	case *apps_v1.StatefulSet:
		return getOrInitialise(obj.Spec.Template.GetAnnotations())
	case *apps_v1.DaemonSet:
		return getOrInitialise(obj.Spec.Template.GetAnnotations())
	case *batch_v1.CronJob:
		return getOrInitialise(obj.Spec.JobTemplate.GetAnnotations())
	}
	return
}

// SetSpecAnnotations - set resource spec template annotations
func (r *GenericResource) SetSpecAnnotations(annotations map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		obj.Spec.Template.SetAnnotations(annotations)
	case *apps_v1.StatefulSet:
		obj.Spec.Template.SetAnnotations(annotations)
	case *apps_v1.DaemonSet:
		obj.Spec.Template.SetAnnotations(annotations)
	case *batch_v1.CronJob:
		obj.Spec.JobTemplate.SetAnnotations(annotations)
	}
}

func getOrInitialise(a map[string]string) map[string]string {
	if a == nil {
		return make(map[string]string)
	}
	return a
}

// GetAnnotations - get resource annotations
func (r *GenericResource) GetAnnotations() (annotations map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getOrInitialise(obj.GetAnnotations())
	case *apps_v1.StatefulSet:
		return getOrInitialise(obj.GetAnnotations())
	case *apps_v1.DaemonSet:
		return getOrInitialise(obj.GetAnnotations())
	case *batch_v1.CronJob:
		return getOrInitialise(obj.GetAnnotations())
	}
	return
}

// SetAnnotations - set resource annotations
func (r *GenericResource) SetAnnotations(annotations map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		obj.SetAnnotations(annotations)
	case *apps_v1.StatefulSet:
		obj.SetAnnotations(annotations)
	case *apps_v1.DaemonSet:
		obj.SetAnnotations(annotations)
	case *batch_v1.CronJob:
		obj.SetAnnotations(annotations)
	}
}

// GetImagePullSecrets - returns secrets from pod spec
func (r *GenericResource) GetImagePullSecrets() (secrets []string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getImagePullSecrets(obj.Spec.Template.Spec.ImagePullSecrets)
	case *apps_v1.StatefulSet:
		return getImagePullSecrets(obj.Spec.Template.Spec.ImagePullSecrets)
	case *apps_v1.DaemonSet:
		return getImagePullSecrets(obj.Spec.Template.Spec.ImagePullSecrets)
	case *batch_v1.CronJob:
		return getImagePullSecrets(obj.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets)
	}
	return
}

// GetImages - returns images used by this resource
func (r *GenericResource) GetImages() (images []string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getContainerImages(obj.Spec.Template.Spec.Containers)
	case *apps_v1.StatefulSet:
		return getContainerImages(obj.Spec.Template.Spec.Containers)
	case *apps_v1.DaemonSet:
		return getContainerImages(obj.Spec.Template.Spec.Containers)
	case *batch_v1.CronJob:
		return getContainerImages(obj.Spec.JobTemplate.Spec.Template.Spec.Containers)
	}
	return
}

// GetInitImages - returns init images used by this resource
func (r *GenericResource) GetInitImages() (images []string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getContainerImages(obj.Spec.Template.Spec.InitContainers)
	case *apps_v1.StatefulSet:
		return getContainerImages(obj.Spec.Template.Spec.InitContainers)
	case *apps_v1.DaemonSet:
		return getContainerImages(obj.Spec.Template.Spec.InitContainers)
	case *batch_v1.CronJob:
		return getContainerImages(obj.Spec.JobTemplate.Spec.Template.Spec.InitContainers)
	}
	return
}

// Containers - returns containers managed by this resource
func (r *GenericResource) Containers() (containers []core_v1.Container) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.Spec.Template.Spec.Containers
	case *apps_v1.StatefulSet:
		return obj.Spec.Template.Spec.Containers
	case *apps_v1.DaemonSet:
		return obj.Spec.Template.Spec.Containers
	case *batch_v1.CronJob:
		return obj.Spec.JobTemplate.Spec.Template.Spec.Containers
	}
	return
}

// InitContainers - returns init containers managed by this resource
func (r *GenericResource) InitContainers() (containers []core_v1.Container) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.Spec.Template.Spec.InitContainers
	case *apps_v1.StatefulSet:
		return obj.Spec.Template.Spec.InitContainers
	case *apps_v1.DaemonSet:
		return obj.Spec.Template.Spec.InitContainers
	case *batch_v1.CronJob:
		return obj.Spec.JobTemplate.Spec.Template.Spec.InitContainers
	}
	return
}

// UpdateContainer - updates container image
func (r *GenericResource) UpdateContainer(index int, image string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		updateDeploymentContainer(obj, index, image)
	case *apps_v1.StatefulSet:
		updateStatefulSetContainer(obj, index, image)
	case *apps_v1.DaemonSet:
		updateDaemonsetSetContainer(obj, index, image)
	case *batch_v1.CronJob:
		updateCronJobContainer(obj, index, image)
	}
}

// UpdateInitContainer - updates init container image
func (r *GenericResource) UpdateInitContainer(index int, image string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		updateDeploymentInitContainer(obj, index, image)
	case *apps_v1.StatefulSet:
		updateStatefulSetInitContainer(obj, index, image)
	case *apps_v1.DaemonSet:
		updateDaemonsetSetInitContainer(obj, index, image)
	case *batch_v1.CronJob:
		updateCronJobInitContainer(obj, index, image)
	}
}

type Status struct {
	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas"`

	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas"`

	// Total number of ready pods targeted by this deployment.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas"`

	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas"`

	// Total number of unavailable pods targeted by this deployment. This is the total number of
	// pods that are still required for the deployment to have 100% available capacity. They may
	// either be pods that are running but not yet available or pods that still have not been created.
	// +optional
	UnavailableReplicas int32 `json:"unavailableReplica"`
}

func (r *GenericResource) GetStatus() Status {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return Status{
			Replicas:            obj.Status.Replicas,
			UpdatedReplicas:     obj.Status.UpdatedReplicas,
			ReadyReplicas:       obj.Status.ReadyReplicas,
			AvailableReplicas:   obj.Status.AvailableReplicas,
			UnavailableReplicas: obj.Status.UnavailableReplicas,
		}
	case *apps_v1.StatefulSet:
		return Status{
			Replicas:            obj.Status.Replicas,
			UpdatedReplicas:     obj.Status.UpdatedReplicas,
			ReadyReplicas:       obj.Status.ReadyReplicas,
			AvailableReplicas:   obj.Status.CurrentReplicas,
			UnavailableReplicas: 0, // N/A
		}
	case *apps_v1.DaemonSet:
		return Status{
			Replicas:            obj.Status.DesiredNumberScheduled,
			UpdatedReplicas:     obj.Status.UpdatedNumberScheduled,
			ReadyReplicas:       obj.Status.NumberReady,
			AvailableReplicas:   obj.Status.NumberAvailable,
			UnavailableReplicas: obj.Status.NumberUnavailable,
		}
	case *batch_v1.CronJob:
		return Status{
			Replicas:            int32(len(obj.Status.Active)),
			UpdatedReplicas:     0,
			ReadyReplicas:       0,
			AvailableReplicas:   0,
			UnavailableReplicas: 0,
		}
	}
	return Status{}
}
