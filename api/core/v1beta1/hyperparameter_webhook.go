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
	"strings"

	"github.com/DataTunerX/utility-server/logging"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	finetunev1beta1 "github.com/DataTunerX/meta-server/api/finetune/v1beta1"
)

func (r *Hyperparameter) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-core-datatunerx-io-v1beta1-hyperparameter,mutating=true,failurePolicy=fail,sideEffects=None,groups=core.datatunerx.io,resources=hyperparameters,verbs=create;update,versions=v1beta1,name=mhyperparameter.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Hyperparameter{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Hyperparameter) Default() {
	// todo Standardised logs
	logging.ZLogger.Infof("Validate default hyperparameter %s/%s", r.Namespace, r.Name)

	//Set the default value to the hyperparameter cr

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-core-datatunerx-io-v1beta1-hyperparameter,mutating=false,failurePolicy=fail,sideEffects=None,groups=core.datatunerx.io,resources=hyperparameters,verbs=create;update,versions=v1beta1,name=vhyperparameter.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Hyperparameter{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Hyperparameter) ValidateCreate() (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate create hyperparameter %s/%s", r.Namespace, r.Name)

	if err := validateScheduler(r.Spec.Parameters.Scheduler); err != nil {
		return nil, err
	}

	if err := validateOptimizer(r.Spec.Parameters.Optimizer); err != nil {
		return nil, err
	}

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Hyperparameter) ValidateUpdate(old runtime.Object) (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate update hyperparameter %s/%s", r.Namespace, r.Name)
	hyperparameter := old.(*Hyperparameter)
	if hyperparameter.Status.ReferenceFinetuneName != nil {
		return nil, fmt.Errorf("hyperparameter %s/%s is referenced by finetune, not allow update",
			hyperparameter.Namespace, hyperparameter.Name)
	}

	if err := validateScheduler(r.Spec.Parameters.Scheduler); err != nil {
		return nil, err
	}

	if err := validateOptimizer(r.Spec.Parameters.Optimizer); err != nil {
		return nil, err
	}

	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Hyperparameter) ValidateDelete() (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate delete hyperparameter %s/%s", r.Namespace, r.Name)
	if r.Status.ReferenceFinetuneName != nil {
		return nil, fmt.Errorf("hyperparameter %s/%s is referenced by finetune, not allow delete",
			r.Namespace, r.Name)
	}
	return nil, nil
}

func validateScheduler(scheduler finetunev1beta1.HyperparameterScheduler) error {
	validSchedulers := map[finetunev1beta1.HyperparameterScheduler]struct{}{
		"cosine":   {},
		"linear":   {},
		"constant": {},
	}

	_, ok := validSchedulers[scheduler]
	if !ok {
		validKeys := make([]string, 0, len(validSchedulers))
		for k := range validSchedulers {
			validKeys = append(validKeys, string(k))
		}
		return fmt.Errorf("invalid scheduler: %s, valid schedulers are: %s", scheduler, strings.Join(validKeys, ", "))
	}
	return nil
}

func validateOptimizer(optimizer finetunev1beta1.HyperparameterOptimizer) error {
	validOptimizers := map[finetunev1beta1.HyperparameterOptimizer]struct{}{
		"adamw_torch": {},
		"adamw_hf":    {},
		"sgd":         {},
		"adafactor":   {},
		"adagrad":     {},
	}

	_, ok := validOptimizers[optimizer]
	if !ok {
		validKeys := make([]string, 0, len(validOptimizers))
		for k := range validOptimizers {
			validKeys = append(validKeys, string(k))
		}
		return fmt.Errorf("invalid optimizer: %s, valid optimizers are: %s", optimizer, strings.Join(validKeys, ", "))
	}
	return nil
}
