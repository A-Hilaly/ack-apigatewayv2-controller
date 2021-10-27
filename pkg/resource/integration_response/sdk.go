// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package integration_response

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.IntegrationResponse{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetIntegrationResponseOutput
	resp, err = rm.sdkapi.GetIntegrationResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetIntegrationResponse", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ContentHandlingStrategy != nil {
		ko.Spec.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		ko.Spec.ContentHandlingStrategy = nil
	}
	if resp.IntegrationResponseId != nil {
		ko.Status.IntegrationResponseID = resp.IntegrationResponseId
	} else {
		ko.Status.IntegrationResponseID = nil
	}
	if resp.IntegrationResponseKey != nil {
		ko.Spec.IntegrationResponseKey = resp.IntegrationResponseKey
	} else {
		ko.Spec.IntegrationResponseKey = nil
	}
	if resp.ResponseParameters != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.ResponseParameters {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		ko.Spec.ResponseParameters = f3
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.ResponseTemplates {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		ko.Spec.ResponseTemplates = f4
	} else {
		ko.Spec.ResponseTemplates = nil
	}
	if resp.TemplateSelectionExpression != nil {
		ko.Spec.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		ko.Spec.TemplateSelectionExpression = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.APIID == nil || r.ko.Status.IntegrationResponseID == nil || r.ko.Spec.IntegrationID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetIntegrationResponseInput, error) {
	res := &svcsdk.GetIntegrationResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Spec.IntegrationID)
	}
	if r.ko.Status.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*r.ko.Status.IntegrationResponseID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateIntegrationResponseOutput
	_ = resp
	resp, err = rm.sdkapi.CreateIntegrationResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateIntegrationResponse", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ContentHandlingStrategy != nil {
		ko.Spec.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		ko.Spec.ContentHandlingStrategy = nil
	}
	if resp.IntegrationResponseId != nil {
		ko.Status.IntegrationResponseID = resp.IntegrationResponseId
	} else {
		ko.Status.IntegrationResponseID = nil
	}
	if resp.IntegrationResponseKey != nil {
		ko.Spec.IntegrationResponseKey = resp.IntegrationResponseKey
	} else {
		ko.Spec.IntegrationResponseKey = nil
	}
	if resp.ResponseParameters != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.ResponseParameters {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		ko.Spec.ResponseParameters = f3
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.ResponseTemplates {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		ko.Spec.ResponseTemplates = f4
	} else {
		ko.Spec.ResponseTemplates = nil
	}
	if resp.TemplateSelectionExpression != nil {
		ko.Spec.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		ko.Spec.TemplateSelectionExpression = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateIntegrationResponseInput, error) {
	res := &svcsdk.CreateIntegrationResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*r.ko.Spec.ContentHandlingStrategy)
	}
	if r.ko.Spec.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Spec.IntegrationID)
	}
	if r.ko.Spec.IntegrationResponseKey != nil {
		res.SetIntegrationResponseKey(*r.ko.Spec.IntegrationResponseKey)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range r.ko.Spec.ResponseParameters {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		res.SetResponseParameters(f4)
	}
	if r.ko.Spec.ResponseTemplates != nil {
		f5 := map[string]*string{}
		for f5key, f5valiter := range r.ko.Spec.ResponseTemplates {
			var f5val string
			f5val = *f5valiter
			f5[f5key] = &f5val
		}
		res.SetResponseTemplates(f5)
	}
	if r.ko.Spec.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*r.ko.Spec.TemplateSelectionExpression)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateIntegrationResponseOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateIntegrationResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateIntegrationResponse", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ContentHandlingStrategy != nil {
		ko.Spec.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		ko.Spec.ContentHandlingStrategy = nil
	}
	if resp.IntegrationResponseId != nil {
		ko.Status.IntegrationResponseID = resp.IntegrationResponseId
	} else {
		ko.Status.IntegrationResponseID = nil
	}
	if resp.IntegrationResponseKey != nil {
		ko.Spec.IntegrationResponseKey = resp.IntegrationResponseKey
	} else {
		ko.Spec.IntegrationResponseKey = nil
	}
	if resp.ResponseParameters != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.ResponseParameters {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		ko.Spec.ResponseParameters = f3
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.ResponseTemplates {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		ko.Spec.ResponseTemplates = f4
	} else {
		ko.Spec.ResponseTemplates = nil
	}
	if resp.TemplateSelectionExpression != nil {
		ko.Spec.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		ko.Spec.TemplateSelectionExpression = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateIntegrationResponseInput, error) {
	res := &svcsdk.UpdateIntegrationResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*r.ko.Spec.ContentHandlingStrategy)
	}
	if r.ko.Spec.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Spec.IntegrationID)
	}
	if r.ko.Status.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*r.ko.Status.IntegrationResponseID)
	}
	if r.ko.Spec.IntegrationResponseKey != nil {
		res.SetIntegrationResponseKey(*r.ko.Spec.IntegrationResponseKey)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f5 := map[string]*string{}
		for f5key, f5valiter := range r.ko.Spec.ResponseParameters {
			var f5val string
			f5val = *f5valiter
			f5[f5key] = &f5val
		}
		res.SetResponseParameters(f5)
	}
	if r.ko.Spec.ResponseTemplates != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range r.ko.Spec.ResponseTemplates {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		res.SetResponseTemplates(f6)
	}
	if r.ko.Spec.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*r.ko.Spec.TemplateSelectionExpression)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteIntegrationResponseOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteIntegrationResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteIntegrationResponse", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteIntegrationResponseInput, error) {
	res := &svcsdk.DeleteIntegrationResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.IntegrationID != nil {
		res.SetIntegrationId(*r.ko.Spec.IntegrationID)
	}
	if r.ko.Status.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*r.ko.Status.IntegrationResponseID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.IntegrationResponse,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
