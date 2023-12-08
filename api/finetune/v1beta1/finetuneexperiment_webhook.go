/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"fmt"

	"github.com/DataTunerX/utility-server/logging"
	"github.com/DataTunerX/utility-server/random"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func (r *FinetuneExperiment) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-finetune-datatunerx-io-v1beta1-finetuneexperiment,mutating=true,failurePolicy=fail,sideEffects=None,groups=finetune.datatunerx.io,resources=finetuneexperiments,verbs=create;update,versions=v1beta1,name=mfinetuneexperiment.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &FinetuneExperiment{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *FinetuneExperiment) Default() {
	logging.ZLogger.Infof("Validate default finetuneExperiment %s/%s", r.Namespace, r.Name)
	if r.Spec.ServeConfig == nil {
		r.Spec.ServeConfig = &ServeConfig{
			NodeSelector: map[string]string{
				"nvidia.com/gpu": "present",
			},
		}
	}
	for i := range r.Spec.FinetuneJobs {
		if r.Spec.FinetuneJobs[i].Name == "" {
			r.Spec.FinetuneJobs[i].Name = fmt.Sprintf("%s-%s", r.Name, random.GenerateRandomString(4))
			logging.ZLogger.Debugf("Set name for finetuneJob %s/%s", r.Namespace, r.Spec.FinetuneJobs[i].Name)
			if r.Spec.FinetuneJobs[i].Spec.FineTune.Name == "" {
				logging.ZLogger.Debugf("Set name for finetune %s/%s", r.Namespace, r.Spec.FinetuneJobs[i].Name)
				r.Spec.FinetuneJobs[i].Spec.FineTune.Name = r.Spec.FinetuneJobs[i].Name
			}
		}
		if r.Spec.ScoringConfig != nil {
			logging.ZLogger.Debugf("Set scoringConfig for finetuneJob %s/%s", r.Namespace, r.Spec.FinetuneJobs[i].Name)
			r.Spec.FinetuneJobs[i].Spec.ScoringConfig = r.Spec.ScoringConfig
		}
		if r.Spec.ServeConfig != nil {
			logging.ZLogger.Debugf("Set serveConfig for finetuneJob %s/%s", r.Namespace, r.Spec.FinetuneJobs[i].Name)
			r.Spec.FinetuneJobs[i].Spec.ServeConfig = r.Spec.ServeConfig
		}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-finetune-datatunerx-io-v1beta1-finetuneexperiment,mutating=false,failurePolicy=fail,sideEffects=None,groups=finetune.datatunerx.io,resources=finetuneexperiments,verbs=create;update,versions=v1beta1,name=vfinetuneexperiment.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &FinetuneExperiment{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *FinetuneExperiment) ValidateCreate() (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate create finetuneExperiment %s/%s", r.Namespace, r.Name)
	if r.Spec.Pending {
		return nil, fmt.Errorf("create finetuneExperiment %s/%s failed,Pending cannot be set to true on creation", r.Namespace, r.Name)
	}
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *FinetuneExperiment) ValidateUpdate(old runtime.Object) (warnings admission.Warnings, err error) {
	oldFinetuneExperiment := old.(*FinetuneExperiment)
	logging.ZLogger.Infof("Validate update finetuneExperiment %s/%s", r.Namespace, r.Name)
	if oldFinetuneExperiment.Spec.Pending != r.Spec.Pending {
		return nil, nil
	}
	return nil, fmt.Errorf("finetuneExperiment updates are not allowed")
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *FinetuneExperiment) ValidateDelete() (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate delete finetuneExperiment %s/%s", r.Namespace, r.Name)
	return nil, nil
}
