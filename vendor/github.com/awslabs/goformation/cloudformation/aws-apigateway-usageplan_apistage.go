package cloudformation

import (
	"encoding/json"
)

// AWSApiGatewayUsagePlan_ApiStage AWS CloudFormation Resource (AWS::ApiGateway::UsagePlan.ApiStage)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html
type AWSApiGatewayUsagePlan_ApiStage struct {

	// ApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
	ApiId *Value `json:"ApiId,omitempty"`

	// Stage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
	Stage *Value `json:"Stage,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_ApiStage) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.ApiStage"
}

func (r *AWSApiGatewayUsagePlan_ApiStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
